package commands

import (
	"fmt"
	"time"
	"io"
	"net/http"
)

func Neofetch() string {

	loc, err := time.LoadLocation("Asia/Kolkata")

	if err != nil {
		loc = time.UTC
	}

	now := time.Now().In(loc)

	currentTime := now.Format("03:04:05 PM IST")

	return fmt.Sprintf(`
 [38;5;205m
        ██╗ ██████╗  ██████╗  ██████╗
        ██║██╔═══██╗██╔════╝ ██╔═══██╗
        ██║██║   ██║██║  ███╗██║   ██║
   ██   ██║██║   ██║██║   ██║██║   ██║
   ╚█████╔╝╚██████╔╝╚██████╔╝╚██████╔╝
    ╚════╝  ╚═════╝  ╚═════╝  ╚═════╝
 [0m
 [38;5;205mjogo [0m@ [38;5;218mportfolio [0m
──────────────────────────────

 [38;5;205mOS [0m:         JogoOS v1.0
 [38;5;205mHost [0m:       Railway TCP Edge
 [38;5;205mKernel [0m:     sshfolio
 [38;5;205mUptime [0m:     forever
 [38;5;205mShell [0m:      jogo-shell
 [38;5;205mTerminal [0m:   SSH
 [38;5;205mTheme [0m:      Pink Noir
 [38;5;205mLocation [0m:   Bengaluru
 [38;5;205mDeveloper [0m:  Jagath Sajjan
 [38;5;205mTime [0m:       %s

 [38;5;205m███ [0m  [38;5;212m███ [0m  [38;5;218m███ [0m  [38;5;225m███ [0m
`, currentTime)
}

func Whoami() string {
	return "jogo"
}

func PWD() string {
	return "/home/jogo"
}

func Weather() string {

	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		"https://wttr.in/Bengaluru?format=v2",
		nil,
	)

	if err != nil {
		return "failed to create request"
	}

	req.Header.Set("User-Agent", "curl")

	resp, err := client.Do(req)

	if err != nil {
		return "failed to fetch weather"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "failed to read weather"
	}

	return string(body)
}
