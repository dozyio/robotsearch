package main

//safe area = 592597

import (
	"container/list"
	"fmt"
	"math"
)

var (
	emp   int = 24
	queue     = list.New()
)

type point interface {
	getX() int
	getY() int
}

type coord struct {
	x int
	y int
}

func (c coord) getX() int {
	return c.x
}
func (c coord) getY() int {
	return c.y
}

func main() {
	x := 0
	y := 0
	maxLength := 0
	count := 0

	//Find y for x = 0. Equals max before a bomb
	for {
		ok := isSafe(x, y)
		if !ok {
			//use previous good y + 1, which is y
			maxLength = y
			//fmt.Println("Max Length: ", maxLength)
			break
		}
		y++
	}

	//Offset area (*2) to avoid negative index and create 2d array
	offset := maxLength * 2
	var visited = make([][]bool, offset)
	for i := 0; i < offset; i++ {
		visited[i] = make([]bool, offset)
	}

	start := coord{x: 0, y: 0}
	queue.PushFront(start)

	//Process area queue
	for queue.Len() > 0 {
		e := queue.Front()
		p := e.Value.(point)
		if visited[p.getX()+maxLength][p.getY()+maxLength] == false {
			visited[p.getX()+maxLength][p.getY()+maxLength] = true
			count++
			checkSurrounding(p)
		}
		queue.Remove(e)
	}

	fmt.Println("Safe area: ", count)
}

//Get (x+1, y) position
func xRight(p point) point {
	return coord{p.getX() + 1, p.getY()}
}

//Get (x-1, y) position
func xLeft(p point) point {
	return coord{p.getX() - 1, p.getY()}
}

//Get (x, y+1) position
func yUp(p point) point {
	return coord{p.getX(), p.getY() + 1}
}

//Get (x, y-1) position
func yDown(p point) point {
	return coord{p.getX(), p.getY() - 1}
}

//Check the surrounding area and add to queue if safe
func checkSurrounding(p point) {
	i := xRight(p)
	if isSafe(i.getX(), i.getY()) {
		queue.PushBack(i)
	}

	i = xLeft(p)
	if isSafe(i.getX(), i.getY()) {
		queue.PushBack(i)
	}

	i = yUp(p)
	if isSafe(i.getX(), i.getY()) {
		queue.PushBack(i)
	}

	i = yDown(p)
	if isSafe(i.getX(), i.getY()) {
		queue.PushBack(i)
	}
}

//Calculate the sum of digits in abs integer
func sumDigits(i int) int {
	sum := 0
	for i = int(math.Abs(float64(i))); i > 0; i = i / 10 {
		sum += (i % 10)
	}
	return sum
}

//Check if coordinates are safe from emp
func isSafe(x int, y int) bool {
	if sumDigits(x)+sumDigits(y) < emp {
		return true
	}
	return false
}
