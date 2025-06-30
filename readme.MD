# Project Backend RESTful API Go - E-Wallet

This backend e-wallet system was developed for Koda Academy's Full-Stack Bootcamp, enhancing a previous iteration with robust features including JWT authentication, transaction processing, and account management

## Prerequisites

Make sure you already install Go to run this project

## How to Run this Project

1. Create a new empty directory for the project and navigate into it
2. Clone this project into the empty current directory:

```
git clone https://github.com/yusufbahtiarr/fgo24-be-weeklytask.git .
```

3. Install dependencies

```
go mod tidy
```

4. Setup .env

Create a .env file in the root folder with the following variables:

```
APP_SECRET=your_jwt_secret_key
DATABASE_URL=postgres://username:password@localhost:5433/ewallet
```

5. Run the project

```
go run main.go
```

## Database Schema

### Tables

- **users**: Stores user information (id, name, email, phone number, password, pin, fullname, phone, picture_image)
- **contact**: Manages user connections by storing friend relationships between accounts.
- **transactions**: Records financial transactions, including top-ups and transfers, with references to sender and receiver users.
- **payment_method**: Catalogs all available payment options along with their respective transfer fees.

### Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    direction LR
    users ||--o{ contacts : has
    users ||--o{ transactions : make
    contacts }o--o| transactions : receives
    transactions ||--o{ payment_methods : uses

    users {
        string id PK
        string fullname
        string email
        string password
        string pin
        int balance
        string profile_image
    }
    transactions {
        string id PK
        string transaction_type
        int amount
        string description
        timestamp created_at
        string sender_id FK
        string receiver_id FK
        string payment_method_id FK
    }
    payment_methods {
        string id PK
        string name
        int transaction_fee
        string payment_image
    }

    contacts {
        string id PK
        int user_id FK
        int friend_id FK
        timestamp created_at
    }

```

## API Documentation

### Auth

| Method | Endpoint         | Description   | Auth Required |
| ------ | ---------------- | ------------- | ------------- |
| POST   | `/auth/register` | Register User | No            |
| POST   | `/auth/login`    | Login User    | No            |

### Users

| Method | Endpoint          | Description                   | Auth Required |
| ------ | ----------------- | ----------------------------- | ------------- |
| GET    | `/users?search`   | Get User by Name for Transfer | Yes           |
| GET    | `/users/balance`  | Get Balance User              | Yes           |
| GET    | `/users/history`  | Get History Transaction User  | Yes           |
| PATCH  | `/users/profile`  | Update Profile                | Yes           |
| PATCH  | `/users/password` | Update Password               | Yes           |
| PATCH  | `/users/pin`      | Update Pin                    | Yes           |

### Transactions

| Method | Endpoint                | Description                 | Auth Required |
| ------ | ----------------------- | --------------------------- | ------------- |
| POST   | `/transaction/topup`    | Create Transaction Topup    | Yes           |
| POST   | `/transaction/transfer` | Create Transaction Transfer | Yes           |

## Dependencies

This project use:

1. gin-gonic from github.com/gin-gonic/gin : for handling HTTP request/response data (gin.Context), for defining middleware and route handlers (gin.HandlerFunc), for organizing routes into groups (gin.RouterGroup) and for managing HTTP routing and server configuration (gin.Engine)
2. jwt v5 from github.com/golang-jwt/jwt/v5 : for creating, parsing and validating JSON Web Tokens (JWT) for authentication and authorization
3. pgx from github.com/jackc/pgx/v5 : for direct database interactions (PostgreSQL)
4. godotenv from github.com/joho/godotenv : for loading environment variables from a .env file into the application
