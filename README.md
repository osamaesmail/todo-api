# todo-api
> Simple to do list API with Echo framework and Gorm (with Postgres) and containerized.

## requirements
* docker
* docker-compose

## build & run
```
docker-compose build & docker-compose run
```

endpoints:

| Method | Route      | Body                                         |
| ------ | ---------- | -------------------------------------------- |
| GET    | /todo      |                                              |
| GET    | /todo/:id  |                                              |
| POST   | /todo      | `{"title": "todo title"}`                    |
| DELETE | /todo/:id  |                                              |
| PUT    | /todo/:id  | `{"title": "todo title", "completed": true}` |