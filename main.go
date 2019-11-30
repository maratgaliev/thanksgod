package main

import (
	"github.com/kyokomi/emoji"
	"github.com/maratgaliev/thanksgod/connectors/github"
	"github.com/maratgaliev/thanksgod/decoder"
)

func main() {
	data := decoder.GetReposList()
	for _, element := range data {
		github.Star(element.Username, element.Name)
		emoji.Println(":sparkling_heart:", element.Name)
	}
}
