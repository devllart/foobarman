package fmtc

import "devllart/foobarman/src/mapsi"

var manyReplaceMap = map[string]string{
	"Оливки":          "%GОливки%C",
	"Coca-Cola":       "%RCoca-Cola%C",
	"Листья":          "%GЛистья%C",
	"Красный биттер":  "%RКрасный биттер%C",
	"красный биттер":  "%Rкрасный биттер%C",
	"Красный":         "%RКрасный%C",
	"красный":         "%Rкрасный%C",
	"синий":           "%Bсиний%C",
	"зелёный":         "%Gзелёный%С",
	"мята":            "%Gмята%C",
	"Мята":            "%GМята%C",
	"лаймовая цедра":  "%Gлаймовая цедра%C",
	"Лаймовая цедра":  "%GЛаймовая цедра%C",
	"лаймовая":        "%Gлаймовая%C",
	"Лаймовая":        "%GЛаймовая%C",
	"лаймовый сок":    "%Gлаймовый сок%C",
	"Лаймовый сок":    "%GЛаймовый сок%C",
	"лаймовый":        "%G%лаймовыйC",
	"Лаймовый":        "%GЛаймовый%C",
	"лайма":           "%Gлайма%C",
	"лайм":            "%Gлайм%C",
	"Лайма":           "%GЛайма%C",
	"Лайм":            "%GЛайм%C",
	"лимонной цедрой": "%Yлимонной цедрой%C",
	"Лимонной цедрой": "%YЛимонной цедрой%C",
	"лимонный сок":    "%Yлимонный сок%C",
	"Лимонный сок":    "%YЛимонный сок%C",
	"лимоный":         "%Yлимоный%C",
	"Лимоный":         "%YЛимоный%C",
	"лимона":          "%Yлимона%C",
	"Лимона":          "%YЛимона%C",
	"лимон":           "%Yлимон%C",
	"Лимон":           "%YЛимон%C",
}

var manyReplace = []string{}

func AddManyReplace(manyReplaceMap map[string]string) {
	// for key, val := range manyReplaceMap {
	// 	manyReplace = append(manyReplace, key, val)
	// }

	newMapsi := mapsi.New(manyReplaceMap)
	newMapsi = newMapsi.SortLongestKeysFirst()

	manyReplace = append(manyReplace, mapsi.MapsiStringsToSliceKeysFirst(&newMapsi)...)
}
