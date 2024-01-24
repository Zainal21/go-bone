-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS history_matchs
(
    id                  CHAR(36) NOT NULL,
    user_id             CHAR(36) NOT NULL,
    match_id            CHAR(36) NOT NULL,
    predicted_result    VARCHAR(10)  NOT NULL,
    actual_result       VARCHAR(10)  NOT NULL,
    created_at          timestamp NULL DEFAULT NULL,
    updated_at          timestamp NULL DEFAULT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS history_matchs;
-- +goose StatementEnd
