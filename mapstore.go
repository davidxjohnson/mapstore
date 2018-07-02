// Package mapstore contains code for managing and persisting an all-purpose map[string]map[string]interface{} data structure
// This type of structure can be used to persist an indexed list of name/value pairs, the value being most anything ...
// - i.e. Contact[a38e2a13-a03d-4643-96c7-d50b3d4f2228][Phone] returns "888-555-1212"
package mapstore

import (
	"net/url"
	"os"
)

// Row represents a dictionary or tupel of interface{}
type Row map[string]interface{} // Field (key) and value pairs

// Table represents a collection of Rows
type Table struct {
	filePath string
	Rows     map[string]Row `json:"rows"`
}

// NewTable function reads data from disk into new Table object
func NewTable(filePath string) (t Table, err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) { // path/to/whatever does not exist
		t.filePath = filePath
		t.Rows = map[string]Row{}
		err = t.CommitTable()
		return
	}
	err = readJSON(&t, filePath)
	if err == nil {
		t.filePath = filePath
	}
	return
}

// CommitTable method writes entire datastore to disk
func (t Table) CommitTable() (err error) {
	err = writeJSON(t, t.filePath) // marshal and write to disk
	return
}

// FindRowByID method gets index to Rows with matching ID
func (t Table) FindRowByID(id string) (r Row, ok bool) {
	r, ok = t.Rows[id]
	return
}

// AddRow method populates a new Row record
func (t Table) AddRow(r Row) (id string, ok bool) {
	id = makeuuid()
	t.Rows[id] = r // Unique ID
	ok = true
	return
}

// UpdateRow method modifies an existing Row record
func (t Table) UpdateRow(id string, r Row) (ok bool) {
	_, ok = t.FindRowByID(id) // Find original Row Info
	if !ok {                  // Return if not found
		return
	}
	// Update Row Info
	t.Rows[id] = r // Update Row
	return
}

// DeleteRow method destroys a Row record
func (t Table) DeleteRow(id string) (ok bool) {
	_, ok = t.FindRowByID(id) // Find Row Info index
	if !ok {                  // Return if Not Found
		return
	}
	// Delete Row Info
	delete(t.Rows, id)
	return
}

// QueryTable method finds matching Row records and returns a Table object with all matching rows.
// For query purposes, each key in url.Values[key] is unique.
// Multiple keys in the map are considered "or" as are also multiple values for a single key.
// Keys in url.Values are table column names (fields) in map[key]map[field]string in the Table object.
func (t Table) QueryTable(pairs url.Values) (dt Table, ok bool) {
	dt.Rows = make(map[string]Row)
Rowloop:
	for id, r := range t.Rows { // Scan Rows
		foundMatch := false
	pairsloop:
		for qk, qa := range pairs { // check each key/value against Row Info
			for _, qv := range qa { // Might be more than one value to check
				if rv, idExists := r[qk]; idExists { // Check for matching key name
					if qv == rv { // Check for matching value
						foundMatch = true  // Got a match
						continue pairsloop // Get next pair as this value matches
					}
				}
			}
			continue Rowloop // Get next Row as no values matched
		}
		if foundMatch { // found one
			dt.Rows[id] = r // add to list
		}
	}
	ok = (len(dt.Rows) > 0)
	return
}
