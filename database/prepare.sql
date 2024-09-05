CREATE TABLE IF NOT EXISTS dicks (
    id serial primary key,
    username text NOT NULL,
    peer_id bigint NOT NULL,
    user_id bigint NOT NULL,
    dick_size bigint default 0,
    issued_at timestamp,
    created_at timestamp default now()
)