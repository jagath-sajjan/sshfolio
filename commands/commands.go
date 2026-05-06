package commands

import (
	"os"
	"strings"
)

func Execute(cmd string) string {

	cmd = strings.TrimSpace(strings.ToLower(cmd))

	switch cmd {

	case "help":
		return `
help
ls
weather
cat 
pwd <file_name>
whoami
neofetch
sudo hire-me
matrix
hack
clear
exit
`

	case "ls":
		return `
about.txt
projects.txt
skills.txt
contact.txt
`

	case "cat about.txt":

		data, _ := os.ReadFile("data/about.txt")
		return string(data)

	case "cat projects.txt":

		data, _ := os.ReadFile("data/projects.txt")
		return string(data)

	case "cat skills.txt":

		data, _ := os.ReadFile("data/skills.txt")
		return string(data)

	case "cat contact.txt":

		data, _ := os.ReadFile("data/contact.txt")
		return string(data)

	case "pwd":
		return PWD()

	case "weather":
	  return Weather()

	case "whoami":
		return Whoami()

	case "neofetch":
		return Neofetch()

	case "sudo hire-me":
		return HireMe()

	case "matrix":
		return Matrix()

	case "hack":
		return Hack()

	case "clear":
		return "__CLEAR__"

	case "exit":
		return "__EXIT__"

	default:
		return "command not found: " + cmd
	}
}
