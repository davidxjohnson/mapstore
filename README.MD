# Mapstore

## A Generic data Structure for In-memory Tables

### Introduction:

This project was started as a learning experiment and evolved into something generally useful. It's a set of methods applied to a flexible table structure: `map[string]map[string]interface{}` ... which is essentially an indexed table of rows. The use of `interface{}` is limited by your imagination.

Example: `Contact[a38e2a13-a03d-4643-96c7-d50b3d4f2228]['Phone']` returns "888-555-1212"

There is no replication, concurrency, locking, transaction processing etc. It's a simple, in-memory table.

Currently being used in the [go-DDD](https://github.com/davidxjohnson/go-DDD) experiment.


### Testing this package:

See contents of [runtest.sh](./runtest.sh)

### Usage:

**Initialize a table from disk:**

```go
// loads stored data into an in-memory table.
// if file does not exist, will create it and return an empty table.
myTable, err = NewTable("./testdata/testdata.json")
if err != nil {
  // bad path, forbidden or invalid file format
}
```

**Commit table to disk:**

```go
// writes table content to path used to populate it.
// this operation is usually called automatically whenever data changes.
err := myTable.CommitTable()
if err != nil {
  // forbidden, i/o level error
}
```

**Add a row to the table:**

```go
// pupulate a Row with data
var myRow = Row{"city": "Hammonton", "email": "lthomas26@go.com", "first": "Janie", "last": "Mitchell", "number": "13", "phone": "718-864-5281", "state": "NJ", "street": "E. Armstrong Rd.", "zip": "08037"}
key, ok := myTable.AddRow(myRow)
if !ok {
  // unusual error
}
```

**Delete a row:**

```go
ok = myTable.DeleteRow("a38e2a13-a03d-4643-96c7-d50b3d4f2228")
if !ok {
  // index value doesn't exist?
}
```

**Update a row:**

```go
myRow["phone"] = "609-555-1212"
ok := myTable.UpdateRow("a38e2a13-a03d-4643-96c7-d50b3d4f2228", myRow)
if ok {
  // index value doesn't exist?
}
```

**Find a row by its index value:**

```go
myRow, ok := myTable.FindRowByID("a38e2a13-a03d-4643-96c7-d50b3d4f2228")
if !ok {
  // index value doesn't exist?
}
```

**Find a row by data values:**
```go
// uses query string value pairs as input, returns Table with matching results
myQuery := url.Values{}
myQuery.Add("state", "VA")
myQuery.Add("state", "NJ")
myNewTable, ok := myTable.QueryTable(myQuery) // find state=NJ or state=VA
if !ok {
  // no matches
}
```
