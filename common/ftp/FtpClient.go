package ftp

import (
	"github.com/kataras/iris/v12/context"
	"github.com/zihao-boy/zihao/common/utils"
	"github.com/zihao-boy/zihao/entity/dto/ls"
	"github.com/zihao-boy/zihao/entity/dto/resources"
	"github.com/zihao-boy/zihao/entity/dto/result"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"

	"gopkg.in/dutchcoders/goftp.v1"
)

// upload file

func UploadFile(filePath string, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.Path, "/") {
		path = resourcesFtpDto.Path
	} else {
		path = "/" + resourcesFtpDto.Path
	}

	if err = ftp.Cwd(path); err != nil {
		return err
	}

	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return err
	}

	defer file.Close()

	if err := ftp.Stor(path, file); err != nil {
		return err
	}
	return nil
}

// upload file

func DownloadFile(resWriter context.ResponseWriter, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.Path, "/") {
		path = resourcesFtpDto.Path
	} else {
		path = "/" + resourcesFtpDto.Path
	}
	resWriter.Header().Set("Content-Type", "application/octet-stream")

	err = ftp.Walk(path, func(path string, info os.FileMode, err error) error {
		_, err = ftp.Retr(path, func(r io.Reader) error {
			io.Copy(resWriter, r)
			return err
		})
		return nil
	})
	resWriter.Flush()
	return nil
}

func ListFile(resourcesFtpDto resources.ResourcesFtpDto) result.ResultDto {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return result.Error(err.Error())
	}

	defer ftp.Close()

	// TLS client authentication
	//config := tls.Config{
	//	InsecureSkipVerify: true,
	//	ClientAuth:         tls.RequestClientCert,
	//}

	//if err = ftp.AuthTLS(&config); err != nil {
	//	return result.Error(err.Error())
	//}

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return result.Error(err.Error())
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}

	//if err = ftp.Cwd(path); err != nil {
	//	return result.Error(err.Error())
	//}

	dirs, err := ftp.List(path)

	if err != nil {
		return result.Error(err.Error())
	}
	var lss = make([]ls.LsDto, 0)
	for _, fil := range dirs {
		lsrs := strings.Split(fil, ";")
		if len(lsrs) == 4 {
			name := strings.Trim(lsrs[3], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			lsDto := ls.LsDto{
				GroupName:    "d",
				Name:         name,
				Privilege:    strings.Split(lsrs[2], "=")[1],
				Size:         0,
				LastModified: strings.Split(lsrs[1], "=")[1],
			}
			lss = append(lss, lsDto)
		}
	}

	for _, fil := range dirs {
		lsrs := strings.Split(fil, ";")
		if len(lsrs) == 5 {
			size, _ := strconv.ParseInt(strings.Split(lsrs[1], "=")[1], 10, 64)
			name := strings.Trim(lsrs[4], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			lsDto := ls.LsDto{
				GroupName:    "-",
				Name:         name,
				Privilege:    strings.Split(lsrs[3], "=")[1],
				Size:         size,
				LastModified: strings.Split(lsrs[2], "=")[1],
			}
			lss = append(lss, lsDto)
		}

	}
	return result.SuccessData(lss)
}

func UploadFileByReader(f multipart.File, resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}

	if strings.Contains(path, "/") {
		pos := strings.LastIndex(path, "/")
		dirStr := path[0:pos]
		dirs := strings.Split(dirStr, "/")
		for i := 0; i < len(dirs); i++ {
			dir := dirs[i]
			if utils.IsEmpty(dir) || "/" == dir {
				continue
			}

			rs, _ := ftp.Stat(dir)
			if len(rs) > 0 {
				ftp.Cwd(dir)
				continue
			}
			ftp.Mkd(dir)
			ftp.Cwd(dir)
		}
	}

	if err := ftp.Stor(path, f); err != nil {
		return err
	}
	return nil
}

func DeleteFile(resourcesFtpDto resources.ResourcesFtpDto) error {
	var (
		ftp *goftp.FTP
		err error
	)
	// connect ftp server
	if ftp, err = goftp.Connect(resourcesFtpDto.Ip + ":" + resourcesFtpDto.Port); err != nil {
		return err
	}

	defer ftp.Close()

	if err = ftp.Login(resourcesFtpDto.Username, resourcesFtpDto.Passwd); err != nil {
		return err
	}

	var path string

	if strings.HasPrefix(resourcesFtpDto.CurPath, "/") {
		path = resourcesFtpDto.CurPath
	} else {
		path = "/" + resourcesFtpDto.CurPath
	}
	if resourcesFtpDto.FileGroupName == "-" {
		err = ftp.Dele(path)
	} else {
		err = deleteDirAndFile(path, ftp)
	}

	if err != nil {
		return err
	}

	return nil
}

// delete dir and file

func deleteDirAndFile(dirPath string, ftp *goftp.FTP) error {

	dirs, err := ftp.List(dirPath)

	if err != nil {
		return err
	}
	for _, fil := range dirs {
		lsrs := strings.Split(fil, ";")
		if len(lsrs) == 4 {
			name := strings.Trim(lsrs[3], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			err = deleteDirAndFile(path.Join(dirPath, name), ftp)
			if err != nil {
				return err
			}
		}

		if len(lsrs) == 5 {
			name := strings.Trim(lsrs[4], " ")
			name = strings.ReplaceAll(name, "\r", "")
			name = strings.ReplaceAll(name, "\n", "")
			err = ftp.Dele(path.Join(dirPath, name))
			if err != nil {
				return err
			}
		}
	}

	err = ftp.Rmd(dirPath)

	if err != nil {
		return err
	}
	return nil
}
