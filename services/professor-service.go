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

func CadastrarProfessor(professor models.Professor) (*models.Professor, *utils.RestErr) {
	var err error
	professor.Senha, err = utils.HashPassword(professor.Senha)
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criptografar senha", err)
	}
	professor.ConfirmarSenha = ""

	var profExist *models.Professor
	err = database.DB.Where("email = ?", professor.Email).First(&profExist).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao verificar se o professor já existe", err)
	}

	if profExist != nil && profExist.Id != "" {
		return nil, utils.NewRestErr(http.StatusBadRequest, "Professor com email já cadastrado", nil)
	}

	err = database.DB.Create(&professor).Error
	if err != nil {
		return nil, utils.NewRestErr(http.StatusInternalServerError, "Erro ao criar professor", err)
	}

	return &professor, nil
}

func Login(login models.Login) (string, *models.Professor, *utils.RestErr) {
	var professor *models.Professor
	err := database.DB.Where("email = ?", login.Email).First(&professor).Error
	if err != nil {
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

func geraToken(professor *models.Professor) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"professor": professor,                             // Dados do professor
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Expiração do token (1 dia)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
