package main

import (
	"fmt"
	"os"

	hidden "github.com/tobychui/goHidden"
)

/*
	goHidden example file
	Author: tobychui

	This script will create a folder and set it to hidden
*/

func main() {
	//Create a test folder
	os.Mkdir("./test", 0775)

	//Set it to hidden
	err := hidden.HideFile("./test")
	if err != nil {
		panic(err)
	}

	//Check if it is hidden
	isHidden, err := hidden.IsHidden("./test", false)
	if err != nil {
		panic(err)
	}

	if isHidden {
		fmt.Println("The test folder is now hidden")
	} else {
		fmt.Println("Unable to hide the test folder")
	}
}
