package main

import (
	"fmt"
	"os"
)

// FileExists: Check the given file exists
func FileExists(_file string) bool {

	fileExists := false

	if len(_file) > 0 {

		// Try to get file info
		_, err := os.Stat(_file)

		// If any errors occur on getting file info, file does not exist
		if os.IsNotExist(err) || err != nil {
			if DEBUG == true {
				fmt.Println(UI_FileNotFound, _file)
			}
			fileExists = false
		} else {
			// File info found - file exists
			if DEBUG == true {
				fmt.Println(UI_FileFound, _file)
			}
			fileExists = true
		}
	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_file:"+_file))
	}

	return fileExists
}

// DeleteFile: Delete the given file
func DeleteFile(_file string) bool {

	var deleteError bool = false
	var err error

	if len(_file) > 0 {
		err = os.Remove(_file)
		if err != nil {
			fmt.Println(UI_DeleteError, _file, err)
			deleteError = true
		} else {
			// File deleted
			if DEBUG == true {
				fmt.Println(UI_FileDeleted, _file)
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
		deleteError = true
	}

	return deleteError
}
