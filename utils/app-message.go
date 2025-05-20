package utils

// AppMessage é uma estrutura padrão de retorno das APIs
type AppMessage struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

// NewAppMessage cria e retorna uma nova instância de AppMessage
//
// Define o status como "success" se o status HTTP estiver entre 200 e 299; caso contrário, define como "error"
// Pode receber erros opcionais no parâmetro variádico `errors`, sendo usado apenas o primeiro elemento
//
// Retorna um ponteiro para AppMessage contendo statusCode, status, mensagem, payload de dados e erro (se houver)
func NewAppMessage(message string, statusCode int, data interface{}, errors ...interface{}) *AppMessage {
	status := "error"
	if statusCode >= 200 && statusCode <= 299 {
		status = "success"
	}

	var errData interface{}
	if len(errors) > 0 {
		errData = errors[0]
	}

	return &AppMessage{
		StatusCode: statusCode,
		Status:     status,
		Message:    message,
		Data:       data,
		Errors:     errData,
	}
}
