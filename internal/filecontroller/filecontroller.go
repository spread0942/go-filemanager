package filecontroller

import (
	"filemanager/internal/filemanager"

	"fmt"
	"os"
)

func CreateFileOrDirectory(sElementType *string, name *string, path *string) {
	elementType, err := filemanager.StringToElementType(*sElementType)

	if elementType == filemanager.Unknown {
		fmt.Println("[+] Invalid type: " + err.Error())
		os.Exit(1)
	} else if elementType == filemanager.File {
		succ, err := filemanager.CreateFile(*name, *path)

		if succ {
			fmt.Println("[+] File created")
			os.Exit(0)
		} else {
			fmt.Println("[-] Error creating file: ", err)
			os.Exit(1)
		}
	} else if elementType == filemanager.Directory {
		succ, err := filemanager.CreateDirectory(*name, *path)

		if succ {
			fmt.Println("[+] Directory created")
			os.Exit(0)
		} else {
			fmt.Println("[-] Error creating directory: ", err)
			os.Exit(1)
		}
	}
}
