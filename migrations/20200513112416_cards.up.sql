CREATE TABLE cards (
    id bigserial not null primary key,
    digits varchar not null unique,
    cvv int not null,
    card_date varchar not null,
    card_owner varchar not null
);