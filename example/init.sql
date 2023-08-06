create table if not exists users(
    id serial primary key ,
    name text not null
);

create table if not exists profiles(
    id serial primary key ,
    user_name text not null,
    user_id int not null,
    foreign key (user_id) references users(id)

);


insert into users (id, name) values(1, 'test1');
insert into users (id, name) values(2,'test2');
insert into profiles (user_name,user_id) values('test_user_1',1);
insert into profiles (user_name,user_id) values('test_user_2',2);