# NilTodo

Небольшой REST API сервис для управления задачами, написанный на Go.
Проект сделан как pet-project с акцентом на чистую архитектуру, работу с HTTP API, PostgreSQL и Docker.

## Возможности

* CRUD операции для задач
* Работа с пользователями
* Swagger документация
* Middleware для HTTP
* PostgreSQL + миграции
* Docker Compose запуск
* Слойная архитектура (transport / service / repository)

## Технологии

* Go
* PostgreSQL
* Docker & Docker Compose
* Swagger
* SQL migrations

## Структура проекта

```text
cmd/                # Точка входа приложения
internal/           # Основная бизнес-логика
migrations/         # SQL миграции
public/             # Статические файлы
/docs               # Swagger документация
```

## Запуск проекта

### 1. Клонирование репозитория

```bash
git clone <repo_url>
cd golang-niltodo
```

### 2. Настройка env

```bash
cp .env.example .env
```

### 3. Запуск

```bash
docker compose up --build
```

## Swagger

После запуска документация доступна по адресу:

```text
http://localhost:8080/swagger/index.html
```

## Что было реализовано

* Проектирование REST API
* Работа с PostgreSQL
* HTTP handlers и middleware
* Валидация запросов
* Обработка ошибок
* Docker окружение
* Swagger документация

## Цель проекта

Практика backend-разработки на Go и построения понятной архитектуры для REST сервисов.
