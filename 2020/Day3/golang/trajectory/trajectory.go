package trajectory

import "aoc.knighttechnology.net/fileParser"

func GetInput(file string) []string {
	return fileParser.Linereader(file)
}

func Move(fromX int, fromY int, rightX int, rightY int, grid []string, count int) (int, int, int) {
	//maxRows := len(grid)
	maxCols := len(grid[0])
	toX := fromX + rightX
	toY := fromY + rightY
	newCount := count
	if grid[toY][toX % maxCols] == '#' {
		newCount++
	}
	return toX, toY, newCount
}

func MoveAll(startX int, startY int, right int, down int, grid []string) (int,int,int) {
	atX := 0
	atY := 0
	count := 0
	for range grid {
		if (atY + down >= len(grid)) {
			break
		}
		atX, atY, count = Move(atX,atY,right,down,grid,count)
		
	}
	return atX, atY, count
}
