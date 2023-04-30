# xakaton-nt

### Запуск проекта

1. Склонировать репозиторий

```
git clone 
```
2. Перейти в папку проекта
```
cd xakaton-nt
```
3. Запустить проект


```
docker-compose up --build -d
```

4. Перейти в браузере по адресу

http://localhost:80/swagger/index.html

### Запуск тестов

1. Перейти в папку проекта
```
cd xakaton-nt
```
2. Запустить тесты


```
go test ./...
```

### Демонстрация работы

1. Перейти в браузере по адресу

http://localhost:80/swagger/index.html

2. Авторизоваться

Берем access_token из ответа на запрос

```
POST /api/v1/auth/login
```

3. Вставить токен в поле Authorize
```
Bearer <token>
```
4. Выбрать нужный метод

5. Нажать Try it out

6. Вставить данные в поля

7. Нажать Execute

8. Посмотреть результат
