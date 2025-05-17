package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AlunoDisciplina struct {
	Id           string    `json:"id" gorm:"primaryKey;column:id"`
	AlunoId      string    `json:"aluno_id" gorm:"not null;column:aluno_id;index:idx_aluno_disciplina_id"`
	DisciplinaId string    `json:"disciplina_id" gorm:"not null;column:disciplina_id;index:idx_aluno_disciplina_id"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Referências
	Aluno      Aluno      `json:"-" gorm:"foreignKey:AlunoId;references:Id"`
	Disciplina Disciplina `json:"-" gorm:"foreignKey:DisciplinaId;references:Id"`
}

func (AlunoDisciplina) TableName() string {
	return "aluno_disciplina"
}

func (ad *AlunoDisciplina) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	ad.Id = uuidStr
	return
}
