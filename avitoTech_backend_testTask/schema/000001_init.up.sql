CREATE TABLE users
(
    id            int       not null unique,
    balance       float        not null
);

CREATE TABLE orders
(
    id            int       not null unique,
    user_id    int references users (id) on delete cascade      not null,
    balance       float        not null
);

CREATE TABLE reserves
(
    id            int      not null unique,
    user_id int references users (id) on delete cascade      not null,
    order_id int references orders (id) on delete cascade      not null,
    balance       float        not null
);

