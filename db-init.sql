CREATE TABLE IF NOT EXISTS items (
    -- GORM standard fields
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,

    -- Custom fields
    name VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL
);

-- Создание индексов для улучшения производительности поиска
CREATE INDEX idx_items_deleted_at ON items (deleted_at);
CREATE INDEX idx_items_name ON items (name);