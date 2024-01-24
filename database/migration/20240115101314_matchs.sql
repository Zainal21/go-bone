-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS matchs
(
    id                CHAR(36)     NOT NULL,
    team1             VARCHAR(255) NOT NULL,
    team2             VARCHAR(255) NOT NULL,
    match_date_time   timestamp NOT NULL,
    created_at        timestamp NULL DEFAULT NULL,
    updated_at        timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS matchs;
-- +goose StatementEnd
