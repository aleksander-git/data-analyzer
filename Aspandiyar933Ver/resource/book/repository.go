package book

import "database/sql"

type Repository struct {
	db *sql.DB // how I even understnd we will use Yandex db but it is just test code 
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

