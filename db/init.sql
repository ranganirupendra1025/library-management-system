 CREATE TABLE IF NOT EXISTS subscription(
    id SERIAL PRIMARY KEY,
    name  VARCHAR(100) NOT NULL,
    duration INTERVAL NOT NULL,
    cost INTEGER NOT NULL
    ); 


CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50)  NOT NULL,
    age INTEGER NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE ,
    subscription_id INTEGER  REFERENCES subscription(id),
    subscription_end_date DATE 
);
CREATE TABLE IF NOT EXISTS book(
    id SERIAL PRIMARY KEY,
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

 

INSERT INTO users(username,age,email_address,password,is_admin) VALUES('admin',34,'admin@gmail.com','adminpass',true);

INSERT INTO subscription(name,duration,cost) VALUES('Free',null, 0),('Monthly','30 days', 500),('Quarterly', '90 days', 1000),('Yearly', '365 days', 2000);
