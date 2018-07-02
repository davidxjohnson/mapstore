package mapstore

import "testing"

func TestReadJSON(t *testing.T) {
	badPath := "/dev/sdc"
	badJSON := "./testdata/badtestdata.json"
	goodJSON := "./testdata/testdata.json"
	testTable := new(Table)
	err := readJSON(testTable, badPath) // test bad path
	if err == nil {
		t.Errorf("TestReadJSON: received false 'ok' reading bad file file '%s'.", badPath)
	}
	err = readJSON(testTable, badJSON) // test malformed JSON
	if err == nil {
		t.Errorf("TestReadJSON: received false 'ok' reading malformed JSON file '%s'.", badJSON)
	}
	err = readJSON(testTable, goodJSON) // test wellformed JSON
	if err != nil {
		t.Errorf("TestReadJSON: Failed to load valid JSON file '%s'- '%s'", goodJSON, err.Error())
	}
}

func TestMakeuuid(t *testing.T) {
	key := makeuuid()
	if key == "" {
		t.Errorf("TestMakeuuid: generated key was blank.")
	}
}

func TestWriteJSON(t *testing.T) {
	err := writeJSON(func() {}, "bogus") // try to write JSON using a bad object
	if err == nil {
		t.Errorf("TestWriteJSON: false 'ok' on write using invalid object func()")
	}
	badPath := "/dev/sdc"
	goodPath := "/tmp/" + makeuuid() + ".json"
	myTable := new(Table)
	err = writeJSON(myTable, badPath) // try to write to a forbidden location
	if err == nil {
		t.Errorf("TestWriteJSON: false 'ok' on write to bogus file '%s'.", badPath)
	}
	err = writeJSON(myTable, goodPath) // try to write to an allowed location
	if err != nil {
		t.Errorf("TestWriteJSON: Failed to write valid object to valid location '%s' - '%s'.", goodPath, err.Error())
	}
}

func TestReadFileStream(t *testing.T) {
	badPath := "./tmp/" + makeuuid() + ".json"
	goodPath := "./testdata/testdata.json"
	_, err := readFileStream(badPath)
	if err == nil {
		t.Errorf("TestWriteJSON: false 'ok' on read of non-existing file %s.", badPath)
	}
	_, err = readFileStream(goodPath)
	if err != nil {
		t.Errorf("TestWriteJSON: Failed to read valid file %s.", goodPath)
	}
}
