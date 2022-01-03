-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Category (
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Category;
-- +goose StatementEnd
