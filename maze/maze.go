package maze

import (
	"fmt"
	"strconv"
	// "log"
	"math/rand"
	"strings"
)

// const wall = rune('▓')
// const empty = rune(' ')
const double_width_maze = true

type TileType rune

const (
	Empty TileType = ' '
	Wall  TileType = '▓'
	Player TileType = '░'
)

func (tt TileType) String() string {
	return strconv.QuoteRune(rune(tt))
}

type MazeCoords struct {
	X int
	Y int
}

type Cell struct {
	Tile    TileType
	Visited bool
}

type Maze struct {
	Cells           [][]Cell
	Steps           []string
	CurrFrame       int
	AnimationPaused bool
	StartPos        MazeCoords
	EndPos          MazeCoords
	PlayerPos       MazeCoords
}

func (m *Maze) InitMazeBase(x, y int) {

	// make sure the dimensions are odd and don't overflow
	if x > 1 && x%2 == 0 {
		x -= 1
	}

	if y > 1 && y%2 == 0 {
		y -= 1
	}

	m.Cells = make([][]Cell, y)

	for i := range y {
		m.Cells[i] = make([]Cell, x)

		for j := range x {
			if i%2 == 0 || j%2 == 0 {
				m.Cells[i][j].Tile = Wall
			} else {
				m.Cells[i][j].Tile = Empty
			}
		}
	}

	m.Steps = []string{m.String()}
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

	m.Cells[begin.Y][begin.X].Tile = Player
	m.Cells[finish.Y][finish.X].Tile = Empty

	m.StartPos = begin
	m.EndPos = finish

	m.MovePlayerTo(m.StartPos)

	m.Steps = append(m.Steps, m.String())
}

func (m *Maze) MakeMazeRDFS() {

	if len(m.Cells) == 0 {
		fmt.Println("Maze has no rows ig")
		m.InitMazeBase(25, 15)
	}

	stack := []MazeCoords{}

	// getting a random cell with odd x and y coords since only those are empty cells
	start_pos := MazeCoords{
		X: (rand.Int()%(len(m.Cells[0])/2))*2 + 1,
		Y: (rand.Int()%(len(m.Cells)/2))*2 + 1,
	}

	m.Cells[start_pos.Y][start_pos.X].Visited = true
	stack = append(stack, start_pos)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		unvisitedNeighbours := getUnvisitedNeighbours(curr, m)
		if len(unvisitedNeighbours) > 0 {
			stack = append(stack, curr)
			next := unvisitedNeighbours[rand.Int()%len(unvisitedNeighbours)]
			wallCoords := MazeCoords{X: (curr.X + next.X) / 2, Y: (curr.Y + next.Y) / 2}
			m.Cells[wallCoords.Y][wallCoords.X].Tile = Empty
			m.Cells[next.Y][next.X].Visited = true
			stack = append(stack, next)

			m.Steps = append(m.Steps, m.String())
		}
	}

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

	var write func(i, j int)

	if double_width_maze {
		write = func(i, j int) {
			b.WriteRune(rune(m.Cells[i][j].Tile))
			b.WriteRune(rune(m.Cells[i][j].Tile)) // write the same sprite again for double width
		}
	} else {
		write = func(i, j int) {
			b.WriteRune(rune(m.Cells[i][j].Tile))
		}
	}

	for i := range len(m.Cells) {
		for j := range len(m.Cells[0]) {
			write(i, j)
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (m *Maze) GetTileType(coords MazeCoords) TileType {
	return m.Cells[coords.Y][coords.X].Tile
}

func (m *Maze) CanPlayerMoveTo(coords MazeCoords) bool {
	isInBounds := (coords.X >= 0 && coords.X < len(m.Cells[0])) && (coords.Y >= 0 && coords.Y < len(m.Cells))
	if !isInBounds {
		panic("AAAAA")
	}
	isEmptyTile := m.GetTileType(coords) == Empty
	mazeFinishedAnimating := m.CurrFrame == len(m.Steps)-1
	return isEmptyTile && isInBounds && mazeFinishedAnimating
}

func AddCoords(c1, c2 MazeCoords) MazeCoords {
	return MazeCoords{X: c1.X + c2.X, Y: c1.Y + c2.Y}
}

func (m *Maze) MovePlayerTo(coords MazeCoords) {
	if !m.CanPlayerMoveTo(coords) {
		return
	}

	m.Cells[m.PlayerPos.Y][m.PlayerPos.X].Tile = Empty

	m.PlayerPos = coords

	m.Cells[coords.Y][coords.X].Tile = Player
}

func (m *Maze) MovePlayerBy(coords MazeCoords) {
	if !m.CanPlayerMoveTo(coords) {
		return
	}

	m.Cells[m.PlayerPos.Y][m.PlayerPos.X].Tile = Empty

	m.PlayerPos = AddCoords(m.PlayerPos, coords)

	m.Cells[coords.Y][coords.X].Tile = Player
}

func (m *Maze) UpdateLastFrame() {
	m.Steps[len(m.Steps)-1] = m.String()
}
