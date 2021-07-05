package remove_islands

import "fmt"

//{0, 1, 1, 1, 0, 0},

var matrix = [6][6]int{
	{1, 0, 0, 0, 0, 0},
	{0, 1, 0, 1, 1, 1},
	{0, 0, 1, 0, 1, 0},
	{1, 1, 0, 0, 1, 0},
	{1, 0, 1, 1, 0, 0},
	{1, 0, 0, 0, 0, 1},
}

var matrix2 = [6][6]int{
	{0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1},
	{0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 0},
	{1, 0, 0, 0, 0, 1},
}

func RemoveIslandsEx() {
	fmt.Println("before:")
	printMatrix()
	removeIslands()
	fmt.Println("after:")
	printMatrix()
}

func removeIslands() {
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[0])-1; j++ {
			var left, right, bot, top int
			for left = -1; left < len(matrix[i]); {
				if left >= 0 && matrix[i][left] == 1 {
					left++
				} else {
					break
				}
			}
			for right = len(matrix[i]); right > 0; {
				if matrix[i][right-1] == 1 {
					right--
				} else {
					break
				}
			}
			for top = -1; top < len(matrix); {
				if top >= 0 && matrix[top][j] == 1 {
					top++
				} else {
					break
				}
			}
			for bot = len(matrix); bot > 0; {
				if matrix[bot-1][j] == 0 {
					bot--
				} else {
					break
				}
			}

			if j == 4 && i == 1 {
				fmt.Println(fmt.Sprintf("l:%v, t:%v, r:%v, b:%v",left,top,right,bot))
			}

			if matrix[i][j] == 1 && j > left && j < right && i > top && j < bot {
				matrix[i][j] = 0
			}
		}
	}
}

func isConnectedBeforeLeft() bool {
	return true
}

func isConnectedBeforeTop() {

}

func isConnectedAfterRight() {

}

func isConnectedAfterBottom() {

}

func printMatrix() {
	for i := 0; i < len(matrix); i++ {
		//for j := 0; j <len(matrix[0]);j++ {
		//}
		fmt.Println(matrix[i])
	}
}
