package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// HiddenFileData: Data structure that describes the file hidden in the container file
type HiddenFileData struct {
	magicNumber int64
	fileName    [FILENAME_LENGTH]byte
	steps       int64
	spacing     int64
}

// Unsteg: Extracts hidden file from given container file
func Unsteg(_containerFile string) (bool, string) {

	var fileUnhidden bool = true
	var containerFileDesc *os.File
	var hiddenFileData HiddenFileData
	var hiddenFileDataRead bool = false
	var hiddenFile string = ""
	var steps int64 = -1
	var spacing int64 = -1
	const START_OFFSET int64 = 2000
	var offset int64 = -1
	var readByte []byte
	var containerFileReadCounter int64 = -1
	var hiddenFileBytesRead int = 0
	var containerFileReadError bool = false
	var hiddenFileTotalBytesRead int = 0
	const BUFFER_SIZE int64 = 1000
	var hiddenFileBuffer []byte
	var finalHiddenFileBuffer []byte
	var hiddenFileDesc *os.File
	var containerFileBytesWritten int = 0
	var err error

	// If file name for container file has been passed in
	if len(_containerFile) > 0 {

		// If container file exists
		if FileExists(_containerFile) {

			// Open container file
			containerFileDesc, err = os.OpenFile(_containerFile, os.O_RDONLY, 0)

			if err == nil {

				// Get hidden file data from container file
				hiddenFileDataRead, hiddenFileData = ReadHiddenFileData(containerFileDesc)

				if hiddenFileDataRead == true {

					if DEBUG == true {
						fmt.Println(fmt.Sprintf(UI_HiddenFileData, hiddenFileData.spacing, hiddenFileData.steps, string(bytes.Trim(hiddenFileData.fileName[:], "\x00"))))
					}

					hiddenFile = string(bytes.Trim(hiddenFileData.fileName[:], "\x00"))
					steps = hiddenFileData.steps
					spacing = hiddenFileData.spacing

					if len(hiddenFile) > 0 && hiddenFile != "" {

						if steps != -1 && spacing != -1 {

							// Set initial offset point to read hidden file from container file
							offset = START_OFFSET + 1

							// Create buffer to store hidden file data
							hiddenFileBuffer = make([]byte, BUFFER_SIZE)

							if int64(len(hiddenFileBuffer)) != BUFFER_SIZE {
								fmt.Println(UI_NoBufferMemory)
								fileUnhidden = false
							} else {

								// Check hidden file does not exist on disk
								if FileExists(hiddenFile) == false {

									// Create hidden file on disk
									hiddenFileDesc, err = os.OpenFile(hiddenFile, os.O_RDWR|os.O_CREATE, 0777)

									if err != nil {
										fmt.Println(fmt.Sprintf(UI_HiddenFileNotCreated, hiddenFile, err))
										fileUnhidden = false
									} else {

										// Loop around container file and read the hidden file data out of it one byte at a time
										for containerFileReadCounter = 0; containerFileReadCounter != steps; containerFileReadCounter++ {

											if UPDATE_UI == true {
												// Update UI
												fmt.Printf("\r" + UI_ExtractingHiddenFile + ".  ")
											}

											// Read a byte from the container file at the current offset point
											readByte = make([]byte, 1)
											hiddenFileBytesRead, err = containerFileDesc.ReadAt(readByte, offset)

											if err != nil {
												if err == io.EOF {
													// If end of container file found, display an error and stop reading data
													fmt.Println(fmt.Sprintf(UI_EOF, _containerFile))
													fileUnhidden = false
													containerFileReadError = true
													break
												} else {
													// Error reading container file - display the error and stop reading data
													fmt.Println(fmt.Sprintf(UI_FileReadError, _containerFile, err))
													fileUnhidden = false
													containerFileReadError = true
													break
												}
											}

											if UPDATE_UI == true {
												// Update UI
												fmt.Printf("\r" + UI_ExtractingHiddenFile + ".. ")
											}

											if hiddenFileBytesRead != 1 {
												fmt.Println(fmt.Sprintf(UI_FileReadError, _containerFile, ""))
												fileUnhidden = false
												containerFileReadError = true
												break
											}

											// Add byte read to hidden file buffer
											hiddenFileBuffer[hiddenFileTotalBytesRead] = readByte[0]
											hiddenFileTotalBytesRead++

											// Increment offset by spacing value
											offset = offset + spacing

											// If buffer is full, write the contents to the hidden file on disk
											if int64(hiddenFileTotalBytesRead) == BUFFER_SIZE {

												containerFileBytesWritten, err = hiddenFileDesc.Write(hiddenFileBuffer)

												if err != nil {
													fmt.Println(fmt.Sprintf(UI_FileWriteError, hiddenFile, err))
													fileUnhidden = false
													containerFileReadError = true
												}

												if int64(containerFileBytesWritten) != BUFFER_SIZE {
													fmt.Println(fmt.Sprintf(UI_FileWriteError, hiddenFile, ""))
													fileUnhidden = false
													containerFileReadError = true
												}

												// Clear hidden file buffer
												hiddenFileBuffer = nil
												hiddenFileBuffer = make([]byte, BUFFER_SIZE)

												// Set buffer counter back to 0
												hiddenFileTotalBytesRead = 0
											}

										} // End byte at a time for loop

										if UPDATE_UI == true {
											// Update UI
											fmt.Printf("\r" + UI_ExtractingHiddenFile + "...")
										}

										if containerFileReadError == false {

											// Write any data left in the buffer to the hidden file
											finalHiddenFileBuffer = make([]byte, hiddenFileTotalBytesRead)

											if len(finalHiddenFileBuffer) != hiddenFileTotalBytesRead {
												fmt.Println(UI_NoBufferMemory)
												fileUnhidden = false
											} else {

												for counter := 0; counter != hiddenFileTotalBytesRead; counter++ {
													finalHiddenFileBuffer[counter] = hiddenFileBuffer[counter]
												}

												containerFileBytesWritten, err = hiddenFileDesc.Write(finalHiddenFileBuffer)

												if err != nil {
													fmt.Println(fmt.Sprintf(UI_FileWriteError, hiddenFile, err))
													fileUnhidden = false
												}

												if containerFileBytesWritten != hiddenFileTotalBytesRead {
													fmt.Println(fmt.Sprintf(UI_FileWriteError, hiddenFile, ""))
													fileUnhidden = false
												}
											}
										} else {
											fmt.Println(fmt.Sprintf(UI_FileWriteError, hiddenFile, ""))
											fileUnhidden = false
										}
									}
								} else {
									fmt.Println(UI_FileExistsError, hiddenFile)
									fileUnhidden = false
								}
							}

							if UPDATE_UI == true {
								// Update UI
								fmt.Printf("\r")
							}

							// Close down file handles for container and hidden files
							containerFileDesc.Close()

							hiddenFileDesc.Sync()
							hiddenFileDesc.Close()

						} else {
							fmt.Println(UI_HiddenFileDataNotFound)
							fileUnhidden = false
						}
					} else {
						fmt.Println(UI_HiddenFileNameNotFound)
						fileUnhidden = false
					}
				} else {
					fmt.Println(UI_HiddenFileDataNotFound)
					fileUnhidden = false
				}
			} else {
				fmt.Println(UI_FileOpenError, err)
				fileUnhidden = false
			}
		} else {
			fmt.Println(UI_FileNotFound, _containerFile)
			fileUnhidden = false
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_containerFile: "+_containerFile))
		fileUnhidden = false
	}

	return fileUnhidden, hiddenFile
}

// ReadHiddenFileData: Reads the hidden file data stored in the container file
func ReadHiddenFileData(_containerFile *os.File) (bool, HiddenFileData) {

	var hiddenFileData HiddenFileData
	var readContainerFileSize []byte
	var hiddenFileDataOffset int64 = -1
	const READ_CONTAINER_FILE_SIZE_BUFFER int = 8
	var dataRead = false
	var err error

	if _containerFile != nil {

		// Create instance of data structure
		hiddenFileData = HiddenFileData{}

		// Create byte array to read data structure into
		readContainerFileSize = make([]byte, READ_CONTAINER_FILE_SIZE_BUFFER)
		if len(readContainerFileSize) != READ_CONTAINER_FILE_SIZE_BUFFER {
			fmt.Println(UI_NoBufferMemory)
		} else {

			// Move to location of data structure in container file and read the data structure into the byte array
			_containerFile.ReadAt(readContainerFileSize, HIDDEN_FILE_DATA_OFFSET)
			hiddenFileDataOffset = int64(binary.BigEndian.Uint64(readContainerFileSize))
			_containerFile.Seek(hiddenFileDataOffset, 0)
			err = binary.Read(_containerFile, binary.LittleEndian, &hiddenFileData.magicNumber)

			if err != nil {
				fmt.Println(UI_ReadMagicNumberFailed, err)
				hiddenFileData.magicNumber = -1
			} else {
				err = binary.Read(_containerFile, binary.LittleEndian, &hiddenFileData.fileName)
				if err != nil {
					fmt.Println(UI_ReadFileNameFailed, err)
				} else {
					err = binary.Read(_containerFile, binary.LittleEndian, &hiddenFileData.steps)
					if err != nil {
						fmt.Println(UI_ReadStepsFailed, err)
						hiddenFileData.steps = -1
					} else {
						err = binary.Read(_containerFile, binary.LittleEndian, &hiddenFileData.spacing)
						if err != nil {
							fmt.Println(UI_ReadSpacingFailed, err)
							hiddenFileData.spacing = -1
						} else {

							// Check if data structure contains magic number
							if DEBUG == true {
								fmt.Println(UI_MagicNumber, hiddenFileData.magicNumber)
							}

							if hiddenFileData.magicNumber == MAGIC_NUMBER {
								if DEBUG == true {
									fmt.Println(UI_MagicNumberFound)
								}
								dataRead = true
							} else {
								fmt.Println(UI_MagicNumberNotFound)
							}
						}
					}
				}
			}
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "containerFile: "+_containerFile.Name()))
	}

	return dataRead, hiddenFileData
}
