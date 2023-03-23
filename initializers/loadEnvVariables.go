package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Função que inicia antes de tudo, a qual via carregar o arquivo .env com as configurações
// Caso erro -> Retorna aviso!
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Erro no arquivo .env")
	}
}
