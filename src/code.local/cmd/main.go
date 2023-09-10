package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	argLen := len(os.Args)
	dirName, err := os.Getwd()
	if err != nil {
		fmt.Println("Get Current Working Directory error, ", err)
		return
	}
	separator := string(os.PathSeparator)

	if argLen == 2 {
		if os.Args[1] == ".." {
			// path包对于处理windows路径不起作用
			index := strings.LastIndex(dirName, separator)
			dirName = dirName[:index]
		} else if os.Args[1][0] == '.' && os.Args[1][1] == '/' {
			relativePath := os.Args[1][2:]
			relativePath = strings.ReplaceAll(relativePath, `/`, `\`) //`
			dirName = dirName + separator + relativePath
		} else {
			dirName = os.Args[1]
		}
	}

	fmt.Println(dirName)
	exec.Command(`cmd`, `/c`, `explorer`, dirName).Start()

}
