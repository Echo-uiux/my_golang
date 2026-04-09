package todo

import (
	"fmt"
	"os"
	"rest/todo"
	"strings"
)

const todoFilename = ".todo.Json"

func main() {
	l := &todo.Todolist{}
	if err := l.Open(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Printf(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")

		l.Add(item)
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
