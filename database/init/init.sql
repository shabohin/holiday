CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS holidays (
    id uuid DEFAULT uuid_generate_v4() constraint users_pk primary key,
    day date NOT NULL,
    event text DEFAULT '' NOT NULL,
    country varchar(255) DEFAULT '',
    is_event boolean DEFAULT false,
    created_at timestamp NOT NULL DEFAULT (now() at time zone 'utc')
);

INSERT INTO holidays VALUES (default, '0001-10-13', 'День без бюстгальтера (No Bra Day)', null, default, default);
INSERT INTO holidays VALUES (default, '0001-10-13', 'День «Побалуйте себя» (Treat Yo` Self Day)', null, default, default);