package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
	lang "github.com/tecnologer/dicegame/language"
)

var (
	term    = termenv.EnvColorProfile()
	keyword = makeFgStyle("211")
	subtle  = makeFgStyle("241")
	dot     = colorFg(" • ", "236")
	lFmt    = lang.GetCurrent()
)

type (
	tui struct {
		optionSelected int
		isSelected     bool
		closing        bool
		playerName     textinput.Model
	}
	frameMsg struct{}
)

func (tui) Init() tea.Cmd {
	return nil
}

func (t tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			t.closing = true
			return t, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !t.isSelected {
		return updateChoices(msg, t)
	}
	return updateChosen(msg, t)
}

func (t tui) View() string {
	var s string
	if t.closing {
		return "\n  See you later!\n\n"
	}
	if !t.isSelected {
		s = choicesView(t)
	} else {
		s = chosenView(t)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

func main() {
	ti := textinput.New()
	ti.Placeholder = lFmt.Sprintf("Playername")
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	initialModel := tui{
		playerName: ti,
	}
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

// Update loop for the first view where you're choosing a task.
func updateChoices(msg tea.Msg, t tui) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			t.optionSelected += 1
			if t.optionSelected > 3 {
				t.optionSelected = 3
			}
		case "k", "up":
			t.optionSelected -= 1
			if t.optionSelected < 0 {
				t.optionSelected = 0
			}
		case "enter":
			t.isSelected = true
			return t, frame()
		}
	}

	return t, nil
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

// Update loop for the second view after a choice has been made
func updateChosen(msg tea.Msg, t tui) (tea.Model, tea.Cmd) {
	switch msg.(type) {

	case frameMsg:

	}

	return t, nil
}

// The first view, where you're choosing a task
func choicesView(t tui) string {
	c := t.optionSelected

	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("SinglePlayer", c == 0),
		checkbox("Multiplayer", c == 1),
	)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	label = lFmt.Sprintf(label)
	if checked {
		return colorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

// Color a string's foreground with the given value.
func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// The second view, after a task has been chosen
func chosenView(t tui) string {
	var msg string

	switch t.optionSelected {
	case 0:
		msg = requestName(t)
	case 1:
		msg = multiplayerOptions(t)
	}

	return msg
}

func requestName(t tui) string {
	t.playerName.Focus()
	return fmt.Sprintf(
		"Enter your name?\n\n%s\n\n%s",
		t.playerName.View(),
		"(esc to quit)",
	) + "\n"
}

func multiplayerOptions(t tui) string {
	c := t.optionSelected + 1
	t.isSelected = false
	tpl := "Multiplayer\n\n"
	tpl += "%s\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("Host game", c == 2),
		checkbox("Join game", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}
