package pythonmoduleapi

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Start() {
	//MetaDataModule\\venv\\Scripts\\python.exe MetaDataModule\\main.py --port=5000 --ip=127.0.0.1
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err.Error())
	}

	path = strings.Replace(path, "FilesWithTag.exe", "", -1)

	cmd := exec.Command(path+`MetaDataModule\venv\Scripts\python.exe `, path+`MetaDataModule\main.py`, `--port=5000`, `--ip=127.0.0.1`)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("[ERROR]", err.Error())
	}
	println("=>", string(output))
	time.Sleep(time.Second * 5)

	metaDataPython, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	metaDataPython.Write([]byte("hello\n"))
}
