# Используем базовый образ Golang
FROM golang:latest

# Обновляем индексы пакетов
RUN apt-get update

# Устанавливаем PostgreSQL client и другие необходимые пакеты
RUN apt-get -y install postgresql-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Создаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы вашего Go-приложения в контейнер
COPY ./ ./

# Устанавливаем зависимости Go
RUN go mod download

# Собираем Go-приложение внутри контейнера
RUN go build -o cs2-bot ./cmd/cs2-bot/main.go

# Определяем команду для запуска приложения
CMD ["./cs2-bot"]