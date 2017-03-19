[![GoDoc](https://godoc.org/github.com/eduncan911/safeincloud?status.svg)](https://godoc.org/github.com/eduncan911/safeincloud)
[![Build Status](https://travis-ci.org/eduncan911/safeincloud.svg?branch=master)](https://travis-ci.org/eduncan911/safeincloud)
[![Coverage Status](https://coveralls.io/repos/github/eduncan911/safeincloud/badge.svg?branch=master)](https://coveralls.io/github/eduncan911/safeincloud?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/eduncan911/safeincloud)](https://goreportcard.com/report/github.com/eduncan911/safeincloud)
[![MIT License](https://img.shields.io/npm/l/mediaelement.svg)](https://eduncan911.mit-license.org/)

# safeincloud
Package safeincloud parses SafeInCloud's exported XML for use in GoLang.

It is intended to be used to help convert from SafeInCloud to another
password manager such as my SafeInCloud-to-LastPass converter:

<a href="https://github.com/eduncan911/sic2lp">https://github.com/eduncan911/sic2lp</a>

### Usage
This is a GoLang library package intended for import.

	$ go get github.com/eduncan911/safeincloud

This package is setup for simple one-liners:

	db, err := safeincloud.ParseFile("/path/to/exported/safeincloud.xml")
	if err != nil {
	    panic(err)
	}
	
	for _, c := range db.Cards {
	    // do what you like with the Card
	}

### Examples
For several more examples, see the GoDocs with embedded examples:

<a href="https://godoc.org/github.com/eduncan911/safeincloud">https://godoc.org/github.com/eduncan911/safeincloud</a>

### Release Notes
1.0.0
* Initial release.

## Table of Contents

* [Imported Packages](#pkg-imports)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-imports">Imported Packages</a>

- [github.com/pkg/errors](https://godoc.org/github.com/pkg/errors)

## <a name="pkg-index">Index</a>
* [type Card](#Card)
* [type Database](#Database)
  * [func ParseFile(filenameAndPath string) (\*Database, error)](#ParseFile)
  * [func ParseReader(r io.Reader) (\*Database, error)](#ParseReader)
* [type Field](#Field)
* [type File](#File)
  * [func (f \*File) MarshalXML(e \*xml.Encoder, start xml.StartElement) error](#File.MarshalXML)
  * [func (f \*File) UnmarshalXML(d \*xml.Decoder, start xml.StartElement) error](#File.UnmarshalXML)
* [type Image](#Image)
* [type Label](#Label)

#### <a name="pkg-examples">Examples</a>
* [Package](#example_)
* [ParseFile](#example_ParseFile)

#### <a name="pkg-files">Package files</a>
[database.go](./database.go) [doc.go](./doc.go) [file.go](./file.go) [parser.go](./parser.go) 

## <a name="Card">type</a> [Card](./database.go#L27-L45)
``` go
type Card struct {
    XMLName  xml.Name `xml:"card"`
    ID       string   `xml:"id,attr"`
    Title    string   `xml:"title,attr"`
    Color    string   `xml:"color,attr"`
    Symbol   string   `xml:"symbol,attr"`
    Notes    string   `xml:"notes"`
    Deleted  bool     `xml:"deleted,attr"`
    Star     bool     `xml:"star,attr"`
    Template bool     `xml:"template,attr"`
    Fields   []Field  `xml:"field"`
    Images   []Image  `xml:"image"`
    Files    []File   `xml:"file"`

    // LabelIDs can be rogue in SafeInCloud, meaning there could be a
    // number that doesn't exist as an actual label.  So, we'll just
    // store labels in Database.Labels instead for mapping.
    LabelIDs []int `xml:"label_id"`
}
```
Card represents a global Card in the database.

## <a name="Database">type</a> [Database](./database.go#L13-L17)
``` go
type Database struct {
    XMLName xml.Name `xml:"database"`
    Cards   []Card   `xml:"card"`
    Labels  []Label  `xml:"label"`
}
```
Database represents an SafeInCloud export.

This struct is not intended encompass an entire SafeInCloud export. It
may not parse back into a functional SafeInCloud XML format. Instead, it
is a subset of most the valuable fields and is intended to be used to
convert to a new password manager.

Card history is also not captured with this version.

### <a name="ParseFile">func</a> [ParseFile](./parser.go#L12)
``` go
func ParseFile(filenameAndPath string) (*Database, error)
```
ParseFile takes a filenameAndPath and returns a Database struct.

### <a name="ParseReader">func</a> [ParseReader](./parser.go#L22)
``` go
func ParseReader(r io.Reader) (*Database, error)
```
ParseReader takes an io.Reader and returns a Database struct.

## <a name="Field">type</a> [Field](./database.go#L48-L53)
``` go
type Field struct {
    XMLName   xml.Name `xml:"field"`
    Name      string   `xml:"name,attr"`
    FieldType string   `xml:"type,attr"`
    Value     string   `xml:",chardata"`
}
```
Field represents a Card.Field.

## <a name="File">type</a> [File](./file.go#L18-L24)
``` go
type File struct {
    XMLName xml.Name `xml:"file"`
    // Name is the file name as uploaded into SafeInCloud.
    Name string `xml:"name,attr"`
    // Value is the byte slice of the file's contents.
    Value []byte
}
```
File represents a Card.File.

TODO: To make this library useful for creating SafeInCloud XML for importing
into SIC, a MarshalXML() will need to be created to zlib compress and base64
encode the results.  At the time of this writing, the library is only
Unmarshalling to be used to move from SafeInCloud to something else.

### <a name="File.MarshalXML">func</a> (\*File) [MarshalXML](./file.go#L66)
``` go
func (f *File) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```
MarshalXML is NOT IMPLEMENTED.

TODO: To make this library useful for creating SafeInCloud XML for importing
into SIC, a MarshalXML() will need to be created to zlib compress and base64
encode the results.  At the time of this writing, the library is only
Unmarshalling to be used to move from SafeInCloud to something else.

### <a name="File.UnmarshalXML">func</a> (\*File) [UnmarshalXML](./file.go#L33)
``` go
func (f *File) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```
UnmarshalXML will base64 decode and zlib decompress the file
attachments to store in the File.Value []byte slice.

## <a name="Image">type</a> [Image](./database.go#L56-L60)
``` go
type Image struct {
    XMLName xml.Name `xml:"image"`
    // Value is Base64 encoded.
    Value []byte `xml:",chardata"`
}
```
Image represents a Card.Image.

## <a name="Label">type</a> [Label](./database.go#L20-L24)
``` go
type Label struct {
    XMLName xml.Name `xml:"label"`
    ID      int      `xml:"id,attr"`
    Name    string   `xml:"name,attr"`
}
```
Label represents a global Label in the database.

- - -
Generated by [godoc2ghmd](https://github.com/eduncan911/godoc2ghmd)