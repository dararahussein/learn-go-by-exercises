/*

main.go — the command-line front end for your task store.

Once `go test ./set06/` is green, fill in runAdd / runList / runDone below, then
use your program for real:

    go run ./set06 add "learn go"
    go run ./set06 add "build something"
    go run ./set06 list
    go run ./set06 done 1
    go run ./set06 list

Tasks get saved to a tasks.json file in whatever directory you run from — open
it and watch it change.

*/

package main

import (
	"fmt"
	"os"
)

const dbPath = "tasks.json"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		return
	}

	store, err := Load(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not load tasks:", err)
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		runAdd(store, args[1:])
	case "list":
		runList(store)
	case "done":
		runDone(store, args[1:])
	default:
		usage()
	}
}

// runAdd should join args into the task text, store.Add(text), store.Save(dbPath),
// and print something like "added #1: learn go".
func runAdd(store *Store, args []string) {
	// TODO: implement the `add` command (you'll want to import "strings" for
	// strings.Join). Right now it does nothing.
}

// runList should print each pending task (store.Pending()) as "#<id> <text>",
// or "no pending tasks" when there are none.
func runList(store *Store) {
	// TODO: implement the `list` command.
}

// runDone should parse args[0] into an int id (import "strconv", strconv.Atoi),
// call store.Complete(id), report any error and exit, otherwise store.Save(dbPath)
// and print "completed #<id>".
func runDone(store *Store, args []string) {
	// TODO: implement the `done` command.
}

func usage() {
	fmt.Println("usage: go run ./set06 <command>")
	fmt.Println("  add <text...>   add a new task")
	fmt.Println("  list            list pending tasks")
	fmt.Println("  done <id>       mark a task complete")
}
