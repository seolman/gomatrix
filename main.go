package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:',.<>?/`~"
)

type model struct {
	screen [][]rune
	width  int
	height int
	speed  int
	colors []lipgloss.Color
}

func initialModel() model {
	return model{
		screen: make([][]rune, 0),
		width:  80,
		height: 20,
		speed:  100,
		colors: []lipgloss.Color{
			lipgloss.Color("1"),
			lipgloss.Color("2"),
			lipgloss.Color("3"),
			lipgloss.Color("4"),
			lipgloss.Color("5"),
			lipgloss.Color("6"),
			lipgloss.Color("7"),
			lipgloss.Color("8"),
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
		case "up":
			if m.speed > 50 {
				m.speed -= 10
			}
		case "down":
			if m.speed < 500 {
				m.speed += 10
			}
		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.screen = make([][]rune, m.height)
		for i := range m.screen {
			m.screen[i] = make([]rune, m.width)
		}
	default:
		for i := 0; i < m.width; i++ {
			if rand.Intn(10) == 0 {
				m.screen[0][i] = rune(chars[rand.Intn(len(chars))])
			} else {
				m.screen[0][i] = ' '
			}
		}

		for i := m.height - 1; i > 0; i-- {
			for j := 0; j < m.width; j++ {
				m.screen[i][j] = m.screen[i-1][j]
			}
		}
	}
	return m, tea.Tick(time.Duration(m.speed)*time.Millisecond, func(t time.Time) tea.Msg {
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
				color := m.colors[rand.Intn(len(m.colors))]
				s += lipgloss.NewStyle().Foreground(color).Render(string(char))
			}
		}
		s += "\n"
	}
	return s
}

func main() {
  rand.New(rand.NewSource(time.Now().UnixNano()))
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
