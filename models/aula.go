package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Aula struct {
	Id              string    `json:"id" gorm:"primaryKey;column:id"`
	DisciplinaId    string    `json:"disciplina_id" gorm:"not null;column:disciplina_id;index"` // FK
	Numero          int       `json:"numero" gorm:"not null;column:numero" binding:"required,gte=1"`
	Data            string    `json:"data" gorm:"not null;column:data" binding:"required,data_valida"`
	QuantidadeHoras int       `json:"quantidade_horas" gorm:"not null;column:quantidade_horas" binding:"required,gte=1"`
	Conteudo        string    `json:"conteudo" gorm:"not null;column:conteudo" binding:"required,min=1,max=1000"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	Disciplina Disciplina  `json:"disciplina" gorm:"foreignKey:DisciplinaId"`
	AlunoAula  []AlunoAula `json:"aluno_aula" gorm:"foreignKey:AulaId;constraint:OnDelete:CASCADE"`
}

func (Aula) TableName() string {
	return "aulas"
}

func (a *Aula) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
