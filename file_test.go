package safeincloud

import (
	"bytes"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalXMLErrors(t *testing.T) {
	t.Parallel()

	// arrange
	data := []struct {
		input  string
		result string
	}{
		{`<fil2e></file>`, "xml.DecodeElement"},
		{`<file name="test.gpg">ABC</file>`, "base64.DecodeString"},
		{`<file name="test.gpg">0x00</file>`, "zlib.NewReader"},
	}

	for _, test := range data {
		b := []byte(test.input)
		r := bytes.NewReader(b)
		d := xml.NewDecoder(r)
		f := File{}

		// act
		err := d.Decode(&f)

		// assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), test.result)
	}
}

func TestMarshalXMLNotImplemented(t *testing.T) {
	t.Parallel()

	// arrange
	f := File{
		Name:  "test.gpg",
		Value: []byte{0x0, 0x0},
	}
	w := bytes.NewBuffer(nil)
	e := xml.NewEncoder(w)

	// act
	err := e.Encode(&f)

	// assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not implemented")
}
