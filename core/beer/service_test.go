package beer_test

import (
	"database/sql"
	"testing"

	"github.com/eminetto/pos-web-go/core/beer"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func newBeer(id int64) *beer.Beer {
	return &beer.Beer{
		ID:    id,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
}

func getDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	// }
	return db
}

func clearAndClose(db *sql.DB, t *testing.T) {
	tx, err := db.Begin()
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Fatalf("Errro criando transação %s", err.Error())
	// }
	_, err = tx.Exec("delete from beer")
	assert.Nil(t, err)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	db.Close()
}

func TestStore(t *testing.T) {
	b := newBeer(1)
	db := getDB(t)
	defer clearAndClose(db, t)
	service := beer.NewService(db)
	err := service.Store(b)
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	// }
}

func TestGet(t *testing.T) {
	db := getDB(t)
	defer clearAndClose(db, t)
	service := beer.NewService(db)
	b := newBeer(1)
	_ = service.Store(b)
	saved, err := service.Get(1)
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Fatalf("Erro buscando %s", err.Error())
	// }
	assert.Equal(t, int64(1), saved.ID)
	// if saved.ID != 1 {
	// 	t.Fatalf("Erro de dados. Esperava %d recebi %d", 1, saved.ID)
	// }
}

func TestGetAll(t *testing.T) {
	db := getDB(t)
	defer clearAndClose(db, t)
	service := beer.NewService(db)
	b1 := newBeer(1)
	b2 := newBeer(2)
	_ = service.Store(b1)
	_ = service.Store(b2)
	saved, err := service.GetAll()
	assert.Nil(t, err)
	// if err != nil {
	// 	t.Fatalf("Erro buscando %s", err.Error())
	// }
	assert.Equal(t, 2, len(saved))
	// if len(saved) != 2 {
	// 	t.Fatalf("Erro de dados. Esperava %d recebi %d", 2, len(saved))
	// }
}

func TestUpdate(t *testing.T) {
	db := getDB(t)
	defer clearAndClose(db, t)
	service := beer.NewService(db)
	b := newBeer(1)
	_ = service.Store(b)
	t.Run("TestUpdate caminho feliz", func(t *testing.T) {
		saved, _ := service.Get(1)
		saved.Name = "Skol"
		err := service.Update(saved)
		if err != nil {
			t.Fatalf("Erro atualizando %s", err.Error())
		}
	})
	t.Run("TestUpdate erro de validação", func(t *testing.T) {
		e := newBeer(0)
		err := service.Update(e)
		if err == nil {
			t.Fatalf("Erro de validação")
		}
	})
}
