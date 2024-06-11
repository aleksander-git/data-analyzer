
CREATE TABLE databases
(
  database_id bigint        NOT NULL GENERATED ALWAYS AS IDENTITY,
  type        varchar(100)  NOT NULL DEFAULT 'YDB',
  name        varchar(100)  NOT NULL,
  descr       varchar(1000),
  dsn         varchar(1000),
  token       varchar(1000),
  PRIMARY KEY (database_id)
);

COMMENT ON TABLE databases IS 'YDBs';

COMMENT ON COLUMN databases.database_id IS 'Database id (autoincrement)';

COMMENT ON COLUMN databases.type IS 'In case we need to use different storages or versions';

COMMENT ON COLUMN databases.name IS 'Storage name';

COMMENT ON COLUMN databases.descr IS 'Description';

COMMENT ON COLUMN databases.dsn IS 'Connection string to connect to storage';

COMMENT ON COLUMN databases.token IS 'Connection token';

CREATE TABLE dataset_statuses
(
  dataseset_status_id int           NOT NULL,
  name                varchar(100)  NOT NULL,
  descr               varchar(1000),
  PRIMARY KEY (dataseset_status_id)
);

COMMENT ON TABLE dataset_statuses IS 'List of available dataset statuses';

COMMENT ON COLUMN dataset_statuses.dataseset_status_id IS 'Status id';

COMMENT ON COLUMN dataset_statuses.name IS 'Status name (uploading, available, failed)';

COMMENT ON COLUMN dataset_statuses.descr IS 'Status description';

CREATE TABLE datasets
(
  dataset_id          bigint        NOT NULL GENERATED ALWAYS AS IDENTITY,
  user_id             bigint        NOT NULL,
  dataseset_status_id int           NOT NULL,
  database_id         bigint        NOT NULL,
  table_name          varchar(1000),
  is_archieved        boolean       NOT NULL DEFAULT false,
  archieved_at        timestamp    ,
  created_at          timestamp     NOT NULL DEFAULT now(),
  PRIMARY KEY (dataset_id)
);

COMMENT ON TABLE datasets IS 'Datasets';

COMMENT ON COLUMN datasets.dataset_id IS 'Dataset id (autoincrement)';

COMMENT ON COLUMN datasets.user_id IS 'Owner of dataset';

COMMENT ON COLUMN datasets.dataseset_status_id IS 'Status id';

COMMENT ON COLUMN datasets.database_id IS 'Database id';

COMMENT ON COLUMN datasets.table_name IS 'Name of the table';

COMMENT ON COLUMN datasets.is_archieved IS 'Dataset is archieved (deleted)';

COMMENT ON COLUMN datasets.archieved_at IS 'When dataset was archieved';

COMMENT ON COLUMN datasets.created_at IS 'When dataset was created';

CREATE TABLE users
(
  user_id                 bigint       NOT NULL GENERATED ALWAYS AS IDENTITY,
  username                varchar(100) NOT NULL,
  password_hash           varchar(100),
  name                    varchar(100),
  surename                varchar(100),
  email                   varchar(100),
  messengers              varchar(100),
  created_at              timestamp    NOT NULL DEFAULT now(),
  is_archieved            boolean      NOT NULL DEFAULT false,
  archieved_at            timestamp   ,
  is_blocked              boolean      NOT NULL DEFAULT false,
  blocked_at              timestamp   ,
  incorrect_pwd_try_count INTEGER      NOT NULL DEFAULT 0,
  PRIMARY KEY (user_id)
);

COMMENT ON TABLE users IS 'User list';

COMMENT ON COLUMN users.user_id IS 'User Id (autoincrement)';

COMMENT ON COLUMN users.username IS 'Unique user text identificator';

COMMENT ON COLUMN users.password_hash IS 'Password hash';

COMMENT ON COLUMN users.name IS 'Person Name';

COMMENT ON COLUMN users.surename IS 'Person Surname';

COMMENT ON COLUMN users.email IS 'Email address';

COMMENT ON COLUMN users.messengers IS 'Messenger addresses';

COMMENT ON COLUMN users.created_at IS 'When user record was added';

COMMENT ON COLUMN users.is_archieved IS 'User record is archieved (deleted)';

COMMENT ON COLUMN users.archieved_at IS 'When user was archieved';

COMMENT ON COLUMN users.is_blocked IS 'User account is blocked';

COMMENT ON COLUMN users.blocked_at IS 'When user account was blocked';

COMMENT ON COLUMN users.incorrect_pwd_try_count IS 'Number of times user entered incorrect password';

ALTER TABLE datasets
  ADD CONSTRAINT FK_users_TO_datasets
    FOREIGN KEY (user_id)
    REFERENCES users (user_id);

ALTER TABLE datasets
  ADD CONSTRAINT FK_dataset_statuses_TO_datasets
    FOREIGN KEY (dataseset_status_id)
    REFERENCES dataset_statuses (dataseset_status_id);

ALTER TABLE datasets
  ADD CONSTRAINT FK_databases_TO_datasets
    FOREIGN KEY (database_id)
    REFERENCES databases (database_id);

CREATE UNIQUE INDEX users_username_idx
  ON users (username ASC);

CREATE INDEX users_email_idx
  ON users (email ASC);

CREATE INDEX datasets_user_idx
  ON datasets (user_id ASC);

CREATE INDEX datasets_status_idx
  ON datasets (dataseset_status_id ASC);
