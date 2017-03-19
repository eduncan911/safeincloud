package safeincloud

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFileOpenError(t *testing.T) {
	t.Parallel()

	// arrange

	// act
	db, err := ParseFile("blah")

	// assert
	assert.Nil(t, db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "os.Open()")
}

func TestParseReaderNilError(t *testing.T) {
	t.Parallel()

	// arrange

	// act
	db, err := ParseReader(nil)

	// assert
	assert.Nil(t, db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "io.Reader")
}

func TestParseReaderDecoderError(t *testing.T) {
	t.Parallel()

	// arrange
	r := strings.NewReader("")

	// act
	db, err := ParseReader(r)

	// assert
	assert.Nil(t, db)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decoder.Decode()")
}

func TestValidateBase64Attachments(t *testing.T) {
	t.Parallel()

	// arrange
	attachments := map[string]string{
		"file":  "the-evolution-of-house-cats.pdf.gpg",
		"image": "cutecat.jpg",
	}

	// act
	db, err := ParseFile("testdata/safeincloud.xml")

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, db)

	found := false
	for _, c := range db.Cards {
		if c.Title == "Attachment Example" {
			found = true
			assert.EqualValues(t, 1, len(c.Files))
			assert.EqualValues(t, 1, len(c.Images))

			// validate the file attachment
			f := c.Files[0]
			assert.EqualValues(t, attachments["file"], f.Name)
			file, err := ioutil.ReadFile("testdata/" + attachments["file"])
			if err != nil {
				panic(err)
			}
			assert.True(t, reflect.DeepEqual(f.Value, file), "file bytes did not match")

			// images cannot be validated as SafeInCloud compresses
			// ALL IMAGES to jpeg 80% regardless of type or size which
			// completely changes the image itself.  :(
			//
		}
	}
	assert.True(t, found, "card title not found 'Attachment Example'")
}
