package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AlunoAula representa a presença ou ausência de um aluno em uma aula específica
//
// É usado para calcular a frequência do aluno na disciplina
type AlunoAula struct {
	Id        string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36)"`
	AulaId    string    `json:"aula_id" gorm:"not null;column:aula_id"`
	AlunoId   string    `json:"aluno_id" gorm:"not null;column:aluno_id" binding:"required"`
	Presenca  bool      `json:"presenca" gorm:"column:presenca;not null" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	Aula  Aula  `json:"-" gorm:"foreignKey:AulaId"`
	Aluno Aluno `json:"aluno" gorm:"foreignKey:AlunoId"`
}

// TableName especifica o nome da tabela do banco de dados para a estrutura AlunoAula
func (AlunoAula) TableName() string {
	return "aluno_aula"
}

// BeforeCreate é usado para o GORM que gera e atribui uma nova string UUID ao campo Id antes de um AlunoAula ser criado
func (aa *AlunoAula) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	aa.Id = uuidStr
	return
}
