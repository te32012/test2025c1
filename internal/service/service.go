package service

import (
	"context"
	"errors"
	"test2025c1/internal/data"
	"test2025c1/internal/model"
)

type Service struct {
	base *data.Postgresql
}

func NewService(uri_database string) (*Service, error) {
	service := Service{}
	base, err := data.NewBase(uri_database)
	if err != nil {
		return nil, err
	}
	service.base = base
	return &service, nil
}

func (service *Service) CreateTask(ctx context.Context, task model.CreateTask) (error) {
	return service.base.CreateTask(ctx, task)
}

func (service *Service) GetTasks(ctx context.Context) ([]model.Task, error) {
	return service.base.GetTasks(ctx)
}

func (service *Service) UpdateTask(ctx context.Context, id int, task model.Task) (error) {
	if task.Status != "new" && task.Status != "in_progress" && task.Status != "done" {
		return errors.New("status not allowed")
	}
	if id < 0 {
		return errors.New("id not allowed")
	}
	err := service.base.UpdateTask(ctx, id, task)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) DeleteTask(ctx context.Context, id int) (error) {
	if id < 0 {
		return errors.New("id not allowed")
	}
	err := service.base.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}