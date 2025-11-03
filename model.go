package main

import (
	"log"
	"maze_gen/maze"
	"os/exec"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type SupportedTerminals int

const (
	Kitty SupportedTerminals = iota
	Gnome
	Konsole
	Other
)

type Animate bool

type Model struct {
	windowWidth  int
	windowHeight int

	theme ColorTheme

	maze         maze.Maze
	gen_steps    []string
	maze_dims    maze.MazeCoords
	end_coords   maze.MazeCoords
	start_coords maze.MazeCoords
}

func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{
		SetCurrentTheme(m.theme),
		tea.SetWindowTitle("Mazy"),
	}

	return tea.Batch(cmds...) // NOTE: set curr theme should be replaced with a function that loads save data and that handles the theme
}

func NewModel() Model {
	return Model{
		theme: DefaultTheme,
	}
}

func ChangeFontSize(term SupportedTerminals, amount int, pos bool) tea.Cmd {
	var term_name string
	var term_cmd []string
	amount_str := strconv.Itoa(amount)

	switch term {
	case Kitty:
		term_name = "kitty"
		sign := "+"
		if !pos {
			sign = "-"
		}
		term_cmd = []string{"@", "set-font-size", "--", sign + amount_str}
	}

	c := exec.Command(term_name, term_cmd...) //nolint:gosec
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return nil
	})
}

func genNewMaze(m *Model) {
	m.maze_dims.X = (m.windowWidth-windowStyle.GetHorizontalFrameSize())/2 - 2
	m.maze_dims.Y = m.windowHeight - windowStyle.GetVerticalFrameSize()
	m.maze.InitMazeBase(m.maze_dims.X, m.maze_dims.Y)
	log.Println("width, height:", m.maze_dims.X, m.maze_dims.Y)
	m.maze.MakeMazeRDFS()
	m.maze.MakeMazeStartEnd()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// log.Println("tea.KeyMsg:", msg.String())
		switch msg.String() {
		case "ctrl+c", "q":
			seq := tea.Sequence(ChangeFontSize(Kitty, 0, true), tea.Quit)
			cmds = append(cmds, seq)
		case "right", "tab":
			// if !m.isTyping {
			// 	m.currTab = TabIndex((int(m.currTab) + 1) % len(m.tabs))
			// }
		case "left", "shift+tab":
			// if !m.isTyping {
			// 	m.currTab = TabIndex((len(m.tabs) + int(m.currTab) - 1) % len(m.tabs))
			// }
		case "ctrl+r": // THERE IS A BUG when the program starts a new maze is made, but then another is made when pressing ctrl+r and then they both get displayed, but after that no new ones are made
			genNewMaze(&m)
			cmds = append(cmds, AnimateSignal(m))
		case "ctrl+up":
			// cmds = append(cmds, ChangeFontSize(&m.terminal, 1, true))
		case "ctrl+down":
			// cmds = append(cmds, ChangeFontSize(&m.terminal, 1, false))
		case "enter":
			// if m.currTab == Home {
			// 	m.isTyping = true
			// 	cmds = append(cmds, tea.ShowCursor)
			// }
		case "esc":
			// if m.isTyping {
			// 	cmds = append(cmds, tea.HideCursor)
			// 	m.isTyping = false
			// 	// stop the test or something
			// }
		default:
			// if m.isTyping {
			// 	HandleTyping(&m, msg.String())
			// }
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		genNewMaze(&m)
	case Animate:
		if bool(msg) == true {
			if m.maze.CurrFrame < len(m.maze.Steps)-1 {
				m.maze.CurrFrame += 1
				cmds = append(cmds, AnimateSignal(m))
			} else {
				m.maze.AnimationPaused = true
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func AnimateSignal(m Model) func()tea.Msg {
	return func() tea.Msg {
		time.Sleep(10 * time.Millisecond)
		return Animate(!m.maze.AnimationPaused)
	}
}
