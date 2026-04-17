// main.go
package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("usage: task-cli <command>")
        return 
    }

    switch args[1] {
        case "add":
            if len(args) < 3 {
                fmt.Println("usage: task-cli add <description>")
                return
            }
            addTasks(args[2])
        case "list":
            if len(args) < 2 {
                fmt.Println("usage: task-cli list")
                return        
            }
            listTasks()
    }
}