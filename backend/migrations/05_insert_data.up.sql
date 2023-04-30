-- INSERT 100 teams
INSERT INTO "team"("id", "team_name", "image") VALUES 
('8143e362-8540-4d89-b6d3-1c4938203ca2','Arsenal','https://upload.wikimedia.org/wikipedia/ru/thumb/5/53/Arsenal_FC.svg/200px-Arsenal_FC.svg.png'),
('716bc850-6f73-4d8d-948a-8cb3b64d6370','Liverpool', 'https://d3j2s6hdd6a7rg.cloudfront.net/images/international/thumb-no-image-2.jpg'),
('f15cb2fe-e042-4f7f-808d-d577b676d2e7','Manchester United', 'https://upload.wikimedia.org/wikipedia/en/thumb/7/7a/Manchester_United_FC_crest.svg/1200px-Manchester_United_FC_crest.svg.png'), 
('47c7cc67-11ed-437f-ad33-0c2c199c60af','Manchester City', 'https://upload.wikimedia.org/wikipedia/en/thumb/e/eb/Manchester_City_FC_badge.svg/1200px-Manchester_City_FC_badge.svg.png'), 
('d041ecf1-8383-4ecc-9661-840b13fb0350','Tottenham Hotspur', 'https://m.media-amazon.com/images/M/MV5BZTEzZTdkN2UtNmE0OS00ZTZiLTk0YWEtZmQyNDQyNDE4NGVlXkEyXkFqcGdeQXVyMjUyNDk2ODc@._V1_.jpg'), 
('4b9c4434-0954-469a-90ca-1df6bc5fad45','Everton', 'https://upload.wikimedia.org/wikipedia/en/thumb/7/7c/Everton_FC_logo.svg/1200px-Everton_FC_logo.svg.png'), 
('89ddf745-2107-4bcc-91e7-cec11c5c7a17','Southampton', 'https://upload.wikimedia.org/wikipedia/en/thumb/c/c9/FC_Southampton.svg/800px-FC_Southampton.svg.png'), 
('4bb3d7e5-4eb0-4256-89cb-892b39c35678','Newcastle United', 'https://upload.wikimedia.org/wikipedia/en/thumb/5/56/Newcastle_United_Logo.svg/1200px-Newcastle_United_Logo.svg.png'), 
('6d29681d-ba78-4d11-8db7-3c090d82f58f','West Ham United', 'https://upload.wikimedia.org/wikipedia/en/thumb/c/c2/West_Ham_United_FC_logo.svg/1200px-West_Ham_United_FC_logo.svg.png'), 
('3642dd5f-9476-4bc4-8186-7c70bf5f440c','Leicester City', 'https://upload.wikimedia.org/wikipedia/en/thumb/2/2d/Leicester_City_crest.svg/1200px-Leicester_City_crest.svg.png'), 
('955285ba-f05c-41a4-be12-441234849d83','Stoke City', 'https://upload.wikimedia.org/wikipedia/en/thumb/2/29/Stoke_City_FC.svg/1200px-Stoke_City_FC.svg.png'), 
('32c35bc3-2401-4646-8219-d758ad086753','Crystal Palace', 'https://i.pinimg.com/originals/80/8d/4e/808d4e6db9a6b74b5d437c9f8429aa1c.jpg'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4','Toulouse', 'https://upload.wikimedia.org/wikipedia/fr/5/58/Toulouse_FC_logo_2021.svg');

INSERT INTO "league" ("id","league_name", "image") VALUES 
('d677c079-dc75-4d57-97a4-066bda707b30','Italiya. Seriya A', 'https://img.championat.com/tournir/15976451711136497650.png'),
('7c501efa-731c-4a20-b4b2-c24268233bd6','Premier Liga', 'https://images.sports.gracenote.com/images/lib/basic/sport/football/competition/logo/300/52.png'),
('a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', 'LaLiga', 'https://thumbs.dreamstime.com/b/algeria-map-collection-253485908.jpg'),
('eff19f85-5025-44fc-ae89-0c1607c23a48', 'Bundesliga', 'https://upload.wikimedia.org/wikipedia/en/thumb/d/df/Bundesliga_logo_%282017%29.svg/1200px-Bundesliga_logo_%282017%29.svg.png'),
('ad6578c9-1a7d-471e-84d4-80d89772a8c7', 'Ligue', 'https://upload.wikimedia.org/wikipedia/commons/f/fb/Ligue1_logo.png');

INSERT INTO "season" ("id", "season_year_from", "season_year_to", "league_id") VALUES 
('51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2023, 2024, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('fe2cf040-2792-4383-a2e4-c2a7c02bf975', 2015, 2016, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('61a42f93-bf67-4ec3-bfa0-6b4253d8704c', 2016, 2017, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('45b84c64-200a-40fe-bf19-4cf799312907', 2017, 2018, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('cbe3ba20-9546-48bc-a97b-a00e17790137', 2018, 2019, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('449ecd26-1c3e-47c4-a1af-469cecd8838c', 2019, 2020, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('9e5c7f17-72b5-4500-b141-e1dde2b953b2', 2020, 2021, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('3753c89c-02d1-4eb4-a524-632953b8cd88', 2021, 2022, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('7f75ecfc-2f3a-430c-9e6a-181dc415fc9e', 2022, 2023, 'd677c079-dc75-4d57-97a4-066bda707b30');

INSERT INTO "season" ("id", "season_year_from", "season_year_to", "league_id") VALUES 
('6178f11f-aee2-47c5-9aba-3d3b13fec14f', 2023, 2024, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('fbfcc5eb-3a5e-4f98-aa47-80d60eac74e5', 2015, 2016, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('4246526e-764f-4ffc-acbd-5cb9d78a7273', 2016, 2017, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('75755dbe-8434-4250-8c82-b25c6defda50', 2017, 2018, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('e5136935-ef39-48e8-ab36-9043a4ec39a3', 2018, 2019, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('c8ceeed9-e8aa-478c-ad3f-6e63eec46b9d', 2019, 2020, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('b29b6a18-aaff-4da4-8847-8d6f477d54d0', 2020, 2021, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('411d1179-e576-4f42-924d-07658b846613', 2021, 2022, '7c501efa-731c-4a20-b4b2-c24268233bd6'),
('092b6144-582a-40d1-8e7f-7032aa976142', 2022, 2023, '7c501efa-731c-4a20-b4b2-c24268233bd6');

INSERT INTO "match"("id", "home_team_id", "away_team_id", "league_id", "season_id", "home_score", "away_score", "m_date") VALUES
('ed5aa7f9-d320-47ea-af0b-8309820f8f96', '8143e362-8540-4d89-b6d3-1c4938203ca2', '716bc850-6f73-4d8d-948a-8cb3b64d6370','d677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2, 1, '2014-08-16 15:00:00'),
('b1a1068f-9103-48b6-a4ec-6184e6180cf6', '47c7cc67-11ed-437f-ad33-0c2c199c60af','4b9c4434-0954-469a-90ca-1df6bc5fad45','d677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('bde4bd2e-9568-4c3a-84f7-9fb615a3ad61','3642dd5f-9476-4bc4-8186-7c70bf5f440c','955285ba-f05c-41a4-be12-441234849d83','d677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('de5a2898-e350-4adc-b6b5-c573fda6e04a', '9107ccda-7e3e-412b-a200-c6daba95a1e4', '32c35bc3-2401-4646-8219-d758ad086753','d677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('4720cd3e-879e-493c-8da3-fef516ff7669', '6d29681d-ba78-4d11-8db7-3c090d82f58f', 'f15cb2fe-e042-4f7f-808d-d577b676d2e7', 'd677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('2d64ceb2-036f-482d-8df5-86c076453097', '32c35bc3-2401-4646-8219-d758ad086753', '955285ba-f05c-41a4-be12-441234849d83', 'd677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('dce37ddd-d2fc-4c4b-9b40-368a116ee938', '3642dd5f-9476-4bc4-8186-7c70bf5f440c', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'd677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('cf73913d-c113-4301-8587-341c5113a132', 'd041ecf1-8383-4ecc-9661-840b13fb0350', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'd677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('8bb7e7b7-37a1-4d8b-b690-f0580db6bdd1', '9107ccda-7e3e-412b-a200-c6daba95a1e4', '47c7cc67-11ed-437f-ad33-0c2c199c60af', 'd677c079-dc75-4d57-97a4-066bda707b30','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00');

INSERT INTO "match"("id", "home_team_id", "away_team_id", "league_id", "season_id", "home_score", "away_score", "m_date") VALUES
('e28f025d-02bc-4cd6-b85a-a48d6ca17ac5', '8143e362-8540-4d89-b6d3-1c4938203ca2', '716bc850-6f73-4d8d-948a-8cb3b64d6370','7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2, 1, '2014-08-16 15:00:00'),
('9fc4df5f-ff3b-439f-b958-0ff7bd60c638', '47c7cc67-11ed-437f-ad33-0c2c199c60af','4b9c4434-0954-469a-90ca-1df6bc5fad45','7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('36b4a840-6030-4816-b15d-29d67a8b2ab3','3642dd5f-9476-4bc4-8186-7c70bf5f440c','955285ba-f05c-41a4-be12-441234849d83','7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('f9a5ee49-9078-4cce-8bc3-08f91ac0af6f', '9107ccda-7e3e-412b-a200-c6daba95a1e4', '32c35bc3-2401-4646-8219-d758ad086753','7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('415868d3-ceba-447c-badf-1e889c78354d', '6d29681d-ba78-4d11-8db7-3c090d82f58f', 'f15cb2fe-e042-4f7f-808d-d577b676d2e7', '7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('27e4c317-aa21-4a85-9d1f-39d8e8ed3fcf', '32c35bc3-2401-4646-8219-d758ad086753', '955285ba-f05c-41a4-be12-441234849d83', '7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('d57b1f11-ea33-425c-89e7-9ed64e3d6f8f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', '7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('fb3bdfd7-1ae0-4593-807e-d6a243a55946', 'd041ecf1-8383-4ecc-9661-840b13fb0350', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', '7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('2effe710-3164-4eaa-ac78-d91e1a92c0fd', '9107ccda-7e3e-412b-a200-c6daba95a1e4', '47c7cc67-11ed-437f-ad33-0c2c199c60af', '7c501efa-731c-4a20-b4b2-c24268233bd6','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00');

INSERT INTO "match"("home_team_id", "away_team_id", "league_id", "season_id", "home_score", "away_score", "m_date") VALUES
('8143e362-8540-4d89-b6d3-1c4938203ca2', '716bc850-6f73-4d8d-948a-8cb3b64d6370','a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2, 1, '2014-08-16 15:00:00'),
('47c7cc67-11ed-437f-ad33-0c2c199c60af','4b9c4434-0954-469a-90ca-1df6bc5fad45','a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c','955285ba-f05c-41a4-be12-441234849d83','a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '32c35bc3-2401-4646-8219-d758ad086753','a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('6d29681d-ba78-4d11-8db7-3c090d82f58f', 'f15cb2fe-e042-4f7f-808d-d577b676d2e7', 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('32c35bc3-2401-4646-8219-d758ad086753', '955285ba-f05c-41a4-be12-441234849d83', 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('d041ecf1-8383-4ecc-9661-840b13fb0350', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '47c7cc67-11ed-437f-ad33-0c2c199c60af', 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00');

INSERT INTO "match"("home_team_id", "away_team_id", "league_id", "season_id", "home_score", "away_score", "m_date") VALUES
('8143e362-8540-4d89-b6d3-1c4938203ca2', '716bc850-6f73-4d8d-948a-8cb3b64d6370','eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2, 1, '2014-08-16 15:00:00'),
('47c7cc67-11ed-437f-ad33-0c2c199c60af','4b9c4434-0954-469a-90ca-1df6bc5fad45','eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c','955285ba-f05c-41a4-be12-441234849d83','eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '32c35bc3-2401-4646-8219-d758ad086753','eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('6d29681d-ba78-4d11-8db7-3c090d82f58f', 'f15cb2fe-e042-4f7f-808d-d577b676d2e7', 'eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('32c35bc3-2401-4646-8219-d758ad086753', '955285ba-f05c-41a4-be12-441234849d83', 'eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('d041ecf1-8383-4ecc-9661-840b13fb0350', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '47c7cc67-11ed-437f-ad33-0c2c199c60af', 'eff19f85-5025-44fc-ae89-0c1607c23a48','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00');

INSERT INTO "match"("home_team_id", "away_team_id", "league_id", "season_id", "home_score", "away_score", "m_date") VALUES
('8143e362-8540-4d89-b6d3-1c4938203ca2', '716bc850-6f73-4d8d-948a-8cb3b64d6370','ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2, 1, '2014-08-16 15:00:00'),
('47c7cc67-11ed-437f-ad33-0c2c199c60af','4b9c4434-0954-469a-90ca-1df6bc5fad45','ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c','955285ba-f05c-41a4-be12-441234849d83','ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '32c35bc3-2401-4646-8219-d758ad086753','ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('6d29681d-ba78-4d11-8db7-3c090d82f58f', 'f15cb2fe-e042-4f7f-808d-d577b676d2e7', 'ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('32c35bc3-2401-4646-8219-d758ad086753', '955285ba-f05c-41a4-be12-441234849d83', 'ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('3642dd5f-9476-4bc4-8186-7c70bf5f440c', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('d041ecf1-8383-4ecc-9661-840b13fb0350', '4bb3d7e5-4eb0-4256-89cb-892b39c35678', 'ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4', '47c7cc67-11ed-437f-ad33-0c2c199c60af', 'ad6578c9-1a7d-471e-84d4-80d89772a8c7','51f0da7a-40a9-4fc3-b51e-e3b511dc608f',2, 1, '2014-08-16 15:00:00');


INSERT INTO "team_league_season"("id", "league_id", "season_id", "team_id") VALUES
('4b6921a4-d497-48c9-b22f-0647b5094eaa', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '8143e362-8540-4d89-b6d3-1c4938203ca2'),
('94deff49-a6c5-4ce1-9e27-fcf29f74fe49', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c'),
('a07a2352-eec6-4518-9695-8c8015cab2dd', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '955285ba-f05c-41a4-be12-441234849d83'),
('eae6035d-00e5-448d-aec9-cdfa2f3f8d47', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '32c35bc3-2401-4646-8219-d758ad086753'),
('c510b5d4-031c-44c7-b133-b009b6104f67', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '4bb3d7e5-4eb0-4256-89cb-892b39c35678'),
('e4909b0b-b55a-4b64-9578-f33872e61100', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 'd041ecf1-8383-4ecc-9661-840b13fb0350'),
('df542cf5-572c-4552-844b-2d17f0fe5819','d677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','9107ccda-7e3e-412b-a200-c6daba95a1e4'),
('2b212295-b1d1-4069-b910-9b7215bd57ff','d677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','47c7cc67-11ed-437f-ad33-0c2c199c60af'),
('f3c21fb5-21c9-4abc-9230-9c03dca6a8ba', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','89ddf745-2107-4bcc-91e7-cec11c5c7a17'),
('5a47d544-afbf-47be-acbf-5260ffa682ff', 'd677c079-dc75-4d57-97a4-066bda707b30', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','f15cb2fe-e042-4f7f-808d-d577b676d2e7');

INSERT INTO "team_league_season"("id", "league_id", "season_id", "team_id") VALUES
('444e57aa-0a2a-4a39-aba3-5f18813babdf', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '8143e362-8540-4d89-b6d3-1c4938203ca2'),
('a01d9acb-07d8-491b-84ff-485abd437f2d', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c'),
('351ee187-02e7-4106-84bc-f6d1b7cf69d2', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '955285ba-f05c-41a4-be12-441234849d83'),
('ce3f5464-567f-4b13-95b0-33799a235cc5', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '32c35bc3-2401-4646-8219-d758ad086753'),
('2b3b0e68-5947-488a-9e03-1357e3cd8df8', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '4bb3d7e5-4eb0-4256-89cb-892b39c35678'),
('6dd057c4-1b18-464c-ade0-a30515f25abd', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 'd041ecf1-8383-4ecc-9661-840b13fb0350'),
('5311b81f-d07d-439f-9b65-714923fdff9f','7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','9107ccda-7e3e-412b-a200-c6daba95a1e4'),
('9c4aec84-cd2a-4003-ab67-4a5221e69810','7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','47c7cc67-11ed-437f-ad33-0c2c199c60af'),
('3acdb270-74b8-42d8-8560-673a7758030c', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','89ddf745-2107-4bcc-91e7-cec11c5c7a17'),
('6a6a1737-dee7-4ec0-8688-560dbb160ae5', '7c501efa-731c-4a20-b4b2-c24268233bd6', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','f15cb2fe-e042-4f7f-808d-d577b676d2e7');


INSERT INTO "team_league_season"("league_id", "season_id", "team_id") VALUES
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '8143e362-8540-4d89-b6d3-1c4938203ca2'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '955285ba-f05c-41a4-be12-441234849d83'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '32c35bc3-2401-4646-8219-d758ad086753'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '4bb3d7e5-4eb0-4256-89cb-892b39c35678'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 'd041ecf1-8383-4ecc-9661-840b13fb0350'),
('a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','9107ccda-7e3e-412b-a200-c6daba95a1e4'),
('a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','47c7cc67-11ed-437f-ad33-0c2c199c60af'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','89ddf745-2107-4bcc-91e7-cec11c5c7a17'),
( 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','f15cb2fe-e042-4f7f-808d-d577b676d2e7');

INSERT INTO "team_league_season"("league_id", "season_id", "team_id") VALUES
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '8143e362-8540-4d89-b6d3-1c4938203ca2'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '955285ba-f05c-41a4-be12-441234849d83'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '32c35bc3-2401-4646-8219-d758ad086753'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '4bb3d7e5-4eb0-4256-89cb-892b39c35678'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 'd041ecf1-8383-4ecc-9661-840b13fb0350'),
('ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','9107ccda-7e3e-412b-a200-c6daba95a1e4'),
('ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','47c7cc67-11ed-437f-ad33-0c2c199c60af'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','89ddf745-2107-4bcc-91e7-cec11c5c7a17'),
( 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','f15cb2fe-e042-4f7f-808d-d577b676d2e7');

INSERT INTO "team_league_season"("league_id", "season_id", "team_id") VALUES
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '8143e362-8540-4d89-b6d3-1c4938203ca2'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '3642dd5f-9476-4bc4-8186-7c70bf5f440c'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '955285ba-f05c-41a4-be12-441234849d83'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '32c35bc3-2401-4646-8219-d758ad086753'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', '4bb3d7e5-4eb0-4256-89cb-892b39c35678'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 'd041ecf1-8383-4ecc-9661-840b13fb0350'),
('eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','9107ccda-7e3e-412b-a200-c6daba95a1e4'),
('eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','47c7cc67-11ed-437f-ad33-0c2c199c60af'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','89ddf745-2107-4bcc-91e7-cec11c5c7a17'),
( 'eff19f85-5025-44fc-ae89-0c1607c23a48', '51f0da7a-40a9-4fc3-b51e-e3b511dc608f','f15cb2fe-e042-4f7f-808d-d577b676d2e7');

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, 'd677c079-dc75-4d57-97a4-066bda707b30', 1);

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, '7c501efa-731c-4a20-b4b2-c24268233bd6', 1);

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, 'a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', 1);

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, 'eff19f85-5025-44fc-ae89-0c1607c23a48', 1);

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, 'ad6578c9-1a7d-471e-84d4-80d89772a8c7', 2);