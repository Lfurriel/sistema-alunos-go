package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"sistema-alunos-go/database"
	"sistema-alunos-go/models"
	"sistema-alunos-go/utils"
	"time"
)

// CadastrarProfessor registra um novo professor no sistema
//
// Criptografa a senha, verifica se já existe um professor com o mesmo e-mail, e então persiste o novo professor
// no banco de dados
//
// Retorna o professor criado (com a senha removida) ou erro em caso de falha
func CadastrarProfessor(professor models.Professor) (*models.Professor, *utils.RestErr) {
	var err error
	professor.Senha, err = utils.CriptografaSenha(professor.Senha)
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criptografar senha", err)
	}
	professor.ConfirmarSenha = ""

	profExiste, restErr := buscaProfessorEmail(professor.Email)
	if restErr != nil && restErr.Err != nil && !errors.Is(restErr.Err, gorm.ErrRecordNotFound) {
		return nil, restErr
	}

	if profExiste != nil && profExiste.Id != "" {
		return nil, utils.NewRestErr(400, "Professor com email já cadastrado", nil)
	}

	err = database.DB.Create(&professor).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criar professor", err)
	}

	professor.Senha = ""
	return &professor, nil
}

// Login autentica um professor com base no e-mail e senha fornecidos
//
// Valida as credenciais, gera um token JWT com os dados do professor, e retorna o token gerado junto
// com os dados do professor
//
// Retorna erro em caso de credenciais inválidas ou falha de autenticação
func Login(login models.Login) (string, *models.Professor, *utils.RestErr) {
	var professor *models.Professor
	if err := database.DB.Where("email = ?", login.Email).First(&professor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, utils.NewRestErr(http.StatusNotFound, "Professor não encontrado", err)
		}
		return "", nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar professor", err)
	}

	if !utils.ComparaSenha(login.Senha, professor.Senha) {
		return "", nil, utils.NewRestErr(http.StatusUnauthorized, "Senha incorreta", nil)
	}

	professor.Senha = ""
	token, err := geraToken(professor)
	if err != nil {
		return "", nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao gerar token", err)
	}
	return token, professor, nil
}

// RemoverProfessor exclui um professor do sistema com base no ID fornecido
//
// Retorna erro caso o professor não exista ou ocorra falha ao remover
func RemoverProfessor(professorId string) *utils.RestErr {
	professor, restErr := buscaProfessor(professorId)
	if restErr != nil {
		return restErr
	}

	if err := database.DB.Delete(&professor).Error; err != nil {
		return utils.NewRestErr(http.StatusInternalServerError, "Erro ao remover professor", err)
	}
	return nil
}

// geraToken cria um token JWT com os dados do professor autenticado com um tempo de expiração de 24 horas
//
// Retorna o token JWT como string ou erro em caso de falha ao assinar
func geraToken(professor *models.Professor) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"professor": professor,                             // Dados do professor
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Expiração do token (1 dia)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// buscaProfessor busca um professor pelo ID
//
// Retorna o professor encontrado ou erro caso não exista ou a consulta falhe
func buscaProfessor(id string) (*models.Professor, *utils.RestErr) {
	var professor models.Professor
	err := database.DB.Where("id = ?", id).First(&professor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Professor não encontrado", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar professor", err)
	}

	return &professor, nil
}

// buscaProfessorEmail busca um professor com base no e-mail
//
// Retorna o professor encontrado ou erro caso não exista ou a consulta falhe
func buscaProfessorEmail(email string) (*models.Professor, *utils.RestErr) {
	var professor models.Professor
	if err := database.DB.Where("email = ?", email).Find(&professor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewRestErr(http.StatusNotFound, "Professor não encontrado", err)
		}
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao buscar professor", err)
	}

	return &professor, nil
}
