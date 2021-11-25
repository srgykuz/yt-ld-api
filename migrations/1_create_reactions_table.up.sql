CREATE TABLE IF NOT EXISTS reactions (
    video_id TEXT PRIMARY KEY,
    likes_count INT NOT NULL DEFAULT 0,
    dislikes_count INT NOT NULL DEFAULT 0
);
