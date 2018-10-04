package main

import "fmt"
import "github.com/montanaflynn/stats"

func main() {
	data := stats.LoadRawData([]interface{}{1.1, "2", 3.0, 4, "5"})
	max, _ := stats.Max(data)
	fmt.Println(max)
}
