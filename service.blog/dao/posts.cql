CREATE KEYSPACE IF NOT EXISTS blog
WITH REPLICATION = {
    'class': 'SimpleStrategy',
    'replication_factor' : 1
};

CREATE TABLE IF NOT EXISTS blog.posts (
    id text,
    title text,
    body text,
    author text,
    created timestamp,
    updated timestamp,
    PRIMARY KEY ((id))
);

