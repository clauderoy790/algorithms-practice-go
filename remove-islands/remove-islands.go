package remove_islands

import "fmt"

var output [][]int

var matrix = [][]int{
	{1, 0, 0, 0, 0, 0},
	{0, 1, 0, 1, 1, 1},
	{0, 0, 1, 0, 1, 0},
	{1, 1, 0, 0, 1, 0},
	{1, 0, 1, 1, 0, 0},
	{1, 0, 0, 0, 0, 1},
}

//var matrix2 = [][]int{
//	{0, 0, 0, 0, 0, 0},
//	{0, 0, 0, 1, 1, 1},
//	{0, 0, 0, 0, 1, 0},
//	{1, 1, 0, 0, 1, 0},
//	{1, 0, 0, 0, 0, 0},
//	{1, 0, 0, 0, 0, 1},
//}

var borderPoints []Point

var neighbors = []Point{
	{0,-1}, 	// top
	{1,0},	// right
	{0,1},	// bot
	{-1,0},	// left
}

type IslandsRemover struct {}

func (isl *IslandsRemover) Resolve() {
	fmt.Println("before:")
	printMatrix(matrix)
	removeIslands()
	fmt.Println("after:")
	printMatrix(output)
}

func removeIslands() {
	output = createSameSizeMatrix(matrix)
	borderPoints = createBorderPointList()

	for _,pt := range borderPoints {
		addIslandAndNeighborsToOutput(pt)
	}
}

func createBorderPointList() []Point {
	var pts []Point

	for x,y := 0,0; x <len(matrix[0]); x++ { // top borderPoints points
		pts = append(pts,Point{x,y})
	}

	for x,y := len(matrix[0])-1,0; y < len(matrix);y++ { // right borderPoints points
		pts = append(pts,Point{x,y})
	}

	for x,y := 0,len(matrix)-1; x < len(matrix[0]);x++ { // bottom borderPoints points
		pts = append(pts,Point{x,y})
	}

	for x,y := 0,0; y < len(matrix);y++ { // left borderPoints points
		pts = append(pts,Point{x,y})
	}
	return pts
}

func addIslandAndNeighborsToOutput(point Point) {
	if !isIslandInOutput(point) && isIsland(point) {
		output[point.x][point.y] = 1
		for _,nb := range neighbors {
			neighborPt := Point{point.x+nb.x,point.y+nb.y}
			addIslandAndNeighborsToOutput(neighborPt)
		}
	}
}

func createSameSizeMatrix(matrix [][]int) [][]int {
	mat := make([][]int, len(matrix))
	for i := range mat {
		mat[i] = make([]int, len(matrix[0]))
	}

	return mat
}

func isIslandInOutput(pt Point) bool {
	return isInBounds(pt) && output[pt.x][pt.y] == 1
}

func isInBounds(pt Point) bool {
	return pt.x >= 0 && pt.x < len(matrix[0]) && pt.y >= 0 && pt.y < len(matrix)
}

func isIsland(pt Point) bool {
	return isInBounds(pt) && matrix[pt.x][pt.y] == 1
}

func printMatrix(ma [][]int) {
	for i := 0; i < len(ma); i++ {
		fmt.Println(ma[i])
	}
}

type Point struct {
	x, y int
}
