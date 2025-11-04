package main

import (
	"log"
	// "os"
	"strings"
	"maze_gen/maze"

	// "github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	doc := strings.Builder{}

	var contents string

	if len(m.maze.Steps) <= m.maze.CurrFrame {
		contents = ""
	} else {
		mazeStr := m.maze.Steps[m.maze.CurrFrame]

		if !m.showPlayer {
			mazeStr = strings.ReplaceAll(mazeStr, maze.Player.String(), maze.Empty.String())
		}

		contents = wallStyle.Render(mazeStr)
	}

	_, err := doc.WriteString(windowStyle.Width(m.windowWidth-windowStyle.GetHorizontalFrameSize()).Height(m.windowHeight-windowStyle.GetVerticalFrameSize()).Render(contents))
	if err != nil {
		log.Println("Error displaying window and contents:", err)
	}

	return docStyle.Render(doc.String())
}
