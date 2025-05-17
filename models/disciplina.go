package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Disciplina struct {
	Id                    string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36)"`
	Nome                  string    `json:"nome" gorm:"not null;column:nome;index" binding:"required,min=1,max=60"`
	ProfessorId           string    `json:"professor_id" gorm:"not null;column:professor_id;index"` // FK
	AnoSemestre           string    `json:"ano_semestre" gorm:"not null;column:ano_semestre" binding:"required,ano_semestre"`
	QuantidadeAlunos      int       `json:"quantidade_alunos" gorm:"not null;column:quantidade_alunos;default:0"`
	QuantidadeProvas      int       `json:"quantidade_provas" gorm:"not null;column:quantidade_provas;default:0"`
	QuantidadeTrabalhos   int       `json:"quantidade_trabalhos" gorm:"not null;column:quantidade_trabalhos;default:0"`
	CargaHorariaPrevista  int       `json:"carga_horaria_prevista" gorm:"not null;column:carga_horaria_prevista" binding:"required,gte=60,lte=120"`
	CargaHorariaRealizada int       `json:"carga_horaria_realizada" gorm:"not null;column:carga_horaria_realizada;default:0"`
	NotaMinima            float64   `json:"nota_minima" gorm:"not null;column:nota_minima" binding:"required,gte=5,lte=10"`
	FrequenciaMinima      float64   `json:"frequencia_minima" gorm:"not null;column:frequencia_minima" binding:"required,gte=70,lte=100"`
	CreatedAt             time.Time `json:"created_at" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt             time.Time `json:"updated_at" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamentos
	Alunos    []Aluno    `json:"alunos,omitempty" gorm:"many2many:aluno_disciplina;foreignKey:Id;joinForeignKey:DisciplinaId;References:Id;joinReferences:AlunoId;constraint:OnDelete:CASCADE"`
	Aulas     []Aula     `json:"aulas,omitempty" gorm:"foreignKey:DisciplinaId;constraint:OnDelete:CASCADE"`
	Professor *Professor `json:"professor,omitempty" gorm:"foreignKey:ProfessorId;constraint:OnDelete:SET NULL"`
}

func (Disciplina) TableName() string {
	return "disciplinas"
}

func (a *Disciplina) BeforeCreate(_ *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
