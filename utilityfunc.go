package mapstore

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// makeuuid function Creates a uuid as a unique record id
func makeuuid() (uuid string) {
	b := make([]byte, 16)
	rand.Read(b)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

// readJSON function gets data from a file into a structure
func readJSON(returnData interface{}, path string) (err error) {
	rawData, err := readFileStream(path)
	if err != nil {
		return
	}
	// convert the raw data into the structure passed by reference to this function
	err = json.Unmarshal(rawData, returnData)
	return
}

// writeJSON function puts data from structure to file
func writeJSON(sourceData interface{}, path string) (err error) {
	rawData, err := json.Marshal(sourceData)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, rawData, 0644)
	return
}

// readFileStream function returns a stream of bytes from a file, or bails if unable
func readFileStream(filePath string) (rawData []byte, err error) {
	// Open file
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fileHandle.Close() // closed this later
	// Read file
	rawData, err = ioutil.ReadAll(fileHandle)
	return
}
