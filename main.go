package main

import (
	// "fmt"
	// "os"
	// "os/exec"
	// "time"
	//
	// "maze_gen/maze"
	"log"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("failed setting the debug log file: %v", err)
	}
	defer f.Close()
	// should check for commandline args
	m := NewModel()

	log.Println()
	log.Println("~~~~~~~~~PROGRAM START~~~~~~~~~")
	log.Println()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal("Unable to run tui:", err)
	}

	// maze_x, maze_y := 35, 25
	//
	// var maze maze.Maze
	// maze.InitMazeBase(maze_x, maze_y)
	//
	// maze.MakeMazeStartEnd()
	// steps := maze.MakeMazeRDFS()
	//
	// fmt.Println()
	// for i := range steps {
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// 	fmt.Print(steps[i])
	// 	time.Sleep(10*time.Millisecond)
	// }
	// fmt.Println()

}

