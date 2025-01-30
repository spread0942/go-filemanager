package main

import (
	"filemanager/internal"
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

	if *createFlag == "file" {
		succ, err := utils.CreateFile(*nameFlag, *pathFlag)

		if (succ) {
			fmt.Println("[+] File created")
			os.Exit(0)
		} else {
			fmt.Println("[-] Error creating file: ", err)
			os.Exit(1)
		}
	} else if *createFlag == "dir" {
		succ, err := utils.CreateDirectory(*nameFlag, *pathFlag)

		if (succ) {
			fmt.Println("[+] Directory created")
			os.Exit(0)
		} else {
			fmt.Println("[-] Error creating directory: ", err)
			os.Exit(1)
		}
	}
	
}