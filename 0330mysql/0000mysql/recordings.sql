create database if not exists recordings;

DROP TABLE IF EXISTS album;
CREATE TABLE album
(
    id     INT AUTO_INCREMENT NOT NULL,
    title  VARCHAR(128)       NOT NULL,
    artist VARCHAR(255)       NOT NULL,
    price  DECIMAL(5, 2)      NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO album (title, artist, price)
VALUES ('Blue Train', 'John Coltrane', 56.99),
       ('Giant Steps', 'John Coltrane', 63.99),
       ('Jeru', 'Gerry Mulligan', 17.99),
       ('Sarah Vaughan', 'Sarah Vaughan', 34.98);

alter table album
    add column `quantity` int not null default 0 comment '数量';

update album
set album.quantity = 5
where id = 3;

select *
from album;

drop table if exists album_order;

create table album_order
(
    id         INT AUTO_INCREMENT NOT NULL,
    album_id   int                not null,
    custom_id  int                not null,
    quantity   int                not null,
    created_at datetime default current_timestamp,
    updated_at datetime default current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    index idx_created_at (`created_at`),
    index idx_updated_at (`updated_at`)
);

select *
from album_order;

show tables;


