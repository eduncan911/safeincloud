// Package safeincloud parses SafeInCloud's exported XML for use in GoLang.
//
// It is intended to be used to help convert from SafeInCloud to another
// password manager such as my SafeInCloud-to-LastPass converter:
//
// https://github.com/eduncan911/sic2lp
//
// Usage
//
// This is a GoLang library package intended for import.
//
//     $ go get github.com/eduncan911/safeincloud
//
// This package is setup for simple one-liners:
//
//     db, err := safeincloud.ParseFile("/path/to/exported/safeincloud.xml")
//     if err != nil {
//         panic(err)
//     }
//
//     for _, c := range db.Cards {
//         // do what you like with the Card
//     }
//
// Examples
//
// For several more examples, see the GoDocs with embedded examples:
//
// https://godoc.org/github.com/eduncan911/safeincloud
//
// Release Notes
//
// 1.0.0
// * Initial release.
//
package safeincloud
