package rotas

import (
	"api/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             "POST",
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
