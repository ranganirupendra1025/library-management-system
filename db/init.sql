 CREATE TABLE IF NOT EXISTS subscription(
    subscription_id SERIAL PRIMARY KEY,
    name  VARCHAR(100) NOT NULL,
    duration INTERVAL NOT NULL,
    cost INTEGER NOT NULL
    ); 


CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL,
    age INTEGER NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE ,
    subscription_id INTEGER  REFERENCES subscription(subscription_id),
    subscription_end_date DATE 
);
CREATE TABLE IF NOT EXISTS book(
    book_id SERIAL PRIMARY KEY,
    genre VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publisher VARCHAR(50) NOT NULL,
    stock_count INTEGER NOT NULL
    );
 

CREATE TABLE IF NOT EXISTS user_book_transaction(
    transaction_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id),
    book_id INTEGER NOT NULL REFERENCES book(book_id),        
    issued_date DATE NOT NULL,
    return_date DATE NOT NULL ,
    fineamt INTEGER NOT NULL,
    book_return BOOLEAN DEFAULT FALSE,
    actual_return_date DATE NOT NULL
    );

 

INSERT INTO users(username,age,email_address,password,is_admin) VALUES('admin',34,'admin@gmail.com','adminpass',true);
