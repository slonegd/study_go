package main

import "fmt"

// Transformer ...
type Transformer interface {
	Transform (int, int) int
}

// Plus ...
type Plus struct {}

// Transform ...
func (p *Plus) Transform (v1, v2 int) int {
	return v1 + v2 
}

// Multiple ...
type Multiple struct {}


// Transform ...
func (m *Multiple) Transform (v1, v2 int) int {
	return v1 * v2 
}

// Task ...
type Task struct {
	transformer Transformer
}

func (t *Task) do(v1, v2 int) {
	fmt.Printf("v1: %v, v2: %v, result:%v\n", v1 , v2, t.transformer.Transform(v1 ,v2))
}

func main() {
	plus := &Plus{}
	plusTask := &Task{plus}
	plusTask.do(1,2)
	plusTask.do(3,4)

	multiple := &Multiple{}
	multipleTask := &Task{multiple}
	multipleTask.do(1,2)
	multipleTask.do(3,4)
}