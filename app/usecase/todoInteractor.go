package usecase

import (
	"github.com/moriwakihikari/clean_architecture_with_todo.git/domain/model"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(t model.Todo) (id int, err error) {
	id, err = interactor.TodoRepository.Store(t)
	return
}

func (interactor *TodoInteractor) Update(t model.Todo) (id int, err error) {
	id, err = interactor.TodoRepository.Update(t)
	return
}

func (interactor *TodoInteractor) Delete(id int) (err error) {
	err = interactor.TodoRepository.Delete(id)
	return
}

func (interactor *TodoInteractor) Todos() (todos model.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoByUserId(id int) (todoList model.Todos, err error) {
	todoList, err = interactor.TodoRepository.FindByUserId(id)
	return
}

func (interactor *TodoInteractor) TodoById(id int) (todo model.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(id)
	return
}
