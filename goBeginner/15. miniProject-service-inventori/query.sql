create database "Inventory";


DROP TABLE IF EXISTS "Transactions" CASCADE;
DROP TABLE IF EXISTS "Items" CASCADE;
DROP TABLE IF EXISTS "Categories" CASCADE;
DROP TABLE IF EXISTS "Locations" CASCADE;
DROP TABLE IF EXISTS "Users" CASCADE;

CREATE TABLE "Users" (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, 
    role VARCHAR(20) CHECK (role IN ('admin', 'staff')) NOT NULL
);

CREATE TABLE "Categories" (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL
);

CREATE TABLE "Locations" (
    location_id SERIAL PRIMARY KEY,
    location_name VARCHAR(100) NOT NULL
);

CREATE TABLE "Items" (
    item_id SERIAL PRIMARY KEY,
    item_code VARCHAR(50) UNIQUE NOT NULL,
    item_name VARCHAR(100) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    category_id INT NOT NULL REFERENCES "Categories"(category_id) ON DELETE CASCADE,
    location_id INT NOT NULL REFERENCES "Locations"(location_id) ON DELETE CASCADE
);

CREATE TABLE "Transactions" (
    transaction_id SERIAL PRIMARY KEY,
    item_id INT NOT NULL REFERENCES "Items"(item_id) ON DELETE CASCADE,
    transaction_type VARCHAR(10) CHECK (transaction_type IN ('in', 'out')) NOT NULL,
    quantity INT NOT NULL,
    timestamp TIMESTAMPTZ DEFAULT NOW(),
    notes TEXT,
    user_id INT REFERENCES "Users"(user_id) ON DELETE SET NULL
);

-- Insert dummy data into Categories table
INSERT INTO "Categories" (category_name) VALUES
('Electronics'),
('Raw Materials'),
('Furniture');

-- Insert dummy data into Locations table
INSERT INTO "Locations" (location_name) VALUES
('Warehouse A'),
('Warehouse B'),
('Warehouse C');

-- Insert dummy data into Users table
INSERT INTO "Users" (username, password, role) VALUES
('admin1', 'hashedpassword1', 'admin'),
('staff1', 'hashedpassword2', 'staff'),
('staff2', 'hashedpassword3', 'staff');

-- Insert dummy data into Items table
INSERT INTO "Items" (item_code, item_name, stock, category_id, location_id) VALUES
('ITEM001', 'Laptop', 50, 1, 1),
('ITEM002', 'TV', 30, 1, 2),
('ITEM003', 'Wood', 100, 2, 3),
('ITEM004', 'Nails', 200, 2, 1),
('ITEM005', 'Table', 40, 3, 2),
('ITEM006', 'Smartphone', 75, 1, 1),
('ITEM007', 'Printer', 20, 1, 3),
('ITEM008', 'Router', 15, 1, 2),
('ITEM009', 'Copper Wire', 250, 2, 1),
('ITEM010', 'Steel Beam', 30, 2, 3),
('ITEM011', 'Desk Chair', 60, 3, 1),
('ITEM012', 'Sofa', 25, 3, 2),
('ITEM013', 'Projector', 10, 1, 1),
('ITEM014', 'Microphone', 40, 1, 2),
('ITEM015', 'Plastic Sheets', 100, 2, 3),
('ITEM016', 'Laptop Stand', 80, 3, 1),
('ITEM017', 'Electric Drill', 22, 2, 2),
('ITEM018', 'File Cabinet', 15, 3, 1),
('ITEM019', 'HDMI Cable', 120, 1, 3),
('ITEM020', 'Battery Pack', 50, 1, 2);

-- Insert dummy data into Transactions table
INSERT INTO "Transactions" (item_id, transaction_type, quantity, timestamp, notes, user_id) VALUES
(1, 'in', 10, '2024-10-01 10:00:00', 'Restock', 1),
(1, 'out', 5, '2024-10-02 11:00:00', 'Shipment to branch', 1),
(2, 'in', 20, '2024-10-03 12:00:00', 'Restock', 1),
(3, 'out', 10, '2024-10-04 09:00:00', 'Shipment to branch B', 2),
(4, 'in', 50, '2024-10-05 14:00:00', 'Stock addition', 2),
(5, 'out', 10, '2024-10-06 15:00:00', 'Shipment to branch C', 1),
(1, 'in', 15, '2024-10-07 16:00:00', 'Additional restock', 1),
(2, 'out', 7, '2024-10-08 17:00:00', 'Replacement for defective unit', 2),
(3, 'in', 40, '2024-10-09 18:00:00', 'Restock for production', 2),
(4, 'out', 20, '2024-10-10 08:00:00', 'Bulk shipment', 1),
(5, 'in', 30, '2024-10-11 09:00:00', 'Restock after sellout', 1),
(1, 'out', 8, '2024-10-12 10:00:00', 'Shipment to customer X', 2),
(2, 'in', 25, '2024-10-13 11:00:00', 'Restock due to increased demand', 1),
(3, 'out', 15, '2024-10-14 12:00:00', 'Shipment of raw material', 2),
(4, 'in', 100, '2024-10-15 13:00:00', 'Large restock', 1),
(5, 'out', 10, '2024-10-16 14:00:00', 'Sales to business partner', 1),
(1, 'in', 5, '2024-10-17 15:00:00', 'Return of demo unit', 2),
(2, 'out', 10, '2024-10-18 16:00:00', 'Replacement for damaged unit', 1),
(3, 'in', 60, '2024-10-19 17:00:00', 'Addition of raw material stock', 2),
(4, 'out', 30, '2024-10-20 18:00:00', 'Bulk shipment to distributor', 1);
