# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Устанавливаем необходимые зависимости для сборки
RUN apk update && apk add --no-cache \
    git \
    nodejs \
    npm

# Копируем модули и устанавливаем зависимости для Go
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бэкенда (Go)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Переход в папку фронтенда и установка его зависимостей
WORKDIR /app/frontend
RUN npm install && npm run generate

# Возвращаемся в корневую директорию
WORKDIR /app

# Этап формирования минимального контейнера
FROM alpine:latest AS runtime

# Устанавливаем только Nginx
RUN apk add --no-cache nginx

# Копируем сгенерированные файлы фронтенда
COPY --from=builder /app/frontend/.output/public /usr/share/nginx/html

# Копируем бинарник бэкенда
COPY --from=builder /app/main /main

# Копируем конфигурацию Nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Делаем точку входа
COPY run.sh /app/run.sh
RUN chmod +x /app/run.sh

# Устанавливаем рабочую директорию
WORKDIR /app

# Открываем порты (пример: 80 для Nginx)
EXPOSE 80

# Устанавливаем команду запуска
CMD ["/app/run.sh"]