build: # Сборка бинарника
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size;
lint: # Прогон линтера
	golangci-lint run
lint-fix: # Автоматическое исправление замечаний линтера
	golangci-lint run --fix
test: # Запуск тестов
	go test ./...