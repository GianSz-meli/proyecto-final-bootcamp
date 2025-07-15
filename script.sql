-- req 1 (localities )

DROP DATABASE bootcamp_project;

CREATE DATABASE IF NOT EXISTS bootcamp_project;

USE bootcamp_project;


CREATE TABLE countries (
   id INT PRIMARY KEY AUTO_INCREMENT,
   country_name VARCHAR(255) NOT NULL
);


CREATE TABLE provinces (
   id INT PRIMARY KEY AUTO_INCREMENT,
   province_name VARCHAR(255) NOT NULL,
   country_id INT NOT NULL,
   FOREIGN KEY (country_id) REFERENCES countries(id)
);


CREATE TABLE localities (
   id INT PRIMARY KEY AUTO_INCREMENT,
   locality_name VARCHAR(255) NOT NULL,
   province_id INT NOT NULL,
   FOREIGN KEY (province_id) REFERENCES provinces(id)
);


CREATE TABLE sellers (
   id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
   cid VARCHAR(255) NOT NULL UNIQUE,
   company_name VARCHAR(255) NOT NULL,
   address VARCHAR(255) NOT NULL,
   telephone VARCHAR(255) NOT NULL,
   locality_id INT NOT NULL,
   FOREIGN KEY (locality_id) REFERENCES localities(id)
);


-- req 2 (carriers y warehouses)
CREATE TABLE carriers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL UNIQUE,
    company_name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    locality_id INT NOT NULL,
    FOREIGN KEY(locality_id) REFERENCES localities(id)
);


CREATE TABLE warehouses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    warehouse_code VARCHAR(255) NOT NULL UNIQUE,
    locality_id INT,
    FOREIGN KEY(locality_id) REFERENCES localities(id)
);

-- req 4 

CREATE TABLE product_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
);


CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_code VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(255) NOT NULL,
    width DECIMAL(10,2) NOT NULL,
    height DECIMAL(10,2) NOT NULL,
    length DECIMAL(10,2) NOT NULL,
    net_weight DECIMAL(10,2) NOT NULL,
    expiration_rate DECIMAL(5,2) NOT NULL,
    recommended_freezing_temperature DECIMAL(5,2) NOT NULL,
    freezing_rate DECIMAL(5,2) NOT NULL,
    product_type_id INT,
    seller_id INT,

    -- Claves foráneas si existen las tablas referenciadas:
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (seller_id) REFERENCES sellers(id)
);


CREATE TABLE products_records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    last_update_date DATETIME(6) NOT NULL,
    purchase_price DECIMAL(19,2) NOT NULL,
    sale_price DECIMAL(19,2) NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);


-- req 3 (product_batches, sections)
CREATE TABLE sections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    section_number VARCHAR(255) NOT NULL,
    current_temperature DECIMAL(10,2) NOT NULL,
    current_capacity INT NOT NULL,
    minimum_temperature DECIMAL(10,2) NOT NULL,
    minimum_capacity INT NOT NULL,
    product_type_id INT NOT NULL,
    warehouse_id INT NOT NULL,

    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);


CREATE TABLE product_batches (
    id INT AUTO_INCREMENT PRIMARY KEY,
    batch_number VARCHAR(255) NOT NULL UNIQUE,
    current_quantity INT NOT NULL,
    current_temperature DECIMAL(10,2) NOT NULL,
    due_date DATE NOT NULL,
    manufacturing_date DATE NOT NULL,
    manufacturing_hour TIME NOT NULL,
    minimum_temperature DECIMAL(10,2) NOT NULL,
    initial_quantity INT NOT NULL,
    product_id INT NOT NULL,
    section_id INT NOT NULL,

    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (section_id) REFERENCES sections(id)
);


-- indices utiles para consulta
-- CREATE INDEX idx_product_batches_batch_number_lookup ON product_batches(batch_number);

-- req 5 (employees, inbound_orders)
CREATE TABLE employees (
    id INT PRIMARY KEY AUTO_INCREMENT,
    card_number_id VARCHAR(50) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    warehouse_id INT NULL,
    
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);
 
CREATE TABLE inbound_orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    order_date DATETIME(6) NOT NULL,
    order_number VARCHAR(255) NOT NULL UNIQUE,
    employee_id INT NOT NULL,
    product_batch_id INT NOT NULL,
    warehouse_id INT NOT NULL,
    
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (product_batch_id) REFERENCES product_batches(id),
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

-- Índices utiles para consultas

-- CREATE INDEX idx_employee_card_number ON employee(card_number_id);
-- CREATE INDEX idx_employee_name ON employee(first_name, last_name);

-- CREATE INDEX idx_inbound_orders_order_number ON inbound_orders(order_number);
-- CREATE INDEX idx_inbound_orders_date ON inbound_orders(order_date);


-- req 6 (buyers, purchase_orders)
CREATE TABLE buyers (
	id INT PRIMARY KEY AUTO_INCREMENT,
	id_card_number VARCHAR(255) UNIQUE NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL
);

CREATE TABLE order_status (
	id INT PRIMARY KEY AUTO_INCREMENT,
	description VARCHAR(255) NOT NULL UNIQUE
);


CREATE TABLE purchase_orders (
	id INT PRIMARY KEY AUTO_INCREMENT,
	order_number VARCHAR(255) UNIQUE NOT NULL,
	order_date DATETIME NOT NULL,
	tracking_code varchar(255) NOT NULL,
	buyer_id INT NOT NULL,
	carrier_id INT NOT NULL,
	order_status_id INT NOT NULL,
	warehouse_id INT NOT NULL,
	
	FOREIGN KEY (buyer_id) REFERENCES buyers(id),
	FOREIGN KEY (carrier_id) REFERENCES carriers(id),
	FOREIGN KEY (order_status_id) REFERENCES order_status(id),
	FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);

