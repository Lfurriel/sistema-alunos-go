package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AlunoMedia struct {
	Id           string    `json:"id" gorm:"primaryKey;column:id"`
	AlunoId      string    `json:"aluno_id" gorm:"not null;column:aluno_id;index:idx_aluno_media_id"`
	DisciplinaId string    `json:"disciplina_id" gorm:"not null;column:disciplina_id;index:idx_aluno_media_id"`
	MediaFinal   float64   `json:"media_final" gorm:"column:media_final"`
	Frequencia   float64   `json:"frequencia" gorm:"column:frequencia"`
	Aprovado     bool      `json:"aprovado" gorm:"column:aprovado"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Referências
	Aluno      Aluno      `json:"-" gorm:"foreignKey:AlunoId;references:Id"`
	Disciplina Disciplina `json:"-" gorm:"foreignKey:DisciplinaId;references:Id"`
}

func (AlunoMedia) TableName() string {
	return "aluno_media"
}

func (am AlunoMedia) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	am.Id = uuidStr
	return
}
