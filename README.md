# go-bank

A simple bank project written in Go.
Aim to cover:
- SQL
  - Schema Design
  - Migration
  - Postgres & its docker operations
  - Generate SQL Go code using sqlc 
  - CRUD operations & Unit tests
  - Transactions
- CI Using GitHub Actions
- RESTful API Using Gin Framework
  - Read config using Viper
  - Mock DB using mockgen
    - Custom Mock Matchers
  - Custom validators in Gin and ValidatorV10
  - Securely store passwords using hash
  - Access tokens using PASETO
  - Login function
