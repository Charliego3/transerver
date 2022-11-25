CREATE TABLE public.regions
(
    id   SERIAL8 PRIMARY KEY NOT NULL,
    code VARCHAR(5)          NOT NULL,
    area VARCHAR(5)          NOT NULL,
    img  VARCHAR             NOT NULL,
    name JSONB               NOT NULL
);
CREATE UNIQUE INDEX regions_code_key ON regions USING BTREE (code);
CREATE INDEX regions_area_key ON regions USING BTREE (area);


CREATE TABLE public.accounts
(
    id        SERIAL8 PRIMARY KEY NOT NULL,
    create_at timestamptz         NOT NULL,
    update_at timestamptz,
    user_id   VARCHAR(21)         NOT NULL,
    region    VARCHAR(5)          NOT NULL,
    username  VARCHAR(30)         NOT NULL,
    area      VARCHAR(5)          NOT NULL,
    phone     VARCHAR(20),
    email     VARCHAR(50),
    avatar    VARCHAR(200),
    password  BYTEA               NOT NULL,
    pwd_level SMALLINT            NOT NULL DEFAULT 0,
    platform  VARCHAR(20)         NOT NULL,
    state     SMALLINT            NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX accounts_user_id_key ON accounts USING BTREE (user_id);
CREATE UNIQUE INDEX accounts_phone_key ON accounts USING BTREE (phone);
CREATE UNIQUE INDEX accounts_email_key ON accounts USING BTREE (email);
