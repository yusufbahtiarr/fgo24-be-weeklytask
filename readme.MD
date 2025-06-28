# ERD E-Wallet

This project, presents an Entity-Relationship Diagram (ERD) for an E-Wallet system using the Mermaid diagramming tool.

```mermaid
erDiagram
    direction LR
    users ||--o{ contacts : has
    users ||--o{ transactions : make
    users ||--o{ sessions : has
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

    sessions {
        string id PK
        token string
        is_active bool
        timestamps created_at
        timesstamp expired_at
        int user_id FK
    }


```
