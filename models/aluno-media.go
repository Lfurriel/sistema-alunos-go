package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AlunoMedia armazena o resultado final de um aluno ao final de uma disciplina
//
// Inclui a média final, frequência e status de aprovação com base nos critérios da disciplina
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

// TableName especifica o nome da tabela do banco de dados para a estrutura AlunoMedia
func (AlunoMedia) TableName() string {
	return "aluno_media"
}

// BeforeCreate é usado para o GORM que gera e atribui uma nova string UUID ao campo Id antes de um AlunoMedia ser criado
func (am *AlunoMedia) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	am.Id = uuidStr
	return
}
