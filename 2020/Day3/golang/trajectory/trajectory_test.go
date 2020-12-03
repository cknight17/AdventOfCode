package trajectory

import "testing"

func TestGetInput(t *testing.T) {
	want := 11
	lines := GetInput("test.txt")
	if got := len(lines); want != got {
		t.Errorf("getInput(%q) = %d, want %d","test.txt",got,want)
   	}
}

func genericMoveTest(lines []string, fromX int, fromY int, moveX int, moveY int, count int, wantX int, wantY int, wantCount int, t *testing.T) (int,int,int) {
	toX, toY, newCount := Move(fromX,fromY,moveX,moveY,lines,count)
	if wantX != toX || wantY != toY || wantCount != newCount {
		t.Errorf("move(%d,%d,%d,%d,lines,%d) = %d %d %d, want %d %d %d",fromX,fromY,moveX,moveY,count,toX,toY,newCount,wantX,wantY,wantCount)
	}
	return toX, toY, newCount  
}

func TestMoveSteps(t *testing.T) {
	lines := GetInput("test.txt")
	atX := 0
	atY := 0
	right := 3
	down := 1
	count := 0
	countAt := []int{0,1,1,2,3,3,4,5,6,7,7}
	for at, _ := range lines {
		atX, atY, count = genericMoveTest(lines,atX,atY,right,down,count,atX + right,atY + down,countAt[at],t)	
	}
}

func TestMoveAll(t *testing.T) {
	lines := GetInput("test.txt")
	gotX,gotY,gotCount := MoveAll(0,0,3,1,lines)
	wantX := 33
	wantY := 11
	wantCount := 7	
	if wantX != gotX || wantY != gotY || wantCount != gotCount {
		t.Errorf("MovaAll(0,0,3,1,lines) = %d %d %d, want %d %d %d",gotX,gotY,gotCount,wantX,wantY,wantCount)
	}
}

func TestMoveAllPart2A(t *testing.T) {
	lines := GetInput("test.txt")
	_,_,got := MoveAll(0,0,1,1,lines)
	want := 2
	if want != got {
		t.Errorf("MovaAll(0,0,1,1,lines) = ? ? %d, want ? ? %d",got,want)
	}
}

func TestMoveAllPart2B(t *testing.T) {
	lines := GetInput("test.txt")
	_,_,got := MoveAll(0,0,3,1,lines)
	want := 7
	if want != got {
		t.Errorf("MovaAll(0,0,3,1,lines) = ? ? %d, want ? ? %d",got,want)
	}
}

func TestMoveAllPart2C(t *testing.T) {
	lines := GetInput("test.txt")
	_,_,got := MoveAll(0,0,5,1,lines)
	want := 3
	if want != got {
		t.Errorf("MovaAll(0,0,5,1,lines) = ? ? %d, want ? ? %d",got,want)
	}
}

func TestMoveAllPart2D(t *testing.T) {
	lines := GetInput("test.txt")
	_,_,got := MoveAll(0,0,7,1,lines)
	want := 4
	if want != got {
		t.Errorf("MovaAll(0,0,7,1,lines) = ? ? %d, want ? ? %d",got,want)
	}
}

func TestMoveAllPart2E(t *testing.T) {
	lines := GetInput("test.txt")
	_,_,got := MoveAll(0,0,1,2,lines)
	want := 2
	if want != got {
		t.Errorf("MovaAll(0,0,1,2,lines) = ? ? %d, want ? ? %d",got,want)
	}
}

func TestMoveAllPart2Final(t *testing.T) {
	lines := GetInput("test.txt")
	input := [][]int{[]int{0,0,1,1},[]int{0,0,3,1},[]int{0,0,5,1},[]int{0,0,7,1},[]int{0,0,1,2}}
	got := 1
	for _, item := range input {
		_, _, count := MoveAll(item[0],item[1],item[2],item[3],lines)
		got = got * count
	}
	want := 336
	if want != got {
		t.Errorf("MovaAll() = ? ? %d, want ? ? %d",got,want)
	}
}