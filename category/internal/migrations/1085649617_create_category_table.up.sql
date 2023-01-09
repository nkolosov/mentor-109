-- CREATE TABLE category (
--     id         int,
--     name       varchar(100) NOT NULL UNIQUE,
--     is_delete  BOOLEAN,
--     created_at DATETIME  DEFAULT CURRENT_TIMESTAMP,
--     delete_at  TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- );

-- CREATE TABLE category (
--                        user_id int NOT NULL AUTO_INCREMENT,
--                        name varchar(255) NOT NULL,
--                        fullname varchar(255),
--                        balance int,
--                        PRIMARY KEY (user_id)
-- );

CREATE TABLE IF NOT EXISTS category (
    id INT NOT NULL,
    name varchar(100) NOT NULL,
    is_delete  BOOLEAN,
    created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    delete_at  TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- id: INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
-- + name: varchar(100) NOT NULL
-- + slug: varchar(100) NOT NULL UNIQUE
-- + is_delete: BOOLEAN
-- + created_at: DATETIME DEFAULT CURRENT_TIMESTAMP
-- + delete_at: TIMESTAMP
-- + updated_at: TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP