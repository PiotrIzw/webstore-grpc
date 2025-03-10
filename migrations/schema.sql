CREATE TABLE accounts (
      id SERIAL PRIMARY KEY,
      username TEXT NOT NULL,
      email TEXT NOT NULL UNIQUE,
      hashed_password TEXT NOT NULL,
      status TEXT NOT NULL DEFAULT 'ACTIVE',
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE preferences (
     user_id INT PRIMARY KEY REFERENCES accounts(id) ON DELETE CASCADE,
     theme TEXT DEFAULT 'light',
     notifications BOOLEAN DEFAULT true,
     locale TEXT DEFAULT 'en'
);

CREATE TABLE roles (
       id SERIAL PRIMARY KEY,
       name TEXT NOT NULL UNIQUE,
       permissions TEXT[] NOT NULL
);

-- Insert default roles
INSERT INTO roles (name, permissions) VALUES
      ('admin', ARRAY['accounts:read', 'accounts:write', 'orders:read', 'orders:write', 'preferences:read', 'preferences:write']),
      ('user', ARRAY['accounts:read', 'orders:read', 'preferences:read']),
      ('guest', ARRAY['accounts:read']);

CREATE TABLE user_roles (
        user_id INT REFERENCES accounts(id) ON DELETE CASCADE,
        role_id INT REFERENCES roles(id) ON DELETE CASCADE,
        PRIMARY KEY (user_id, role_id)
);

CREATE TABLE orders (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES accounts(id) ON DELETE CASCADE,
        total DECIMAL(10, 2) NOT NULL,
        status TEXT NOT NULL DEFAULT 'PENDING',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
         order_id INT REFERENCES orders(id) ON DELETE CASCADE,
         product_id TEXT NOT NULL,
         quantity INT NOT NULL,
         price DECIMAL(10, 2) NOT NULL,
         PRIMARY KEY (order_id, product_id)
);

CREATE TABLE files (
       id TEXT PRIMARY KEY,
       name TEXT NOT NULL,
       type TEXT NOT NULL,
       size BIGINT NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);