# 📚 Sistema de Gerenciamento de Alunos - API em Go

Este projeto é uma API RESTful desenvolvida em Go (Golang) com o objetivo de gerenciar entidades educacionais como
alunos, professores, disciplinas, aulas e avaliações. O sistema é voltado para uso **local**, com banco de dados **PostgreSQL**, e segue boas práticas de modelagem, autenticação com **JWT**, e uso do ORM **GORM**.

---

## 💡 Proposta do Projeto

A proposta é reestruturar um antigo projeto em C voltado ao gerenciamento de alunos, convertendo-o em uma **API moderna
e escalável**, com rotas HTTP e persistência em banco de dados. Os principais objetivos incluem:

- Cadastro e autenticação de professores
- Cadastro de alunos, disciplinas, aulas e avaliações
- Controle de presenças e notas por disciplina e aula
- Relacionamentos sólidos entre entidades via GORM

---

## 👨‍💻 Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) – framework web
- [GORM](https://gorm.io/) – ORM para Go
- [PostgreSQL](https://www.postgresql.org/) – banco de dados
- [JWT](https://jwt.io/) – autenticação

---

## 🛠️ Como executar o projeto localmente

### 1. Configure o arquivo .env

Crie um arquivo .env na raiz com o seguinte conteúdo:

```text
DATABASE_HOST=localhost
DATABASE_USER=seu_usuario
DATABASE_PASS=sua_senha
DATABASE_NAME=nome_do_banco
DATABASE_PORT=5432
DATABASE_SSL=false
PORT=porta_do_servidor
JWT_SECRET=sua_chave_secreta_super_segura
```

### 2. Instale as dependências

```bash
  go mod tidy
```

### 3. Rode a aplicação

```bash
  go run cmd/main.go
```

A aplicação deve subir na porta `localhost:3333` (ou conforme definido no seu .env).

---

## 🧭 Endpoints Disponíveis

### Criar um novo professor

**POST** /professor

- Body (JSON):

```json
{
  "nome": "Geraldo da Silva",
  "email": "geraldo@example.com",
  "senha": "senha#Forte123",
  "confirmar_senha": "senha#Forte123"
}
```

### Login de professor

**POST** /professor/login

- Body (JSON):

```json
{
  "email": "geraldo@example.com",
  "senha": "senha#Forte123"
}
```