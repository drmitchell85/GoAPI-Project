IF NOT EXISTS(SELECT * FROM pg_database WHERE name = 'projectonedb')
BEGIN
    CREATE DATABASE projectonedb;
END;
GO