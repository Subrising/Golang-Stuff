DROP SCHEMA public CASCADE;

CREATE SCHEMA public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

SELECT * FROM pg_extension;

CREATE TABLE owner(
  owner_id	          UUID        DEFAULT gen_random_uuid(),
  owner_name			    CHAR(50) 		not null,
  PRIMARY KEY (owner_id)
);

CREATE TABLE document(
  document_id          UUID           DEFAULT gen_random_uuid(),
  owner_id	           UUID,
  doc_type			       CHAR(50) 		  not null,
  status               CHAR(50)       not null,
  doc_hash             CHAR(100)      not null,
  PRIMARY KEY (document_id),
  FOREIGN KEY (owner_id) references owner (owner_id)
);

CREATE TABLE page(
  document_id         UUID,
  page_id	            UUID          DEFAULT gen_random_uuid(),
  status              CHAR(50)      not null,
  PRIMARY KEY (page_id),
  FOREIGN KEY (document_id) references document (document_id)
);

CREATE TABLE text_annotation(
  text_id             UUID          DEFAULT gen_random_uuid(),
  page_id	            UUID,
  description			    CHAR(1000) 	  not null,
  locale			        CHAR(50) 		  not null,
  xcoMin                FLOAT         not null,
  xcoMax                FLOAT         not null,
  ycoMin                FLOAT         not null,
  ycoMax                FLOAT         not null,
  PRIMARY KEY (text_id),
  FOREIGN KEY (page_id) references page (page_id)
);

CREATE TABLE label_annotation(
  label_id            UUID          DEFAULT gen_random_uuid(),
  page_id	            UUID,
  description			    CHAR(200) 		not null,
  score			          FLOAT 		  not null,
  PRIMARY KEY (label_id),
  FOREIGN KEY (page_id) references page (page_id)
);

CREATE TABLE crop_hint(
  crop_id             UUID          DEFAULT gen_random_uuid(),
  page_id	            UUID,
  confidence			    CHAR(50)      not null,
  importance_fraction FLOAT		      not null,
  xcoMin                FLOAT         not null,
  xcoMax                FLOAT         not null,
  ycoMin                FLOAT         not null,
  ycoMax                FLOAT         not null,
  PRIMARY KEY (crop_id),
  FOREIGN KEY (page_id) references page (page_id)
);

CREATE TABLE logs(
  log_id             UUID         DEFAULT gen_random_uuid(),
  page_id            UUID,
  type               char(50),
  time               TIMESTAMP,
  PRIMARY KEY (log_id),
  FOREIGN KEY (page_id) references page (page_id)
);

SELECT * FROM Owner;


