-- +goose Up
ALTER TABLE results ADD CONSTRAINT unique_params_ad UNIQUE (params_id, ad_id);

-- +goose Down
ALTER TABLE results DROP CONSTRAINT unique_params_ad;