CREATE TABLE IF NOT EXISTS tables (
       id SERIAL PRIMARY KEY,
       capacity INTEGER NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS time_slots (
        id SERIAL PRIMARY KEY ,
        start_time TIME NOT NULL ,
        end_time TIME NOT NULL ,
        UNIQUE (start_time, end_time)
);

CREATE TABLE IF NOT EXISTS reservation (
       id SERIAL PRIMARY KEY,
       token VARCHAR(10) NOT NULL UNIQUE,
       name VARCHAR NOT NULL,
       email VARCHAR NOT NULL,
       status VARCHAR NOT NULL,
       table_no INTEGER REFERENCES tables (id) ON DELETE CASCADE,
       time_slot_id INTEGER REFERENCES time_slots (id) ON DELETE CASCADE ,
       reservation_date DATE NOT NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1
            FROM pg_constraint
            WHERE conname = 'unique_table_reservation'
        ) THEN
            ALTER TABLE reservation
                ADD CONSTRAINT unique_table_reservation UNIQUE (reservation_date, table_no, time_slot_id);
        END IF;
    END $$;