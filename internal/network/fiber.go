package network

import (
	"context"
	"encoding/json"
	"log/slog"
	"strconv"
	"test2025c1/internal/model"
	"test2025c1/internal/service"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	fiber *fiber.App
	service *service.Service
	uri string
}

func NewServer(uri_server, uri_database string) (*Server, error) {
	server := &Server{}
	server.uri = uri_server
	server.fiber = fiber.New()
	service, err := service.NewService(uri_database)
	if err != nil {
		slog.Error("error creating service")
		return nil, err
	}
	server.service = service
	server.fiber.Post("tasks", func(c fiber.Ctx) error {
		var task model.CreateTask
		json.Unmarshal(c.Body(), &task)
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		err := server.service.CreateTask(ctx, task)
		if err != nil {
			c.Status(500)
			return err
		}
		return nil
	})
	server.fiber.Get("tasks", func(c fiber.Ctx) error {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		tasks, err := server.service.GetTasks(ctx)
		if err != nil {
			c.Status(500)
			return err
		}
		b, err := json.Marshal(tasks)
		if err != nil {
			c.Status(500)
			return err
		}
		c.Write(b)
		return nil
	})
	server.fiber.Put("tasks/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id")) 
		if err != nil {
			c.Status(400)
			return err
		}
		var task model.Task
		err = json.Unmarshal(c.Body(), &task)
		if err != nil {
			c.Status(400)
			return err
		}
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

		server.service.UpdateTask(ctx, id, task)
		return nil
	})
	server.fiber.Delete("tasks/:id", func(c fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id")) 
		if err != nil {
			c.Status(400)
			return err
		}
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		err = server.service.DeleteTask(ctx, id)
		if err != nil {
			c.Status(500)
			return err
		}
		return nil
	})
	return server, nil
}
func (server *Server) Run() {
	slog.Error("ошибка сервера", slog.Any("error = ",(server.fiber.Listen(server.uri))))
}