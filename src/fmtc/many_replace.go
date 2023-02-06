package fmtc

var manyReplaceMap = map[string]string{
	"Оливки":         "%GОливки%C",
	"Coca-Cola":      "%RCoca-Cola%C",
	"Листья":         "%GЛистья%C",
	"красный":        "%Rкрасный%C",
	"синий":          "%Bсиний%C",
	"зелёный":        "%Gзелёный%С",
	"мята":           "%Gмята%C",
	"Мята":           "%GМята%C",
	"лаймовая цедра": "%Gлаймовая цедра%C",
	"Лаймовая цедра": "%GЛаймовая цедра%C",
	"%Gлаймовая%C":   "Лаймовая",
	"%GЛаймовая%C":   "%GЛаймовая цедра%C",
	"лаймовый сок":   "%Gлаймовый сок%C",
	"Лаймовый сок":   "%GЛаймовый сок%C",
	"лаймовый":       "%G%лаймовыйC",
	"Лаймовый":       "%GЛаймовый%C",
	"лайма":          "%Gлайма%C",
	"лайм":           "%Gлайм%C",
	"Лайма":          "%GЛайма%C",
	"Лайм":           "%GЛайм%C",
	"лимон":          "%Yлимон%C",
	"Лимон":          "%YЛимон%C",
}

var manyReplace = []string{}

func AddManyReplace(manyReplaceMap map[string]string) {
	for key, val := range manyReplaceMap {
		manyReplace = append(manyReplace, key)
		manyReplace = append(manyReplace, val)
	}
}
