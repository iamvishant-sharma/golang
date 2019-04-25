package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute(){
	out, err := exec.Command("/bin/sh", "-c", "ssh non-prod pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Command executed successfully")
	output := string(out[:])
	fmt.Printf(output)
}

func main(){
	if runtime.GOOS == "windows" {
		fmt.Println("Cannot execute command on windows")
	} else {
		execute()
	}
}
