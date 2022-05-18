package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Ideas struct {
	Index int
	Idea  string
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func writeToFile(path string, list *[]Ideas) {

	bytes, err := json.Marshal(&list)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = ioutil.WriteFile(path, bytes, 0744)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func readFromFile(path string, list *[]Ideas) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, list)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func printList(list *[]Ideas) {

	if len(*list) == 0 {
		fmt.Println("Ideas List is empty...")
	}

	for i := 0; i < len(*list); i++ {
		fmt.Printf("%d. %s\n", (*list)[i].Index, (*list)[i].Idea)
	}
}

func removeFromList(index int, list *[]Ideas) *[]Ideas {
	newList := make([]Ideas, 0)
	newList = append(newList, (*list)[:index]...)
	newList = append(newList, (*list)[index+1:]...)

	for i := index; i < len(newList); i++ {
		newList[i].Index = i + 1
	}

	return &newList

}

func editList(index int, text string, list *[]Ideas) {
	(*list)[index].Idea = text
}

func main() {

	var path string = "ideas.json"
	var list []Ideas
	var listPointer = &list
	var write = true

	viewcmd := flag.NewFlagSet("view", flag.ExitOnError)
	newcmd := flag.NewFlagSet("new", flag.ExitOnError)
	delcmd := flag.NewFlagSet("del", flag.ExitOnError)
	editcmd := flag.NewFlagSet("edit", flag.ExitOnError)

	idea := newcmd.String("I", "", "The idea to add to the Ideas List")
	delIndex := delcmd.Int("i", -1, "The index of the idea to remove from Ideas List")
	editIndex := editcmd.Int("i", -1, "The index of the ideas to edit")
	editText := editcmd.String("t", "", "The text to replace")

	if len(os.Args) < 2 {
		fmt.Println("Subcommand is requiried...")
		fmt.Println("Try: view new del edit")
		os.Exit(1)
	}

	//Check if list.json exist
	if !fileExist(path) {
		writeToFile(path, listPointer)
	}

	readFromFile(path, listPointer)

	switch os.Args[1] {
	case "view":
		viewcmd.Parse(os.Args[2:])

		write = false
	case "new":
		newcmd.Parse(os.Args[2:])
		list = append(list, Ideas{len(list) + 1, *idea})
		listPointer = &list
	case "del":
		delcmd.Parse(os.Args[2:])
		if *delIndex == -1 {
			fmt.Println("Invalid Index...")
			os.Exit(1)
		}
		listPointer = removeFromList(*delIndex-1, listPointer)
	case "edit":
		editcmd.Parse(os.Args[2:])
		if *editIndex == -1 {
			fmt.Println("Invalid Index...")
			os.Exit(1)
		}
		editList(*editIndex-1, *editText, listPointer)
	}

	printList(listPointer)

	if write {
		writeToFile(path, listPointer)
	}

}
