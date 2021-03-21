package main

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	te "github.com/muesli/termenv"
)

const focusedTextColor = "205"

var (
	color               = te.ColorProfile().Color
	focusedPrompt       = te.String("> ").Foreground(color("205")).String()
	blurredPrompt       = "> "
	focusedSubmitButton = "[ " + te.String("Submit").Foreground(color("205")).String() + " ]"
	blurredSubmitButton = "[ " + te.String("Submit").Foreground(color("240")).String() + " ]"

	keyword       textinput.Model
	location      textinput.Model
	packageInputs = [][]textinput.Model{}

	baseurl    = "https://www.indeed.com/jobs?"
	baselimit  = "&limit=50"
	maxresults = 100
)

func UserInput() {
	if err := tea.NewProgram(initialModel()).Start(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}

	for _, val := range packageInputs {
		url := ""
		for i, v := range val {
			// fmt.Printf("Index: %d, Value: %+v\n", i, v.Value())
			if i == 0 {
				url += "q=" + strings.ReplaceAll(v.Value(), " ", "+") + "&"
			} else {
				// Keyword can have multiple locations, comma is the delimiter
				loc := strings.Split(v.Value(), ",")
				for _, l := range loc {
					// trim excess space before replacing with urlencoded spaces
					l = strings.TrimSpace(l)
					//append to main slice
					urlSlice = append(urlSlice, baseurl+url+"l="+strings.ReplaceAll(l, " ", "%2C+")+baselimit)
				}
			}
		}
		url = baseurl + url + baselimit
	}
}

func initialModel() model {
	keyword = textinput.NewModel()
	keyword.Placeholder = "Keyword eg: Cart Pusher"
	keyword.Focus()
	keyword.Prompt = focusedPrompt
	keyword.TextColor = focusedTextColor
	keyword.CharLimit = 32

	location = textinput.NewModel()
	location.Placeholder = "Location eg: Huntsville AL, Boulder CO"
	location.Prompt = blurredPrompt
	location.CharLimit = 64

	return model{0, keyword, location, blurredSubmitButton}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Cycle between inputs
		case "tab", "shift+tab", "enter", "up", "down", "ctrl+j", "ctrl+k":
			input := []textinput.Model{
				m.keywordInput,
				m.locationInput,
			}

			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, start over.
			if s == "enter" && m.index == len(input) {
				packageInputs = append(packageInputs, input)
				// keyword.Reset()
				// location.Reset()
				m = initialModel()
				return m, m.Init()
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" || s == "ctrl+k" {
				m.index--
			} else {
				m.index++
			}

			if m.index > len(input) {
				m.index = 0
			} else if m.index < 0 {
				m.index = len(input)
			}

			for i := 0; i <= len(input)-1; i++ {
				if i == m.index {
					// Set focused state
					input[i].Focus()
					input[i].Prompt = focusedPrompt
					input[i].TextColor = focusedTextColor
					continue
				}
				// Remove focused state
				input[i].Blur()
				input[i].Prompt = blurredPrompt
				input[i].TextColor = ""
			}

			m.keywordInput = input[0]
			m.locationInput = input[1]

			if m.index == len(input) {
				m.submitButton = focusedSubmitButton
			} else {
				m.submitButton = blurredSubmitButton
			}

			return m, nil
		}
	}

	// Handle character input and blinks
	m, cmd = updateInputs(msg, m)

	return m, cmd
}

// Pass messages and models through to text input components. Only text inputs
// with Focus() set will respond, so it's safe to simply update all of them
// here without any further logic.
func updateInputs(msg tea.Msg, m model) (model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.keywordInput, cmd = m.keywordInput.Update(msg)
	cmds = append(cmds, cmd)

	m.locationInput, cmd = m.locationInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	s := "\nSubmit to submit query.\n\nPress Ctrl+c or ESC to search on Indeed!\n\n"

	inputs := []string{
		m.keywordInput.View(),
		m.locationInput.View(),
	}

	for i := 0; i < len(inputs); i++ {
		s += inputs[i]
		if i < len(inputs)-1 {
			s += "\n"
		}
	}

	s += "\n\n" + m.submitButton + "\n"

	return s
}
