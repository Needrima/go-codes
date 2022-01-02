package main

import (
	"errors"
	"fmt"
	"log"
)

type matrix struct {
	layout [][]int
}

var (
	ErrUnequalRowOrColumn    = "matrix row or column unequal"
	ErrMinRowAndColumnLength = "row and column number must be atleast 1"
	ErrRowNotEqualColumn     = "row not equal column"
)

// creates a new matrix with all its values set to 0 (i.e zero matrix)
func newMatrix(rowNum, colNum int) matrix {
	if rowNum == 0 || colNum == 0 {
		log.Println(ErrMinRowAndColumnLength)
		return matrix{}
	}

	l := [][]int{}

	for i := 0; i < rowNum; i++ {
		col := []int{}

		for j := 0; j < colNum; j++ {
			col = append(col, 0)
		}

		l = append(l, col)
	}

	return matrix{layout: l}
}

// sets the value of an element given its row and column positions
// with positions from 0 to n-1
func (m *matrix) set(rPos, cPos, value int) {
	m.layout[rPos][cPos] = value
}

//adds two matrices
func (m matrix) add(m2 matrix) (matrix, error) {
	row1, col1 := getRowAndColumnNum(m)

	row2, col2 := getRowAndColumnNum(m2)

	if row1 != row2 || col1 != col2 {
		return matrix{}, errors.New(ErrUnequalRowOrColumn)
	}

	sum := newMatrix(row1, col1) //sum will have same dimension as base matrices

	for i := 0; i < row1; i++ {
		for j := 0; j < col1; j++ {
			sum.set(i, j, m.layout[i][j]+m2.layout[i][j])
		}
	}

	return sum, nil
}

func (m matrix) multiply(m2 matrix) (matrix, error) {
	row1, col1 := getRowAndColumnNum(m) // dim = row1 x col1

	row2, col2 := getRowAndColumnNum(m2) // dim = row2 x col2

	if col1 != row2 { // product is defined only if number of columns in m is equal number of rows in m2
		return matrix{}, errors.New(ErrRowNotEqualColumn)
	}

	product := newMatrix(row1, col2) // dim = row1 x col2

	for i := range m.layout {

		for _, v := range m2.layout {

			for k := range v {
				column := getColumnFromMatrix(m2, k)

				sum := performMultiplicationOperaton(m.layout[i], column)

				product.set(i, k, sum)
			}

		}

	}

	return product, nil
}

func (m *matrix) print() {
	for _, v := range m.layout {
		fmt.Println(v)
	}
}

// get the number of rows and columns in a matrix respectively
func getRowAndColumnNum(m matrix) (int, int) {
	var colLength int

	for _, v := range m.layout {
		for i := range v {
			colLength = i + 1
		}
	}

	return len(m.layout), colLength
}

func getColumnFromMatrix(m matrix, colNum int) []int {
	column := []int{}

	for _, v := range m.layout {
		column = append(column, v[colNum])
	}

	return column
}

func performMultiplicationOperaton(row, col []int) int {
	var sum int

	for i := range row {
		sum += row[i] * col[i]
	}

	return sum
}

func main() {
	// m1 := newMatrix(2, 3)
	// m2 := newMatrix(2, 3)

	// m1.set(0, 0, 4)
	// m1.set(0, 1, -3)
	// m1.set(0, 2, 5)
	// m1.set(1, 0, -2)
	// m1.print()
	// fmt.Println("---------------")

	// m2.set(0, 0, 4)
	// m2.set(0, 1, -3)
	// m2.set(0, 2, 5)
	// m2.set(1, 0, -2)
	// m2.print()
	// fmt.Println("------------------")

	// add, _ := m1.add(m2)
	// add.print()

	// m1 := newMatrix(2, 2)
	// m2 := newMatrix(2, 3)

	// m1.set(0, 0, 2)
	// m1.set(0, 1, -1)
	// m1.set(1, 0, 1)
	// m1.set(1, 1, 0)

	// m2.set(0, 0, 1)
	// m2.set(0, 1, 0)
	// m2.set(0, 2, 1)
	// m2.set(1, 0, 2)
	// m2.set(1, 1, 4)
	// m2.set(1, 2, 5)

	// m1.print()
	// fmt.Println("----------------------")
	// m2.print()
	// fmt.Println("----------------------")

	// product, err := m1.multiply(m2)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// product.print()

}
