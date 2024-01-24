INSERT INTO users (
    id, 
    name, 
    email, 
    password, 
    status,
    role
) VALUES (UUID(), 'administrator', 'admin@mindpredictor.com','$2y$10$pmU8V7pCRnrrVJMQG8fqYuDH92V2pE7HQo.b3BKhraA9s3CRur8Jy', 'active', 'admin'), 
        (UUID(), 'example_user', 'user@mindpredictor.com','$2y$10$pmU8V7pCRnrrVJMQG8fqYuDH92V2pE7HQo.b3BKhraA9s3CRur8Jy', 'active', 'user');
