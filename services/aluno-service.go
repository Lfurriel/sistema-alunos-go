package services

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
)

// CadastrarAluno insere um novo aluno no banco.
//
// Ele verifica se já existe um aluno com o mesmo e-mail.
// Se não houver, ativa o aluno e o salva.
//
// Retorna o aluno cadastrado ou um erro, caso ocorra falha na verificação ou na criação.
func CadastrarAluno(aluno models.Aluno) (*models.Aluno, *utils.RestErr) {
	alunoExiste, restErr := buscaAlunoEmail(aluno.Email)
	if restErr != nil && restErr.Err != nil && !errors.Is(restErr.Err, gorm.ErrRecordNotFound) {
		return nil, restErr
	}

	if alunoExiste != nil && alunoExiste.Id != "" {
		return nil, utils.NewRestErr(400, "Aluno com email já cadastrado", nil)
	}

	aluno.Ativo = true
	if err := database.DB.Create(&aluno).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao cadastrar aluno", err)
	}
	return &aluno, nil
}

// AtualizarAluno ativa ou desativa o aluno no banco de dados
//
// Ela primeiramente busca o aluno no banco de dados para verificar a existência
// Então verifica se o aluno já se encontra no status que o usuário esta tentando atualizar
// Caso não esteja, o aluno é ativado/desativado e salva no banco
//
// As disciplinas em que o aluno está matriculado é atualizada com a quantidade de alunos matriculados para mais (caso
// esteja ativando o aluno) ou menos (caso contrário)
//
// Retorna o aluno com o novo status ou algum erro durante o processo
func AtualizarAluno(alunoId string, ativo bool) (*models.Aluno, *utils.RestErr) {
	aluno, restErr := buscaAluno(alunoId)
	if restErr != nil {
		return nil, restErr
	}

	if ativo && aluno.Ativo {
		return nil, utils.NewRestErr(400, "Aluno já está ativo", nil)
	}

	if !ativo && !aluno.Ativo {
		return nil, utils.NewRestErr(400, "Aluno já está desativado", nil)
	}

	var alunoDisciplinas []models.AlunoDisciplina
	if err := database.DB.Where("aluno_id = ?", alunoId).Find(&alunoDisciplinas).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplinas do aluno", err)
	}

	for _, ad := range alunoDisciplinas {
		if err := atualizaQuantidadeAlunos(ad.DisciplinaId, ativo); err != nil {
			return nil, err
		}
	}

	aluno.Ativo = ativo
	if err := database.DB.Save(&aluno).Error; err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar aluno", err)
	}
	return aluno, nil
}

// RemoverAluno apaga o registro da tabela "alunos" no banco de dados
//
// Ela primeiramente busca o aluno no banco para verificar a sua existência e então o remove
// Retorna (caso ocorra) erro durante o processo de remoção do aluno
func RemoverAluno(id string) *utils.RestErr {
	aluno, restErr := buscaAluno(id)
	if restErr != nil {
		return restErr
	}

	var alunoDisciplinas []models.AlunoDisciplina
	if err := database.DB.Where("aluno_id = ?", id).Find(&alunoDisciplinas).Error; err != nil {
		return utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar disciplinas do aluno", err)
	}

	for _, ad := range alunoDisciplinas {
		if err := atualizaQuantidadeAlunos(ad.DisciplinaId, false); err != nil {
			return err
		}
	}

	if err := database.DB.Delete(&aluno).Error; err != nil {
		return utils.NewRestErr(http.StatusInternalServerError, "Erro ao remover aluno", err)
	}
	return nil
}

// buscaAluno busca um aluno pelo ID
//
// Retorna o aluno encontrado ou erro caso não exista ou a consulta falhe
func buscaAluno(id string) (*models.Aluno, *utils.RestErr) {
	var aluno models.Aluno

	err := database.DB.Where("id = ?", id).First(&aluno).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Aluno não encontrado", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aluno", err)
	}

	return &aluno, nil
}

// buscaAlunoEmail busca um aluno com base no e-mail
//
// Retorna o aluno encontrado ou erro caso não exista ou a consulta falhe
func buscaAlunoEmail(email string) (*models.Aluno, *utils.RestErr) {
	var aluno models.Aluno
	if err := database.DB.Where("email = ?", email).Find(&aluno).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Aluno não encontrado", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar aluno", err)
	}

	return &aluno, nil
}

// atualizaQuantidadeAlunos busca um disciplina com base no parâmetro 'id' e então atualiza a quantidade de alunos
// matriculados na disciplina em questão
//
// Retorna erro caso ocurra durante a busca da disciplina no banco de dados ou a atualização
func atualizaQuantidadeAlunos(id string, soma bool) *utils.RestErr {
	disciplina, restErr := buscaDisciplina(id)
	if restErr != nil {
		return restErr
	}
	if soma {
		disciplina.QuantidadeAlunos += 1
	} else {
		disciplina.QuantidadeAlunos -= 1
	}

	if err := database.DB.Save(&disciplina).Error; err != nil {
		return utils.NewRestErr(http.StatusInternalServerError, "Erro ao atualizar quantidade de alunos", err)
	}

	return nil
}
