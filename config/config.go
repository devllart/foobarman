package config

import "golang.org/x/text/language"

// init
var LangTag language.Tag
var err error

var Lang string = "ru"

func init() {
	LangTag, err = language.Parse(Lang)
}
