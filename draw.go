package main

import (
	//"time"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	tm.MoveCursor(12, 12)
	line := tm.NewLineChart(12, 12)
	tm.Println(line)
	tm.Flush()
}
