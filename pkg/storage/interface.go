package storage

import "sftasks/v2/pkg/storage/postgres"

type Interface interface {
	Tasks(int, int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
}
