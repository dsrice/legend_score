-- +goose Up
CREATE TABLE games (
                       id int AUTO_INCREMENT NOT NULL COMMENT 'ゲームID'
    , user_id INT NOT NULL COMMENT 'ユーザーID'
    , name VARCHAR(100) COMMENT 'ゲーム名称'
    , score INT NOT NULL COMMENT 'スコア'
    , count INT COMMENT 'ゲーム数'
    , game_date DATE COMMENT '投球日'
    , created_at datetime DEFAULT now() NOT NULL COMMENT '作成日'
    , updated_at datetime DEFAULT now() NOT NULL COMMENT '更新日'
    , deleted_flg boolean DEFAULT false NOT NULL COMMENT '削除フラグ'
    , deleted_at datetime COMMENT '削除日'
    , CONSTRAINT games_PKC PRIMARY KEY (id)
) COMMENT 'ゲーム情報' ;

ALTER TABLE games
    ADD CONSTRAINT games_FK1 FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE if exists games CASCADE;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd