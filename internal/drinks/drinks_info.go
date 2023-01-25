package drinks

import (
	"devllart/foobarman/internal/config"
	"fmt"
	"strings"
)

type DrinkInfo struct {
	Type          string
	Taste         string
	Alc           float64
	AviableVolume []float64
	Prices        []float64
	Description   string
}

func (drink DrinkInfo) PrettyDescription() {
	if config.ShowDescription {
		lines := strings.Split(drink.Description, "\n")

		for _, line := range lines {
			fmt.Printf("      | %s\n", line)
		}
	}
}

var AviableDrinks = map[string]DrinkInfo{
	"Эбботтс биттер": {
		Type:          "Добавка/Ароматический биттер",
		Taste:         "Горький",
		Alc:           41.5,
		AviableVolume: []float64{.1},
		Prices:        []float64{20.50},
		Description: `
Старинный биттер из бобов тонка производили в Балтиморе еще в 1865 году. 
Спустя 60 лет после запрета на употребление тонка в США, 
оригинальная рецептура была воссоздана экспертом по биттерам Бобом Петри
с помощью газохроматографического анализа закрытой бутылки тех времен.
`,
	},
	"Апероль": {
		Type:          "Слабоалкогольное/Апероль",
		Taste:         "Горький",
		Alc:           22,
		AviableVolume: []float64{0.75, 1},
		Prices:        []float64{27.99, 34.71},
		Description: `
Знаменитый итальянский аперитив производят с 1919 года по секретной рецептуре братьев Барбьери из Падуи.
В состав напитка входит более 30 компонентов, включая цедру горьких апельсинов, травы и ревень.`,
	},
	"Просекко": {
		Type:          "Слабоалкогольное/Сухое вино",
		Taste:         "Полусладкое",
		AviableVolume: []float64{0.75},
		Prices:        []float64{14.94},
		Description: `
Согласно Апелласьону 2009 года, Просекко могут носить лишь вина с виноградников данного региона.
Собранный в ручную виноград Глера подвергается брожению в вакуумной таре, что позволяет сохранять его вкус намного дольше.
Игристое вино отличается выразительным ароматом с оттенками весенних белых цветов, зеленого яблока и цитрусовых плодов,
обладает притяным вкусом с нежным фруктовым наполнением и тонкой минеральностью в послевкусии.
Идеально в гастрономическом сочетании с морепродуктами и твёрдыми сырами.`,
	},
}
