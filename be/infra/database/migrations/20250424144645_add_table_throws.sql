-- +goose Up
CREATE TABLE throws (
                        id int auto_increment NOT NULL COMMENT '投球ID'
    , user_id INT NOT NULL COMMENT 'ユーザーID'
    , game_id INT NOT NULL COMMENT 'ゲームID'
    , frame_id INT NOT NULL COMMENT 'フレームID'
    , throw_count INT NOT NULL COMMENT 'フレーム投球回数'
    , throw_score INT NOT NULL COMMENT '投球スコア'
    , strike_flag BOOLEAN DEFAULT false NOT NULL COMMENT 'ストライクフラグ'
    , spare_flag BOOLEAN DEFAULT false NOT NULL COMMENT 'スペアフラグ'
    , split_flag BOOLEAN DEFAULT false NOT NULL COMMENT 'スプリットフラグ'
    , pin_1 INT NOT NULL COMMENT '1ピン結果'
    , pin_2 INT NOT NULL COMMENT '2ピン結果'
    , pin_3 INT NOT NULL COMMENT '3ピン結果'
    , pin_4 INT NOT NULL COMMENT '4ピン結果'
    , pin_5 INT NOT NULL COMMENT '5ピン結果'
    , pin_6 INT NOT NULL COMMENT '6ピン結果'
    , pin_7 INT NOT NULL COMMENT '7ピン結果'
    , pin_8 INT NOT NULL COMMENT '8ピン結果'
    , pin_9 INT NOT NULL COMMENT '9ピン結果'
    , pin_10 INT NOT NULL COMMENT '10ピン結果'
    , created_at datetime DEFAULT now() NOT NULL COMMENT '作成日'
    , updated_at datetime DEFAULT now() NOT NULL COMMENT '更新日'
    , deleted_flg boolean DEFAULT false NOT NULL COMMENT '削除フラグ'
    , deleted_at datetime COMMENT '削除日'
    , CONSTRAINT throws_PKC PRIMARY KEY (id)
) COMMENT '投球情報' ;

ALTER TABLE throws
    ADD CONSTRAINT throws_FK1 FOREIGN KEY (game_id) REFERENCES games(id);

ALTER TABLE throws
    ADD CONSTRAINT throws_FK2 FOREIGN KEY (frame_id) REFERENCES frames(id);

ALTER TABLE throws
    ADD CONSTRAINT throws_FK3 FOREIGN KEY (user_id) REFERENCES users(id);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE if exists throws CASCADE;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd