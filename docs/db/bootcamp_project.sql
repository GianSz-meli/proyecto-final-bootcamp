-- req 1 (localities )
DROP DATABASE bootcamp_project;
CREATE DATABASE IF NOT EXISTS bootcamp_project;
USE bootcamp_project;
CREATE TABLE countries (
                           id INT PRIMARY KEY AUTO_INCREMENT,
                           country_name VARCHAR(255) NOT NULL
);
INSERT INTO countries (country_name) VALUES
                                         ('Argentina'),
                                         ('Brazil'),
                                         ('United States'),
                                         ('Spain'),
                                         ('France'),
                                         ('Italy'),
                                         ('Germany'),
                                         ('Mexico'),
                                         ('Chile'),
                                         ('Peru');
CREATE TABLE provinces (
                           id INT PRIMARY KEY AUTO_INCREMENT,
                           province_name VARCHAR(255) NOT NULL,
                           country_id INT NOT NULL,
                           FOREIGN KEY (country_id) REFERENCES countries(id)
);
INSERT INTO provinces (province_name, country_id) VALUES
                                                      ('Buenos Aires', 1),
                                                      ('Cordoba', 1),
                                                      ('Santa Fe', 1),
                                                      ('Rio de Janeiro', 2),
                                                      ('Sao Paulo', 2),
                                                      ('California', 3),
                                                      ('Texas', 3),
                                                      ('Madrid', 4),
                                                      ('Ile-de-France', 5),
                                                      ('Lombardy', 6);
CREATE TABLE localities (
                            id INT PRIMARY KEY AUTO_INCREMENT,
                            locality_name VARCHAR(255) NOT NULL,
                            province_id INT NOT NULL,
                            FOREIGN KEY (province_id) REFERENCES provinces(id)
);
INSERT INTO localities (locality_name, province_id) VALUES
                                                        ('La Plata', 1),
                                                        ('Rosario', 3),
                                                        ('Cordoba Capital', 2),
                                                        ('Copacabana', 4),
                                                        ('Campinas', 5),
                                                        ('Los Angeles', 6),
                                                        ('Houston', 7),
                                                        ('Getafe', 8),
                                                        ('Paris', 9),
                                                        ('Milan', 10);
CREATE TABLE sellers (
                         id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
                         cid VARCHAR(255) NOT NULL UNIQUE,
                         company_name VARCHAR(255) NOT NULL,
                         address VARCHAR(255) NOT NULL,
                         telephone VARCHAR(255) NOT NULL,
                         locality_id INT NOT NULL UNIQUE,
                         FOREIGN KEY (locality_id) REFERENCES localities(id)
);
INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES
                                                                             ('CID001', 'Acme Fresh Foods', '123 Main Ave', '111-1111', 1),
                                                                             ('CID002', 'Brazilian Foods Ltda', '456 Copacabana', '222-2222', 4),
                                                                             ('CID003', 'Califruit USA', '789 LA St', '333-3333', 6),
                                                                             ('CID004', 'Cordoba Vendors', '321 Cordoba', '444-4444', 3),
                                                                             ('CID005', 'Rosario Goods', '654 Rosario', '555-5555', 2),
                                                                             ('CID006', 'Paris Gourmet', '987 Paris Rd', '666-6666', 9),
                                                                             ('CID007', 'Lombardy Cheese', '123 Milan Ave', '777-7777', 10),
                                                                             ('CID008', 'Sao Paulo Foods', '456 Campinas', '999-9999', 5),
                                                                             ('CID09', 'Madrid Imports', 'XYZ Getafe', '000-0000', 8);


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
INSERT INTO carriers (cid, company_name, address, telephone, locality_id) VALUES
                                                                              ('CAR001', 'QuickTrans', '101 Route', '623-0000', 1),
                                                                              ('CAR002', 'RapidoExpress', '102 Route', '623-0001', 2),
                                                                              ('CAR003', 'GlobalCarrier', '103 Route', '623-0002', 3),
                                                                              ('CAR004', 'AzulLogistics', '104 Route', '623-0003', 4),
                                                                              ('CAR005', 'RedTransport', '105 Route', '623-0004', 5),
                                                                              ('CAR006', 'BlueSky', '106 Route', '623-0005', 6),
                                                                              ('CAR007', 'RoadRunner', '107 Route', '623-0006', 7),
                                                                              ('CAR008', 'EuroCarrier', '108 Route', '623-0007', 9),
                                                                              ('CAR009', 'MilanFreight', '109 Route', '623-0008', 10),
                                                                              ('CAR010', 'TransMadrid', '110 Route', '623-0009', 8);
CREATE TABLE warehouses (
                            id INT AUTO_INCREMENT PRIMARY KEY,
                            address VARCHAR(255) NOT NULL,
                            telephone VARCHAR(255) NOT NULL,
                            warehouse_code VARCHAR(255) NOT NULL UNIQUE,
                            minimum_capacity INT UNSIGNED NOT NULL,
                            minimum_temperature DECIMAL(10,2) NOT NULL,
                            locality_id INT,
                            FOREIGN KEY(locality_id) REFERENCES localities(id) ON DELETE SET NULL
);
INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature, locality_id) VALUES
                                                                                                                    ('Warehouse 1, La Plata', '5001', 'W1', 100, -18.00, 1),
                                                                                                                    ('Warehouse 2, Rosario', '5002', 'W2', 120, -20.00, 2),
                                                                                                                    ('Warehouse 3, Cordoba', '5003', 'W3', 80, -15.00, 3),
                                                                                                                    ('Warehouse 4, Copacabana', '5004', 'W4', 150, -19.00, 4),
                                                                                                                    ('Warehouse 5, Campinas', '5005', 'W5', 90, -16.00, 5),
                                                                                                                    ('Warehouse 6, LA', '5006', 'W6', 200, -21.00, 6),
                                                                                                                    ('Warehouse 7, Houston', '5007', 'W7', 110, -18.50, 7),
                                                                                                                    ('Warehouse 8, Paris', '5008', 'W8', 170, -22.00, 9),
                                                                                                                    ('Warehouse 9, Milan', '5009', 'W9', 160, -17.00, 10),
                                                                                                                    ('Warehouse 10, Getafe', '5010', 'W10', 130, -14.00, 8);
-- req 4
CREATE TABLE product_types (
                               id INT AUTO_INCREMENT PRIMARY KEY,
                               description VARCHAR(255) NOT NULL
);
INSERT INTO product_types (description) VALUES
                                            ('Fruits'),
                                            ('Vegetables'),
                                            ('Dairy'),
                                            ('Meat'),
                                            ('Frozen Meals'),
                                            ('Beverages'),
                                            ('Bakery'),
                                            ('Seafood'),
                                            ('Snacks'),
                                            ('Condiments');
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
INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES
                                                                                                                                                                                      ('P001', 'Apple', 10.0, 8.0, 12.0, 0.20, 0.15, -2.00, 0.05, 1, 1),
                                                                                                                                                                                      ('P002', 'Banana', 11.0, 6.0, 15.0, 0.25, 0.18, -3.00, 0.06, 1, 1),
                                                                                                                                                                                      ('P003', 'Broccoli', 13.0, 10.0, 13.0, 0.18, 0.12, -4.00, 0.04, 2, 2),
                                                                                                                                                                                      ('P004', 'Milk', 10.0, 25.0, 10.0, 1.00, 0.30, -0.50, 0.08, 3, 6),
                                                                                                                                                                                      ('P005', 'Cheese', 8.0, 5.0, 20.0, 0.35, 0.16, -5.00, 0.03, 3, 7),
                                                                                                                                                                                      ('P006', 'Beef', 25.0, 10.0, 40.0, 2.00, 0.20, -15.00, 0.20, 4, 4),
                                                                                                                                                                                      ('P007', 'Chicken Nuggets', 14.0, 12.0, 15.0, 0.40, 0.22, -17.00, 0.18, 5, 5),
                                                                                                                                                                                      ('P008', 'Orange Juice', 8.0, 30.0, 8.0, 1.20, 0.17, -1.00, 0.10, 6, 3),
                                                                                                                                                                                      ('P009', 'Baguette', 6.0, 5.0, 60.0, 0.30, 0.10, -10.00, 0.01, 7, 6),
                                                                                                                                                                                      ('P010', 'Tuna', 15.0, 6.0, 15.0, 0.60, 0.25, -17.00, 0.19, 8, 8);
CREATE TABLE products_records (
                                  id INT AUTO_INCREMENT PRIMARY KEY,
                                  last_update_date DATETIME(6) NOT NULL,
                                  purchase_price DECIMAL(19,2) NOT NULL,
                                  sale_price DECIMAL(19,2) NOT NULL,
                                  product_id INT NOT NULL,
                                  FOREIGN KEY (product_id) REFERENCES products(id)
);
INSERT INTO products_records (last_update_date, purchase_price, sale_price, product_id) VALUES
                                                                                            ('2024-06-01 10:00:00.000000', 10.00, 15.00, 1),
                                                                                            ('2024-06-02 11:00:00.000000', 12.00, 18.00, 2),
                                                                                            ('2024-06-03 12:00:00.000000', 8.50, 11.00, 3),
                                                                                            ('2024-06-04 13:00:00.000000', 20.00, 27.00, 4),
                                                                                            ('2024-06-05 14:00:00.000000', 30.00, 45.00, 5),
                                                                                            ('2024-06-06 15:00:00.000000', 50.00, 75.00, 6),
                                                                                            ('2024-06-07 16:00:00.000000', 25.00, 32.00, 7),
                                                                                            ('2024-06-08 17:00:00.000000', 15.00, 21.00, 8),
                                                                                            ('2024-06-09 18:00:00.000000', 4.50, 7.00, 9),
                                                                                            ('2024-06-10 19:00:00.000000', 10.50, 15.00, 10);
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
INSERT INTO sections (section_number, current_temperature, current_capacity, minimum_temperature, minimum_capacity, product_type_id, warehouse_id) VALUES
                                                                                                                                                       ('S1', -2.00, 100, -5.00, 80, 1, 1),
                                                                                                                                                       ('S2', -3.00, 120, -6.00, 110, 1, 2),
                                                                                                                                                       ('S3', -4.00, 90, -5.00, 70, 2, 3),
                                                                                                                                                       ('S4', -5.00, 150, -6.00, 130, 3, 4),
                                                                                                                                                       ('S5', -17.00, 140, -18.00, 120, 4, 5),
                                                                                                                                                       ('S6', -21.00, 180, -22.00, 160, 5, 6),
                                                                                                                                                       ('S7', -18.50, 90, -20.00, 80, 3, 7),
                                                                                                                                                       ('S8', -22.00, 120, -23.00, 110, 5, 8),
                                                                                                                                                       ('S9', -17.00, 100, -18.00, 90, 8, 9),
                                                                                                                                                       ('S10', -14.00, 110, -15.00, 100, 6, 10);
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
INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, manufacturing_date, manufacturing_hour, minimum_temperature, initial_quantity, product_id, section_id) VALUES
                                                                                                                                                                                                       ('BATCH01', 100, -2.00, '2024-07-01', '2024-06-01', '10:00:00', -4.00, 120, 1, 1),
                                                                                                                                                                                                       ('BATCH02', 120, -3.00, '2024-07-05', '2024-06-02', '11:00:00', -5.00, 130, 2, 2),
                                                                                                                                                                                                       ('BATCH03', 90, -4.00, '2024-07-10', '2024-06-03', '12:00:00', -6.00, 100, 3, 3),
                                                                                                                                                                                                       ('BATCH04', 70, -5.00, '2024-07-12', '2024-06-04', '13:00:00', -7.00, 80, 4, 4),
                                                                                                                                                                                                       ('BATCH05', 60, -13.00, '2024-08-01', '2024-06-05', '14:00:00', -14.00, 70, 5, 5),
                                                                                                                                                                                                       ('BATCH06', 180, -21.00, '2024-08-07', '2024-06-06', '15:00:00', -22.00, 200, 6, 6),
                                                                                                                                                                                                       ('BATCH07', 80, -18.50, '2024-08-13', '2024-06-07', '16:00:00', -19.00, 90, 7, 7),
                                                                                                                                                                                                       ('BATCH08', 110, -22.00, '2024-08-20', '2024-06-08', '17:00:00', -23.00, 120, 8, 8),
                                                                                                                                                                                                       ('BATCH09', 85, -17.00, '2024-09-01', '2024-06-09', '18:00:00', -18.00, 100, 9, 9),
                                                                                                                                                                                                       ('BATCH10', 95, -14.00, '2024-09-10', '2024-06-10', '19:00:00', -15.00, 110, 10, 10);
-- req 5 (employees, inbound_orders)
CREATE TABLE employees (
                           id INT PRIMARY KEY AUTO_INCREMENT,
                           card_number_id VARCHAR(50) NOT NULL UNIQUE,
                           first_name VARCHAR(100) NOT NULL,
                           last_name VARCHAR(100) NOT NULL,
                           warehouse_id INT NULL,

                           FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);
INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES
                                                                                ('EMP001', 'John', 'Doe', 1),
                                                                                ('EMP002', 'Ana', 'Silva', 2),
                                                                                ('EMP003', 'Lucas', 'Smith', 3),
                                                                                ('EMP004', 'Carla', 'Pereira', 4),
                                                                                ('EMP005', 'Pedro', 'Santos', 5),
                                                                                ('EMP006', 'Marie', 'Dubois', 8),
                                                                                ('EMP007', 'Giulia', 'Bianchi', 9),
                                                                                ('EMP008', 'Sara', 'Gonzalez', 10),
                                                                                ('EMP009', 'Leonel', 'Perez', 7),
                                                                                ('EMP010', 'Tom', 'Taylor', 6);
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
INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES
                                                                                                       ('2024-06-01 10:00:00.000000', 'IN0001', 1, 1, 1),
                                                                                                       ('2024-06-02 11:00:00.000000', 'IN0002', 2, 2, 2),
                                                                                                       ('2024-06-03 12:00:00.000000', 'IN0003', 3, 3, 3),
                                                                                                       ('2024-06-04 13:00:00.000000', 'IN0004', 4, 4, 4),
                                                                                                       ('2024-06-05 14:00:00.000000', 'IN0005', 5, 5, 5),
                                                                                                       ('2024-06-06 15:00:00.000000', 'IN0006', 6, 6, 8),
                                                                                                       ('2024-06-07 16:00:00.000000', 'IN0007', 7, 7, 9),
                                                                                                       ('2024-06-08 17:00:00.000000', 'IN0008', 8, 8, 10),
                                                                                                       ('2024-06-09 18:00:00.000000', 'IN0009', 9, 9, 7),
                                                                                                       ('2024-06-10 19:00:00.000000', 'IN0010', 10, 10, 6);
-- req 6 (buyers, purchase_orders)
CREATE TABLE buyers (
                        id INT PRIMARY KEY AUTO_INCREMENT,
                        id_card_number VARCHAR(255) UNIQUE NOT NULL,
                        first_name VARCHAR(255) NOT NULL,
                        last_name VARCHAR(255) NOT NULL
);
INSERT INTO buyers (id_card_number, first_name, last_name) VALUES
                                                               ('BUY001', 'Antonio', 'Ramirez'),
                                                               ('BUY002', 'Julia', 'Mendez'),
                                                               ('BUY003', 'Sofia', 'Lopez'),
                                                               ('BUY004', 'Fernando', 'Gutierrez'),
                                                               ('BUY005', 'Veronica', 'Blanco'),
                                                               ('BUY006', 'Adrian', 'Garcia'),
                                                               ('BUY007', 'Candela', 'Torres'),
                                                               ('BUY008', 'Nicolas', 'Castro'),
                                                               ('BUY009', 'Camila', 'Sanchez'),
                                                               ('BUY010', 'Ramon', 'Fernandez');
CREATE TABLE order_status (
                              id INT PRIMARY KEY AUTO_INCREMENT,
                              description VARCHAR(255) NOT NULL UNIQUE
);
INSERT INTO order_status (description) VALUES
                                           ('Pending'),
                                           ('Processing'),
                                           ('Dispatched'),
                                           ('Delivered'),
                                           ('Cancelled'),
                                           ('On Hold'),
                                           ('Returned'),
                                           ('Completed'),
                                           ('Payment Pending'),
                                           ('Awaiting Pickup');
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
INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id) VALUES
                                                                                                                               ('PO0001', '2024-06-11 10:00:00', 'TRK001', 1, 1, 1, 1),
                                                                                                                               ('PO0002', '2024-06-12 11:00:00', 'TRK002', 2, 2, 2, 2),
                                                                                                                               ('PO0003', '2024-06-13 12:00:00', 'TRK003', 3, 3, 3, 3),
                                                                                                                               ('PO0004', '2024-06-14 13:00:00', 'TRK004', 4, 4, 4, 4),
                                                                                                                               ('PO0005', '2024-06-15 14:00:00', 'TRK005', 5, 5, 5, 5),
                                                                                                                               ('PO0006', '2024-06-16 15:00:00', 'TRK006', 6, 6, 6, 6),
                                                                                                                               ('PO0007', '2024-06-17 16:00:00', 'TRK007', 7, 7, 7, 7),
                                                                                                                               ('PO0008', '2024-06-18 17:00:00', 'TRK008', 8, 8, 8, 8),
                                                                                                                               ('PO0009', '2024-06-19 13:00:00', 'TRK009', 9, 9, 9, 9),
                                                                                                                               ('PO0010', '2024-06-20 14:00:00', 'TRK010', 10, 10, 10, 10);

CREATE TABLE order_details (
                               id INT PRIMARY KEY AUTO_INCREMENT,
                               cleanliness_status VARCHAR(255) NOT NULL,
                               quantity INT NOT NULL,
                               temperature DECIMAL(19,2) NOT NULL,
                               product_record_id INT NOT NULL,
                               purchase_order_id INT NOT NULL,

                               FOREIGN KEY (product_record_id) REFERENCES products_records(id),
                               FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id)
);

INSERT INTO order_details (cleanliness_status, quantity, temperature, product_record_id, purchase_order_id) VALUES
                                                                                                                ('Washed',      25, 4.50, 1, 1),
                                                                                                                ('Unwashed',    10, 6.20, 2, 1),
                                                                                                                ('Sanitized',   12, 3.80, 3, 2),
                                                                                                                ('Needs washing',  5, 7.00, 4, 2),
                                                                                                                ('Washed',      8, 5.60, 5, 3),
                                                                                                                ('Partially washed', 7, 8.10, 1, 3),
                                                                                                                ('Contaminated', 2, 9.20, 2, 4),
                                                                                                                ('Fresh',      20, 2.10, 3, 4),
                                                                                                                ('Washed',      15, 4.00, 4, 5),
                                                                                                                ('Sanitized',    9, 3.30, 5, 5);
