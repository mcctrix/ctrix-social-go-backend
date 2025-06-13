-- Enable the uuid-ossp extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- User

CREATE TABLE IF NOT EXISTS user_auth (
    id VARCHAR(50) PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    username VARCHAR(30) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS user_profile (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    first_name VARCHAR(30),
    last_name VARCHAR(30),
    avatar VARCHAR(25),
    profile_picture VARCHAR(200),
    last_seen TIMESTAMP
);

CREATE TABLE IF NOT EXISTS follows (
    follower_id text NOT NULL,  
    following_id text NOT NULL, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (follower_id, following_id),
    FOREIGN KEY (follower_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (following_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS user_additional_info (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth ON DELETE CASCADE ON UPDATE CASCADE,
    hobbies TEXT[],
    relation_status VARCHAR(12),
    dob DATE,
    bio VARCHAR(250),
    gender VARCHAR(6)
);

CREATE TABLE IF NOT EXISTS user_settings (
    id VARCHAR(50) PRIMARY KEY,
    FOREIGN KEY (id) REFERENCES user_auth ON DELETE CASCADE ON UPDATE CASCADE,
    hide_post TEXT[],
    hide_story TEXT[],
    block_user TEXT[],
    show_online BOOLEAN
);

-- User End

-- Posts
CREATE TABLE IF NOT EXISTS user_posts (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    creator_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    group_id VARCHAR(50),
    text_content TEXT,
    media_attached TEXT[]
);

CREATE TABLE IF NOT EXISTS user_post_like (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES user_posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    like_type VARCHAR(20),
    UNIQUE (post_id, user_id)
);

CREATE TABLE IF NOT EXISTS user_post_comments (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    post_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(post_id) REFERENCES user_posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    creator_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    content text,
    giff text,
    nested_comments TEXT[]
);

CREATE TABLE IF NOT EXISTS user_post_comment_like (
    user_id VARCHAR(50),
    FOREIGN KEY(user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    comment_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(comment_id) REFERENCES user_post_comments(id) ON DELETE CASCADE ON UPDATE CASCADE,
    like_type VARCHAR(20),
    UNIQUE (comment_id, user_id)
);

-- Posts End

-- Messenger
CREATE TABLE IF NOT EXISTS private_chats (
    id VARCHAR(50) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    firstUserID VARCHAR(50) NOT NULL,
    secondUserID VARCHAR(50) NOT NULL,
    isDeleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (firstUserID) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (secondUserID) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (firstUserID, secondUserID)
);

 CREATE TABLE IF NOT EXISTS private_chat_messages (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    private_chat_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    message_type VARCHAR(20),
    message VARCHAR(250) NOT NULL,
    isDeleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (private_chat_id) REFERENCES private_chats(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS group_chats (
    id VARCHAR(50) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    usersList TEXT[],
    adminList TEXT[],
    owner_id VARCHAR(50) NOT NULL,
    isDeleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (owner_id) REFERENCES user_auth(id)
);

CREATE TABLE IF NOT EXISTS group_chat_messages (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_chat_id VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    message_type VARCHAR(20),
    message VARCHAR(250) NOT NULL,
    isDeleted BOOLEAN DEFAULT FALSE,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (group_chat_id) REFERENCES group_chats(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS group_chat_member_requests (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_chat_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (group_chat_id) REFERENCES group_chats(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS message_reaction (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    message_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    reaction VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (message_id) REFERENCES group_chat_messages(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS group_chat_join_requests (
    id VARCHAR(50) PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_chat_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (group_chat_id) REFERENCES group_chats(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES user_auth(id) ON DELETE CASCADE ON UPDATE CASCADE
);
-- Messenger End

