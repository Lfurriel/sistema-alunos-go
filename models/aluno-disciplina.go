package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// AlunoDisciplina representa a matrícula de um aluno em uma disciplina
//
// Essa entidade intermedia a relação many-to-many entre alunos e disciplinas
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

// TableName especifica o nome da tabela do banco de dados para a estrutura AlunoDisciplina
func (AlunoDisciplina) TableName() string {
	return "aluno_disciplina"
}

// BeforeCreate é usado para o GORM que gera e atribui uma nova string UUID ao campo Id antes de um AlunoDisciplina ser criado
func (ad *AlunoDisciplina) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	ad.Id = uuidStr
	return
}
