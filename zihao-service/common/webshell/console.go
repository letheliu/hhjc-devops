package webshell

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/kataras/iris/v12/websocket"
	"net"
	"time"
	"unicode/utf8"

	gossh "golang.org/x/crypto/ssh"
)

type ptyRequestMsg struct {
	Term     string
	Columns  uint32
	Rows     uint32
	Width    uint32
	Height   uint32
	Modelist string
}

type SshLoginModel struct {
	Addr     string
	UserName string
	Pwd      string
	PemKey   string
	PtyCols  uint32
	PtyRows  uint32
}

var (
	channels = make(map[string]gossh.Channel)
	clients  = make(map[string]*gossh.Client)
)

func sshConnect(login SshLoginModel) (client *gossh.Client, ch gossh.Channel, err error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = login.UserName
	if login.Pwd == "" {
		return
	} else {
		config.Auth = []gossh.AuthMethod{gossh.Password(login.Pwd)}
	}
	config.HostKeyCallback = func(hostname string, remote net.Addr, key gossh.PublicKey) error { return nil }
	client, err = gossh.Dial("tcp", login.Addr, config)
	if err != nil {
		return
	}
	channel, incomingRequests, err := client.Conn.OpenChannel("session", nil)
	if err != nil {
		return
	}
	go func() {
		for req := range incomingRequests {
			if req.WantReply {
				req.Reply(false, nil)
			}
		}
	}()
	modes := gossh.TerminalModes{
		gossh.ECHO:          1,
		gossh.TTY_OP_ISPEED: 14400,
		gossh.TTY_OP_OSPEED: 14400,
	}
	var modeList []byte
	for k, v := range modes {
		kv := struct {
			Key byte
			Val uint32
		}{k, v}
		modeList = append(modeList, gossh.Marshal(&kv)...)
	}
	modeList = append(modeList, 0)
	req := ptyRequestMsg{
		Term:     "xterm",
		Columns:  login.PtyCols,
		Rows:     login.PtyRows,
		Width:    login.PtyCols * 8,
		Height:   login.PtyRows * 8,
		Modelist: string(modeList),
	}
	ok, err := channel.SendRequest("pty-req", true, gossh.Marshal(&req))
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("e001")
		return
	}
	ok, err = channel.SendRequest("shell", true, nil)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("e002")
		return
	}
	ch = channel
	return
}

type Request struct {
	Operate   string `json:"operate"`
	SshKey    string `json:"sshKey"`
	TypeIp    string `json:"typeIp"`
	Command   string `json:"command"`
	WinWidth  int64  `json:"winWidth"`
	WinHeight int64  `json:"winHeight"`
}

func WebSocketHandler(data []byte, connId string, nsConn *websocket.NSConn) {
	var channel gossh.Channel
	var client *gossh.Client
	//done := make(chan bool, 2)
	req := Request{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return
	}
	if req.Operate == "connect" {

		if !checkUserToken(req.SshKey) {
			return
		}
		loginInfo := getServerInfo(req.TypeIp)
		if loginInfo.Addr == "" {
			return
		}
		client, channel, err = sshConnect(loginInfo)
		if err != nil {
			return
		}

		channels[connId] = channel
		clients[connId] = client

	} else {
		channel = channels[connId]

		if channel == nil {
			return
		}
		if _, err := channel.Write([]byte(req.Command)); nil != err {
			return
		}
	}
	if req.Operate != "connect" {
		return
	}
	go func() {
		for {
			defer func() {
				//done <- true
			}()
			br := bufio.NewReader(channel)
			buf := []byte{}
			t := time.NewTimer(time.Millisecond * 100)
			defer t.Stop()
			r := make(chan rune)
			go func() {
				for {
					x, size, err := br.ReadRune()
					if err != nil {
						return
					}
					if size > 0 {
						r <- x
					}
				}
			}()
			for {
				select {
				case <-t.C:
					if len(buf) != 0 {
						//err = ws.WriteMessage(websocket.TextMessage, buf)
						//returnMsg(buf)
						mg := websocket.Message{
							Body:     buf,
							IsNative: true,
						}
						nsConn.Conn.Write(mg)
						buf = []byte{}

					}
					t.Reset(time.Millisecond * 100)
				case d := <-r:
					if d != utf8.RuneError {
						p := make([]byte, utf8.RuneLen(d))
						utf8.EncodeRune(p, d)
						buf = append(buf, p...)
					} else {
						buf = append(buf, []byte("@")...)
					}
				}
			}
		}
	}()
	//<-done

}

func CloseSsh(connId string) {
	if _, ok := channels[connId]; ok {
		channels[connId].Close()
		delete(channels, connId)
	}

	if _, ok := clients[connId]; ok {
		clients[connId].Close()
		delete(clients, connId)
	}
}

func checkUserToken(token string) bool {
	return true
}

func getServerInfo(hostId string) SshLoginModel {

	return SshLoginModel{
		UserName: "root",
		Pwd:      "Hjx965067",
		Addr:     "175.27.221.224:22",
		PtyCols:  100,
		PtyRows:  100,
	}
}
