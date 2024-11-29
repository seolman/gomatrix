package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	width  = 80
	height = 20
	chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
)

type model struct {
	screen [][]rune
}

func initialModel() model {
	screen := make([][]rune, height)
	for i := range screen {
		screen[i] = make([]rune, width)
	}
	return model{screen: screen}
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
		}
	case tea.WindowSizeMsg:
		// Handle window resizing if needed
	default:
		for i := 0; i < width; i++ {
			if rand.Intn(10) == 0 {
				m.screen[0][i] = rune(chars[rand.Intn(len(chars))])
			} else {
				m.screen[0][i] = ' '
			}
		}

		for i := height - 1; i > 0; i-- {
			for j := 0; j < width; j++ {
				m.screen[i][j] = m.screen[i-1][j]
			}
		}
	}
	return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) View() string {
	var s string
	for _, row := range m.screen {
		for _, char := range row {
			if char == 0 {
				s += " "
			} else {
				s += string(char)
			}
		}
		s += "\n"
	}
	return s
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
