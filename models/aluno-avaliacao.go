package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type AlunoAvaliacao struct {
	Id          string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AlunoId     string    `json:"aluno_id" gorm:"type:varchar(36);not null" binding:"required"`
	AvaliacaoId string    `json:"avaliacao_id" gorm:"type:varchar(36);not null" binding:"required"`
	Nota        float64   `json:"nota" gorm:"not null" binding:"required,gte=0,lte=10"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamentos
	Aluno     Aluno     `gorm:"foreignKey:AlunoId;references:Id;constraint:OnDelete:CASCADE"`
	Avaliacao Avaliacao `gorm:"foreignKey:AvaliacaoId;references:Id;constraint:OnDelete:CASCADE"`
}

func (AlunoAvaliacao) TableName() string {
	return "aluno_avaliacao"
}

func (aa *AlunoAvaliacao) BeforeCreate(tx *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	aa.Id = uuidStr
	return
}
