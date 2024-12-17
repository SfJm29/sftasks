package storage

import "sftasks/v2/pkg/storage/postgres"

type Interface interface {
	TasksByID(int) ([]postgres.Task, error)
	TasksByAuthor(int) ([]postgres.Task, error)
	TasksByLabel(int) ([]postgres.Task, error)
	TasksUpdateByID(int, postgres.Task) (int, error)
	TaskDelete(int) (int, error)
	NewTask(postgres.Task) (int, error)
}
