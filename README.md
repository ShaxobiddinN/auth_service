# Blogpost
## It consists of commits that show step by step evolution of a REST API project in Golang.

Migrate DB:
<br>
```migrate -path ./storage/migration -database 'postgres://admin:postgres@localhost:5432/auth_service_db?sslmode=disable' up```