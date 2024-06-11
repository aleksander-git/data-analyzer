-- +goose Up
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN surename TO surname;
ALTER TABLE datasets RENAME COLUMN dataseset_status_id TO dataset_status_id;
COMMENT ON COLUMN datasets.dataset_status_id IS 'Status id';

ALTER TABLE dataset_statuses RENAME COLUMN dataseset_status_id TO dataset_status_id;
COMMENT ON COLUMN datasets.dataset_status_id IS 'Status id';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users RENAME COLUMN surname TO surename;
ALTER TABLE datasets RENAME COLUMN dataset_status_id TO dataseset_status_id;
ALTER TABLE dataset_statuses RENAME COLUMN dataset_status_id TO dataseset_status_id;

-- +goose StatementEnd