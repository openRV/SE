CREATE TABLE Doc(
    docsId VARCHAR not NULL PRIMARY KEY,
    docsName VARCHAR not NULL,
    docsFile blob,
    author VARCHAR,
    createDate VARCHAR,
    lastUpdate VARCHAR,
    docsType VARCHAR,
    viewCounts INTEGER,
    open BOOLEAN
);

CREATE TABLE Dir(
    dirId VARCHAR not NULL PRIMARY KEY,
    dirName VARCHAR not NULL,
    owner VARCHAR not NULL,
    createDate VARCHAR,
    lastView VARCHAR
);

CREATE TABLE Tree(
    dirId VARCHAR not NULL,
    root BOOLEAN not NULL,
    subType VARCHAR not NULL, -- dir | doc
    subId VARCHAR not NULL -- sub dir or sub doc id
);

CREATE TABLE History(
    userName VARCHAR not NULL,
    operator VARCHAR not NULL, -- delete | view | download | move
    opetatorType VARCHAR not NULL, -- dir | doc
    itemId VARCHAR not NULL
);

CREATE TABLE Trash(
    itemType VARCHAR , -- dir | doc
    itemId VARCHAR,
    owner VARCHAR,
    PRIMARY KEY(itemType, itemId)
);

CREATE TABLE User(
    userName VARCHAR not NULL PRIMARY KEY,
    password VARCHAR not NULL,
    registDate VARCHAR,
    role CHAR(5),--区分管理员与普通用户
    avatar VARCHAR
);