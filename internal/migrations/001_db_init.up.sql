create table urls(
    id BIGINT,
    identifier text not null unique,
    real_url text not null,
    usages INTEGER not null default 0,
    created_at timestamp not null default now(),
    expires_at timestamp not null default now()
    );