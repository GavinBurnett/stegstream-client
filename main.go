// stegstream-client project main.go
package main

import (
	"fmt"
	"os"
)

// Main: program entry point
func main() {

	var hiddenFileExtracted bool = false
	var hiddenFileName string

	exitCode := 0

	// Hardcode command line arguments for testing
	// testArgs := []string{"", ""}
	// testArgs[0] = "./stegstream-client"
	// testArgs[1] = "http://localhost:8080/Audio"
	// os.Args = testArgs

	if os.Args != nil {

		if DEBUG == true {
			fmt.Println(len(os.Args), UI_Arguments, os.Args)
		}

		if len(os.Args) == 1 {
			// No user arguments given - display help
			fmt.Println(UI_Help)
		}
		if len(os.Args) == 2 {
			if IsStringHelpArgument(os.Args[1]) {
				// User has given help argument - display help
				fmt.Println(UI_Help)
			} else {
				// User has given server URL as argument

				// Get container file from server URL
				if GetServerFile(os.Args[1]) == true {
					// Extract hidden file from container file
					hiddenFileExtracted, hiddenFileName = Unsteg(TEMP_FILE)

					// Wipe container file
					if FileExists(TEMP_FILE) {

						if WipeFile(TEMP_FILE) == false {
							exitCode = 1
							fmt.Println(fmt.Sprintf(UI_WipeFileError, TEMP_FILE))
						}
					}

					if hiddenFileExtracted == true {
						fmt.Println(UI_ExtractedHiddenFile, hiddenFileName)
					} else {
						exitCode = 1
						fmt.Println(UI_FailedToExtractHiddenFile)
					}
				} else {

					// Delete container file
					if FileExists(TEMP_FILE) {
						if DeleteFile(TEMP_FILE) == false {
							// File deleted
						} else {
							exitCode = 1
							fmt.Println(UI_DeleteError, TEMP_FILE)
						}
					}

					exitCode = 1
					fmt.Println(UI_FailedToDownloadContainerFile)
				}
			}
		}
		if len(os.Args) == 3 {

			if IsStringHelpArgument(os.Args[1]) || IsStringHelpArgument(os.Args[2]) {
				// User has given help argument - display help
				fmt.Println(UI_Help)
			} else {
				// User has given local command and local file as arguments

				if os.Args[1] == UI_LocalFile {

					if FileExists(os.Args[2]) {

						// Extract hidden file from container file
						hiddenFileExtracted, hiddenFileName = Unsteg(os.Args[2])

						if hiddenFileExtracted == true {
							fmt.Println(UI_ExtractedHiddenFile, hiddenFileName)
						} else {
							exitCode = 1
							fmt.Println(UI_FailedToExtractHiddenFile)
						}

					} else {
						exitCode = 1
						fmt.Println(UI_FileNotFound, os.Args[2])
					}

				} else {
					exitCode = 1
					fmt.Println(UI_InvalidArgs)
				}
			}
		}
		if len(os.Args) > 3 {
			// Too many arguments - display error
			exitCode = 1
			fmt.Println(UI_InvalidArgs)
		}

	} else {
		// No arguments
		exitCode = 1
		fmt.Println(UI_NoParametersGiven)
	}

	os.Exit(exitCode)
}

// IsStringHelpArgument: Returns true if given string is a help argument, false if it is not
func IsStringHelpArgument(_theString string) bool {

	isHelpArgument := false

	if len(_theString) > 0 {

		switch _theString {
		case "?":
			isHelpArgument = true
		case "/?":
			isHelpArgument = true
		case "-?":
			isHelpArgument = true
		case "--?":
			isHelpArgument = true
		case "h":
			isHelpArgument = true
		case "/h":
			isHelpArgument = true
		case "-h":
			isHelpArgument = true
		case "--h":
			isHelpArgument = true
		case "help":
			isHelpArgument = true
		case "/help":
			isHelpArgument = true
		case "-help":
			isHelpArgument = true
		case "--help":
			isHelpArgument = true
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_theString:"+_theString))
	}

	return isHelpArgument
}
