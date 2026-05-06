package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
)

type model strcut {
	cursor int
	choices []string
}

var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		BorderStyle(lipgloss.RoundedBorder()).
		Padding(1, 2)

var selectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")).
		Bold(true)

func initialModel() model {
	return model{
		choices: []string{
			"About",
			"Projects",
			"Skills",
			"Github",
			"Contact",
			"Exit",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
				return m, tea.Quit

		case "up", "k":
				if m.cursor > 0 {
						m.cursor
			}

		case "down", "j":
				if m.cursor < len(m.choices)-1 {
						m.cursor++
			}

		case "enter":
				selected := m.choices[m.cursor]

				if selected == "Exit" {
						return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Jogo's sshStation")
	s += "\n\n"

	for i, choice := range m.choices {

		cursor := " "

		if m.cursor == i {
				cursor = ">"
				s += selectedStyle.Render(fmt.Sprintf("%s %s\n", cursor, choice))
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	}

	s += "\n↑ ↓ navigate • enter select • q quit"

	return s
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	return initialModel(), []tea.ProgramOption{
		tea.WithAltScreen(),
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "22097"
	}

	server, err := wish.NewServer(
		wish.WithAddress("0.0.0.0:"+port),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithNoClientAuth(),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting SSH server on port", port)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
