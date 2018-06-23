package mapstore

import (
	"net/url"
	"testing"
)

var myTable Table // global var representing our test table
var addKey string
var addFieldName = "email"
var addRow = Row{"city": "Hammonton", "email": "lthomas26@go.com", "first": "Janie", "last": "Mitchell", "number": "13", "phone": "718-864-5281", "state": "NJ", "street": "E. Armstrong Rd.", "zip": "08037"}
var addFieldValue interface{}

func TestNewTable(t *testing.T) {
	filePath := "/tmp/" + makeuuid() + ".json"
	_, ok := NewTable(filePath) // test creation of new table storage file
	if !ok {
		t.Errorf("TestInit: refused to create file '%s'", filePath)
	}
	// read the test data for the rest of the test cases
	filePath = "./testdata/testdata.json"
	myTable, ok = NewTable(filePath) // myTable is global
	if !ok {
		t.Errorf("TestNewTable: Table object create failed using filePath '%s'", filePath)
	}
}

func TestAdd(t *testing.T) {
	var myRow = addRow
	key, ok := myTable.AddRow(myRow)
	if !ok {
		t.Errorf("TestAdd: failed to add new row %+v", myRow)
	}
	addKey = key // store in global var for later use
	addFieldValue = myRow[addFieldName]
}

func TestFindById(t *testing.T) {
	var id = addKey
	var fieldName = addFieldName
	var fieldValue = addFieldValue
	myRow, ok := myTable.FindRowByID(id)
	if !ok {
		t.Errorf("TestFindById: myRow[%s]: not found.", id)
	}
	if myRow[fieldName] != fieldValue {
		t.Errorf("TestFindById: myRow[%s][%s]: not does not match \"%s\".", id, fieldName, fieldValue)
	}
}

func TestUpdate(t *testing.T) {
	var id = addKey
	bogusID := makeuuid()
	var fieldName = addFieldName
	var fieldValue = "jboyd12@telegraph.co.uk"
	myRow := addRow
	myRow[fieldName] = fieldValue
	ok := myTable.UpdateRow(bogusID, myRow)
	if ok {
		t.Errorf("TestUpdate: myRow[%s]: was found, but shouldn't have (key is bogus).", bogusID)
	}
	ok = myTable.UpdateRow(id, myRow)
	if !ok {
		t.Errorf("TestUpdate: myRow[%s]: not found.", id)
	}
	if myRow[fieldName] != fieldValue {
		t.Errorf("TestUpdate: myRow[%s][%s]: not does not match \"%s\".", id, fieldName, fieldValue)
	}
}

func TestDelete(t *testing.T) {
	var id = addKey
	bogusID := makeuuid()
	ok := myTable.DeleteRow(bogusID)
	if ok {
		t.Errorf("TestDelete: myRow[%s]: returned false 'ok for deletion of bogus key.", bogusID)
	}
	ok = myTable.DeleteRow(id)
	if !ok {
		t.Errorf("TestDelete: myRow[%s]: was not found for deletion.", id)
	}
	_, ok = myTable.FindRowByID(id)
	if ok {
		t.Errorf("TestDelete: myRow[%s]: was found (should be deleted).", id)
	}
}

func TestQuery(t *testing.T) {
	myQuery := url.Values{}
	myQuery.Add("state", "VA")
	myQuery.Add("state", "NJ")
	_, ok := myTable.QueryTable(myQuery)
	if !ok {
		t.Errorf("TestQuery: No results found.")
	}
}

func TestCommit(t *testing.T) {
	ok := myTable.CommitTable()
	if !ok {
		t.Errorf("TestCommit: failed to write '%s'.", myTable.filePath)
	}
}
