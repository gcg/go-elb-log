package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "elbla"
	app.Usage = "point to your logs and I will do the stuff"

	app.Action = func(c *cli.Context) {

		filePresent := c.Args().Present()

		if filePresent == false {
			println("Please supply a file. You can use help command to get help.")
			os.Exit(1)
		}

		p := process(c.Args().First(), c)

		if p == false {
			panic("EOF")
		}
	}

	app.Run(os.Args)
}

func process(file string, c *cli.Context) bool {

	println("Ok let me read the file: ", file)

	filePath, _ := filepath.Abs(file)

	f, err := os.Open(filePath)

	if err != nil {
		println("Error on opening the file: ", filePath)
		return false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile("([^ ]*) ([^ ]*) ([^ ]*):([0-9]*) ([^ ]*):([0-9]*) ([.0-9]*) ([.0-9]*) ([.0-9]*) (-|[0-9]*) (-|[0-9]*) (-|[0-9]*) (-|[0-9]*) \"([^ ]*) ([^ ]*) (- |[^ ]*)\" \"(.*?)\" ([^ ]*) ([^ ]*)")

	for scanner.Scan() {

		line := scanner.Text()

		println(line)

		cols := re.Split(line, -1)
		fmt.Printf("%v", cols)
		println("")
		println("Line 1 ended.")
		os.Exit(2)
	}

	return true
}
