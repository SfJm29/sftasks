package postgres

//"github.com/jackc/pgx/v4/pgxpool"
import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(connstr string) (*Storage, error) {

	db, err := pgxpool.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}
	return &s, nil
}

type User struct {
	ID   int
	Name string
}

type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

type Label struct {
	ID   int
	Name string
}

func (s *Storage) Tasks(taskID, authorID, labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
SELECT 
id,
opened,
closed,
author_id,
assigned_id,
title,
content
FROM tasks
WHERE
($1 = 0 OR id = $1) AND
($2 = 0 OR author_id = $2) AND
($3 = 0 OR label_id = $3)
ORDER BY id;
`,
		taskID,
		authorID,
		labelID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task

	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
	INSERT INTO tasks (title, content)
	VALUES ($1, $2) RETURNING id;
	`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}
