create table users (
  id uuid not null,
  email varchar(255) not null,
  password text not null,
  role varchar(255) not null,
  primary key (id),
  constraint email_unique unique (email)
);

