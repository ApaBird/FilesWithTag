package pythonmoduleapi

import (
	"fmt"
	"net"
)

type MetaDataModule struct {
	Port string
	Ip   string
	conn net.Conn
}

func NewModule(ip, port string) MetaDataModule {
	module := MetaDataModule{
		Port: port,
		Ip:   ip,
	}
	var err error
	module.conn, err = net.Dial("tcp", ip+":"+port)
	if err != nil {
		panic(err)
	}

	return module
}

func (module MetaDataModule) SendData(data string) {
	module.conn.Write([]byte(data + "\n"))
}

func (module MetaDataModule) Scan() string {
	ans := ""
	for {
		buf := make([]byte, 1)
		_, err := module.conn.Read(buf)
		if err != nil {
			fmt.Println("[ERROR]", err.Error())
			return ""
		}
		if string(buf[0]) == "\n" {
			return ans
		}
		ans += string(buf[0])
	}
}
