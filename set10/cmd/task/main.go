package main

import (
	"fmt"
	"os"

	"learn-go-by-exercises/set10/task"
)

const dbPath = "tasks.json"

func main() {
	store, err := task.Load(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "add":
		runAdd(store, os.Args[2:])
	case "list":
		runList(store)
	case "done":
		runDone(store, os.Args[2:])
	default:
		usage()
	}
}

func runAdd(store *task.Store, args []string)  { /* TODO: join, Add, Save, print */ }
func runList(store *task.Store)                { /* TODO: print pending tasks */ }
func runDone(store *task.Store, args []string) { /* TODO: parse ID, Complete, Save */ }

func usage() { fmt.Println("usage: task <add|list|done>") }
