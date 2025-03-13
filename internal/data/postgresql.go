package data

import (
	"context"
	"fmt"
	"test2025c1/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

const q1 = "insert into tasks(title,  description) values ($1, $2);"
const q2 = "update tasks set title=$2, description=$3, status=$4, updated_at=now() where id=$1;"
const q3 = "delete from tasks where id=$1;"
const q4 = "select id, title, description, status, created_at, updated_at from tasks;"
type Postgresql struct {
	pool *pgxpool.Pool
}

func NewBase(uri string) (*Postgresql, error) {
	pool, err := pgxpool.New(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return &Postgresql{pool: pool}, nil
}



func (postgres *Postgresql) CreateTask(ctx context.Context, task model.CreateTask) (error) {
	row, err := postgres.pool.Query(ctx, q1, task.Title, task.Description)
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (postgres *Postgresql) GetTasks(ctx context.Context) (ans []model.Task, err error) {
	rows, e := postgres.pool.Query(ctx, q4)
	if e != nil {
		err = fmt.Errorf("ошибка запроса select \n %w", e)
		return
	}

	for rows.Next() {
		var task model.Task
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		ans = append(ans, task)
	}
	fmt.Println(ans)
	fmt.Println(err)
	rows.Close()
	return
}

func (postgres *Postgresql) UpdateTask(ctx context.Context, id int, task model.Task) (error) {
	tx, err := postgres.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	rows, err := tx.Query(ctx, q2, id, task.Title, task.Description, task.Status)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	rows.Close()
	tx.Commit(ctx)
	return nil
}

func (postgres *Postgresql) DeleteTask(ctx context.Context, id int) (error) {
	tx, err := postgres.pool.Begin(context.Background())
	if err != nil {
		return err
	}
	rows, err := tx.Query(ctx, q3, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	rows.Close()
	tx.Commit(ctx)
	return nil
}