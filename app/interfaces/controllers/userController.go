package controllers

// ここのレイヤではリクエストに対するルーティングを実装します。

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"
	"github.com/moriwakihikari/clean_architecture_with_todo.git/interfaces/database"
	"github.com/moriwakihikari/clean_architecture_with_todo.git/usecase"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

// NewUserControllerでdatabase.Sqlhandlerを引数に持つのでInteractorを返して
// Usecaseレイヤとinfrastructure/databaseを紐づけられます。
// 次にルーティングを実装します。
func NewUserController(sqlHandler database.Sqlhandler) *UserController {
	fmt.Println(sqlHandler)
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				Sqlhandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var user model.User
	if err := json.Unmarshal(b, &user); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := controller.Interactor.Users()
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, users)
}

func (controller *UserController) Show(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/user/get/")
    _, string_id := filepath.Split(sub)
    if string_id != "" {
		var id int
		id, _ = strconv.Atoi(string_id)
		user, err := controller.Interactor.UserById(id)
		if err != nil {
			ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}
		ResponseOk(w, user)
	}
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var user model.User
	if err := json.Unmarshal(b, &user); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := controller.Interactor.Update(user)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	ResponseOk(w, id)
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
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