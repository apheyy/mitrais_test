CREATE TABLE IF NOT EXISTS employee (
    id serial not null,
    first_name varchar(255),
    last_name varchar(255),
    email varchar(255),
    hire_date timestamp,
    constraint employee_pkey primary key (id),
);
