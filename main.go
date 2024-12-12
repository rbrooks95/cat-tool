package main

import (
	"io"
	"log"
	"os"
)
func main()  {
	if len(os.Args) < 2 {
		log.Fatal("usage: cat-tool for file")
	}
	for _, fileName := range os.Args[1:]{
		if err := catFile(fileName); err != nil {
			log.Fatal(err)
		}
	}
}


func catFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	
	
	_, err = io.Copy(os.Stdout, f)
	return err
}