package service

import (
	"errors"
	"todo/model"
	"todo/repository"
)

type TodoService struct {
	RepoTodo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{RepoTodo: repo}
}

func (is *TodoService) GetTodosService(limit, page int, searchThread string) ([]model.Todo, int, error) {
	offset := (page - 1) * limit

	todos, totalTodos, err := is.RepoTodo.GetTodosWithPagination(limit, offset, searchThread)
	if err != nil {
		return nil, 0, errors.New("failed to retrieve todos")
	}

	return todos, totalTodos, nil
}

func (is *TodoService) AddTodoService(userId int, thread, status string) error {
	if thread == "" {
		return errors.New("todo name cannot be empty")
	}
	if status == "" {
		return errors.New("todo code cannot be empty")
	}

	todo := model.Todo{
		UserID: userId,
		Thread: thread,
		Status: status,
	}
	_, err := is.RepoTodo.Create(&todo)
	if err != nil {
		return errors.New("failed to create todo: " + err.Error())
	}

	// fmt.Println("Successfully added todo with ID:", todo.ID)
	return nil
}

func (is *TodoService) UpdateStatusService(id uint16, status string) error {

	if id == 0 {
		return errors.New("todo code cannot be empty")
	}
	if status < "" {
		return errors.New("stock cannot be negative")
	}

	exists, err := is.RepoTodo.TodoExists(id)
	if err != nil {
		return errors.New("failed to check todo existence: " + err.Error())
	}
	if !exists {
		return errors.New("todo not found")
	}

	err = is.RepoTodo.UpdateStatus(status, id)
	if err != nil {
		return errors.New("failed to update stock: " + err.Error())
	}

	return nil
}
