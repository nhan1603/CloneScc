INSERT INTO public.users (id, display_name, email, password, user_role) VALUES
(250, 'UTest1', 'utest1@example.com', '123456', 'SECURITY_GUARD'),
(251, 'UTest2', 'utest2@example.com', '123456', 'SECURITY_GUARD'),
(252, 'UTest3', 'utest3@example.com', '123456', 'SECURITY_GUARD');

INSERT INTO public.device_tokens (id, user_id, device_token, platform) VALUES
(150, 250, 'qUsCeUjQ5Z6010', 'ios'),
(151, 251, 'eUjQ5Z6010qUsC', 'ios');
