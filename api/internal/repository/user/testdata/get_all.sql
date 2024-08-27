TRUNCATE TABLE users CASCADE;
INSERT INTO users(id, display_name, email, password, user_role)
VALUES (100, 'John', 'john@scc.com', 'test', 'OPERATION_USER');
