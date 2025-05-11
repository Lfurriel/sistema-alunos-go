package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Avaliacao struct {
	Id            string    `json:"id" gorm:"primaryKey;column:id"`
	DisciplinaId  string    `json:"disciplina_id" gorm:"not null;column:disciplina_id"` // FK
	Nome          string    `json:"nome" gorm:"not null;column:nome" binding:"required,min=1,max=60"`
	Tipo          string    `json:"tipo" gorm:"not null;column:tipo" binding:"required,oneof=P T"`
	DataAvaliacao string    `json:"data_avaliacao" gorm:"not null;column:data_avaliacao" binding:"required,data_valida"`
	Peso          float64   `json:"peso" gorm:"not null;column:peso" binding:"required,gte=0,lte=1"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	Disciplina      Disciplina       `gorm:"foreignKey:DisciplinaId"`
	AlunoAvaliacoes []AlunoAvaliacao `gorm:"foreignKey:AvaliacaoId"`
}

func (Avaliacao) TableName() string {
	return "avaliacoes"
}

func (a Avaliacao) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
