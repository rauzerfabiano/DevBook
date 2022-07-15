package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a respostas de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON
func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(body); erro != nil {
			log.Fatal(erro)
		}
	}
}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
