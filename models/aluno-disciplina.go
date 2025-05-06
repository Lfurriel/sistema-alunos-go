package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AlunoDisciplina struct {
	Id           string    `json:"id" gorm:"primaryKey;column:id"`
	AlunoId      string    `json:"alunoId" gorm:"not null;column:aluno_id"`
	DisciplinaId string    `json:"disciplinaId" gorm:"not null;column:disciplina_id"`
	Presenca     bool      `json:"presenca" gorm:"column:presenca;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Referências
	Aluno      Aluno      `json:"-" gorm:"foreignKey:AlunoId;references:Id"`
	Disciplina Disciplina `json:"-" gorm:"foreignKey:DisciplinaId;references:Id"`
}

func (AlunoDisciplina) TableName() string {
	return "aluno_disciplina"
}

func (ad *AlunoDisciplina) BeforeCreate(tx *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	ad.Id = uuidStr
	return
}
