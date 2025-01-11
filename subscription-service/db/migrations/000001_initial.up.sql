CREATE TABLE subscription (
      id SERIAL PRIMARY KEY,
      email VARCHAR(50) NOT NULL UNIQUE,
      status BOOLEAN NOT NULL DEFAULT true,
      started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      ended_at TIMESTAMP
);
