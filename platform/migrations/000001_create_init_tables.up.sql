-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET GLOBAL time_zone ="Asia/Ho_Chi_Minh";

-- Create books items
CREATE TABLE item (
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    title VARCHAR(255) NOT NULL
);
-- Add indexes
CREATE INDEX active_item ON item(id) ;