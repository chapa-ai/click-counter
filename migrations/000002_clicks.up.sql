CREATE TABLE clicks (
    banner_id INT REFERENCES banners(id) ON DELETE CASCADE,
    timestamp TIMESTAMPTZ NOT NULL,
    count INT DEFAULT 0,
    UNIQUE (banner_id, timestamp)
);