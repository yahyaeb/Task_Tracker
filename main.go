// main.go
package main

import (
    "fmt"
    "os"
    "strconv"
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
            if err := addTasks(args[2]); err != nil {
                fmt.Println("error:", err)
            }
        case "list":
            if err := listTasks(); err != nil {
                fmt.Println("error:", err)
            }
        case "delete":
            if len(args) < 3 {
                fmt.Println("usage: task-cli delete <Id>")
                return
            }
            Id, err := strconv.Atoi(args[2])
            if err != nil {
                fmt.Println("Error", err)
                return
            }
            delTask(Id)
        default:
            fmt.Println("unknown command:", args[1])
    }
}