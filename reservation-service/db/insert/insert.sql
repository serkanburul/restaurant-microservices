INSERT INTO time_slots (start_time, end_time) VALUES
                                                  ('15:30:00', '16:45:00'),
                                                  ('16:45:00', '18:00:00'),
                                                  ('18:00:00', '19:15:00'),
                                                  ('19:15:00', '20:30:00'),
                                                  ('20:30:00', '21:45:00'),
                                                  ('21:45:00', '23:00:00');


DO $$
    DECLARE
        i INT;
        j INT;
    BEGIN
        FOR i IN 1..5 LOOP
            FOR j IN 1..10 LOOP
                INSERT INTO tables (capacity) VALUES (i);
            END LOOP;
        END LOOP;
    END $$;