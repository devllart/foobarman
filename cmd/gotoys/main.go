package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotoys"
	app.Commands = []cli.Command{
		{
			Name:  "mv",
			Usage: "move file and fix imports",
			Action: func(c *cli.Context) {
				moveGo(c.Args().Get(0), c.Args().Get(1))
			},
		},
	}
	app.Run(os.Args)
	app.Usage = "Golang tolls, but we called it go toys"
}

func moveGo(oldLocation, newLocation string) {
	fmt.Println(oldLocation, newLocation)

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}

}

func replaceInFile(listFiles []string, oldText, newText string) {
	for _, pathFile := range listFiles {
		input, err := ioutil.ReadFile(pathFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		output := bytes.ReplaceAll(input, []byte(oldText), []byte(newText))

		if err = ioutil.WriteFile(pathFile, output, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
