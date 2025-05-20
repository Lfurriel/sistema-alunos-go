package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword gera um hash seguro a partir de uma senha
//
// Retorna a senha criptografada como string ou um erro, caso a operação falhe.
func CriptografaSenha(senha string) (string, error) {
	saltRounds := 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(senha), saltRounds)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparaSenha verifica se uma senha em texto puro corresponde ao hash recebido.
//
// Retorna true se a senha coincidir com o hash, ou false caso contrário.
func ComparaSenha(senha, senhaHashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(senhaHashed), []byte(senha))
	return err == nil
}
