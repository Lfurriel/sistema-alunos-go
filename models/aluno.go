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
	Ativo     bool      `json:"ativo,omitempty" gorm:"column:ativo;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamento
	AlunoDisciplina []AlunoDisciplina `json:"aluno_disciplina,omitempty" gorm:"foreignKey:AlunoId;constraint:OnDelete:CASCADE"`
	AlunoAvaliacao  []AlunoAvaliacao  `json:"aluno_avaliacao,omitempty" gorm:"foreignKey:AlunoId;constraint:OnDelete:CASCADE"`
	AlunoAula       []AlunoAula       `json:"aluno_aula,omitempty" gorm:"foreignKey:AlunoId;constraint:OnDelete:CASCADE"`
}

func (Aluno) TableName() string {
	return "alunos"
}

func (a *Aluno) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
