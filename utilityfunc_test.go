package mapstore

import "testing"

func TestReadJSON(t *testing.T) {
	badJSON := "./testdata/badtestdata.json"
	goodJSON := "./testdata/testdata.json"
	testTable := new(Table)
	ok := readJSON(testTable, badJSON) // test malformed JSON
	if ok {
		t.Errorf("TestReadJSON: received false 'ok' reading malformed JSON file '%s'", badJSON)
	}
	ok = readJSON(testTable, goodJSON) // test wellformed JSON
	if !ok {
		t.Errorf("TestReadJSON: Failed to load valid JSON file '%s'", goodJSON)
	}
}

func TestMakeuuid(t *testing.T) {
	key := makeuuid()
	if key == "" {
		t.Errorf("TestMakeuuid: generated key was blank.")
	}
}

func TestWriteJSON(t *testing.T) {
	ok := writeJSON(func() {}, "bogus") // try to write JSON using a bad object
	if ok {
		t.Errorf("TestWriteJSON: false 'ok' on write using invalid object func()")
	}
	badPath := "/dev/sdc"
	goodPath := "/tmp/" + makeuuid() + ".json"
	myTable := new(Table)
	ok = writeJSON(myTable, badPath) // try to write to a forbidden location
	if ok {
		t.Errorf("TestWriteJSON: false 'ok' on write to bogus file '%s'.", badPath)
	}
	ok = writeJSON(myTable, goodPath) // try to write to an allowed location
	if !ok {
		t.Errorf("TestWriteJSON: Failed to write valid object to valid location '%s'.", goodPath)
	}
}

func TestReadFileStream(t *testing.T) {
	badPath := "./tmp/" + makeuuid() + ".json"
	goodPath := "./testdata/testdata.json"
	_, ok := readFileStream(badPath)
	if ok {
		t.Errorf("TestWriteJSON: false 'ok' on read of non-existing file %s.", badPath)
	}
	_, ok = readFileStream(goodPath)
	if !ok {
		t.Errorf("TestWriteJSON: Failed to read valid file %s.", goodPath)
	}
}
