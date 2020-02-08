package beer_test

import (
	"database/sql"
	"testing"

	"github.com/eminetto/pos-web-go/core/beer"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	defer db.Close()
	service := beer.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %d", 1, saved.ID)
	}
}

//@todo implementar testes para as outras funções
