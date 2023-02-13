package server

import (
	"errors"
	"fmt"
	"github.com/letheliu/hhjc-devops/common/innerNet/encrypt"
	"github.com/letheliu/hhjc-devops/common/innerNet/header"
	"github.com/letheliu/hhjc-devops/common/innerNet/io"
	"github.com/letheliu/hhjc-devops/entity/dto/innerNet"
	"net"
	"strings"
	"time"
)

var USERCHANBUFFERSIZE = 1024
var READBUFFERSIZE = 65535

var UserPrivileges []*innerNet.InnerNetPrivilegeDto

type User struct {
	Client        string
	Protocol      string
	RemoteTunIp   string
	LocalTunIp    string
	Token         string
	Key           string
	TunToConnChan chan string
	ConnToTunChan chan string
	Conn          net.Conn
	Logout        func(client string)
	HeartbeatTime time.Time
	Privileges    []*innerNet.InnerNetPrivilegeDto
}

func NewUser(client string, protocol string, tun string, token string, conn net.Conn, logout func(string)) *User {
	key := string(encrypt.GetAESKey([]byte(token)))
	var privs []*innerNet.InnerNetPrivilegeDto
	for _, privilege := range UserPrivileges {
		if privilege.Token != token {
			continue
		}
		privs = append(privs, privilege)
	}

	return &User{
		Client:        client,
		Protocol:      protocol,
		LocalTunIp:    tun,
		RemoteTunIp:   tun,
		Token:         token,
		Key:           key,
		TunToConnChan: make(chan string, USERCHANBUFFERSIZE),
		ConnToTunChan: make(chan string, USERCHANBUFFERSIZE),
		Conn:          conn,
		Logout:        logout,
		HeartbeatTime: time.Now().Add(60 * time.Second),
		Privileges:    privs,
	}
}

func (user *User) Start() {
	encryptKey := encrypt.GetAESKey([]byte(user.Token))
	//read from client, write to channel
	go func() {
		for {
			var err error
			var data []byte

			if user.Protocol == "tcp" {
				data, err = io.ReadPacket(user.Conn)
			}

			if err != nil {
				fmt.Println("io.ReadPacket err", err)
				user.Close()
				return
			}

			//fmt.Println("接受到数据 data",data)

			if ln := len(data); ln > 0 {
				data, err = encrypt.DecryptAES(data, encryptKey)
				if err == nil {
					if proto, src, dst, err := header.GetBase(data); err == nil {
						err = user.checkSrcHostPrivilege(dst)
						if err != nil {
							continue
						}
						remoteIp, _ := header.ParseAddr(src)
						user.RemoteTunIp = remoteIp
						//user, ok := lm.Users[client];
						Snat(data, user.LocalTunIp)
						user.ConnToTunChan <- string(data)
						fmt.Println("From %v client: client:%v, protocol:%v, len:%v, src:%v, dst:%v", user.Protocol, user.Client, proto, ln, src, dst)
						continue
					}
				}
				ipData := string(data)
				//fmt.Println("接受到数据 ipData",ipData)
				if ipData != "ping" {
					continue
				}
				user.HeartbeatTime = time.Now().Add(60 * time.Second)
				data, _ = encrypt.EncryptAES(data, encryptKey)
				_, err = io.WritePacket(user.Conn, data)
				//fmt.Println("回写ping ipData",ipData,user.HeartbeatTime)

				if err != nil {
					fmt.Println("deal data err", err)
				}

			}
		}
	}()

	//read from channel, write to client
	go func() {
		for {
			datas, ok := <-user.TunToConnChan
			if !ok {
				user.Close()
				return
			}
			data := []byte(datas)
			if ln := len(data); ln > 0 {
				if proto, src, dst, err := header.GetBase(data); err == nil {
					Dnat(data, user.RemoteTunIp)
					if endata, err := encrypt.EncryptAES(data, encryptKey); err == nil {
						if user.Protocol == "tcp" {
							_, err = io.WritePacket(user.Conn, endata)

						}

						if err != nil {
							user.Close()
							return
						}
						fmt.Println("To %v client: client:%v, protocol:%v, len:%v, src:%v, dst:%v", user.Protocol, user.Client, proto, ln, src, dst)
					}
				}
			}
		}
	}()

	//check heart beat
	go func() {
		for {
			time.Sleep(60 * time.Second)
			if user.HeartbeatTime.After(time.Now()) {
				continue
			}
			fmt.Println("检测现成", user.HeartbeatTime, time.Now())
			user.Close()
			return
		}

	}()
}

func (user *User) Close() {
	user.Logout(user.Client)
	go func() {
		defer func() {
			recover()
		}()
		close(user.TunToConnChan)
	}()

	go func() {
		defer func() {
			recover()
		}()
		close(user.ConnToTunChan)
	}()

	go func() {
		defer func() {
			recover()
		}()
		user.Conn.Close()
	}()

}

func (user *User) checkSrcHostPrivilege(dst string) error {
	if user.Privileges == nil || len(user.Privileges) < 1 {
		return nil
	}
	if !strings.Contains(dst, ":") {
		return nil
	}

	dsts := strings.Split(dst, ":")
	dstIp := dsts[0]
	dstPort := dsts[1]
	for _, privilege := range user.Privileges {
		if privilege.TargetUserId == dstIp && privilege.TargetPort == dstPort {
			return nil
		}
		if privilege.TargetUserId == dstIp && privilege.TargetPort == "*" {
			return nil
		}
	}

	return errors.New("no privilege")

}
