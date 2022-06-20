package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.CarregarVariaveisDeAmbiente()

	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
