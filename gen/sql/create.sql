--- Tag
CREATE TABLE IF NOT EXISTS tag (
  tag_id SERIAL PRIMARY KEY,
  tag_name varchar(100) NOT NULL,

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

--- Writer
CREATE TABLE IF NOT EXISTS writer (
  writer_id SERIAL PRIMARY KEY,
  writer_username varchar(100) NOT NULL,
  writer_email varchar(400) NOT NULL,

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

--- Editor
CREATE TABLE IF NOT EXISTS editor (
  editor_id SERIAL PRIMARY KEY,
  editor_username varchar(100) NOT NULL,
  editor_email varchar(400) NOT NULL,

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

--- Webmaster
CREATE TABLE IF NOT EXISTS webmaster (
  webmaster_id SERIAL PRIMARY KEY,
  webmaster_username varchar(100) NOT NULL,
  webmaster_email varchar(400) NOT NULL,

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

--- Website
CREATE TABLE IF NOT EXISTS website (
  website_id SERIAL PRIMARY KEY,
  website_url varchar(2048) NOT NULL,
  webmaster_id int,
  CONSTRAINT fk_webmaster FOREIGN KEY(webmaster_id) REFERENCES webmaster(webmaster_id),

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);
--- Host Admin
CREATE TABLE IF NOT EXISTS hosting_admin (
  hosting_admin_id SERIAL PRIMARY KEY,
  hosting_admin_username varchar(100) NOT NULL,
  hosting_admin_email varchar(400) NOT NULL,

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

--- Hosting
CREATE TABLE IF NOT EXISTS hosting (
  hosting_id SERIAL PRIMARY KEY,
  hosting_name varchar(2048) NOT NULL,

  hosting_admin_id int,
  CONSTRAINT fk_hosting_admin FOREIGN KEY(hosting_admin_id) REFERENCES hosting_admin(hosting_admin_id),

  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);


--- Blog
CREATE TABLE IF NOT EXISTS blog (
  blog_id SERIAL PRIMARY KEY,
  blog_title varchar(100) NOT NULL,
  blog_url varchar(2048) NOT NULL UNIQUE,
  blog_status varchar(100) NOT NULL,
  blog_approved_at timestamp,
  blog_published_at timestamp,

  blog_writer_id int,
  CONSTRAINT fk_blog_writer FOREIGN KEY(blog_writer_id) REFERENCES writer(writer_id),

  blog_editor_id int,
  CONSTRAINT fk_blog_editor FOREIGN KEY(blog_editor_id) REFERENCES editor(editor_id),

  blog_website_id int,
  CONSTRAINT fk_blog_website FOREIGN KEY(blog_website_id) REFERENCES website(website_id),

  blog_created_at timestamp NOT NULL DEFAULT NOW(),
  blog_updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS blog_tag (
  tag_id int,
  blog_id int,

  primary key (tag_id, blog_id),

  CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tag(tag_id),
  CONSTRAINT fk_blog FOREIGN KEY(blog_id) REFERENCES blog(blog_id)
);
