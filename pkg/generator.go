// orcgen generates files from HTML -
// any static webpage can be informed, or even the HTML itself.
// The file will be generated according the choosen extension.
package orcgen

import (
	"log"

	"github.com/luabagg/orc-generator/internal"
	"github.com/luabagg/orc-generator/internal/generator"
)

const (
	// Valid extension types constants.

	// PDF const.
	PDF = internal.PDF
	// PNG const.
	PNG = internal.PNG
	// JPEG const.
	JPEG = internal.JPEG
)

var (
	// FullPage sets the pages to be converted. If false, only the first page is selected.
	FullPage bool = true
)

// New starts a new orc-generator - Director contains the available methods.
//
// ext is the extension to be converted to (use the defined constants).
//
// Connect and Close are used for the Browser connection controll.
// ConvertWebpage and ConvertHTML are used for page conversion.
//
// There are a set of setters for specific config.
func New(ext internal.Ext) *internal.Director {
	var gen generator.Generator
	if gen = internal.Build(ext, FullPage); gen == nil {
		log.Fatal("Generator not found.")
	}

	return internal.NewDirector(gen).Connect()
}
