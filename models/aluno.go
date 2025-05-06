package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Aluno struct {
	Id        string    `json:"id,omitempty" gorm:"primaryKey;column:id;type:varchar(36);not null"`
	Nome      string    `json:"nome,omitempty" gorm:"type:varchar(60);column:nome;not null" binding:"required,min=1,max=60"`
	Email     string    `json:"email,omitempty" gorm:"type:text;column:email;not null" binding:"required,email"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	AlunoDisciplinas []AlunoDisciplina `gorm:"foreignKey:AlunoId;constraint:OnDelete:CASCADE"`
	AlunoAvaliacoes  []AlunoAvaliacao  `gorm:"foreignKey:AlunoId"`
	AlunoAula        []AlunoAula       `json:"aluno_aula" gorm:"foreignKey:AlunoId"`
}

func (Aluno) TableName() string {
	return "alunos"
}

func (a *Aluno) BeforeCreate(tx *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
