ALTER TABLE users
    ALTER COLUMN dob
        TYPE INTEGER
        USING  extract(epoch from dob)::INTEGER;