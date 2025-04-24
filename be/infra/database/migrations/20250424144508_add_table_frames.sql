-- +goose Up
CREATE TABLE frames (
                        id int auto_increment NOT NULL COMMENT 'フレームID'
    , user_id INT NOT NULL COMMENT 'ユーザーID'
    , game_id INT NOT NULL COMMENT 'ゲームID'
    , frame_count DECIMAL NOT NULL COMMENT 'フレーム数'
    , frame_score INT COMMENT 'フレームスコア'
    , strike_flag BOOLEAN DEFAULT false COMMENT 'ストライクフラグ'
    , spare_flag BOOLEAN DEFAULT false COMMENT 'スペアフラグ'
    , created_at datetime DEFAULT now() NOT NULL COMMENT '作成日'
    , updated_at datetime DEFAULT now() NOT NULL COMMENT '更新日'
    , deleted_flg boolean DEFAULT false NOT NULL COMMENT '削除フラグ'
    , deleted_at datetime COMMENT '削除日'
    , CONSTRAINT frames_PKC PRIMARY KEY (id)
) COMMENT 'フレーム情報' ;

ALTER TABLE frames
    ADD CONSTRAINT frames_FK1 FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;

ALTER TABLE frames
    ADD CONSTRAINT frames_FK2 FOREIGN KEY (game_id) REFERENCES games(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE if exists frames CASCADE;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd