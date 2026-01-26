package main

import (
	"fmt"
)

func Intro() {
	fmt.Println("Welcome to the Harvester")
	fmt.Println("Enter any DOC URL and i'll download any page under it")
	fmt.Println("All data will be saved into a .txt files in /data/\n")
}

func TakeInput() (string, error) {
	var URL string
	fmt.Print("Enter DOC URL: ")
	_, err := fmt.Scanln(&URL)
	return URL, err
}
