package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
)

type Task struct {
	Id          int    `json:"id"`
	User_Id     int    `json:"user_Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type Tasks struct {
	Tasks []*Task `json:"Tasks"`
}

func (t *Task) Create(task *Task) error {
	exists := checkIfUserExists(task.User_Id)
	if exists == false {
		return errors.New("user with given ID not found")
	}
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening tasks file: ", err)
		return err
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	task.Id = uuid.New().ClockSequence()

	if err != nil {
		return err
	}

	err = file.Truncate(0)
	file.Seek(0, 0)

	if err != nil {
		fmt.Println("Error during file trunc:", err)
	}
	var TasksData Tasks
	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, &TasksData)
		if err != nil {
			fmt.Println("Error unmarshaling JSON: ", err)
		}
	}

	var appTasks []*Task
	for _, value := range TasksData.Tasks {
		appTasks = append(appTasks, &Task{Id: value.Id, User_Id: value.User_Id, Title: value.Title, Description: value.Description})
	}
	appTasks = append(appTasks, &Task{Id: task.Id, User_Id: task.User_Id, Title: task.Title, Description: task.Description})
	jsonData, err := json.MarshalIndent(Tasks{appTasks}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *Task) Read(user_Id int) (interface{}, error) {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return nil, err
	}
	userTasks := findUserTasks(tasks.Tasks, user_Id)
	if userTasks != nil {
		for i := 0; i < len(userTasks); i++ {
			fmt.Println("Task found: ")
			fmt.Println("Task ID: ", userTasks[i].Id)
			fmt.Println("Task user ID: ", userTasks[i].User_Id)
			fmt.Println("Task title: ", userTasks[i].Title)
			fmt.Println("Task description: ", userTasks[i].Description)
			fmt.Println("---------------------------------------")
		}
		return userTasks, nil
	} else {
		return nil, errors.New("Tasks for given user not found\n")
	}

}

func (t *Task) ReadAll(display bool) (interface{}, error) {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var tasks Tasks
	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error during unmarshalling JSON: ", err)
	}
	if display {
		for i := 0; i < len(tasks.Tasks); i++ {
			task := tasks.Tasks[i]
			fmt.Println("--------------------")
			fmt.Println("Task ID: ", task.Id)
			fmt.Println("User assigned: ", task.User_Id)
			fmt.Println("Task title: ", task.Title)
			fmt.Println("Task description: ", task.Description)
		}
	}
	return tasks, err
}
func (t *Task) Update(id int, data *Task) error {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return err
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}

	task := findTask(tasks.Tasks, id)
	if task == nil {
		fmt.Println("Task not found")
		return nil
	}
	if data.User_Id != 0 {
		exists := checkIfUserExists(data.User_Id)
		if exists {
			task.User_Id = data.User_Id
		} else {
			errors.New("user with given ID not found")
		}
	}

	if data.Title != "" {
		task.Title = data.Title
	}

	if data.Description != "" {
		task.Description = data.Description
	}

	jsonData, err := json.MarshalIndent(Tasks{tasks.Tasks}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}

	err = file.Truncate(0)
	file.Seek(0, 0)

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}
func (t *Task) Delete(id int) error {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error reading file: ", err)
	}
	byteValue, _ := io.ReadAll(file)
	var tasks Tasks

	err = json.Unmarshal(byteValue, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return err
	}

	task := findTask(tasks.Tasks, id)
	if task == nil {
		fmt.Println("User not found")
		return nil
	}
	var updatedTasks []*Task
	for _, task := range tasks.Tasks {
		if task.Id == id {
			continue
		} else {
			updatedTasks = append(updatedTasks, &Task{Id: task.Id, User_Id: task.User_Id, Title: task.Title, Description: task.Description})
		}
	}

	jsonData, err := json.MarshalIndent(Tasks{updatedTasks}, "", "\t")

	if err != nil {
		fmt.Println("Error serializing object: ", err)
		return err
	}

	_, err = file.Write(jsonData)

	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return err
	}
	return nil
}

func findTask(tasks []*Task, id int) *Task {
	for _, task := range tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func findUserTasks(allTasks []*Task, id int) []*Task {
	var tasks []*Task
	for i := 0; i < len(allTasks); i++ {
		task := allTasks[i]
		if task.User_Id == id {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
func checkIfTaskExists(id int) bool {
	task := &Task{}
	var tasks, _ = task.ReadAll(false)
	for i := 0; i < len(tasks.(Tasks).Tasks); i++ {
		if tasks.(Tasks).Tasks[i].Id == id {
			return true
		}
	}
	return false
}
