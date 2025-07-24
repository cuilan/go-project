CREATE DATABASE `test_db` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

grant alter, alter routine, create, create routine, create temporary tables, create view, delete, drop,
event, execute, index, insert, lock tables, references, select, show view, trigger, update on test_db.* to root;
