-- +goose Up
-- +goose StatementBegin
INSERT INTO public.dataset_statuses(
	dataseset_status_id, name, descr)
	VALUES (1, 'new', 'New dataset, not ready');
INSERT INTO public.dataset_statuses(
	dataseset_status_id, name, descr)
	VALUES (2, 'uploading', 'Dataset is being uploaded');
	
INSERT INTO public.dataset_statuses(
	dataseset_status_id, name, descr)
	VALUES (3, 'available', 'Dataset ready to use');
	
INSERT INTO public.dataset_statuses(
	dataseset_status_id, name, descr)
	VALUES (4, 'failed', 'Dataset is in error state');
	
	
	
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM dataset_statuses where dataseset_status_id in (1,2,3,4);
-- +goose StatementEnd
