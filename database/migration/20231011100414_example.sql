-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS example
(
    user_id char(36)     NOT NULL,
    name    VARCHAR(255) NOT NULL,
    age     INT          NOT NULL,
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS example;
-- +goose StatementEnd
