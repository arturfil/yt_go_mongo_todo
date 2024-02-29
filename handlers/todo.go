package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-mongo-todos/services"
)

var todo services.Todo

// healthCheck - simple function to test api if its working
func healthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
    todos, err := todo.GetAllTodos()
    if err != nil {
        log.Println(err)
        return 
    }


	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(todos)	
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    todo, err := todo.GetTodoById(id)
    if err != nil {
        log.Println(err)
        return 
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(todo)	
}

func createTodo(w http.ResponseWriter, r *http.Request) {

    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        log.Fatal(err)
    }

    err = todo.InsertTodo(todo)
    if err != nil {
        errorRes := Response {
            Msg: "Error",
            Code: 304,
        }
        json.NewEncoder(w).Encode(errorRes)
        return 
    }

    res := Response {
        Msg: "Succesfully Created Todo",
        Code: 201,
    }

    jsonStr, err := json.Marshal(res)
    if err != nil {
        log.Fatal(err)
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        log.Println(err)
        return 
    }

    _, err = todo.UpdateTodo(id, todo)
    if err != nil {
        errorRes := Response {
            Msg: err.Error(),
            Code: 500,
        }
        jsonStr, err := json.Marshal(errorRes)
        if err != nil {
            log.Fatal(err)
        }   

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(errorRes.Code)
        w.Write(jsonStr)
        return
    }

    res := Response {
        Msg: "Succesfully updated",
        Code: 200,
    }

    jsonStr, err := json.Marshal(res)
    if err != nil {
        log.Fatal(err)
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    err := todo.DeleteTodo(id)
    if err != nil {
        errorRes := Response {
            Msg: "Error",
            Code: 304,
        }
        json.NewEncoder(w).Encode(errorRes)
        w.WriteHeader(errorRes.Code)
        return 
    }

    res := Response {
        Msg: "Succesfully deleted",
        Code: 200,
    }

    jsonStr, err := json.Marshal(res)
    if err != nil {
        log.Fatal(err)
    }

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

