package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
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
