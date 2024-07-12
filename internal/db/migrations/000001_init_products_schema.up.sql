DROP TYPE IF EXISTS categoryType CASCADE;
CREATE TYPE categoryType AS ENUM (
    'Lanche', 'Acompanhamento', 'Bebida', 'Sobremesa'
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL,
    name TEXT NOT NULL,
    category categoryType NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT NOT NULL,
    is_available BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();