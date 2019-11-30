package decoder

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/maratgaliev/thanksgod/models"
)

type tomlConfig struct {
	SolveMeta solveMeta `toml:"solve-meta"`
}

type solveMeta struct {
	InputImports []string `toml:"input-imports"`
}

func GetUsername(domain, username string) string {
	domainName := strings.Split(domain, ".")[0]
	if domain == "golang.org" {
		return domainName
	} else {
		return username
	}
}

func BuildRepo(url string) models.Repository {
	parsed := strings.Split(url, "/")
	domain, username, reponame := parsed[0], parsed[1], parsed[2]
	return models.Repository{Username: GetUsername(domain, username), Name: reponame, Domain: domain}
}

func GetReposList() []models.Repository {
	var config tomlConfig

	if _, err := toml.DecodeFile("Gopkg.lock", &config); err != nil {
		fmt.Println(err)
	}

	var repos []models.Repository
	data := config.SolveMeta.InputImports

	for _, element := range data {
		item := BuildRepo(element)
		repos = append(repos, item)
	}
	return repos
}
