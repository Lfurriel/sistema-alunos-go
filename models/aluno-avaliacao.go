package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AlunoAvaliacao representa a nota que um aluno tirou em uma determinada avaliação.
//
// Também referencia a disciplina à qual a avaliação pertence
type AlunoAvaliacao struct {
	Id           string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	AlunoId      string    `json:"aluno_id" gorm:"type:varchar(36);not null" binding:"required"`
	AvaliacaoId  string    `json:"avaliacao_id" gorm:"type:varchar(36);not null"`
	DisciplinaId string    `json:"disciplina_id" gorm:"type:varchar(36);not null"`
	Nota         float64   `json:"nota" gorm:"not null" binding:"required,gte=0,lte=10"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamentos
	Aluno     *Aluno     `json:"-" gorm:"foreignKey:AlunoId;references:Id"`
	Avaliacao *Avaliacao `json:"-" gorm:"foreignKey:AvaliacaoId;references:Id"`
}

// TableName especifica o nome da tabela do banco de dados para a estrutura AlunoAvaliacao
func (AlunoAvaliacao) TableName() string {
	return "aluno_avaliacao"
}

// BeforeCreate é usado para o GORM que gera e atribui uma nova string UUID ao campo Id antes de um AlunoAvaliacao ser criado
func (aa *AlunoAvaliacao) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	aa.Id = uuidStr
	return
}
