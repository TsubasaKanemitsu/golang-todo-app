package transfer

import (
	"github.com/TsubasaKanemitsu/golang-todo-app/domain/model"
	"github.com/TsubasaKanemitsu/golang-todo-app/pkg/dbmodels"
)

func ToTodoTaskInfoEntity(m *model.TodoTaskInfo) *dbmodels.MTodoTask {
	return (*dbmodels.MTodoTask)(m)
}

func ToTodoTaskInfoDomain(e *dbmodels.MTodoTask) *model.TodoTaskInfo {
	return (*model.TodoTaskInfo)(e)
}
