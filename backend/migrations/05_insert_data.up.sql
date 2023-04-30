-- INSERT 100 teams
INSERT INTO "team"("id", "team_name") VALUES 
('8143e362-8540-4d89-b6d3-1c4938203ca2','Arsenal'), 
('716bc850-6f73-4d8d-948a-8cb3b64d6370','Liverpool'), 
('f15cb2fe-e042-4f7f-808d-d577b676d2e7','Manchester United'), 
('47c7cc67-11ed-437f-ad33-0c2c199c60af','Manchester City'), 
('d041ecf1-8383-4ecc-9661-840b13fb0350','Tottenham Hotspur'), 
('4b9c4434-0954-469a-90ca-1df6bc5fad45','Everton'), 
('89ddf745-2107-4bcc-91e7-cec11c5c7a17','Southampton'), 
('4bb3d7e5-4eb0-4256-89cb-892b39c35678','Newcastle United'), 
('6d29681d-ba78-4d11-8db7-3c090d82f58f','West Ham United'), 
('3642dd5f-9476-4bc4-8186-7c70bf5f440c','Leicester City'), 
('955285ba-f05c-41a4-be12-441234849d83','Stoke City'), 
('32c35bc3-2401-4646-8219-d758ad086753','Crystal Palace'),
('9107ccda-7e3e-412b-a200-c6daba95a1e4','Liverpool');

INSERT INTO "league" ("id","league_name") VALUES 
('d677c079-dc75-4d57-97a4-066bda707b30','Italiya. Seriya A'),
('7c501efa-731c-4a20-b4b2-c24268233bd6','Premier Liga'),
('a0a5ed6b-0e8a-4ab5-aef5-a1cd13c1bba7', 'LaLiga'),
('eff19f85-5025-44fc-ae89-0c1607c23a48', 'Bundesliga'),
('ad6578c9-1a7d-471e-84d4-80d89772a8c7', 'Ligue');

INSERT INTO "season" ("id", "season_year_from", "season_year_to", "league_id") VALUES 
('51f0da7a-40a9-4fc3-b51e-e3b511dc608f', 2014, 2015, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('fe2cf040-2792-4383-a2e4-c2a7c02bf975', 2015, 2016, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('61a42f93-bf67-4ec3-bfa0-6b4253d8704c', 2016, 2017, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('45b84c64-200a-40fe-bf19-4cf799312907', 2017, 2018, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('cbe3ba20-9546-48bc-a97b-a00e17790137', 2018, 2019, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('449ecd26-1c3e-47c4-a1af-469cecd8838c', 2019, 2020, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('9e5c7f17-72b5-4500-b141-e1dde2b953b2', 2020, 2021, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('3753c89c-02d1-4eb4-a524-632953b8cd88', 2021, 2022, 'd677c079-dc75-4d57-97a4-066bda707b30'),
('7f75ecfc-2f3a-430c-9e6a-181dc415fc9e', 2022, 2023, 'd677c079-dc75-4d57-97a4-066bda707b30');

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

INSERT INTO tur("soni", "league_id", "current_tur") VALUES
(14, 'd677c079-dc75-4d57-97a4-066bda707b30', 1);