create table
  users (
    id uuid not null,
    email varchar(255) not null,
    password text not null,
    role varchar(255) not null,
    primary key (id),
    constraint email_unique unique (email)
  );

create table
  posts (
    id uuid not null,
    title text not null,
    content text not null,
    created_by_id uuid not null,
    comments_count int not null default 0,
    created_at timestamp not null default now(),
    updated_at timestamp null,
    deleted_at timestamp null,
    primary key (id),
    constraint fk_created_by_id foreign key (created_by_id) references users (id)
  );

create table
  comments (
    id uuid not null,
    content varchar(512) not null,
    post_id uuid not null,
    likes_count int not null default 0,
    created_by_id uuid not null,
    parent_comment_id uuid null,
    hierarchy_id uuid not null,
    created_at timestamp not null default now(),
    updated_at timestamp null,
    deleted_at timestamp null,
    primary key (id),
    constraint fk_post_id foreign key (post_id) references posts (id),
    constraint fk_created_by_id foreign key (created_by_id) references users (id),
    constraint fk_parent_comment_id foreign key (parent_comment_id) references comments (id)
  );

create table
  comment_likes (
    comment_id uuid not null,
    user_id uuid not null,
    primary key (comment_id, user_id),
    constraint fk_comment_id foreign key (comment_id) references comments (id),
    constraint fk_user_id foreign key (user_id) references users (id)
  );