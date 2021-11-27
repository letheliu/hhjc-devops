package shell

import (
	"github.com/zihao-boy/zihao/entity/dto/host"
	"golang.org/x/crypto/ssh"
)

func ExecShell(host host.HostDto, cmd string) error {
	client, err := ssh.Dial("tcp", host.Ip, &ssh.ClientConfig{
		User:            host.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(host.Passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	//defer client.Close()

	if err != nil {
		return err
	}
	session, err := client.NewSession()
	defer session.Close()
	defer client.Close()

	// 使用内存
	session.Output(cmd)

	return nil

}
