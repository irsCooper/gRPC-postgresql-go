INSERT INTO apps (id, name, seacret)
VALUES (1, 'test', 'test-secret')
ON CONFLICT DO NOTHING;