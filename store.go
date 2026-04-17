package main

import (
	"os"
	"fmt"
	"time"
	"math"
	"encoding/json"
)

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile("file.json")
    if err != nil {
        return nil, err
    }
	var slice[] Task
	err = json.Unmarshal(data, &slice)
	if err != nil {
		return nil, err
	}
	return slice, err
}

func addTasks(task string) (error) {
	// loading tasks from the json file and saving them in the tasks slice.
	tasks, err := loadTasks()
    if err != nil {
        fmt.Println("error:", err)
        return err
    }
	max_value := math.MinInt
	//find the highest id value
	for _, task := range tasks {
		if task.Id > max_value {
			max_value = task.Id
		}
	}
	
	// creating a newTask struct slice, then adding the new task
	newTask := Task {
		Id: max_value + 1,
		Description: task,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//appending the new task to the loaded tasks list
	tasks = append(tasks, newTask)
	// convert the tasks, to a json format
	data, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		return err
	}
	// write the converted tasks list to the json file
	err = os.WriteFile("file.json", data, 06440)
	if err != nil {
			fmt.Println("error:", err)
			return err
		}
	return nil
}


func listTasks() (error) {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
        fmt.Println("No tasks found")
        return nil
    }
    for i, task := range tasks {
        fmt.Printf("%d) [%s] %s\n", i+1, task.Status, task.Description)
        fmt.Printf("   ID: %d\n", task.Id)
        fmt.Printf("   Created: %s\n", task.CreatedAt.Format("2006-01-02 15:04:05"))
        fmt.Printf("   Updated: %s\n", task.UpdatedAt.Format("2006-01-02 15:04:05"))
    }
	return nil
}

func delTask(id int) (error) {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found to delete")
		return nil
	}
	var newList[] Task 
	for _, task := range tasks {
		if task.Id != id {
			newList = append(newList , task)
		}
	}
	data, err := json.MarshalIndent(newList, "", "	")
	if err != nil {
		return err
	}
	// write the converted tasks list to the json file
	err = os.WriteFile("file.json", data, 06440)
	if err != nil {
			fmt.Println("error:", err)
			return err
		}
	return nil
}