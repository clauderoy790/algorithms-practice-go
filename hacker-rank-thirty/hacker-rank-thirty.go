package hacker_rank_thirty

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HackerRankThirty struct{}

func (hr *HackerRankThirty) Resolve() {
	dayEleven()
}

func scannerExample() {
	scanner := bufio.NewScanner(os.Stdin)

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	scanner.Scan()
	nb, _ := strconv.ParseUint(scanner.Text(), 10, 8)
	scanner.Scan()
	fl, _ := strconv.ParseFloat(scanner.Text(), 8)
	scanner.Scan()
	str := scanner.Text()

	nb2 := nb + i
	nb3 := fl + d
	fmt.Println(nb2)
	fmt.Printf("%.1f\n", nb3)
	fmt.Printf("%v%v\n", s, str)
}

func moduloString() {
	fmt.Println("enter nb")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	for i := 0; i < nb; i++ {
		beg, end := "", ""
		scanner.Scan()
		text := scanner.Text()
		for i, r := range text {
			if i%2 == 0 {
				beg += string(r)
			} else {
				end += string(r)
			}
		}
		fmt.Printf("%v %v\n", beg, end)
	}
}

func daySeven() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Text()
	//nb, err := strconv.Atoi(scanner.Text())
	//if err != nil {
	//	panic(err)
	//}
	//sl := make([]int,nb)

	scanner.Scan()
	strs := strings.Split(scanner.Text(), " ")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	fmt.Println(strings.Join(strs, " "))
}

func dayEight() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("error scanning number: ", err)
	}

	m := make(map[string]int)
	for i := 0; i < nb; i++ {
		scanner.Scan()
		sl := strings.Split(scanner.Text(), " ")
		if len(sl) != 2 {
			panic("need to only have one space in the string")
		}
		v, err := strconv.Atoi(sl[1])
		if err != nil {
			panic(err)
		}
		m[sl[0]] = v
	}

	for scanner.Scan() {
		key := scanner.Text()
		if len(key) == 0 {
			break
		}
		if v, ok := m[key]; ok {
			fmt.Printf("%v=%v\n", key, v)
		} else {
			fmt.Println("Not found")
		}

	}
}

func dayNine() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	var fac func(n int) int
	fac = func(n int) int {
		if n <= 0 {
			return 0
		} else if n == 1 {
			return 1
		} else {
			return n * fac(n-1)
		}
	}

	res := fac(nb)
	fmt.Println(res)
}

func dayTen() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nb, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	biStr := strconv.FormatInt(int64(nb), 2)
	maxOne, adjacentOne := 0, 0
	for _, r := range biStr {
		isOne := rune(r) == '1'
		if isOne {
			adjacentOne++
		} else {
			adjacentOne = 0
		}
		if adjacentOne > maxOne {
			maxOne = adjacentOne
		}
	}
	fmt.Println(maxOne)
}
const (
	INT_MAX = int32(^uint32(0) >> 1)
	INT_MIN = ^INT_MAX
)
func dayEleven() {
	hWidth := 3
	//matrix := [][]int32{
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 1, 0, 0, 0, 0},
	//	{1, 1, 1, 0, 0, 0},
	//	{0, 0, 2, 4, 4, 0},
	//	{0, 0, 0, 2, 0, 0},
	//	{0, 0, 1, 2, 4, 0},
	//}
	matrix := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{-1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}
	biggestSum := INT_MIN
	points := getHourglassPoints(matrix, hWidth)

	for _, p := range points {
		sum := sumHourglass(matrix, p, hWidth)
		if sum > biggestSum {
			biggestSum = sum
		}
	}
	fmt.Println("the biggest sum for the matrix is: ", biggestSum)
}

func sumHourglass(matrix [][]int32, p Point, width int) (sum int32) {
	removedPoints := 0
	pointIncrement := 1
	for y := 0; y < width; y++ { // Loop the height
		for x := 0; x < width;x++ {
			if x >= removedPoints && x < width - removedPoints {
				sum += matrix[p.y +y][p.x+x]
			}
		}
		if width-removedPoints*2 == 1 { // change direction
			pointIncrement *= -1
		}
		removedPoints += pointIncrement
	}
	return
}

func getHourglassPoints(matrix [][]int32, width int) []Point {
	var pts []Point
	matWidth, matHeight := len(matrix[0]), len(matrix)
	for y := 0; y < matHeight;y++ {
		for x := 0; x < matWidth;x++ {
			if x+width-1 < matWidth && y+width-1 < matHeight {
				pts = append(pts, Point{x, y})
			}
		}
	}
	return pts
}

type Point struct {
	x, y int
}
