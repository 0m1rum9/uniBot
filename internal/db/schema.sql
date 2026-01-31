
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username TEXT,
    chat_id BIGINT,
    last_message_id BIGINT,
    login TEXT,
    password TEXT
);
