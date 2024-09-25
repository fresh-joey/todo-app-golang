package models

var tasks []string

func GetTasks() []string {
    return tasks
}

func AddTask(task string) {
    tasks = append(tasks, task)
}
