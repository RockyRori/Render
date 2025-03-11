USE render;

-- 插入示例游戏数据
INSERT INTO games (title, description, price, release_date) VALUES
('Cyberpunk 2077', '开放世界角色扮演游戏，设定在赛博朋克未来都市', 59.99, '2020-12-10'),
('The Witcher 3: Wild Hunt', '史诗奇幻冒险RPG，扮演猎魔人杰洛特', 39.99, '2015-05-19'),
('Elden Ring', '魂系开放世界RPG，由宫崎英高与乔治·R·R·马丁合作打造', 59.99, '2022-02-25'),
('Hollow Knight', '类银河恶魔城动作游戏，拥有深度探索和战斗系统', 14.99, '2017-02-24');

-- 插入示例用户数据
INSERT INTO users (username, email, password_hash) VALUES
('player1', 'player1@example.com', 'hashedpassword1'),
('player2', 'player2@example.com', 'hashedpassword2'),
('player3', 'player3@example.com', 'hashedpassword3');

-- 插入示例订单数据
INSERT INTO orders (user_id, game_id, status) VALUES
(1, 1, 'completed'),
(2, 2, 'pending'),
(3, 3, 'completed'),
(1, 4, 'cancelled');