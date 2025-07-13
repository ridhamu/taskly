package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ridhamu/taskly/internal"
)

const filename = "tasks.json"

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: task-cli [add|update|delete|list|mark-done|mark-in-progress]")
	}

	command := args[1]
	tasks, _ := internal.LoadTasks(filename)

	// iterate possible command
	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			return
		}
		desc := args[2]
		id := len(tasks) + 1
		task := internal.Task{
			Id:          id,
			Description: desc,
			Status:      internal.Todo,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		tasks = append(tasks, task)
		internal.SaveTasks(filename, tasks)
		fmt.Println("success adding task, taskId: ", task.Id)
	case "update":
		if len(args) < 4 {
			fmt.Println("Usage: task-cli update <id> <new description>")
			return
		}
		id, _ := strconv.Atoi(args[2])
		desc := args[3]
		updated := false
		for i := range tasks {
			if id == tasks[i].Id {
				tasks[i].Description = desc
				tasks[i].UpdatedAt = time.Now()
				updated = true
				break
			}
		}
		if updated {
			internal.SaveTasks(filename, tasks)
			fmt.Printf("Task %d updated!\n", id)
		} else {
			fmt.Printf("Task with id %d not found!\n", id)
		}
	case "list":
		// tasks, err := internal.LoadTasks(filename)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
		// fmt.Println(tasks)
		if len(tasks) > 0 {
			for _, v := range tasks {
				printTask(&v)
			}
		} else {
			fmt.Println("list are empty")
		}
	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}

		id, _ := strconv.Atoi(args[2])
		newListTask := []internal.Task{}
		deleted := false
		for _, task := range tasks {
			if task.Id != id {
				newListTask = append(newListTask, task)
			} else {
				deleted = true
			}
		}
		if deleted {
			internal.SaveTasks(filename, newListTask)
			fmt.Printf("Task %d successfully deleted!\n", id)
		} else {
			fmt.Printf("Unable to delete task %d\n", id)
		}
	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(args[2])

		if tasks[id].Status == internal.Done {
			fmt.Println("status already done!")
			return
		} else {
			for i := range tasks {
				if tasks[i].Id == id {
					tasks[i].Status = internal.Done
					tasks[i].UpdatedAt = time.Now()
					internal.SaveTasks(filename, tasks)
					printTask(&tasks[i])
					return
				}
			}
		}
		fmt.Printf("Task with id %d not found!\n", id)

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(args[2])

		if tasks[id].Status == internal.InProgress {
			fmt.Println("status already in progress!")
			return
		} else {
			for i := range tasks {
				if tasks[i].Id == id {
					tasks[i].Status = internal.InProgress
					tasks[i].UpdatedAt = time.Now()
					internal.SaveTasks(filename, tasks)
					printTask(&tasks[i])
					return
				}
			}
		}
		fmt.Printf("Task with id %d not found!\n", id)
	default:
		fmt.Println("Unknown command", command)
	}
}

func printTask(task *internal.Task) {
	fmt.Printf("[%d] %s | %s | CreatedAt: %s | UpdatedAt: %s\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822))
}
