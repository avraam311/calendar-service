## Обзор

Сервис для работы с календарем событий.
Стэк: Golang, Chi, Docker, PostgreSQL.

---

## Запуск

### 1. Сканируйте репозиторий

```bash
git clone https://github.com/avraam311/calendar-service
cd calendar-service
````

### 2. Создайте `.env` файл

Скопируйте шаблонный файл .env.example и заполните его нужными данными:

```bash
cp .env.example .env
```

Редактируйте `.env` как нужно.

---

### 3. Запуск проекта через docker-compose

Чтобы собрать и запустить проект

```bash
make up
```

---

### 4 Доступ к API

Запросы по адресу:

```
http://localhost:8080
```

---------

## Примечания

* Убедитесь, что Docker и docker-compose установлены на вашей ос
* `.env` файл настроен правильно.
* Backend слушает на порту `8080`.
* Напишите "make down", чтобы остановить работу системы