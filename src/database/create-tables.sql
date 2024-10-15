DROP TABLE IF EXISTS url_mapping;
CREATE TABLE url_mapping (
    id              INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    short_url       VARCHAR(255) NOT NULL UNIQUE,
    long_url        TEXT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accessed_at     TIMESTAMP NULL DEFAULT NULL,
    access_count    INT DEFAULT 0
);