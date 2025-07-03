-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Clear existing data (if any)
TRUNCATE TABLE user_post_comment_like CASCADE;
TRUNCATE TABLE user_post_comments CASCADE;
TRUNCATE TABLE user_post_like CASCADE;
TRUNCATE TABLE user_posts CASCADE;
TRUNCATE TABLE follow CASCADE;
TRUNCATE TABLE user_additional_info CASCADE;
TRUNCATE TABLE user_settings CASCADE;
TRUNCATE TABLE user_profile CASCADE;
TRUNCATE TABLE user_auth CASCADE;

-- Insert Users (Auth)
INSERT INTO user_auth (id, email, username, password, created_at, updated_at) VALUES
-- Special users
('u1', 'ctrix@example.com', 'ctrix', '9887', '2020-01-15 10:00:00', '2020-01-15 10:00:00'),
('u2', 'anmol@example.com', 'anmol', '123456789', '2020-03-20 14:30:00', '2020-03-20 14:30:00'),
-- Random users
('u3', 'john@example.com', 'john_doe', 'password123', '2021-06-10 09:15:00', '2021-06-10 09:15:00'),
('u4', 'sarah@example.com', 'sarah_smith', 'password123', '2021-08-25 16:45:00', '2021-08-25 16:45:00'),
('u5', 'mike@example.com', 'mike_wilson', 'password123', '2022-02-05 11:20:00', '2022-02-05 11:20:00'),
('u6', 'emma@example.com', 'emma_brown', 'password123', '2022-04-18 13:10:00', '2022-04-18 13:10:00'),
('u7', 'david@example.com', 'david_jones', 'password123', '2023-01-30 08:50:00', '2023-01-30 08:50:00'),
('u8', 'lisa@example.com', 'lisa_taylor', 'password123', '2023-03-15 15:25:00', '2023-03-15 15:25:00'),
('u9', 'james@example.com', 'james_miller', 'password123', '2023-07-22 10:40:00', '2023-07-22 10:40:00'),
('u10', 'anna@example.com', 'anna_white', 'password123', '2023-09-05 14:15:00', '2023-09-05 14:15:00');

-- Insert User Profiles
INSERT INTO user_profile (id, first_name, last_name, avatar, profile_picture, last_seen, verified_user) VALUES
('u1', 'Ctrix', 'Developer', 'itachi', '', NOW() - INTERVAL '2 hours',true), -- Active recently
('u2', 'Anmol', 'Sharma', 'kakashi', '', NOW() - INTERVAL '2 days',false), -- Active few days ago
('u3', 'John', 'Doe', 'crawk', '', NOW() - INTERVAL '2 months',false), -- Active few months ago
('u4', 'Sarah', 'Smith', 'sasuke', '', NOW() - INTERVAL '1 year',false), -- Active last year
('u5', 'Mike', 'Wilson', 'gojo', '', NOW() - INTERVAL '3 hours',false), -- Active recently
('u6', 'Emma', 'Brown', 'zoro', '', NOW() - INTERVAL '1 month',false), -- Active last month
('u7', 'David', 'Jones', 'itachi', '', NOW() - INTERVAL '6 months',false), -- Active half year ago
('u8', 'Lisa', 'Taylor', 'kakashi', '', NOW() - INTERVAL '4 hours',false), -- Active recently
('u9', 'James', 'Miller', 'crawk', '', NOW() - INTERVAL '3 months',false), -- Active few months ago
('u10', 'Anna', 'White', 'sasuke', '', NOW() - INTERVAL '2 weeks',false); -- Active few weeks ago

-- Insert User Additional Info
INSERT INTO user_additional_info (id, hobbies, relation_status, dob, bio, gender) VALUES
('u1', ARRAY['Coding', 'Gaming', 'Photography'], 'Single', '1995-01-01', 'Full-stack developer and tech enthusiast', 'Male'),
('u2', ARRAY['Reading', 'Travel', 'Music'], 'Single', '1996-02-02', 'Software engineer and music lover', 'Male'),
('u3', ARRAY['Sports', 'Cooking'], 'Married', '1990-03-03', 'Sports enthusiast and family man', 'Male'),
('u4', ARRAY['Dancing', 'Painting'], 'Married', '1991-04-04', 'Artist and dancer', 'Female'),
('u5', ARRAY['Fishing', 'Hiking'], 'Single', '1992-05-05', 'Nature lover and adventurer', 'Male'),
('u6', ARRAY['Yoga', 'Meditation'], 'Single', '1993-06-06', 'Yoga instructor and wellness coach', 'Female'),
('u7', ARRAY['Photography', 'Travel'], 'Married', '1994-07-07', 'Travel photographer', 'Male'),
('u8', ARRAY['Cooking', 'Baking'], 'Married', '1995-08-08', 'Professional chef', 'Female'),
('u9', ARRAY['Gaming', 'Coding'], 'Single', '1996-09-09', 'Game developer', 'Male'),
('u10', ARRAY['Dancing', 'Singing'], 'Single', '1997-10-10', 'Professional dancer', 'Female');

-- Insert User Settings with varying configurations
INSERT INTO user_settings (id, hide_post, hide_story, block_user, show_online) VALUES
-- Ctrix: Hides some posts, blocks no one, shows online
('u1', ARRAY['p11', 'p12'], ARRAY[]::TEXT[], ARRAY[]::TEXT[], true),
-- Anmol: Hides no posts, blocks one user, shows online
('u2', ARRAY[]::TEXT[], ARRAY[]::TEXT[], ARRAY['u7'], true),
-- John: Hides some posts, blocks one user, doesn't show online
('u3', ARRAY['p1', 'p2', 'p3'], ARRAY[]::TEXT[], ARRAY['u8'], false),
-- Sarah: Hides many posts, blocks no one, doesn't show online
('u4', ARRAY['p11', 'p12', 'p13', 'p14', 'p15'], ARRAY[]::TEXT[], ARRAY[]::TEXT[], false),
-- Mike: Hides no posts, blocks two users, shows online
('u5', ARRAY[]::TEXT[], ARRAY[]::TEXT[], ARRAY['u9', 'u10'], true),
-- Emma: Hides some posts, blocks one user, doesn't show online
('u6', ARRAY['p21', 'p22'], ARRAY[]::TEXT[], ARRAY['u3'], false),
-- David: Hides no posts, blocks no one, shows online
('u7', ARRAY[]::TEXT[], ARRAY[]::TEXT[], ARRAY[]::TEXT[], true),
-- Lisa: Hides some posts, blocks one user, doesn't show online
('u8', ARRAY['p31', 'p32'], ARRAY[]::TEXT[], ARRAY['u4'], false),
-- James: Hides no posts, blocks two users, shows online
('u9', ARRAY[]::TEXT[], ARRAY[]::TEXT[], ARRAY['u5', 'u6'], true),
-- Anna: Hides some posts, blocks one user, doesn't show online
('u10', ARRAY['p35', 'p36'], ARRAY[]::TEXT[], ARRAY['u2'], false);

-- Insert Follows (Random following relationships with varying counts)
INSERT INTO follow (follower_id, following_id, created_at) VALUES
('u2', 'u1', '2020-04-01 10:00:00'),
('u3', 'u1', '2020-05-15 14:30:00'),
('u4', 'u1', '2020-06-20 09:45:00'),
('u5', 'u1', '2020-07-10 16:20:00'),
('u1', 'u2', '2020-04-02 11:00:00'),
('u1', 'u3', '2020-05-16 15:30:00'),
('u1', 'u4', '2020-06-21 10:45:00'),

-- Anmol (u2) - 3 followers, 2 following
('u6', 'u2', '2020-08-05 13:15:00'),
('u7', 'u2', '2020-09-12 17:30:00'),
('u2', 'u5', '2020-07-11 16:20:00'),

-- John (u3) - 2 followers, 4 following
('u8', 'u3', '2020-10-20 11:45:00'),
('u3', 'u4', '2020-06-22 10:45:00'),
('u3', 'u6', '2020-08-06 13:15:00'),
('u3', 'u9', '2020-11-15 14:30:00'),

-- Sarah (u4) - 4 followers, 3 following
('u7', 'u4', '2020-09-13 17:30:00'),
('u10', 'u4', '2020-12-01 15:45:00'),
('u4', 'u3', '2020-06-23 10:45:00'),
('u4', 'u8', '2020-10-21 11:45:00'),

-- Mike (u5) - 1 follower, 1 following
('u5', 'u4', '2020-07-12 16:20:00'),

-- Emma (u6) - 2 followers, 1 following
('u2', 'u6', '2020-08-05 13:15:00'),
('u9', 'u6', '2020-11-16 14:30:00'),

-- David (u7) - 2 followers, 3 following
('u2', 'u7', '2020-09-12 17:30:00'),
('u4', 'u7', '2020-09-13 17:30:00'),
('u7', 'u10', '2020-12-02 15:45:00'),

-- Lisa (u8) - 1 follower, 2 following
('u3', 'u8', '2020-10-20 11:45:00'),
('u8', 'u4', '2020-10-22 11:45:00'),

-- James (u9) - 2 followers, 1 following
('u6', 'u9', '2020-11-16 14:30:00'),

-- Anna (u10) - 1 follower, 1 following
('u4', 'u10', '2020-12-01 15:45:00'),
('u10', 'u7', '2020-12-02 15:45:00');

-- Insert Posts (10 posts per user with different timestamps)
-- Ctrix's posts (u1)
INSERT INTO user_posts (id, creator_id, created_at, updated_at, text_content, media_attached) VALUES
('p1', 'u1', '2020-02-01 10:00:00', '2020-02-01 11:00:00', 'Just launched my new project! #coding #development', ARRAY['https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg']),
('p2', 'u1', '2020-03-15 14:30:00', '2020-03-15 15:30:00', 'Beautiful sunset view from my office', ARRAY['https://i.pinimg.com/736x/81/c0/0f/81c00ffa83ba2e80ee984ddffde27599.jpg']),
('p3', 'u1', '2020-04-20 09:45:00', '2020-04-20 10:45:00', 'Working on some exciting features', NULL),
('p4', 'u1', '2020-05-10 16:20:00', '2020-05-10 17:20:00', 'Coffee and code ☕', NULL),
('p5', 'u1', '2020-06-25 11:15:00', '2020-06-25 12:15:00', 'New tech stack implementation', NULL),
('p6', 'u1', '2020-07-30 13:40:00', '2020-07-30 14:40:00', 'Weekend coding session', NULL),
('p7', 'u1', '2020-08-15 15:50:00', '2020-08-15 16:50:00', 'Debugging time!', NULL),
('p8', 'u1', '2020-09-20 10:25:00', '2020-09-20 11:25:00', 'Project milestone achieved', NULL),
('p9', 'u1', '2020-10-05 14:10:00', '2020-10-05 15:10:00', 'Learning new technologies', NULL),
('p10', 'u1', '2020-11-10 16:35:00', '2020-11-10 17:35:00', 'Team meeting day', NULL);

-- Anmol's posts (u2)
INSERT INTO user_posts (id, creator_id, created_at, updated_at, text_content, media_attached) VALUES
('p11', 'u2', '2020-04-01 10:00:00', '2020-04-01 11:00:00', 'Morning workout complete!', NULL),
('p12', 'u2', '2020-05-15 14:30:00', '2020-05-15 15:30:00', 'Beautiful day for a hike', ARRAY['https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg']),
('p13', 'u2', '2020-06-20 09:45:00', '2020-06-20 10:45:00', 'Coding session with friends', NULL),
('p14', 'u2', '2020-07-10 16:20:00', '2020-07-10 17:20:00', 'New project ideas', NULL),
('p15', 'u2', '2020-08-25 11:15:00', '2020-08-25 12:15:00', 'Weekend plans', NULL),
('p16', 'u2', '2020-09-30 13:40:00', '2020-09-30 14:40:00', 'Tech conference day', NULL),
('p17', 'u2', '2020-10-15 15:50:00', '2020-10-15 16:50:00', 'Learning new frameworks', NULL),
('p18', 'u2', '2020-11-20 10:25:00', '2020-11-20 11:25:00', 'Team collaboration', NULL),
('p19', 'u2', '2020-12-05 14:10:00', '2020-12-05 15:10:00', 'Code review time', NULL),
('p20', 'u2', '2021-01-10 16:35:00', '2021-01-10 17:35:00', 'Project deployment day', NULL);

-- Random posts for other users (u3-u10)
INSERT INTO user_posts (id, creator_id, created_at, updated_at, text_content, media_attached) VALUES
('p21', 'u3', '2021-07-01 10:00:00', '2021-07-01 11:00:00', 'Morning coffee ☕', NULL),
('p22', 'u3', '2021-08-15 14:30:00', '2021-08-15 15:30:00', 'Working from home', ARRAY['https://i.pinimg.com/736x/81/c0/0f/81c00ffa83ba2e80ee984ddffde27599.jpg']),
('p23', 'u4', '2021-09-20 09:45:00', '2021-09-20 10:45:00', 'Beautiful day!', NULL),
('p24', 'u4', '2021-10-10 16:20:00', '2021-10-10 17:20:00', 'Weekend vibes', ARRAY['https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg']),
('p25', 'u5', '2022-03-01 10:00:00', '2022-03-01 11:00:00', 'New project started', NULL),
('p26', 'u5', '2022-04-15 14:30:00', '2022-04-15 15:30:00', 'Team meeting', NULL),
('p27', 'u6', '2022-05-20 09:45:00', '2022-05-20 10:45:00', 'Yoga session', NULL),
('p28', 'u6', '2022-06-10 16:20:00', '2022-06-10 17:20:00', 'Meditation time', NULL),
('p29', 'u7', '2023-02-01 10:00:00', '2023-02-01 11:00:00', 'Photography day', NULL),
('p30', 'u7', '2023-03-15 14:30:00', '2023-03-15 15:30:00', 'Travel plans', NULL),
('p31', 'u8', '2023-04-20 09:45:00', '2023-04-20 10:45:00', 'Cooking class', NULL),
('p32', 'u8', '2023-05-10 16:20:00', '2023-05-10 17:20:00', 'New recipe', NULL),
('p33', 'u9', '2023-08-01 10:00:00', '2023-08-01 11:00:00', 'Gaming session', NULL),
('p34', 'u9', '2023-08-15 14:30:00', '2023-08-15 15:30:00', 'New game release', NULL),
('p35', 'u10', '2023-09-20 09:45:00', '2023-09-20 10:45:00', 'Dance practice', NULL),
('p36', 'u10', '2023-09-25 16:20:00', '2023-09-25 17:20:00', 'Performance day', NULL);

-- Insert Comments (2-6 random comments per post with different timestamps)
INSERT INTO user_post_comments (id, post_id, creator_id, created_at, updated_at, content) VALUES
-- Comments on Ctrix's posts
('c1', 'p1', 'u2', '2020-02-01 10:30:00', '2020-02-01 10:30:00', 'Amazing project!'),
('c2', 'p1', 'u3', '2020-02-01 11:00:00', '2020-02-01 11:00:00', 'Great work!'),
('c3', 'p2', 'u4', '2020-03-15 15:00:00', '2020-03-15 15:00:00', 'Beautiful view!'),
('c4', 'p2', 'u5', '2020-03-15 15:30:00', '2020-03-15 15:30:00', 'Nice office!'),
('c5', 'p3', 'u6', '2020-04-20 10:00:00', '2020-04-20 10:00:00', 'Looking forward to it!'),
('c6', 'p3', 'u7', '2020-04-20 10:30:00', '2020-04-20 10:30:00', 'Can''t wait to see!'),

-- Comments on Anmol's posts
('c7', 'p11', 'u1', '2020-04-01 10:30:00', '2020-04-01 10:30:00', 'Great job!'),
('c8', 'p11', 'u3', '2020-04-01 11:00:00', '2020-04-01 11:00:00', 'Keep it up!'),
('c9', 'p12', 'u4', '2020-05-15 15:00:00', '2020-05-15 15:00:00', 'Beautiful scenery!'),
('c10', 'p12', 'u5', '2020-05-15 15:30:00', '2020-05-15 15:30:00', 'Amazing view!'),

-- Comments on other posts
('c11', 'p21', 'u1', '2021-07-01 10:30:00', '2021-07-01 10:30:00', 'Enjoy your coffee!'),
('c12', 'p21', 'u2', '2021-07-01 11:00:00', '2021-07-01 11:00:00', 'Morning vibes!'),
('c13', 'p22', 'u3', '2021-08-15 15:00:00', '2021-08-15 15:00:00', 'WFH life!'),
('c14', 'p22', 'u4', '2021-08-15 15:30:00', '2021-08-15 15:30:00', 'Same here!'),
('c15', 'p23', 'u5', '2021-09-20 10:00:00', '2021-09-20 10:00:00', 'Beautiful day indeed!'),
('c16', 'p23', 'u6', '2021-09-20 10:30:00', '2021-09-20 10:30:00', 'Perfect weather!');

-- Insert Post Likes (with different timestamps)
INSERT INTO user_post_like (user_id, post_id) VALUES
('u1', 'p11'),
('u1', 'p12'),
('u2', 'p1'),
('u2', 'p2'),
('u3', 'p3'),
('u3', 'p4'),
('u4', 'p5'),
('u4', 'p6'),
('u5', 'p7'),
('u5', 'p8');

-- Insert Comment Likes (with different timestamps)
INSERT INTO user_post_comment_like (user_id, comment_id) VALUES
('u1', 'c1'),
('u1', 'c2'),
('u2', 'c3'),
('u2', 'c4'),
('u3', 'c5'),
('u3', 'c6'),
('u4', 'c7'),
('u4', 'c8'),
('u5', 'c9'),
('u5', 'c10');

-- Messenger Section
-- Clear existing messenger data
TRUNCATE TABLE message_reaction CASCADE;
TRUNCATE TABLE group_chat_member_requests CASCADE;
TRUNCATE TABLE group_chat_messages CASCADE;
TRUNCATE TABLE group_chats CASCADE;
TRUNCATE TABLE private_chat_messages CASCADE;
TRUNCATE TABLE private_chats CASCADE;

-- Insert Private Chats (2-7 chats per user)
INSERT INTO private_chats (id, created_at, firstUserID, secondUserID, isDeleted) VALUES
-- Ctrix's private chats
('pc1', '2020-02-01 10:00:00', 'u1', 'u2', false),
('pc2', '2020-03-15 14:30:00', 'u1', 'u3', false),
('pc3', '2020-04-20 09:45:00', 'u1', 'u4', false),
('pc4', '2020-05-10 16:20:00', 'u1', 'u5', false),
('pc5', '2020-06-25 11:15:00', 'u1', 'u6', false),
('pc6', '2020-07-30 13:40:00', 'u1', 'u7', false),
('pc7', '2020-08-15 15:50:00', 'u1', 'u8', false),

-- Anmol's private chats (excluding duplicates with u1)
('pc9', '2020-03-15 14:30:00', 'u2', 'u3', false),
('pc10', '2020-04-20 09:45:00', 'u2', 'u4', false),
('pc11', '2020-05-10 16:20:00', 'u2', 'u5', false),
('pc12', '2020-06-25 11:15:00', 'u2', 'u6', false),

-- John's private chats (excluding duplicates with u1 and u2)
('pc15', '2020-05-10 16:20:00', 'u3', 'u4', false),
('pc16', '2020-06-25 11:15:00', 'u3', 'u5', false),
('pc17', '2020-07-30 13:40:00', 'u3', 'u6', false),
('pc18', '2020-08-15 15:50:00', 'u3', 'u7', false),
('pc19', '2020-09-20 10:25:00', 'u3', 'u8', false),

-- Sarah's private chats (excluding duplicates with u1, u2, and u3)
('pc23', '2020-07-30 13:40:00', 'u4', 'u5', false),
('pc24', '2020-08-15 15:50:00', 'u4', 'u6', false);

-- Insert Private Chat Messages (4-20 messages per chat)
INSERT INTO private_chat_messages (id, private_chat_id, user_id, created_at, message_type, message, isDeleted) VALUES
-- Ctrix-Anmol chat messages
('pcm1', 'pc1', 'u1', '2020-02-01 10:00:00', 'text', 'Hey Anmol!', false),
('pcm2', 'pc1', 'u2', '2020-02-01 10:01:00', 'text', 'Hi Ctrix!', false),
('pcm3', 'pc1', 'u1', '2020-02-01 10:02:00', 'text', 'How are you?', false),
('pcm4', 'pc1', 'u2', '2020-02-01 10:03:00', 'text', 'I''m good, thanks!', false),
('pcm5', 'pc1', 'u1', '2020-02-01 10:04:00', 'media', 'https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg', false),
('pcm6', 'pc1', 'u2', '2020-02-01 10:05:00', 'text', 'Nice picture!', false),
('pcm7', 'pc1', 'u1', '2020-02-01 10:06:00', 'giff', 'https://i.pinimg.com/736x/81/c0/0f/81c00ffa83ba2e80ee984ddffde27599.jpg', false),
('pcm8', 'pc1', 'u2', '2020-02-01 10:07:00', 'text', 'Haha, that''s funny!', false),

-- Ctrix-John chat messages
('pcm9', 'pc2', 'u1', '2020-03-15 14:30:00', 'text', 'Hey John!', false),
('pcm10', 'pc2', 'u3', '2020-03-15 14:31:00', 'text', 'Hi Ctrix!', false),
('pcm11', 'pc2', 'u1', '2020-03-15 14:32:00', 'text', 'How''s the project going?', false),
('pcm12', 'pc2', 'u3', '2020-03-15 14:33:00', 'text', 'Going well, thanks!', false),
('pcm13', 'pc2', 'u1', '2020-03-15 14:34:00', 'media', 'https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg', false),
('pcm14', 'pc2', 'u3', '2020-03-15 14:35:00', 'text', 'Looking good!', false);

-- Insert Group Chats (1-4 per user)
INSERT INTO group_chats (id, created_at, usersList, adminList, owner_id, isDeleted) VALUES
-- Tech Team Group
('gc1', '2020-02-01 10:00:00', ARRAY['u1', 'u2', 'u3', 'u4'], ARRAY['u1', 'u2'], 'u1', false),
-- Project Alpha Group
('gc2', '2020-03-15 14:30:00', ARRAY['u1', 'u2', 'u5', 'u6'], ARRAY['u1', 'u5'], 'u1', false),
-- Design Team Group
('gc3', '2020-04-20 09:45:00', ARRAY['u3', 'u4', 'u7', 'u8'], ARRAY['u3', 'u4'], 'u3', false),
-- Marketing Group
('gc4', '2020-05-10 16:20:00', ARRAY['u5', 'u6', 'u9', 'u10'], ARRAY['u5', 'u6'], 'u5', false),
-- Social Group
('gc5', '2020-06-25 11:15:00', ARRAY['u1', 'u2', 'u7', 'u8', 'u9', 'u10'], ARRAY['u1', 'u2', 'u7'], 'u1', false),
-- Gaming Group (deleted)
('gc6', '2020-07-30 13:40:00', ARRAY['u3', 'u4', 'u9', 'u10'], ARRAY['u3', 'u4'], 'u3', true);

-- Insert Group Chat Messages (2-4 messages per user in each group)
INSERT INTO group_chat_messages (id, group_chat_id, created_at, updated_at, message_type, message, isDeleted, user_id) VALUES
-- Tech Team Group messages
('gcm1', 'gc1', '2020-02-01 10:00:00', '2020-02-01 10:00:00', 'text', 'Welcome to Tech Team!', false, 'u1'),
('gcm2', 'gc1', '2020-02-01 10:01:00', '2020-02-01 10:01:00', 'text', 'Thanks for having me!', false, 'u2'),
('gcm3', 'gc1', '2020-02-01 10:02:00', '2020-02-01 10:02:00', 'text', 'Excited to work with everyone!', false, 'u3'),
('gcm4', 'gc1', '2020-02-01 10:03:00', '2020-02-01 10:03:00', 'text', 'Let''s make it great!', false, 'u4'),
('gcm5', 'gc1', '2020-02-01 10:04:00', '2020-02-01 10:04:00', 'media', 'https://i.pinimg.com/736x/d4/83/56/d48356c7b841e2f58525c86eebbc46e8.jpg', false, 'u1'),
('gcm6', 'gc1', '2020-02-01 10:05:00', '2020-02-01 10:05:00', 'text', 'Great team photo!', false, 'u2'),

-- Project Alpha Group messages
('gcm7', 'gc2', '2020-03-15 14:30:00', '2020-03-15 14:30:00', 'text', 'Project Alpha kickoff!', false, 'u1'),
('gcm8', 'gc2', '2020-03-15 14:31:00', '2020-03-15 14:31:00', 'text', 'Ready to start!', false, 'u2'),
('gcm9', 'gc2', '2020-03-15 14:32:00', '2020-03-15 14:32:00', 'text', 'Let''s do this!', false, 'u5'),
('gcm10', 'gc2', '2020-03-15 14:33:00', '2020-03-15 14:33:00', 'giff', 'https://i.pinimg.com/736x/81/c0/0f/81c00ffa83ba2e80ee984ddffde27599.jpg', false, 'u6');

-- Insert Message Reactions (only for group chat messages)
INSERT INTO message_reaction (id, message_id, user_id, reaction) VALUES
('mr1', 'gcm1', 'u2', 'like'),
('mr2', 'gcm2', 'u3', 'love'),
('mr3', 'gcm3', 'u1', 'like'),
('mr4', 'gcm4', 'u4', 'haha'),
('mr5', 'gcm5', 'u2', 'love'),
('mr6', 'gcm6', 'u3', 'like');

-- Insert Group Chat Member Requests
INSERT INTO group_chat_member_requests (id, group_chat_id, user_id) VALUES
('gcmr1', 'gc1', 'u5'),
('gcmr2', 'gc2', 'u7'),
('gcmr3', 'gc3', 'u9'),
('gcmr4', 'gc4', 'u1'); 

INSERT INTO bookmark (user_id, post_id) VALUES
('u1', 'p1'),
('u2', 'p2'),
('u2', 'p3'),
('u3', 'p3'),
('u4', 'p4'),
('u5', 'p5');
