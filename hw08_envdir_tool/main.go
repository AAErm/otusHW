package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("wona be more then 2 args")
		return
	}

	env, err := ReadDir(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Exit(RunCmd(os.Args[2:], env))
}
