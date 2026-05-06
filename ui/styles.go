package ui

const (
	Reset = "\033[0m"

	Pink = "\033[38;5;205m"
	LightPink = "\033[38;5;218m"
	DarkPink = "\033[38;5;168m"

	White = "\033[97m"
	Red = "\033[31m"

	Bold = "\033[1m"
)

func Prompt() string {
	return Pink + Bold + "jogo@portfolio:~$ " + Reset
}

func ErrorText(text string) string {
	return Red + Bold + text + Reset
}

func Header(text string) string {
	return DarkPink + Bold + text + Reset
}

func Output(text string) string {
	return White + text + Reset
}
