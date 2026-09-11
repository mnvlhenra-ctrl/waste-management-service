create database wms_db;

DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS invoices CASCADE;
DROP TABLE IF EXISTS houses CASCADE;
DROP TABLE IF EXISTS wastes CASCADE;
DROP TABLE IF EXISTS users CASCADE;


-- 1. USERS (Sesuai syarat instruktur)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'USER',
    -- 'ADMIN' atau 'USER'
    balance DECIMAL(12, 2) NOT NULL DEFAULT 0 -- Menyimpan saldo untuk fitur Top Up
);
-- 2. HOUSES / KOMPLEK LISTINGS (Main Entity)
-- Memenuhi syarat: id, name, services, costs, category
-- 1 user = 1 rumah
-- house_number diisi user saat register
CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    house_number VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50) NOT NULL,
    -- services diubah fungsinya menjadi Jadwal Pengambilan
    services VARCHAR(255) NOT NULL DEFAULT 'Senin: Organik, Kamis: Anorganik',
    costs DECIMAL(12, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);
-- 3. INVOICES (Tagihan Bulanan - untuk fitur "Details A udah bayar, B belum")
CREATE TABLE invoices (
    id SERIAL PRIMARY KEY,
    house_id INT NOT NULL,
    billing_month VARCHAR(7) NOT NULL,
    -- Format: 'YYYY-MM' (Contoh: '2026-09')
    amount DECIMAL(12, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'UNPAID',
    -- 'UNPAID', 'PAID'
    due_date TIMESTAMP NOT NULL,
    FOREIGN KEY (house_id) REFERENCES houses(id) ON DELETE CASCADE
);
-- 4. TRANSACTIONS (Sesuai syarat instruktur)
-- Memenuhi syarat: id, name, amount, user_id, transaction_date, status
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    invoice_id INT,
    -- Boleh NULL jika ini adalah transaksi "Top Up"
    name VARCHAR(255) NOT NULL,
    -- Contoh: "Top Up Saldo" atau "Bayar Iuran Sep 2026"
    amount DECIMAL(12, 2) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    -- 'TOP_UP' atau 'PAYMENT'
    status VARCHAR(50) NOT NULL,
    -- 'SUCCESS', 'FAILED'
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (invoice_id) REFERENCES invoices(id)
);

-- 5. WASTES
CREATE TABLE wastes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    stock_availability DECIMAL(12, 2) NOT NULL DEFAULT 0,
    costs DECIMAL(12, 2) NOT NULL,
    category VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);