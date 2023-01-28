package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

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
				oldPath := c.Args().Get(0)
				newPath := c.Args().Get(1)
				moveGo(oldPath, newPath)

				if strings.HasPrefix(oldPath, "./") {
					oldPath = oldPath[2:]
				}
				if strings.HasPrefix(newPath, "./") {
					newPath = newPath[2:]
				}
				files := getGoFilesList(".")
				fmt.Println(files)

				replaceInFile(files, oldPath, newPath)
			},
		},
	}
	app.Run(os.Args)
	app.Usage = "Golang tolls, but we called it go toys"
}

func moveGo(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func replaceInFile(listFiles []string, oldText, newText string) {
	for _, pathFile := range listFiles {
		data, err := ioutil.ReadFile(pathFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := strings.Split(string(data), "\n")
		output := []string{}

		inBlockImport := false
		for _, line := range lines {
			if inBlockImport && strings.Contains(line, ")") {
				inBlockImport = false
			} else if strings.Contains(line, "import") && strings.Contains(line, "(") {
				inBlockImport = true
			} else if strings.HasPrefix(line, "import") || inBlockImport {
				output = append(output, strings.Replace(line, oldText, newText, -1))
				continue
			}
			output = append(output, line)
		}

		if err = ioutil.WriteFile(pathFile, []byte(strings.Join(output, "\n")), 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func getGoFilesList(path string) []string {
	listFiles := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			listFiles = append(listFiles, getGoFilesList(path+"/"+file.Name())...)

		} else if bool, err := regexp.Match(`.go`, []byte(file.Name())); err == nil && bool {
			listFiles = append(listFiles, path+"/"+file.Name())
		}
	}

	return listFiles
}
