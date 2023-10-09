-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions
(
    id                   serial PRIMARY KEY,
    user_vk_id           int    NOT NULL,
    sessions_date        timestamp with time zone,
    solutions            text
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
