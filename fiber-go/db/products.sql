CREATE DATABASE IF NOT EXISTS golang;

CREATE TABLE IF NOT EXISTS golang.category(
    `ID`        INT NOT NULL AUTO_INCREMENT,
    `Category`  VARCHAR(32),
    `Description`       VARCHAR(1000) NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS golang.product (
    `ID`                INT        NOT NULL AUTO_INCREMENT,
    `Name`              VARCHAR(100)  NOT NULL,
    `Description`       VARCHAR(1000) NULL,
    `Price`             DECIMAL(9,2)  NOT NULL,
    `Quantity`          INT           NOT NULL DEFAULT 0,
    `Category_ID`       INT 
    `Status`            BOOL          NOT NULL DEFAULT TRUE,       
    `Creation_time`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP
    PRIMARY KEY (ID)
    FOREIGN KEY (Category_ID) REFERENCES golang.category(id)
);

CREATE TABLE IF NOT EXISTS golang.order(
    `ID`                BIGINT NOT NULL AUTO_INCREMENT,
    `Customer_ID`       BIGINT NOT NULL,
    `Status_payment`    BOOL,
    `Product_id`        INT NOT NULL,
    `Creation_time`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `Modification_time` DATETIME ON UPDATE CURRENT_TIMESTAMP
    PRIMARY KEY (ID)
    FOREIGN KEY (Product_id) REFERENCES golang.product (id)
    FOREIGN KEY (Customer_ID) REFERENCES golang.customer(id)
);

CREATE TABLE IF NOT EXISTS golang.customer(
    `ID`         INT NOT NULL AUTO_INCREMENT,
    `Name`       varchar(128) NOT NULL,
    `Documents`  DECIMAL(9,2) NOT NULL UNIQUE
    PRIMARY KEY (ID)

);