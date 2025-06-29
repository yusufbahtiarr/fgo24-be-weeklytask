
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  fullname VARCHAR(255) DEFAULT '',
  phone VARCHAR(255) DEFAULT '',
  email VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(120) NOT NULL,
  pin CHAR(6) NOT NULL CHECK (pin ~ '^\d{6}$'),
  balance BIGINT,
  profile_image VARCHAR(255)
);

CREATE TABLE contacts (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  friend_id INT REFERENCES users(id),
  created_at TIMESTAMP(0) DEFAULT NOW(),
  CONSTRAINT unique_contact UNIQUE (user_id, friend_id),
  CHECK (user_id != friend_id)
);

CREATE TABLE payment_methods (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  transaction_fee INT,
  payment_image VARCHAR(255)
);

CREATE TYPE type_trans AS ENUM ('topup', 'transfer', 'send');

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  transaction_type type_trans,
  amount BIGINT,
  description VARCHAR(255),
  created_at TIMESTAMP(0) DEFAULT NOW(),
  sender_id INT REFERENCES users(id),
  receiver_id INT REFERENCES users(id),
  payment_method_id INT REFERENCES payment_methods(id) 
);

CREATE TABLE sessions (
  id SERIAL PRIMARY KEY,
  token VARCHAR(255) UNIQUE NOT NULL,
  is_active BOOLEAN DEFAULT TRUE, 
  created_at TIMESTAMP(0) DEFAULT NOW(),
  expired_at TIMESTAMP(0) DEFAULT NOW() + INTERVAL '1 day',
  user_id INT REFERENCES users(id)
);