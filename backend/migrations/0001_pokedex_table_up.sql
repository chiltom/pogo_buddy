create table pokedex (
    id varchar(255) primary key,
    formId varchar(255) not null,
    dex_num int not null,
    generation int not null,
    name varchar(255) not null,
    primary_type varchar(255) not null,
    secondary_type varchar(255)
);