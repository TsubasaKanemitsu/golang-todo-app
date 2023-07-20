-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    m_todo_task (
        id SERIAL,
        title VARCHAR(256) NOT NULL,
        contents TEXT,
        assignee VARCHAR(32),
        label VARCHAR(16),
        status VARCHAR(16) NOT NULL,
        start_date TIMESTAMP DEFAULT NOW () NOT NULL,
        end_date TIMESTAMP NOT NULL,
        created_at TIMESTAMP DEFAULT NOW (),
        updated_at TIMESTAMP,
        CONSTRAINT id PRIMARY KEY (id)
    );

COMMENT ON TABLE m_todo_task IS 'TODOタスクマスタ';

COMMENT ON COLUMN m_todo_task.id IS 'ID';

COMMENT ON COLUMN m_todo_task.title IS 'Todoタイトル';

COMMENT ON COLUMN m_todo_task.contents IS '内容';

COMMENT ON COLUMN m_todo_task.assignee IS '担当者';

COMMENT ON COLUMN m_todo_task.label IS 'ラベル';

COMMENT ON COLUMN m_todo_task.status IS '進捗状況';

COMMENT ON COLUMN m_todo_task.start_date IS '開始予定日';

COMMENT ON COLUMN m_todo_task.end_date IS '終了予定日';

COMMENT ON COLUMN m_todo_task.created_at IS '作成時刻';

COMMENT ON COLUMN m_todo_task.updated_at IS '更新時刻';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_todo_task;

-- +goose StatementEnd