package coderetreat2023

import "fmt"

type Grid [][]bool

func (g Grid) String() string {

	aliveString := "*"
	deadString := "-"
	ret := ""

	stringToPrint := make(map[bool]string)
	stringToPrint[true] = aliveString
	stringToPrint[false] = deadString

	for _, row := range g {
		for _, val := range row {
			ret += fmt.Sprintf(stringToPrint[val] + " ")
		}
		ret += "\n"
	}

	return ret
}

func ParseGridInput() Grid {
	rowLen := 0
	colLen := 0
	fmt.Scanf("%d %d", &rowLen, &colLen)

	inp := 0
	grid := make([][]bool, rowLen)
	for i := 0; i < rowLen; i++ {
		grid[i] = make([]bool, colLen)
		for j := 0; j < colLen; j++ {
			fmt.Scanf("%d", &inp)
			if inp == 1 {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	return grid
}
