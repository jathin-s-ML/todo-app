package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct{
	Id int
	Name string
	Completed bool
}

var tasks[] Task
var nextid int =1

func addtask(name string){
	
	tasks=append(tasks, Task{Id: nextid,Name: name,Completed: false})
	fmt.Println("task added :",name)
	nextid++
}

func listoftasks()  {
	if len(tasks)==0{
		fmt.Println("there are no tasks to display")
		return
	}else{
		for _,i:=range tasks{
			fmt.Printf("ID: %d | Task: %s | Status: %t\n", i.Id, strings.TrimSpace(i.Name), i.Completed)
		}
	}
}

func deletetask(taskid int){
	for i:=range tasks{
		if tasks[i].Id==taskid{
			fmt.Println("removing task :",tasks[i].Name)
			tasks=append(tasks[:i],tasks[i+1:]... )
			fmt.Println("task removed successfully")
			return
		}
	}
	fmt.Println("task not found")
	
}

func markascompleted(taskid int){
	for i :=range tasks{
		if tasks[i].Id==taskid{
			fmt.Println("updating the task as completed")
			tasks[i].Completed=true
			fmt.Println("task updated")
			return
		}
	}
	fmt.Println("task not found")
}

func main() {
	fmt.Println("welcome to TODO application")
	var task string
	var choice int
	reader :=bufio.NewReader(os.Stdin)

	
	for{
		fmt.Println()
		fmt.Println("choose an option \n1. add task\n2. list all tasks\n3. mark a task as completed\n4. delete a task\n5. exit todo application")
		fmt.Println("enter the choice")
		fmt.Scan(&choice)
		switch choice{
				case 1:{
					fmt.Println("enter the task to add")
					task,_=reader.ReadString('\n')
					addtask(task)
				}
				case 2:{
					listoftasks()
				}	
				case 3:{
					fmt.Println("enter the id of the task to mark as completed")
					var id int
					fmt.Scan(&id)
					markascompleted(id)
				}
				case 4:{
					fmt.Println("enter the id of the task to delete ")
					var id int
					fmt.Scan(&id)
					deletetask(id)
				}
				case 5:{
					fmt.Println("exit the todo application")
					os.Exit(0)
				}
				default  :{
					fmt.Println("invalid choice")
				}
		}
	}
}