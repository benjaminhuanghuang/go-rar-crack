package main

import (
	"fmt"
	"os/exec"
)

// const RAR_FILE = "50.rar"
const RAR_FILE = "unarMac.zip"
const OUTPUT = "./temp"
const DIR_FILE = ""

func cmdshell(rarFile string, pwd string) {
	options := []string{"-f", "-output-directory", OUTPUT, rarFile}
	if len(pwd) > 0 {
		options = append(options, "-p", pwd)
	}
	cmd := exec.Command("./unar", options...) //解压出来保存 D/test 上
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s\n", out)
	} else {
		if len(out) == 248 { //len 248 为成功，每个人不同
			fmt.Printf("密码为：%s \n", pwd)
		}
	}
}

func main() {
	cmdshell(RAR_FILE, "")
}
