package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eminetto/pos-web-go/core/beer"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_getAllBeer(t *testing.T) {
	b1 := &beer.Beer{
		ID:    10,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	b2 := &beer.Beer{
		ID:    20,
		Name:  "Skol",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	assert.Nil(t, err)
	assert.Nil(t, clearDB(db))
	service := beer.NewService(db)
	assert.Nil(t, service.Store(b1))
	assert.Nil(t, service.Store(b2))

	handler := getAllBeer(service)
	r := mux.NewRouter()
	r.Handle("/v1/beer", handler)
	req, err := http.NewRequest("GET", "/v1/beer", nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result []*beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, b1.ID, result[0].ID)
	assert.Equal(t, b2.ID, result[1].ID)
}

type BeerServiceMock struct{}

func (t BeerServiceMock) GetAll() ([]*beer.Beer, error) {
	b1 := &beer.Beer{
		ID:    10,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	b2 := &beer.Beer{
		ID:    20,
		Name:  "Skol",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	return []*beer.Beer{b1, b2}, nil
}

func (t BeerServiceMock) Get(ID int64) (*beer.Beer, error) {
	b1 := &beer.Beer{
		ID:    ID,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}
	return b1, nil
}
func (t BeerServiceMock) Store(b *beer.Beer) error {
	return nil
}

func (t BeerServiceMock) Update(b *beer.Beer) error {
	return nil
}

func (t BeerServiceMock) Remove(ID int64) error {
	return nil
}

func Test_getAllBeerWithMock(t *testing.T) {
	service := &BeerServiceMock{}
	handler := getAllBeer(service)
	r := mux.NewRouter()
	r.Handle("/v1/beer", handler)

	req, err := http.NewRequest("GET", "/v1/beer", nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result []*beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, int64(10), result[0].ID)
	assert.Equal(t, int64(20), result[1].ID)
}

func Test_getBeerWithMock(t *testing.T) {
	service := &BeerServiceMock{}
	handler := getBeer(service)
	r := mux.NewRouter()
	r.Handle("/v1/beer/{id}", handler)

	req, err := http.NewRequest("GET", "/v1/beer/10", nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var result *beer.Beer
	err = json.NewDecoder(rr.Body).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), result.ID)
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer")
	tx.Commit()
	return err
}
