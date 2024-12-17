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

func (s *Storage) Tasks(taskID int) ([]Task, error) {
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
	($1 = 0 OR id = $1)
	ORDER BY id;
`,
		taskID,
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

func (s *Storage) TasksByAuthor(authorID int) ([]Task, error) {
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
	author_id = $1
	ORDER BY id;
`,
		authorID,
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

func (s *Storage) TasksByLabel(labelID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT 
	tasks.id,
	opened,
	closed,
	author_id,
	assigned_id,
	title,
	content
	FROM tasks
	JOIN tasks_labels ON tasks_labels.task_id = tasks.id
	JOIN labels ON tasks_labels.label_id = labels.id
	WHERE labels.id = $1
	ORDER BY id;
`,
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

func (s *Storage) TasksUpdateByID(taskID int, t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
	UPDATE tasks

	SET 
	opened = $1,
	closed = $2,
	author_id = $3,
	assigned_id = $4,
	title = $5,
	content = $6
	
	WHERE id = $7

	RETURNING id;
`,
		t.Opened,
		t.Closed,
		t.AuthorID,
		t.AssignedID,
		t.Title,
		t.Content,
		taskID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err

}

func (s *Storage) TaskDelete(taskID int) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
	DELETE 
	FROM
	tasks
	
	WHERE id = $1

	RETURNING id;
`,
		taskID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
