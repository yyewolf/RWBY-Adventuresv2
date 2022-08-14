package game

import (
	"fmt"
	"math/rand"
	"rwby-adventures/microservices"

	uuid "github.com/satori/go.uuid"
)

func NewDungeon(height, width int) *Dungeon {
	d := &Dungeon{
		Height: height,
		Width:  width,
		Position: &PlayerPosition{
			X: 1,
			Y: 1,
		},
		Rewards: &microservices.DungeonReward{},
		Health:  150,
	}
	d.GenerateMaze()
	return d
}

func (d *Dungeon) MovePlayer(direction int) (end bool) {
	if d.Position.X+dx[direction] < 0 || d.Position.X+dx[direction] >= d.Width {
		return
	}
	if d.Position.Y+dy[direction] < 0 || d.Position.Y+dy[direction] >= d.Height {
		return
	}
	newCell := d.Grid[d.Position.Y+dy[direction]][d.Position.X+dx[direction]]
	oldCell := d.Grid[d.Position.Y][d.Position.X]

	if oldCell.Type == tileMoney {
		d.Rewards.Lien += oldCell.Amount
		oldCell.Type = tileFloor
		oldCell.Amount = 0
		oldCell.Message = ""
	}

	if oldCell.Type == tileEnnemy {
		if oldCell.Amount <= 0 {
			oldCell.Type = tileFloor
			oldCell.Amount = 0
			oldCell.Message = ""
		}
	}

	if oldCell.Type == tileArm {
		d.Rewards.Arms += 1
		oldCell.Type = tileFloor
		oldCell.Amount = 0
		oldCell.Message = ""
	}

	if oldCell.Type == tileMinion {
		d.Rewards.Minions += 1
		oldCell.Type = tileFloor
		oldCell.Amount = 0
		oldCell.Message = ""
	}

	if newCell.Type == tileWall {
		return
	}

	if newCell.Type == tileTPWall {
		if d.InSecretRoom {
			d.SecretRoom[1][1].Message = findWall
			d.InSecretRoom = false
			d.Position = d.PreviousPos
			d.Grid = d.Temp
			d.Temp = nil
		} else {
			d.InSecretRoom = true
			d.PreviousPos = &PlayerPosition{
				X: d.Position.X,
				Y: d.Position.Y,
			}
			d.Position = &PlayerPosition{
				X: 1,
				Y: 1,
			}
			d.Temp = d.Grid
			d.Grid = d.SecretRoom
		}
		return
	}

	d.Position.Y += dy[direction]
	d.Position.X += dx[direction]

	if newCell.Type == tileEnnemy {
		newCell.Amount -= 1
		d.Health -= newCell.Damages
		if d.Health <= 0 {
			d.Health = 0
			d.Win = false
			return true
		}
	}

	if newCell.Type == tileEscape {
		d.Win = true
		return true
	}

	return
}

func (d *Dungeon) Init() {
	d.Grid = make([][]*DungeonCell, d.Height)
	for i := 0; i < d.Height; i++ {
		d.Grid[i] = make([]*DungeonCell, d.Width)
		for j := 0; j < d.Width; j++ {
			d.Grid[i][j] = &DungeonCell{}
			d.Grid[i][j].Generate()
		}
		if i == 0 || i == d.Height-1 {
			for j := 0; j < d.Width; j++ {
				d.Grid[i][j].Type = tileWall
				if j%2 == 1 {
					d.Grid[i][j].GenerateWall()
				}
			}
		} else {
			for j := 0; j < d.Width; j++ {
				if j == 0 || j == d.Width-1 {
					d.Grid[i][j].Type = tileWall
					if i%2 == 1 {
						d.Grid[i][j].GenerateWall()
					}
				}
			}
		}
	}
}

func (d *Dungeon) GenerateSecretRoom() {
	var height = 6
	var width = 6
	d.SecretRoom = make([][]*DungeonCell, height)
	for i := 0; i < height; i++ {
		d.SecretRoom[i] = make([]*DungeonCell, width)
		for j := 0; j < width; j++ {
			d.SecretRoom[i][j] = &DungeonCell{}
			d.SecretRoom[i][j].Generate()
			// Regenerates for double luck
			if d.SecretRoom[i][j].Type == tileFloor {
				d.SecretRoom[i][j].Generate()
			}
		}
		if i == 0 || i == height-1 {
			for j := 0; j < width; j++ {
				d.SecretRoom[i][j].Type = tileWall
			}
		} else {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					d.SecretRoom[i][j].Type = tileWall
				}
			}
		}
	}
	d.SecretRoom[1][1].Message = findWall
	d.SecretRoom[height-1][width/2].Type = tileTPWall
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
				if row != miny+doorposi {
					d.Grid[row][minx+cutposi] = &DungeonCell{Type: tileWall}
				}
			}
		} else {
			for col := minx; col < maxx+1; col++ {
				if col != minx+doorposi {
					d.Grid[miny+cutposi][col] = &DungeonCell{Type: tileWall}
				}
			}
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

	// make exit
	d.Grid[d.Height-2][d.Width-1].Type = tileEscape
	d.Grid[d.Height-2][d.Width-1].Message = "You escaped!"

	// remove what's under the player
	d.Grid[1][1] = &DungeonCell{
		Type: tileFloor,
	}
	d.GenerateSecretRoom()
}

func (d *Dungeon) GetSmallGrid(width, height int) [][]*DungeonCell {
	var smallGrid [][]*DungeonCell
	x := d.Position.X - width/2
	y := d.Position.Y - height/2

	for i := y; i < y+height; i++ {
		currentRow := make([]*DungeonCell, width)
		for j := x; j < x+width; j++ {
			currentRow[j-x] = &DungeonCell{}
			if i < 0 || i >= d.Height || j < 0 || j >= d.Width {
				currentRow[j-x].Type = tileVoid
			} else {
				currentRow[j-x] = d.Grid[i][j]
			}
		}
		smallGrid = append(smallGrid, currentRow)
	}

	return smallGrid
}

func (c *DungeonCell) Generate() {
	c.ID = uuid.NewV4().String()
	rng := rand.Float64() * 100

	if rng < 6.5 && rng > 0 {
		c.Type = tileMoney
		c.Amount = rand.Intn(100) + 50
		c.Message = fmt.Sprintf(findMoney, c.Amount)
		return
	}

	rng -= 6.5

	if rng < 3 && rng > 0 {
		c.Type = tileEnnemy
		c.Amount = 3
		c.Damages = rand.Intn(15) + 7
		c.Message = fmt.Sprintf(findEnnemy, c.Damages)
		return
	}

	rng -= 3

	if rng < 1 && rng > 0 {
		c.Type = tileArm
		c.Message = findArm
		return
	}

	rng -= 1

	if rng < 1 && rng > 0 {
		c.Type = tileMinion
		c.Message = findMinion
		return
	}

	rng -= 1

	if rng < 0.5 && rng > 0 {
		c.Type = tileAmbrosius
		c.Choices = generateChoices(2)
		return
	}

	rng -= 0.5

	c.Type = tileFloor
}

func (c *DungeonCell) GenerateWall() {
	c.ID = uuid.NewV4().String()
	rng := rand.Float64() * 100

	if rng < 1 && rng > 0 {
		c.Type = tileTPWall
		return
	}
}
