package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"
	"github.com/moriwakihikari/clean_architecture_with_todo.git/interfaces/database"
	"github.com/moriwakihikari/clean_architecture_with_todo.git/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(handler database.Sqlhandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				Sqlhandler: handler,
			},
		},
	}
}

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todoReceptor struct {
		Title   string `json:"title"`
		Note    string `json:"note"`
		DayTime string `json:"day_time"`
		UserID  int    `json:"user_id"`
	}
	if err := json.Unmarshal(b, &todoReceptor); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	formattedTime := parseTime(todoReceptor.DayTime)
	todo := model.Todo{
		Title:   todoReceptor.Title,
		Note:    todoReceptor.Note,
		DayTime: formattedTime,
		UserID:  todoReceptor.UserID,
	}
	id, err := controller.Interactor.Add(todo)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todoReceptor struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Note    string `json:"note"`
		DayTime string `json:"day_time"`
		UserID  int    `json:"user_id"`
	}
	if err := json.Unmarshal(b, &todoReceptor); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	formattedTime := parseTime(todoReceptor.DayTime)
	todo := model.Todo{
		ID:      todoReceptor.ID,
		Title:   todoReceptor.Title,
		Note:    todoReceptor.Note,
		DayTime: formattedTime,
		UserID:  todoReceptor.UserID,
	}
	id, err := controller.Interactor.Update(todo)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *TodoController) Index(w http.ResponseWriter, r *http.Request) {
	todoList, err := controller.Interactor.Todos()
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, todoList)
}

func (controller *TodoController) Show(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := controller.Interactor.TodoById(req.ID)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, todo)
}

func (controller *TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = controller.Interactor.Delete(req.ID)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, "Success!")
}

func (controller *TodoController) Search(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		UserID int `json:"user_id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	todoList, err := controller.Interactor.TodoByUserId(req.UserID)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, todoList)
}

func ResponseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func parseTime(date string) (formattedTime time.Time) {
	formattedTime, _ = time.Parse("20060102", date)
	return
}
