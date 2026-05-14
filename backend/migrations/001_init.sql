CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS figures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    name VARCHAR(255) NOT NULL,
    character_name VARCHAR(255),
    series_name VARCHAR(255),
    manufacturer VARCHAR(255),
    figure_type VARCHAR(100),
    scale VARCHAR(50),

    status VARCHAR(50) NOT NULL DEFAULT 'wishlist',

    price NUMERIC(12, 2) DEFAULT 0,
    deposit NUMERIC(12, 2) DEFAULT 0,
    balance NUMERIC(12, 2) DEFAULT 0,

    preorder_start_date DATE,
    preorder_deadline DATE,
    release_date DATE,
    payment_due_date DATE,
    arrival_date DATE,

    shop_name VARCHAR(255),
    note TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT figures_status_check CHECK (
        status IN (
            'wishlist',
            'preordered',
            'deposit_paid',
            'balance_due',
            'paid',
            'shipped',
            'arrived',
            'cancelled',
            'sold'
        )
    )
);

CREATE INDEX IF NOT EXISTS idx_figures_user_id ON figures(user_id);
CREATE INDEX IF NOT EXISTS idx_figures_status ON figures(status);
CREATE INDEX IF NOT EXISTS idx_figures_payment_due_date ON figures(payment_due_date);
CREATE INDEX IF NOT EXISTS idx_figures_release_date ON figures(release_date);