
create table tasks (
    id serial primary key,
    title text not null,
    description text,
    status text check (status in ('new', 'in_progress', 'done')) default 'new',
    created_at timestamp default now(),
    updated_at timestamp default now()
);