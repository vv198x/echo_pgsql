ALTER TABLE users
    ALTER COLUMN dob
        TYPE timestamp with time zone
        USING to_timestamp(dob);