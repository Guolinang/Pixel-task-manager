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
    --     2,        
    --     'false',
    --     'pied piper',
    --     1,
    --     'super coolь compression algorithm',
    --     'cod',
    --     'charisma',
    --     '2023-11-02',
    --     '0010001' ,
    --     'first|second',
    --     'gvthrfcbjdnkmkjfnmvdlc ,.xs;.za/.dvflkmhn bnkjmvlaQA;/;.BL,GKM',
    --     'false'
    -- ), (

    --     2,        
    --     'true',
    --     'hooli',
    --     3,
    --     'asdasdasdression algorithm',
    --     'caod',
    --     'money',
    --     '2070-11-12',
    --     '1011011' ,
    --     'first',
    --     'rfdgtcvhbjnkml,;.',
    --     'true'

    
         2,        
        'true',
        'hooli22',
        3,
        'xbcvbcvbm',
        'caxvod',
        'str',
        '2025-02-15',
        '1011011' ,
        'first',
        'rfdgtcvhbjnkml,;.',
        'false'    
    ),(
        2,        
        'false',
        'hooli333',
        2,
        'hfgddfas',
        'aaaaaaaacaxvod',
        'str',
        '2025-02-15',
        '1111111' ,
        'asdas',
        'aaaaarfdgfasfatcvhbjnkml,;.',
        'false'      
    ),(
  2,        
        'true',
        'aaaahooli333',
        2,
        'hsdafgddfas',
        'aaaaaaaacaxvod',
        'str',
        '2021-02-15',
        '1111111' ,
        'asdas',
        'aaaaarfdgfasfatcvhbjnkml,;.',
        'false'   
    ),
    (
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
        'false'
    )
    