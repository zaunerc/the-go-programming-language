package main

import "fmt"
import "github.com/OGFris/GoStats"

func main() {
	data := []float64{1, 2, 3}
	_, max := GoStats.Max(data)
	fmt.Println(max)
}
