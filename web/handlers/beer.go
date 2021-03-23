package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/pos-web-go/core/beer"
	"github.com/gorilla/mux"
)

//a função recebe como terceiro parâmetro a interface
//ou seja, ela pode receber qualquer coisa que implemente a interface
//isso é muito útil para escrevermos testes, ou podermos substituir toda a
//implementação da regra de negócios
func MakeBeerHandlers(r *mux.Router, n *negroni.Negroni, service beer.UseCase) {
	r.Handle("/v1/beer", n.With(
		negroni.Wrap(getAllBeer(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(getBeer(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/beer", n.With(
		negroni.Wrap(storeBeer(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(updateBeer(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/v1/beer/{id}", n.With(
		negroni.Wrap(removeBeer(service)),
	)).Methods("DELETE", "OPTIONS")
}

/*
Para testar:
curl http://localhost:4000/v1/beer
*/
func getAllBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@TODO este código está duplicado em todos os handlers. Pergunta: como podemos melhorar isso?
		w.Header().Set("Content-Type", "application/json")
		all, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		//vamos converter o resultado em JSON e gerar a resposta
		err = json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro convertendo em JSON"))
			return
		}
	})
}

/*
Para testar:
curl http://localhost:4000/v1/beer/1
*/
func getBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@TODO este código está duplicado em todos os handlers. Pergunta: como podemos melhorar isso?
		w.Header().Set("Content-Type", "application/json")

		//vamos pegar o ID da URL
		//na definição do protocolo http, os parâmetros são enviados no formato de texto
		//por isso precisamos converter em int64
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		b, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError(err.Error()))
			return
		}
		//vamos converter o resultado em JSON e gerar a resposta
		err = json.NewEncoder(w).Encode(b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro convertendo em JSON"))
			return
		}
	})
}

/*
Para testar:
curl -v -X "POST" "http://localhost:4000/v1/beer" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "id": 7,
  "name": "Skol",
  "type": 1,
  "style":2
}'
*/
func storeBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@TODO este código está duplicado em todos os handlers. Pergunta: como podemos melhorar isso?
		w.Header().Set("Content-Type", "application/json")

		//vamos pegar os dados enviados pelo usuário via body
		var b beer.Beer
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		//@TODO precisamos validar os dados antes de salvar na base de dados. Pergunta: Como fazer isso?
		err = service.Store(&b)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

/*
Para testar:
curl -X "PUT" "http://localhost:4000/v1/beer/2" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "name": "Alterada",
  "type": 3,
  "style":1
}'
*/
func updateBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@TODO exercício: implementar este handler
	})
}

/*
Para testar:
curl -X "DELETE" "http://localhost:4000/v1/beer/2" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json'
*/
func removeBeer(service beer.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//@TODO exercício: implementar este handler
	})
}
