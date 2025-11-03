package main

import (
	"log"
	// "os"
	"strings"

	// "github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	doc := strings.Builder{}

	var contents string

	contents = m.maze.String()

	_, err := doc.WriteString(windowStyle.Width(m.windowWidth-windowStyle.GetHorizontalFrameSize()).Height(m.windowHeight-windowStyle.GetVerticalFrameSize()).Render(contents))
	log.Println("m.windowWidth-windowStyle.GetHorizontalFrameSize(): ", m.windowWidth-windowStyle.GetHorizontalFrameSize())
	log.Println("m.windowHeight-windowStyle.GetVerticalFrameSize(): ", m.windowHeight-windowStyle.GetVerticalFrameSize())
	if err != nil {
		log.Println("Error displaying window and contents:", err)
	}

	return docStyle.Render(doc.String())
}
