package main

import (
	"fmt"
	"testing"
)

func TestUnsteg(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		containerFile  string
		expectedResult bool
		hideFile       string
	}{
		{"NoParameterData", "", false, ""},
		{"ContainerFileDoesNotExist", "", false, ""},
		{"ContainerFileEmpty", "EmptyFile", false, ""},
		{"NoHiddenFileData", "1000ByteContainerFile", false, ""},
		{"ValidFile", "Waves.mp3", true, "HideFile.txt"},
	}

	// Set up test data
	CreateEmptyFile("EmptyFile")
	CreateFile("1000ByteContainerFile", 1000)

	// Write name of function being tested to test results file
	LogResult("Unsteg")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result, hideFile := Unsteg(currentTest.containerFile)

			if result != currentTest.expectedResult || hideFile != currentTest.hideFile {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t %s Expected: %t %s", currentTest.containerFile, result, hideFile, currentTest.expectedResult, currentTest.hideFile) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t %s Expected: %t %s", currentTest.containerFile, result, hideFile, currentTest.expectedResult, currentTest.hideFile) + " - PASS")
			}
		})
	}

	// Clean up test data
	DeleteFile("EmptyFile")
	DeleteFile("1000ByteContainerFile")
	DeleteFile("HideFile.txt")
}
