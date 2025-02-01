package main

import (
	"filemanager/internal/filecontroller"
	"os"

	"flag"
	"fmt"
)

func main()  {
	createFlag := flag.String("create", "", "Specify type: a file or dir to create")
	pathFlag := flag.String("path", ".", "Specify the path where to create the file or directory")
	nameFlag := flag.String("name", "", "Specify the file or directory name")

	flag.Parse()

	if *createFlag == "" {
		fmt.Println("[-] Error: You must specify a file or directory to create using -create flag.")
		os.Exit(1)
	}

	filecontroller.CreateFileOrDirectory(createFlag, nameFlag, pathFlag)
	
	
}