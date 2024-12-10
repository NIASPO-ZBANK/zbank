-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS money
(
    user_id VARCHAR(36) PRIMARY KEY NOT NULL DEFAULT 'GOLBUTSA-1337-1487-911Z-Salla4VO2022',
    money   INT                     NOT NULL DEFAULT 0,
    deposit INT                     NOT NULL DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS money;
-- +goose StatementEnd
