wla-micro-server
================
Простой JSON WEB-сервер на GOLANG, сделанный на основе пакета [Chi](https://github.com/go-chi/chi/v5).

## Структура
+ actions.go - Описываются все действия
+ controller.go - Контейнер для хранения действий
+ middlewares.go - Место для хранения middleware всех запросов
+ model.go - Место для хранения всех структур с которыми ведется работа в экшенах
+ requests.go - Место для хранения обработчиков данных перед их передачей в действие
+ router.go - Место где настраиваются все обрабатываемые роуты
+ server.go - Инициализация слушателя портов

## Создание контроллера
Для создания контроллера необходимо вызвать функцию ```NewController()``` из файла controller.go, пример:
```go
controller := NewController()
```
Контроллер является общим хранилищем тегированных действий и структур для парсинга JSON реквеста. 
Для записи действий в контроллер можно использовать его встроенные функции типа:
+ AddAction
+ AddStructFunc
+ SetActions
+ RemoveAction
+ UnsetActions

Или передать контроллер в конструктор роутов ```router.go``` в функцию ```GetRoutes```, 
в которой можно описать все действия, промежуточные функции и реквесты, пример:
```go
func GetRoutes(c *controller) func(router *chi.Mux) {
	return func(router *chi.Mux) {
		router.Group(func(r chi.Router) {
			r.With(
				jsonInCtxStruct(c.GetStructFunc("POST.")),
				RootRequest,
			).Post("/", c.GetAction("POST."))
		})

		router.Get("/", c.GetAction("."))
		router.Get("/home", c.GetAction(".home"))
	}
}
```
### middlewares и requests
#### middlewares
В этих строчках:
```go
r.With(
    jsonInCtxStruct(c.GetStructFunc("POST.")),
    RootRequest,
)
```
Используются функции из файла ```middlewares.go``` а именно ```jsonInCtxStruct```, 
которая парсит входящий JSON файл и преобразует его в структуру полученную из действия (GetStructFunc контроллера) 
и записывает данное значение в контекст с ключом ```json```
#### requests
Функция ```RootRequest``` определена в ```requests.go``` реквесты необходимы для валидации данных и 
выбрасывания ошибок в случае исключений 

# Start
Для старта сервера необходимо прописать всего несколько строк:
```go
controller := NewController()
SetActions(controller)
StartServe("", 1111, controller)
```





