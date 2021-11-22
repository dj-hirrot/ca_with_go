---- create ----
CREATE DATABASE IF NOT EXISTS cawithgo;

USE cawithgo;

CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

---- insert ----
insert into
    users(name, age)
values
    ("山本五十六", 51);

insert into
    users(name, age)
values
    ("阿南惟幾", 53);

insert into
    users(name, age)
values
    ("米内光政", 55);

insert into
    users(name, age)
values
    ("東條英機", 57);

insert into
    users(name, age)
values
    ("鈴木貫太郎", 59);