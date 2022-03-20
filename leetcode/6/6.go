package leetcode

import "strings"

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var simulatedString []strings.Builder
	for i := 0; i < numRows; i++ {
		simulatedString = append(simulatedString, strings.Builder{})
	}
	var walker ConverterWalker
	walker = newZigZagWalker(numRows)
	for _, char := range s {
		pointer := walker.CurrentPoint()
		simulatedString[pointer].WriteRune(char)
		walker.Walk()
	}
	result := strings.Builder{}
	for i := 0; i < numRows; i++ {
		result.WriteString(simulatedString[i].String())
	}
	return result.String()
}

type ConverterWalker interface {
	CurrentPoint() int
	Walk()
}

type zigZagWalker struct {
	state            zigZagWalkerState
	pointer          int
	numberOfRows     int
}

func (w *zigZagWalker) CurrentPoint() int {
	return w.pointer
}

func (w *zigZagWalker) Walk() {
	w.state.Move(w)
}

func (w *zigZagWalker) changeState(state zigZagWalkerState) {
	w.state = state
}

func newZigZagWalker(numberOfRows int) *zigZagWalker {
	return &zigZagWalker{
		state:            &zigZagWalkerStateDown{},
		pointer:          0,
		numberOfRows:     numberOfRows,
	}
}

type zigZagWalkerState interface {
	Move(walker *zigZagWalker)
}

type zigZagWalkerStateDown struct {
}

func (s *zigZagWalkerStateDown) Move(walker *zigZagWalker) {
	if walker.pointer == walker.numberOfRows-1 {
		walker.changeState(&zigZagWalkerStateUp{})
		walker.pointer--
	} else {
		walker.pointer++
	}
}

type zigZagWalkerStateUp struct {
}

func (s *zigZagWalkerStateUp) Move(walker *zigZagWalker) {
	if walker.pointer == 0 {
		walker.changeState(&zigZagWalkerStateDown{})
		walker.pointer++
	} else {
		walker.pointer--
	}
}
