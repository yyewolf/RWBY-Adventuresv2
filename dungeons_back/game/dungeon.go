package game

import "math/rand"

const (
	tileFloor = iota
	tileWall
	tileFow
	tilePlayer
)

var (
	dx = [4]int{1, 0, -1, 0}
	dy = [4]int{0, 1, 0, -1}
)

type PlayerPosition struct {
	X int
	Y int
}

type DungeonCell struct {
	Type int `json:"type"`
}

type Dungeon struct {
	Grid     [][]DungeonCell
	Height   int
	Width    int
	Position PlayerPosition
}

func NewDungeon(height, width int) *Dungeon {
	d := &Dungeon{
		Height: height,
		Width:  width,
		Position: PlayerPosition{
			X: 1,
			Y: 1,
		},
	}
	d.GenerateMaze()
	d.Grid[d.Position.Y][d.Position.X].Type = tilePlayer
	return d
}

func (d *Dungeon) MovePlayer(direction int) {
	if d.Grid[d.Position.Y+dy[direction]][d.Position.X+dx[direction]].Type == tileFloor {
		d.Grid[d.Position.Y][d.Position.X].Type = tileFloor
		d.Position.Y += dy[direction]
		d.Position.X += dx[direction]
		d.Grid[d.Position.Y][d.Position.X].Type = tilePlayer
	}
}

func (d *Dungeon) Init() {
	d.Grid = make([][]DungeonCell, d.Height)
	for i := 0; i < d.Height; i++ {
		d.Grid[i] = make([]DungeonCell, d.Width)
		if i == 0 || i == d.Height-1 {
			for j := 0; j < d.Width; j++ {
				d.Grid[i][j].Type = tileWall
			}
		} else {
			for j := 0; j < d.Width; j++ {
				if j == 0 || j == d.Width-1 {
					d.Grid[i][j].Type = tileWall
				}
			}
		}
	}
}

func (d *Dungeon) GenerateMaze() {
	d.Init()

	var stack [][][]int
	stack = append(stack, [][]int{{1, 1}, {d.Height - 2, d.Width - 2}})

	for len(stack) > 0 {
		currentRegion := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		miny := currentRegion[0][0]
		maxy := currentRegion[1][0]
		minx := currentRegion[0][1]
		maxx := currentRegion[1][1]
		height := maxy - miny + 1
		width := maxx - minx + 1

		if height <= 1 || width <= 1 {
			continue
		}

		var cutdirection int

		if width < height {
			cutdirection = 1
		} else if width > height {
			cutdirection = 0
		} else {
			if width == 2 {
				continue
			}
			cutdirection = rand.Intn(2)
		}

		directions := []int{height, width}

		// make cut
		//   select cut position (can't be completely on the edge of the region)
		cutlength := directions[(cutdirection+1)%2]
		if cutlength < 3 {
			continue
		}
		var r []int
		for i := 1; i < cutlength; i += 2 {
			r = append(r, i)
		}
		cutposi := 0
		if len(r) == 1 {
			cutposi = r[0]
		} else {
			cutposi = r[rand.Intn(len(r)-1)]
		}
		r = []int{}
		for i := 0; i < directions[cutdirection]; i += 2 {
			r = append(r, i)
		}
		doorposi := 0
		if len(r) == 1 {
			doorposi = r[0]
		} else {
			doorposi = r[rand.Intn(len(r)-1)]
		}
		if cutdirection == 0 {
			for row := miny; row < maxy+1; row++ {
				d.Grid[row][minx+cutposi].Type = tileWall
			}
			d.Grid[miny+doorposi][minx+cutposi].Type = tileFloor
		} else {
			for col := minx; col < maxx+1; col++ {
				d.Grid[miny+cutposi][col].Type = tileWall
			}
			d.Grid[miny+cutposi][minx+doorposi].Type = tileFloor
		}
		if cutdirection == 0 {
			var firstArea = [][]int{{miny, minx}, {maxy, minx + cutposi - 1}}
			var secondArea = [][]int{{miny, minx + cutposi + 1}, {maxy, maxx}}
			stack = append(stack, firstArea)
			stack = append(stack, secondArea)
		} else {
			var firstArea = [][]int{{miny, minx}, {miny + cutposi - 1, maxx}}
			var secondArea = [][]int{{miny + cutposi + 1, minx}, {maxy, maxx}}
			stack = append(stack, firstArea)
			stack = append(stack, secondArea)
		}
	}
}

func (d *Dungeon) GetSmallGrid(width, height int) [][]DungeonCell {
	var smallGrid [][]DungeonCell
	x := d.Position.X - width/2
	y := d.Position.Y - height/2
	for i := y; i < y+height; i++ {
		smallGrid = append(smallGrid, d.Grid[i][x:x+width])
	}
	return smallGrid
}
