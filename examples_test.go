package safeincloud_test

import (
	"fmt"

	"github.com/eduncan911/safeincloud"
)

func ExampleParseFile() {

	// using SafeInCloud is a simple 1 liner.
	db, err := safeincloud.ParseFile("testdata/safeincloud.xml")
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	// iterate and pull out data
	var files []safeincloud.File
	var images []safeincloud.Image
	for _, c := range db.Cards {
		if len(c.Files) != 0 {
			for _, f := range c.Files {
				files = append(files, f)
			}
		}
		if len(c.Images) != 0 {
			for _, i := range c.Images {
				images = append(images, i)
			}
		}
	}
	fmt.Println("Total Files:", len(files), "Images:", len(images))
	// Output:
	// Total Files: 1 Images: 1
}
