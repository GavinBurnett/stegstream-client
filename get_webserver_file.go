package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// GetServerFile: Downloads the container file from the website
func GetServerFile(_url string) bool {

	var serverFile *os.File
	var err error
	var serverResponse *http.Response
	var gotFile bool = false

	if len(_url) > 0 {

		fmt.Printf("\r" + UI_GettingFile)

		// Check URL is valid
		_, err = url.ParseRequestURI(_url)
		if err != nil {
			fmt.Println(UI_InvalidURL)
		} else {

			// Create the container file
			serverFile, err = os.Create(TEMP_FILE)
			if err != nil {
				fmt.Println(UI_FileCreateError, err)
			} else {

				// Get the container file data from the server
				serverResponse, err = http.Get(_url)
				if err != nil {
					fmt.Println(UI_ServerError, err)
				} else {

					// Check server response
					if serverResponse.StatusCode != http.StatusOK {
						fmt.Println(UI_ServerError, err)
					} else {

						// Write container file data to the file
						_, err = io.Copy(serverFile, serverResponse.Body)
						if err != nil {
							fmt.Println(UI_FileCopyError, err)
						} else {
							// Got file from server
							fmt.Println("\r" + UI_GotFile + "       ")
							gotFile = true
						}
					}
					serverResponse.Body.Close()
				}
			}

			serverFile.Close()

		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_url: "+_url))
	}

	return gotFile
}
