package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readAndRemove()  {
	files, err := ioutil.ReadDir("./")
	
	if err != nil {
		panic(err)
	}
	
	for _, file := range files {
		if file.Name()[0] == '.' {
			continue
		}

		if file.Name() == "node_modules" {
			err := os.Remove(file.Name())
			
			if err != nil {
				panic(err)
			}

			continue
		}

		if file.IsDir() {
			err := os.Chdir(fmt.Sprintf("./%s", file.Name()))
			
			if err != nil {
				panic(err)
			}

			readAndRemove()
		}
	}
}

func main() {
	if len(os.Args) == 2 {
		os.Chdir(os.Args[1])
	}

	readAndRemove()
}