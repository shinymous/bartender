package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, satatusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(satatusCode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

func Error(w http.ResponseWriter, satatusCode int, erro error) {
	JSON(w, satatusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
