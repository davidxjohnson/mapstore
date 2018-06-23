package mapstore

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
func readJSON(returnData interface{}, path string) (ok bool) {
	rawData, ok := readFileStream(path)
	// convert the raw data into the structure passed by reference to this function
	err := json.Unmarshal(rawData, returnData)
	if err != nil {
		log.Print("readJSON Unmarshal: " + err.Error())
		ok = false
	}
	return
}

// writeJSON function puts data from structure to file
func writeJSON(sourceData interface{}, path string) (ok bool) {
	rawData, err := json.Marshal(sourceData)
	if err != nil {
		log.Print("writeJSON Marshal: '" + path + "' " + err.Error())
		ok = false
		return
	}
	err = ioutil.WriteFile(path, rawData, 0644)
	if err != nil {
		log.Print("writeJSON WriteFile: '" + path + "' " + err.Error())
		ok = false
		return
	}
	ok = true
	return
}

// readFileStream function returns a stream of bytes from a file, or bails if unable
func readFileStream(filePath string) (rawData []byte, ok bool) {
	// Open file
	fileHandle, err := os.Open(filePath)
	if err != nil {
		ok = false
		log.Print("readFileStream: " + err.Error())
		return
	}
	defer fileHandle.Close() // closed this later
	// Read file
	rawData, err = ioutil.ReadAll(fileHandle)
	if err != nil { // nearly impossible to throw this error and just as hard to test
		ok = false
		log.Print("readFileStream: " + err.Error())
		return
	}
	ok = true
	return
}
