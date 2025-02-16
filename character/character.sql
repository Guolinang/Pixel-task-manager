CREATE TABLE if not exists character (
    userid integer primary key,
    levelC integer,
    expC integer,
    maxexpC integer,
    hpC integer,
    maxhpC integer,
    strC integer,
    intC integer,
    charC integer,
    wisC integer,
    cnstC integer,
    foreign key (userid) references users (id)
)


