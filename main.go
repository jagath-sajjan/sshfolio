package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	terminal "golang.org/x/term"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"

	"github.com/jagath-sajjan/sshfolio/boot"
	"github.com/jagath-sajjan/sshfolio/commands"
	"github.com/jagath-sajjan/sshfolio/ui"
)

func sessionHandler(s ssh.Session) {

	model := ui.NewModel()

	boot.Boot(s)

	term := terminal.NewTerminal(s, "")

	for {

		ui.RenderPrompt(s)

		input, err := term.ReadLine()

		if err != nil {
			return
		}

		input = strings.TrimSpace(input)

		model.AddHistory(input)

		response := commands.Execute(input)

		switch response {

		case "__CLEAR__":

			model.ClearHistory()

			fmt.Fprint(s, "\033[2J\033[H")

			continue

		case "__EXIT__":

			fmt.Fprintln(s, "\nbye.\n")

			return
		}

		if strings.HasPrefix(response, "command not found") {

			ui.RenderError(s, response)

		} else {

			ui.RenderOutput(s, response)
		}
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

		wish.WithPasswordAuth(

			func(ctx ssh.Context, password string) bool {

				return true
			},
		),

		wish.WithMiddleware(

			func(next ssh.Handler) ssh.Handler {

				return func(s ssh.Session) {

					sessionHandler(s)
				}
			},
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("JogoOS SSH running on port", port)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
