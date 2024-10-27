package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool){
	if len(s.vals) > 0{
		s.vals.
	}
}

func main() {
}

func speed() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	thing()
}

func thing() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}
