GoCashier API

GoCashier API adalah RESTful API untuk aplikasi kasir sederhana yang
dibangun dengan Golang, Gin, dan Clean Architecture.

============================== TECH STACK
============================== - Golang - Gin Web Framework - SQL
Database (MySQL/PostgreSQL) - Swaggo Swagger - Clean Architecture

============================== FEATURES ============================== -
Category Management - Product Management - Transaction Checkout -
Transaction Reports (Daily & Range) - Swagger API Documentation - JSON
REST API

============================== PROJECT STRUCTURE
============================== gocashier.db/ - api/ - handler/ -
router.go - internal/ - models/ - repository/ - services/ - docs/ -
main.go - go.mod

============================== INSTALLATION
==============================

1.  Clone Repository git clone
    https://github.com/yourusername/gocashier.git cd gocashier

2.  Install Dependencies go mod tidy

3.  Setup Database CREATE DATABASE gocashier;

4.  Run App go run main.go

Server: http://localhost:8080

============================== SWAGGER DOCS
==============================

Generate docs: swag init –parseInternal

Open Swagger UI: http://localhost:8080/swagger/index.html

============================== API ENDPOINTS
==============================

CATEGORIES POST /api/categories GET /api/categories GET
/api/categories/{id} PUT /api/categories/{id} DELETE
/api/categories/{id}

PRODUCTS POST /api/products GET /api/products GET /api/products/{id} PUT
/api/products/{id} DELETE /api/products/{id} GET
/api/products/{id}/detail

TRANSACTIONS POST /api/transaction/checkout GET
/api/transaction/report/today GET /api/transaction/report

============================== EXAMPLE CREATE PRODUCT
============================== { “name”: “Coffee”, “price”: 15000,
“category_id”: 1 }

============================== EXAMPLE CHECKOUT
============================== { “items”: [ { “product_id”: 1, “qty”: 2
} ] }

============================== ARCHITECTURE
==============================

Handler: - HTTP handling - Input validation

Service: - Business logic

Repository: - Database access

Models: - Entity structs

============================== FUTURE IMPROVEMENTS
============================== - JWT Authentication - Unit Testing -
Docker Support - Logging Middleware

============================== AUTHOR ==============================
GoCashier Project

============================== LICENSE ==============================
MIT License
