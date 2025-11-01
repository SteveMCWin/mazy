package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

const wall = rune('█')
const empty = rune(' ')

type MazeCoords struct {
	X int
	Y int
}

type Cell struct {
	Sprite  rune
	Visited bool
}

type Maze struct {
	Cells [][]Cell
}

func main() {

	maze_x, maze_y := 35, 25

	var maze Maze
	maze.InitMazeBase(maze_x, maze_y)

	// fmt.Print(MazeStr(maze))

	maze.MakeMazeStartEnd()
	steps := maze.MakeMazeRDFS()

	fmt.Println()
	for i := range steps {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Print(steps[i])
		time.Sleep(30*time.Millisecond)
	}
	// fmt.Print(maze.String())
	fmt.Println()

}

func (m *Maze) InitMazeBase(x, y int) {
	m.Cells = make([][]Cell, y)

	for i := range y {
		m.Cells[i] = make([]Cell, x)

		for j := range x {
			if i%2 == 0 || j%2 == 0 {
				m.Cells[i][j].Sprite = wall
			} else {
				m.Cells[i][j].Sprite = empty
			}
		}
	}
}

func (m *Maze) MakeMazeStartEnd(coords ...MazeCoords) {
	var begin, finish MazeCoords
	if len(coords) == 2 {
		begin = coords[0]
		finish = coords[1]
	} else {
		begin = MazeCoords{X: 1, Y: 0}
		finish = MazeCoords{X: len(m.Cells[0]) - 2, Y: len(m.Cells) - 1}
	}

	m.Cells[begin.Y][begin.X].Sprite = empty
	m.Cells[finish.Y][finish.X].Sprite = empty
}

func (m *Maze) MakeMazeRDFS() []string {

	if len(m.Cells) == 0 {
		fmt.Println("Maze has no rows ig")
		m.InitMazeBase(25, 15)
	}

	steps := []string{ m.String() }

	stack := []MazeCoords{}

	m.Cells[1][1].Visited = true
	stack = append(stack, MazeCoords{1, 1})

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		unvisitedNeighbours := getUnvisitedNeighbours(curr, m)
		if len(unvisitedNeighbours) > 0 {
			stack = append(stack, curr)
			next := unvisitedNeighbours[rand.Int()%len(unvisitedNeighbours)]
			wallCoords := MazeCoords{X: (curr.X + next.X) / 2, Y: (curr.Y + next.Y) / 2}
			m.Cells[wallCoords.Y][wallCoords.X].Sprite = empty
			m.Cells[next.Y][next.X].Visited = true
			stack = append(stack, next)

			steps = append(steps, m.String())
		}
	}

	fmt.Println("len(steps): ", len(steps))
	return steps
}

func getUnvisitedNeighbours(curr MazeCoords, maze *Maze) []MazeCoords {
	res := make([]MazeCoords, 0)
	if curr.X > 1 && !maze.Cells[curr.Y][curr.X-2].Visited {
		res = append(res, MazeCoords{X: curr.X - 2, Y: curr.Y})
	}
	if curr.X < len(maze.Cells[0])-2 && !maze.Cells[curr.Y][curr.X+2].Visited {
		res = append(res, MazeCoords{X: curr.X + 2, Y: curr.Y})
	}
	if curr.Y > 1 && !maze.Cells[curr.Y-2][curr.X].Visited {
		res = append(res, MazeCoords{X: curr.X, Y: curr.Y - 2})
	}
	if curr.Y < len(maze.Cells)-2 && !maze.Cells[curr.Y+2][curr.X].Visited {
		res = append(res, MazeCoords{X: curr.X, Y: curr.Y + 2})
	}

	return res
}

func (m *Maze) String() string {
	var b strings.Builder
	for i := range len(m.Cells) {
		for j := range len(m.Cells[0]) {
			b.WriteRune(m.Cells[i][j].Sprite)
			b.WriteRune(m.Cells[i][j].Sprite) // write the same sprite again for double width
		}
		b.WriteRune('\n')
	}

	return b.String()
}
