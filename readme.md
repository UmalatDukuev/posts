Система для добавления и чтения постов и комментариев с использованием GraphQL

Перед запуском приложения прописывается команда для выбора режима хранения данных (in-memory или в локальной бд)

для Windows сначала прописать команду:
$env:STORAGE="postgresql"
или
$env:STORAGE="inmemory"

потом go run server.go

для Linux:
STORAGE=postgresql go run server.go
или
STORAGE=inmemory go run server.go
