package main

import (
	"fmt"

	"github.com/hashicorp/go-memdb"
)

// Person create a sample struct
type Person struct {
	Email string
	Name  string
	Age   int
}

func main() {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert a new person
	p := &Person{
		Email: "joe@aol.com",
		Name:  "Joe",
		Age:   30,
	}
	if err := txn.Insert("person", p); err != nil {
		panic(err)
	}

	// Commit the transaction
	txn.Commit()

	txn = db.Txn(true)

	if err := txn.Insert("person", &Person{
		Email: "shunkkb@awa.fm",
		Name:  "Shun",
		Age:   27,
	}); err != nil {
		panic(err)
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// Lookup by email
	//raw, err := txn.First("person", "id", "joe@aol.com")
	raw, err := txn.First("person", "id", "shunkkb@awa.fm")
	if err != nil {
		panic(err)
	}

	// Say hi!
	fmt.Printf("Hello %s!", raw.(*Person).Name)
}
