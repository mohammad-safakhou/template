package repository

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"template/domain/backend/adapters/repository/models"
	"template/domain/backend/core/ports"
)

type helloRepository struct {
	db *sql.DB
}

func NewHelloRepository(db *sql.DB) ports.HelloRepository {
	return helloRepository{db: db}
}

func (hp helloRepository) Get() string {
	r, _ := models.Hellos().One(context.TODO(), hp.db)
	return r.Value.String
}

func (hp helloRepository) Save(input string) {
	var helloModel models.Hello
	helloModel.Value = null.NewString(input, true)
	helloModel.Insert(context.TODO(), hp.db, boil.Infer())
}
