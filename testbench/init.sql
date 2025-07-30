-- init.sql
-- This script will be executed when MySQL container starts

-- Create the performance test database
CREATE DATABASE IF NOT EXISTS performance_test;

-- Use the database
USE performance_test;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY,
    profile_data VARCHAR(255) NOT NULL
);

-- Insert sample data
INSERT IGNORE INTO users (id, profile_data) VALUES 
    (1, 'user1_profile_data_with_some_content'),
    (2, 'user2_profile_data_with_different_content'),
    (3, 'user3_profile_data_with_more_content'),
    (4, 'user4_profile_data_with_additional_content'),
    (5, 'user5_profile_data_with_extended_content');

-- Grant privileges (optional, but good practice)
GRANT ALL PRIVILEGES ON performance_test.* TO 'root'@'%';
FLUSH PRIVILEGES;
