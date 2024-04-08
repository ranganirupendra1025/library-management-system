 DROP TABLE user_book_transaction;
DROP TABLE users;
 DROP TABLE subscription;
 DROP TABLE book;

 
 CREATE TABLE IF NOT EXISTS subscription(
    id SERIAL PRIMARY KEY,
    name  VARCHAR(100) NOT NULL,
    duration INTEGER NOT NULL ,
    cost INTEGER NOT NULL
    ); 


CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL,
    age INTEGER NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE ,
    subscription_id INTEGER  REFERENCES subscription(id) ,
    subscription_end_date DATE
);
CREATE TABLE IF NOT EXISTS book(
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publisher VARCHAR(50) NOT NULL,
    stock_count INTEGER NOT NULL
    );
 

CREATE TABLE IF NOT EXISTS user_book_transaction(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    book_id INTEGER NOT NULL REFERENCES book(id),        
    issued_date DATE NOT NULL,
    return_date DATE NOT NULL ,
    fineamt INTEGER NOT NULL,
    book_return BOOLEAN DEFAULT FALSE,
    actual_return_date DATE NOT NULL
    );
INSERT INTO subscription(name,duration,cost) VALUES('Monthly',30, 500),('Quarterly', 90, 1000),('Yearly', 365, 2000);
INSERT INTO users(username,age,email_address,password,is_admin,subscription_id,subscription_end_date) VALUES('admin',34,'admin@gmail.com','adminpass',true,3,'2006-01-02');

--DROP TABLE subscription