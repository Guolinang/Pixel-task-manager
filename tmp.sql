CREATE TABLE if not exists tasks (
    userid integer,
    id serial primary key,
    isImportant boolean,
    taskname varchar,
    difficulty integer,
    sdescription text,
    tasktype varchar,
    taskstats varchar,
    deadline date,
    taskrepeat  varchar,
    subtask varchar,
    fdescription text,
    done boolean,
    foreign key (userid) references users (id)
);

INSERT INTO tasks (  
    userid ,   
    isImportant ,
    taskname ,
    difficulty ,
    sdescription ,
    tasktype ,
    taskstats ,
    deadline ,
    taskrepeat  ,
    subtask ,
    fdescription ,
    done
    ) VALUES (
        2,        
        'false',
        'pied piper',
        1,
        'super coolь compression algorithm',
        'cod',
        'charisma',
        '2023-11-02',
        '0010001' ,
        'first|second',
        'gvthrfcbjdnkmkjfnmvdlc ,.xs;.za/.dvflkmhn bnkjmvlaQA;/;.BL,GKM',
        'false'
    ), (

        2,        
        'true',
        'hooli',
        3,
        'asdasdasdression algorithm',
        'caod',
        'money',
        '2070-11-12',
        '1011011' ,
        'first',
        'rfdgtcvhbjnkml,;.',
        'true'


    )