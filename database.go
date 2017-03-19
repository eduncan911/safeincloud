package safeincloud

import "encoding/xml"

// Database represents an SafeInCloud export.
//
// This struct is not intended encompass an entire SafeInCloud export. It
// may not parse back into a functional SafeInCloud XML format. Instead, it
// is a subset of most the valuable fields and is intended to be used to
// convert to a new password manager.
//
// Card history is also not captured with this version.
type Database struct {
	XMLName xml.Name `xml:"database"`
	Cards   []Card   `xml:"card"`
	Labels  []Label  `xml:"label"`
}

// Label represents a global Label in the database.
type Label struct {
	XMLName xml.Name `xml:"label"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}

// Card represents a global Card in the database.
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

// Field represents a Card.Field.
type Field struct {
	XMLName   xml.Name `xml:"field"`
	Name      string   `xml:"name,attr"`
	FieldType string   `xml:"type,attr"`
	Value     string   `xml:",chardata"`
}

// Image represents a Card.Image.
type Image struct {
	XMLName xml.Name `xml:"image"`
	// Value is Base64 encoded.
	Value []byte `xml:",chardata"`
}
