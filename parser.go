package safeincloud

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/pkg/errors"
)

// ParseFile takes a filenameAndPath and returns a Database struct.
func ParseFile(filenameAndPath string) (*Database, error) {
	f, err := os.Open(filenameAndPath)
	if err != nil {
		return nil, errors.Wrap(err, "os.Open() returned error")
	}
	defer f.Close()
	return ParseReader(f)
}

// ParseReader takes an io.Reader and returns a Database struct.
func ParseReader(r io.Reader) (*Database, error) {
	if r == nil {
		return nil, errors.New("io.Reader cannot be nil")
	}
	sic := &Database{}
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(sic); err != nil {
		return nil, errors.Wrap(err, "decoder.Decode() returned an error")
	}
	return sic, nil
}
