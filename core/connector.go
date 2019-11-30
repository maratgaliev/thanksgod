package core

type Connector interface {
	Star(username, reponame string)
}
