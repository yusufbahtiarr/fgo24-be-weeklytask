
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

INSERT INTO users (fullname, email, password, pin, balance, profile_image) VALUES
('Andi Pratama', 'andi.pratama@gmail.com', 'andi1', '123456', 5000000, 'andi_profile.jpg'),
('Budi Santoso', 'budi.santoso@yahoo.co.id', 'budi2', '789012', 2500000, NULL),
('Citra Lestari', 'citra.lestari@outlook.com', 'citra3', '345678', 10000000, 'citra_profile.png'),
('Dewi Sari', 'dewi.sari@gmail.com', 'dewi4', '901234', '750000', NULL),
('Eko Wahyudi', 'eko.wahyudi@hotmail.com', 'eko5', '567890', 3000000, 'eko_profile.jpg');


INSERT INTO sessions (token, user_id) VALUES
('token_1_abc123', 1),
('token_2_def456', 2),
('token_3_ghi789', 3),
('token_4_jkl012', 4),
('token_5_mno345', 5);

INSERT INTO payment_methods (name, transaction_fee, payment_image) VALUES
('Bank Transfer (BCA)', 2500, 'https://placehold.co/100x50/0000FF/FFFFFF?text=BCA'),
('OVO', 1000, 'https://placehold.co/100x50/800080/FFFFFF?text=OVO'),
('Gopay', 800, 'https://placehold.co/100x50/008000/FFFFFF?text=GOPAY'),
('Dana', 750, 'https://placehold.co/100x50/007FFF/FFFFFF?text=DANA'),
('Virtual Account (Mandiri)', 3000, 'https://placehold.co/100x50/FF8C00/FFFFFF?text=MandiriVA');


INSERT INTO contacts (user_id, friend_id) VALUES
(1, 2),
(1, 3),
(2, 1),
(2, 3),
(3, 1),
(3, 4),
(4, 1),
(4, 5),
(5, 2),
(5, 3);
