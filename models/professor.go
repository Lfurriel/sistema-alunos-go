package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Professor struct {
	Id             string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36)"`
	Nome           string    `json:"nome" gorm:"not null;column:nome" binding:"required,min=1,max=60"`
	Email          string    `json:"email" gorm:"not null;column:email;uniqueIndex:idx_unique_email_prof" binding:"required,email"`
	Senha          string    `json:"senha,omitempty" gorm:"not null;column:senha" binding:"required,senha_forte"`
	ConfirmarSenha string    `json:"confirmar_senha,omitempty" gorm:"-" binding:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	Disciplinas []Disciplina `json:"disciplinas" gorm:"foreignKey:ProfessorId;constraint:OnDelete:CASCADE"`
}

func (Professor) TableName() string {
	return "professores"
}

func (p *Professor) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	p.Id = uuidStr
	return
}

type Login struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}
