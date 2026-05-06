package ui

type Model struct {
	CurrentDir string
	Username string
	Hostname string
	History []string
}

func NewModel() *Model {
	
		return &Model{
			CurrentDir: "home/jogo",
		  Username: "jogo",
			Hostname: "portfolio",
			History: []string{},
	}
} 
