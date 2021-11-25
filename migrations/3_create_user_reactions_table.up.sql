CREATE TABLE IF NOT EXISTS user_reactions (
    user_id INT NOT NULL REFERENCES users(id),
    video_id TEXT NOT NULL,
    has_like BOOLEAN NOT NULL DEFAULT false,
    has_dislike BOOLEAN NOT NULL DEFAULT false
);
