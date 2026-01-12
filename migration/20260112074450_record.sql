-- +goose Up
-- +goose StatementBegin
Create table record(
    id serial not null,
    user_id uuid not null,
    service_name text not null,
    price int not null,
    start_date date not null,
    primary key(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop table record;
-- +goose StatementEnd
