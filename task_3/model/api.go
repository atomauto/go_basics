package model

type Fail struct {
	OriginalResult int //правильный ответ на тест с точки зрения API сервера
	ExternalResult int //наш неправильный ответ
}

type ApiRequestAnswer struct {
	Percent int    `json:"percent"` //процент выполненных тестов
	Fails   []Fail `json:"fails"`   //массив ошибок
}

type Results struct {
	Payload interface{} `json:"payload"` //данные полученные для решения задачи
	Results interface{} `json:"results"` //результаты полученные при решении задачи
}

type ApiRequestData struct {
	UserName string   `json:"user_name"` //username из telegram бота
	Task     string   `json:"task"`      //название задачи
	Results  *Results `json:"results"`   //массив с результатами и входными данными
}
