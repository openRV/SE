CREATE TABLE Doc(
    docsId INTEGER not NULL PRIMARY KEY,
    docsName VARCHAR not NULL,
    docsFile blob,
    author VARCHAR,
    createDate VARCHAR,
    lastUpdate VARCHAR,
    docsType VARCHAR,
    viewCounts INTEGER
);

CREATE TABLE Dir(
    dirId INTEGER not NULL PRIMARY KEY,
    dirName VARCHAR not NULL,
    owner integer,
    lastView VARCHAR,
    createDate VARCHAR,
    subDir INTEGER
);

CREATE TABLE User(
    userName VARCHAR not NULL PRIMARY KEY,
    password VARCHAR not NULL,
    registDate VARCHAR,
    role CHAR(5),
    avatar VARCHAR
);