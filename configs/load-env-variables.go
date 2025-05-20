package configs

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv carrega variáveis de ambiente de um arquivo .env usando o pacote godotenv e registra uma mensagem se o arquivo não for encontrado.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Não encontrado o arquivo .env")
	}
}
