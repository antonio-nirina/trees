package main

import (
	"fmt"
	// "flag"
	// "log"
	"io/ioutil"

	"github.com/fatih/color"
)

// go build -o /path/go/bin
func mainff1() {
	var array []string
	files, _ := ioutil.ReadDir(".")
	green := color.New(color.FgRed)
	read := color.New(color.FgGreen)
	c := color.New(color.FgCyan)
	// indent := "├──"
	// nextIndent := " │

	for _, val := range files {
		if val.IsDir() {
			read.Println(val.Name())
			path := fmt.Sprintf("%s%s", "./", val.Name())
			files1, _ := ioutil.ReadDir(path)
			for _, sub := range files1 {
				c.Println(" |----", sub.Name())
				if sub.IsDir() {
					path2 := fmt.Sprintf("%s%s%s%s", "./", val.Name(), "/", sub.Name())
					files2, _ := ioutil.ReadDir(path2)
					for _, sub2 := range files2 {
						c.Println("     |----", sub2.Name())
					}
				} else {
					c.Println("     |****", sub.Name())
				}
			}
		} else {
			array = append(array, val.Name())
		}

	}

	if len(array) > 0 {
		for _, res := range array {
			green.Println(res)
		}
	}
}
