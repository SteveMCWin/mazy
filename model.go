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

	showPlayer bool

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
		showPlayer: true,
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

func (m *Model) genNewMaze() {
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
		case "ctrl+r":
			m.maze.CurrFrame = 0
			m.maze.AnimationPaused = false
			m.genNewMaze()
			m.maze.MovePlayerTo(m.maze.StartPos)
			cmds = append(cmds, AnimateSignal(m))
		case " ":
			m.maze.AnimationPaused = !m.maze.AnimationPaused
			cmds = append(cmds, AnimateSignal(m))
		case "h", "left":
			log.Printf("Pressed %s\n", msg.String())
			log.Printf("Adding {%d, %d} and {%d, %d}\n", m.maze.PlayerPos.X, m.maze.PlayerPos.Y, -1, 0)
			newPos := maze.AddCoords(m.maze.PlayerPos, maze.MazeCoords{X:-1, Y: 0})
			log.Printf("New pos: {%d, %d}", newPos.X, newPos.Y)

			m.maze.MovePlayerTo(newPos)
			// log.Printf("New player pos: {%d, %d}", m.maze.PlayerPos.X, m.maze.PlayerPos.Y)
			m.maze.UpdateLastFrame()
		case "j", "down":
			log.Printf("Pressed %s\n", msg.String())
			log.Printf("Adding {%d, %d} and {%d, %d}\n", m.maze.PlayerPos.X, m.maze.PlayerPos.Y, 0, 1)
			newPos := maze.AddCoords(m.maze.PlayerPos, maze.MazeCoords{X: 0, Y: 1})
			log.Printf("New pos: {%d, %d}", newPos.X, newPos.Y)
			m.maze.MovePlayerTo(newPos)
			// log.Printf("New player pos: {%d, %d}", m.maze.PlayerPos.X, m.maze.PlayerPos.Y)
			m.maze.UpdateLastFrame()
		case "k", "up":
			log.Printf("Pressed %s\n", msg.String())
			log.Printf("Adding {%d, %d} and {%d, %d}\n", m.maze.PlayerPos.X, m.maze.PlayerPos.Y, 0, -1)
			newPos := maze.AddCoords(m.maze.PlayerPos, maze.MazeCoords{X: 0, Y:-1})
			log.Printf("New pos: {%d, %d}", newPos.X, newPos.Y)
			m.maze.MovePlayerTo(newPos)
			// log.Printf("New player pos: {%d, %d}", m.maze.PlayerPos.X, m.maze.PlayerPos.Y)
			m.maze.UpdateLastFrame()
		case "l", "right":
			log.Printf("Pressed %s\n", msg.String())
			log.Printf("Adding {%d, %d} and {%d, %d}\n", m.maze.PlayerPos.X, m.maze.PlayerPos.Y, 1, 0)
			newPos := maze.AddCoords(m.maze.PlayerPos, maze.MazeCoords{X: 1, Y: 0})
			log.Printf("New pos: {%d, %d}", newPos.X, newPos.Y)
			m.maze.MovePlayerTo(newPos)
			// log.Printf("New player pos: {%d, %d}", m.maze.PlayerPos.X, m.maze.PlayerPos.Y)
			m.maze.UpdateLastFrame()
		case "ctrl+up":
			// cmds = append(cmds, ChangeFontSize(&m.terminal, 1, true))
		case "ctrl+down":
			// cmds = append(cmds, ChangeFontSize(&m.terminal, 1, false))
		case "enter":
			m.maze.CurrFrame = len(m.maze.Steps)-1
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
		m.genNewMaze()
		m.maze.AnimationPaused = true
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
