package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Disciplina struct {
	Id                    string    `json:"id" gorm:"primaryKey;column:id;type:varchar(36)"`
	Nome                  string    `json:"nome" gorm:"not null;column:nome" binding:"required,min=1,max=60"`
	ProfessorId           string    `json:"professor_id" gorm:"not null;column:professor_id" binding:"required"` // FK
	AnoSemestre           string    `json:"ano_semestre" gorm:"not null;column:ano_semestre" binding:"required,ano_semestre"`
	QuantidadeAlunos      int       `json:"quantidade_alunos" gorm:"not null;column:quantidade_alunos" binding:"-"`
	QuantidadeProvas      int       `json:"quantidade_provas" gorm:"not null;column:quantidade_provas" binding:"-"`
	QuantidadeTrabalhos   int       `json:"quantidade_trabalhos" gorm:"not null;column:quantidade_trabalhos" binding:"-"`
	CargaHorariaPrevista  int       `json:"carga_horaria_prevista" gorm:"not null;column:carga_horaria_prevista" binding:"required,gte=60,lte=120"`
	CargaHorariaRealizada int       `json:"carga_horaria_realizada" gorm:"not null;column:carga_horaria_realizada" binding:"-"`
	NotaMinima            float64   `json:"nota_minima" gorm:"not null;column:nota_minima" binding:"required,gte=5,lte=10"`
	FrequenciaMinima      float64   `json:"frequencia_minima" gorm:"not null;column:frequencia_minima" binding:"required,gte=70,lte=100"`
	CreatedAt             time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt             time.Time `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at;not null"`

	// Relacionamentos
	Alunos    []Aluno   `gorm:"many2many:aluno_disciplinas;foreignKey:Id;joinForeignKey:DisciplinaId;References:Id;joinReferences:AlunoId"`
	Aulas     []Aula    `json:"aulas" gorm:"foreignKey:DisciplinaId;constraint:OnDelete:CASCADE"`
	Professor Professor `gorm:"foreignKey:ProfessorId;constraint:OnDelete:SET NULL"`
}

func (Disciplina) TableName() string {
	return "disciplinas"
}

func (a *Disciplina) BeforeCreate(tx *gorm.DB) (err error) {
	uuidStr := uuid.New().String()
	a.Id = uuidStr
	return
}
