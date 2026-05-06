package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
)

var (
	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true)

	outputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)
)

type model struct {
	input   textinput.Model
	history []string
}

func initialModel() model {
	ti := textinput.New()

	ti.Placeholder = "type a command..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	history := []string{
		"",
		"██╗ ██████╗  ██████╗  ██████╗ ",
		"██║██╔═══██╗██╔════╝ ██╔═══██╗",
		"██║██║   ██║██║  ███╗██║   ██║",
		"██║██║   ██║██║   ██║██║   ██║",
		"██║╚██████╔╝╚██████╔╝╚██████╔╝",
		"╚═╝ ╚═════╝  ╚═════╝  ╚═════╝ ",
		"",
		"Welcome to JogoOS v1.0",
		"Type 'help' to begin.",
		"",
	}

	return model{
		input:   ti,
		history: history,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func runCommand(cmd string) string {
	cmd = strings.TrimSpace(strings.ToLower(cmd))

	switch cmd {

	case "help":
		return `
Available commands:

about
projects
skills
github
contact
discord
reddit
server
whoami
neofetch
pwd
clear
exit
`

	case "about":
		return `
Designer, developer, organizer.

Background includes:
- Freelancing
- UI/UX
- Software development
- Community and event work

Focused on building:
- cleaner coordination
- sharper management
- reliable systems around people and projects

Based in Bengaluru.
`

	case "projects":
		return `
Projects in circulation:

1. Spoorthi Delicacy
   Brand & web
   https://www.spoorthidelicacy.in/

2. OpenBMTC
   Transit tooling
   https://github.com/jagath-sajjan/OPENBMTC

3. Fills Game
   Interactive web
   https://fillsgame.vercel.app/

4. Caleox Space Forum
   Community platform
   https://github.com/jagath-sajjan/caleox-spaceforum/releases/tag/v1.0.0
`

	case "skills":
		return `
Skills & Areas:

- UI/UX
- Software Development
- Freelancing
- Community Building
- Event Coordination
- Product Organization
- Management Systems
`

	case "github":
		return `
GitHub:
https://github.com/jagath-sajjan
`

	case "contact":
		return `
Mail:
jagathsajjan227@gmail.com

Portfolio:
https://jogodevs.vercel.app
`

	case "discord":
		return `
Discord:
jogohere
`

	case "reddit":
		return `
Reddit:
https://www.reddit.com/user/Cool_jagath/
`

	case "server":
		return `
The Bengaluru Hub:
https://discord.gg/Pg2pjBcx
`

	case "whoami":
		return "jogo"

	case "pwd":
		return "/home/jogo"

	case "neofetch":
		return `
OS: JogoOS 1.0
Host: Railway
Shell: sshfolio
Terminal: Bubble Tea
Location: Bengaluru
Edition: The Jogo Gazette
`

	case "clear":
		return "__CLEAR__"

	case "exit":
		return "__EXIT__"

	case "":
		return ""

	default:
		return "command not found: " + cmd
	}
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "enter":

			cmd := m.input.Value()

			m.history = append(
				m.history,
				promptStyle.Render("jogo@portfolio:~$ ")+cmd,
			)

			result := runCommand(cmd)

			if result == "__CLEAR__" {
				m.history = []string{}
			} else if result == "__EXIT__" {
				return m, tea.Quit
			} else if result != "" {

				if strings.HasPrefix(result, "command not found") {
					m.history = append(
						m.history,
						errorStyle.Render(result),
					)
				} else {
					m.history = append(
						m.history,
						outputStyle.Render(result),
					)
				}
			}

			m.input.SetValue("")
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m model) View() string {

	content := strings.Join(m.history, "\n")

	return fmt.Sprintf(
		"%s\n\n%s %s",
		content,
		promptStyle.Render("jogo@portfolio:~$"),
		m.input.View(),
	)
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

		wish.WithPasswordAuth(func(ctx ssh.Context, password string) bool {
			return true
		}),

		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting JogoOS SSH server on port", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
