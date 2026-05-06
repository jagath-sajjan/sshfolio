package ui

func (m *Model) AddHistory(text string) {

	m.History = append(m.History, text)

	if len(m.History) > 20 {
		m.History  = m.History[len(m.History)-20:]
	}
}

func (m *Model) ClearHistory() {
	m.History = []string{}
}
