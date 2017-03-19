package safeincloud

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/xml"
	"github.com/pkg/errors"
	"io"
)

// File represents a Card.File.
//
// TODO: To make this library useful for creating SafeInCloud XML for importing
// into SIC, a MarshalXML() will need to be created to zlib compress and base64
// encode the results.  At the time of this writing, the library is only
// Unmarshalling to be used to move from SafeInCloud to something else.
type File struct {
	XMLName xml.Name `xml:"file"`
	// Name is the file name as uploaded into SafeInCloud.
	Name string `xml:"name,attr"`
	// Value is the byte slice of the file's contents.
	Value []byte
}

type rawFile struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

// UnmarshalXML will base64 decode and zlib decompress the file
// attachments to store in the File.Value []byte slice.
func (f *File) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v rawFile
	if err := d.DecodeElement(&v, &start); err != nil {
		return errors.Wrap(err, "xml.DecodeElement returned an error")
	}
	f.Name = v.Name // name

	decoded, err := base64.StdEncoding.DecodeString(v.Value)
	if err != nil {
		return errors.Wrap(err, "base64.DecodeString returned an error")
	}

	br := bytes.NewReader(decoded)
	zr, err := zlib.NewReader(br)
	if err != nil {
		return errors.Wrap(err, "zlib.NewReader returned an error")
	}
	defer zr.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, zr); err != nil {
		return errors.Wrap(err, "could not zlib read the compressed data")
	}
	f.Value = buf.Bytes() // value
	return nil
}

// MarshalXML is NOT IMPLEMENTED.
//
// TODO: To make this library useful for creating SafeInCloud XML for importing
// into SIC, a MarshalXML() will need to be created to zlib compress and base64
// encode the results.  At the time of this writing, the library is only
// Unmarshalling to be used to move from SafeInCloud to something else.
func (f *File) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return errors.New("not implemented")
}
