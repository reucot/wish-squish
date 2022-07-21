-- Роли пользователей
CREATE TABLE roles
(
    id   serial       not null primary key,
    name varchar(255) not null
);

-- Пользователи
CREATE TABLE users
(
    id            serial                                         not null primary key,
    name          varchar(255)                                   not null,
    surname       varchar(255)                                   not null,
    email         varchar(255)                                   not null,
    password_hash varchar(255)                                   not null,
    created_at    timestamp                                      not null default CURRENT_TIMESTAMP,
    updated_at    timestamp                                      not null default CURRENT_TIMESTAMP,
    verified_at   timestamp,
    role_id       int constraint fk_roles references roles (id) on delete cascade not null
);

-- Настройки пользователя
CREATE TABLE user_settings
(
    id              serial                                         not null primary key,
    user_id         int constraint fk_users references users (id)  on delete cascade not null,
    handle          varchar(255)                                   not null,
    wish_list_name  varchar(255)                                   not null,
    alias           varchar(255)                                   not null
);

-- Желания
CREATE TABLE wishes
(
    id              serial                                         not null primary key,
    user_id         int constraint fk_users references users (id)  on delete cascade not null,
    name            varchar(255)                                   not null,
    price           varchar(255)                                   not null,
    is_allow_repeat boolean                                        not null default false,
    created_at      timestamp                                      not null default CURRENT_TIMESTAMP,
    updated_at      timestamp                                      not null default CURRENT_TIMESTAMP
);

-- Категории желаний
CREATE TABLE categories
(
    id              serial                                         not null primary key,
    user_id         int constraint fk_users references users (id)  on delete cascade not null,
    name            varchar(255)                                   not null,
    created_at      timestamp                                      not null default CURRENT_TIMESTAMP,
    updated_at      timestamp                                      not null default CURRENT_TIMESTAMP
);

-- Категория-желание
CREATE TABLE wishes_categories
(
    id              serial                                                   not null primary key,
    wish_id         int constraint fk_wishes     references wishes     (id)  on delete cascade not null,
    category_id     int constraint fk_categories references categories (id)  on delete cascade not null
);

-- Движение денежных средств
CREATE TABLE transaction
(
    id              serial                                                   not null primary key,
    wish_id         int constraint fk_wishes     references wishes     (id)  on delete cascade not null,
    category_id     int constraint fk_categories references categories (id)  on delete cascade not null
);

-- Движение денежных средств
CREATE TABLE photos
(
    id              serial           not null primary key,
    model_id        int              not null,
    name            varchar(255)     not null,
    path            varchar(255)     not null,
    type            varchar(255)     not null,
    created_at      timestamp        not null default CURRENT_TIMESTAMP,
    updated_at      timestamp        not null default CURRENT_TIMESTAMP
);