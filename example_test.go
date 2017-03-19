package safeincloud_test

import (
	"fmt"

	"github.com/eduncan911/safeincloud"
)

func Example() {

	// using SafeInCloud is a simple 1 liner.
	db, err := safeincloud.ParseFile("testdata/safeincloud.xml")

	// all errors are wrapped with full stack trace
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	// use the returned safeincloud.Database struct as you like.
	fields, deleted, templates, stars := 0, 0, 0, 0
	for _, c := range db.Cards {
		fields = fields + len(c.Fields)
		if c.Deleted {
			deleted++
		}
		if c.Template {
			templates++
		}
		if c.Star {
			stars++
		}
	}
	fmt.Println("Total Cards:", len(db.Cards), "Fields:", fields, "Labels:", len(db.Labels))
	fmt.Println("Deleted:", deleted, "Template:", templates, "Stars:", stars)
	// Output:
	// Total Cards: 19 Fields: 61 Labels: 7
	// Deleted: 1 Template: 8 Stars: 2
}
