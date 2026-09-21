package ui

type Model struct {
	CurrentDir string
	Username   string
	Hostname   string
	History    []string
}

func NewModel() *Model {

	return &Model{
		CurrentDir: "home/jagath",
		Username:   "jagath",
		Hostname:   "portfolio",
		History:    []string{},
	}
}
