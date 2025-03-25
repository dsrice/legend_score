-- +goose Up
CREATE TABLE user_tokens (
                             id int AUTO_INCREMENT NOT NULL COMMENT 'ID'
    , user_id INT NOT NULL COMMENT 'ユーザーID'
    , token VARCHAR(300) NOT NULL COMMENT 'トークン'
    , refresh_token VARCHAR(300) NOT NULL COMMENT 'リフレッシュトークン'
    , created_at datetime DEFAULT now() NOT NULL COMMENT '作成日'
    , updated_at datetime DEFAULT now() NOT NULL COMMENT '更新日'
    , deleted_flg boolean DEFAULT false NOT NULL COMMENT '削除フラグ'
    , deleted_at datetime COMMENT '削除日'
    , CONSTRAINT user_tokens_PKC PRIMARY KEY (id)
) COMMENT 'ユーザートークン情報' ;

ALTER TABLE user_tokens
    ADD CONSTRAINT user_tokens_FK1 FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE user_tokens;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd