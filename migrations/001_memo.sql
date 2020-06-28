-- +migrate Up
CREATE TABLE memo
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    memo       VARCHAR(1024) NOT NULL
);

COMMENT ON TABLE  memo IS 'Memo';
COMMENT ON COLUMN memo.id IS 'id';
COMMENT ON COLUMN memo.created_at IS 'create date';
COMMENT ON COLUMN memo.updated_at IS 'update date';
COMMENT ON COLUMN memo.deleted_at IS 'delete date';
COMMENT ON COLUMN memo.memo IS 'memo';

-- +migrate Down
DROP TABLE memo;