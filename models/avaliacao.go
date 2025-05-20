package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Avaliacao representa uma prova ou trabalho relacionado a uma disciplina
//
// Cada avaliação possui um tipo ("P" para prova, "T" para trabalho), uma data, peso e pertence a uma disciplina específica
type Avaliacao struct {
	Id            string    `json:"id" gorm:"primaryKey;column:id"`
	DisciplinaId  string    `json:"disciplina_id" gorm:"not null;column:disciplina_id;index:idx_avaliacao"` // FK
	Nome          string    `json:"nome" gorm:"not null;column:nome" binding:"required,min=1,max=60"`
	Tipo          string    `json:"tipo" gorm:"not null;column:tipo" binding:"required,oneof=P T"`
	DataAvaliacao string    `json:"data_avaliacao" gorm:"not null;column:data_avaliacao" binding:"required,data_valida"`
	Peso          float64   `json:"peso" gorm:"not null;column:peso" binding:"required,gte=0,lte=1"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	Disciplina      *Disciplina      `json:"disciplina,omitempty" gorm:"foreignKey:DisciplinaId"`
	AlunoAvaliacoes []AlunoAvaliacao `json:"aluno_avaliacoes,omitempty" gorm:"foreignKey:AvaliacaoId;constraint:OnDelete:CASCADE"`
}

// TableName especifica o nome da tabela do banco de dados para a estrutura Avaliacao
func (Avaliacao) TableName() string {
	return "avaliacoes"
}

// BeforeCreate é usado para o GORM que gera e atribui uma nova string UUID ao campo Id antes de uma Avaliacao ser criada
func (a *Avaliacao) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
