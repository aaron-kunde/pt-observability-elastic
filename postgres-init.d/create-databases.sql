CREATE database app1;
GRANT ALL PRIVILEGES ON DATABASE app1 TO postgres;
CREATE database app2;
GRANT ALL PRIVILEGES ON DATABASE app2 TO postgres;
CREATE database app3;
GRANT ALL PRIVILEGES ON DATABASE app3 TO postgres;
CREATE database app4;
GRANT ALL PRIVILEGES ON DATABASE app4 TO postgres;
\c app4
CREATE TABLE public.data_entities (id serial, data varchar(255), PRIMARY KEY (id));