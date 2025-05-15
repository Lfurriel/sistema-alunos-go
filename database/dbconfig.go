package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sistema-alunos-go/models"
	"time"
)

var DB *gorm.DB

func ConectaBD() {
	var errConnection error

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslRequired := os.Getenv("DATABASE_SSL")

	ssl := "disable"
	if sslRequired == "true" {
		ssl = "require"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, name, port, ssl,
	)

	DB, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Configurações recomendadas
		SkipDefaultTransaction:                   true,  // performance: desativa transações implícitas
		PrepareStmt:                              true,  // performance: prepara e reutiliza statements
		DisableForeignKeyConstraintWhenMigrating: false, // mantém integridade
	})

	if errConnection != nil {
		log.Panic("Erro ao conectar com banco de dados:", errConnection)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Panic("Erro ao obter instância SQL do GORM:", err)
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	migrate()

	fmt.Println("Banco de dados conectado")
}

func migrate() {
	err := DB.AutoMigrate(
		&models.Professor{},
		&models.Aluno{},
		&models.Disciplina{},
		&models.Avaliacao{},
		&models.Aula{},
		&models.AlunoDisciplina{},
		&models.AlunoAvaliacao{},
		&models.AlunoAula{},
		&models.AlunoMedia{},
	)

	if err != nil {
		log.Fatalf("Erro ao realizar AutoMigrate: %v", err)
	}

	fmt.Println("Migrações aplicadas com sucesso.")
}
