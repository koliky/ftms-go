DROP TABLE IF EXISTS APP_ROLE;
DROP TABLE IF EXISTS APP_USER;

CREATE TABLE APP_USER (
    ID int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    ADDRESS varchar(255),
    STATUS varchar(255),
    CREATE_DATE datetime,
    DEPARTMENT varchar(255),
    EMAIL varchar(255),
    EMPLOYEE_ID varchar(255) NOT NULL UNIQUE,
    FIRST_NAME varchar(255),
    IMAGE_PROFILE varchar(255),
    LAST_NAME varchar(255),
    PASSWORD varchar(255),
    PHONE_NUMBER varchar(255),
    SEX varchar(255),
    SHIFT varchar(255),
    START_DATE datetime,
    USERNAME varchar(255) NOT NULL UNIQUE
);

CREATE TABLE APP_ROLE (
    ID int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    CREATE_DATE datetime,
    ROLE_NAME varchar(255) NOT NULL,
    FK_APP_USER_ID int NOT NULL,
    FOREIGN KEY (FK_APP_USER_ID) REFERENCES APP_USER(ID)
);