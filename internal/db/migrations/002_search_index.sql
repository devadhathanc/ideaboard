CREATE INDEX IF NOT EXISTS idx_tasks_fts
    ON tasks
    USING GIN (to_tsvector('english', title));
