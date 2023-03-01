CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTOINCREMENT NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL CHECK(LENGTH(email) <= 50),
    name VARCHAR(32) UNIQUE NOT NULL CHECK(LENGTH(nick) <= 32),
    hashpass VARCHAR(50) NOT NULL CHECK(LENGTH(hashpass) = 32)
    );

CREATE TABLE IF NOT EXISTS communities (
    id INT PRIMARY KEY AUTOINCREMENT NOT NULL,

    title TEXT NOT NULL CHECK(LENGTH(title) <= 50),
    description TEXT NOT NULL,

    UNIQUE(title),
    );

CREATE TABLE IF NOT EXISTS subscribers (
    user_id INT NOT NULL,
    community_id INT NOT NULL DEFAULT 0,
    permission INT NOT NULL DEFAULT 3,
    );

CREATE TABLE IF NOT EXISTS posts (
    id INT PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INT NOT NULL,
    community_id INT NOT NULL DEFAULT 0,

    created_at DATE NOT NULL DEFAULT SELECT TIME('now'),
    upgrated_at DATE NOT NULL DEFAULT SELECT TIME('now'),

    title TEXT NOT NULL CHECK(LENGTH(title) <= 50),
    content TEXT NOT NULL,

    version INT NOT NULL DEFAULT 1,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (community_id) REFERENCES community(id)
    );

CREATE TABLE IF NOT EXISTS comments (
    id INT PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id INT NOT NULL,
    user_id INT NOT NULL,

    created_at DATE NOT NULL DEFAULT SELECT TIME('now'),

    content TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
    );

CREATE TABLE IF NOT EXISTS votes (
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    UNIQUE (post_id, user_id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id));

PRAGMA foreign_keys = ON;