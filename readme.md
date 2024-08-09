# Why Hook? 🪝
this is the same api as Rocketseat project but with GORM because ORM is cool and query builder sucks 

it's my first try building a GO api from zero. 

## Running the aplication 🐳

**SetUp the aplication Database**
Note: be carefull of having .env variables like the env.example
```
docker compose up
```

```
docker compose start
```

**Migrate Database**
```
go run migrations/migrate.go
```

**Run API**

```
go run main.go
```

## Tech Stack ⚙️

<p align="center">
    <img src="https://img.shields.io/badge/go-1.22.4-gray?labelColor=gray&color=blue">
    <img src="https://img.shields.io/badge/gin-1.10.0-red?labelColor=gray&color=red">
    <img src="https://img.shields.io/badge/gorm-1.25.11-orange?labelColor=gray&color=orange">
</p>