-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_points
(
    id                  CHAR(36) NOT NULL,
    user_id             CHAR(36) NOT NULL,
    point               INT NOT NULL,
    created_at          timestamp NULL DEFAULT NULL,
    updated_at          timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_points;
-- +goose StatementEnd
