package handlers

import (
    "net/http"
    "html/template"
    "todo-app/models"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    tasks := models.GetTasks()
    tmpl.Execute(w, tasks)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
    task := r.FormValue("task")
    models.AddTask(task)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
