package design

import (
	dsl "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

// API describes the global properties of the API server.
var _ = dsl.API("todo", func() {
	dsl.Title("Todoタスク管理サービス")
	dsl.Description("HTTP service for todo management")
	dsl.Server("todo", func() {
		dsl.Host("localhost", func() { dsl.URI("http://localhost:8088") })
	})
})

// Service describes a service
var _ = dsl.Service("todoservice", func() {
	dsl.Description("Todoタスク管理サービス")
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("X-Shared-Secret", "X-Authorization", "content-type")
		cors.Methods("GET", "POST", "PUT", "DELETE", "OPTIONS")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})

	// CREATE
	dsl.Method("addTodoTask", func() {
		dsl.Description("Todoタスクを追加する。")
		dsl.Payload(func() {
			dsl.Attribute("title", dsl.String, "Todoタスクのタイトル")
			dsl.Attribute("contents", dsl.String, "Todoタスクの説明")
			dsl.Attribute("label", dsl.String, "Todoタスクのラベル")
			dsl.Attribute("asignee", dsl.String, "タスクを割り当てられた人の名前")
			dsl.Attribute("start_date", dsl.String, "Todoタスクの開始予定日", func() {
				dsl.Format(dsl.FormatDate)
			})
			dsl.Attribute("end_date", dsl.String, "Todoタスクの終了予定日", func() {
				dsl.Format(dsl.FormatDate)
			})
			dsl.Required("title", "start_date", "end_date")
		})
		dsl.HTTP(func() {
			dsl.POST("/add")
			dsl.Response(dsl.StatusOK)
		})
		dsl.Result(dsl.Boolean)
	})
	// READ
	dsl.Method("GetTodoTask", func() {
		dsl.Description("指定したTodoタスクの詳細を取得する。")
		dsl.Payload(func() {
			dsl.Attribute("id", dsl.Int, "Todo task id")
			dsl.Required("id")
		})
		dsl.HTTP(func() {
			dsl.GET("/todo/{id}")
			dsl.Response(dsl.StatusOK)
		})
		dsl.Result(todoTaskInfo, "Todo task")
	})
	dsl.Method("GetTodoTaskList", func() {
		dsl.Description("Todoタスク一覧を取得する。")
		dsl.HTTP(func() {
			dsl.GET("/todolist")
			dsl.Response(dsl.StatusOK)
		})
		dsl.Result(dsl.ArrayOf(todoTaskTitleList), "Todo task list")
	})
	// UPDATE
	dsl.Method("UpdateTodoTask", func() {
		dsl.Description("指定したTodoタスクを更新する。")
		dsl.Payload(func() {
			dsl.Attribute("id", dsl.Int, "Todo task id")
			dsl.Required("id")
		})
		dsl.HTTP(func() {
			dsl.PUT("/todo/{id}")
			dsl.Response(dsl.StatusOK)
		})
		dsl.Result(dsl.Boolean)
	})
	// DELETE
	dsl.Method("DeleteTodoTask", func() {
		dsl.Description("指定したTodoタスクを削除する。")
		dsl.Payload(func() {
			dsl.Attribute("id", dsl.Int, "Todo task id")
			dsl.Required("id")
		})
		dsl.HTTP(func() {
			dsl.DELETE("/todo/{id}")
			dsl.Response(dsl.StatusOK)
		})
		dsl.Result(dsl.Boolean)
	})
})

var todoTaskInfo = dsl.Type(("todoTaskInfo"), func() {
	dsl.Description("Todoタスクの情報")
	dsl.Attribute("id", dsl.Int, "TodoタスクのユニークID")
	dsl.Attribute("title", dsl.String, "Todoタスクのタイトル")
	dsl.Attribute("contents", dsl.String, "Todoタスクの説明")
	dsl.Attribute("status", dsl.String, "Todoタスクの進捗状況")
	dsl.Attribute("label", dsl.String, "Todoタスクのラベル")
	dsl.Attribute("asignee", dsl.String, "タスクを割り当てられた人の名前")
	dsl.Attribute("start_date", dsl.String, "Todoタスクの開始予定日", func() {
		dsl.Format(dsl.FormatDate)
	})
	dsl.Attribute("end_date", dsl.String, "Todoタスクの終了予定日", func() {
		dsl.Format(dsl.FormatDate)
	})
	dsl.Required("id", "title", "status")
})

var todoTaskTitleList = dsl.Type(("todoTaskTitleList"), func() {
	dsl.Description("Todoタスクの情報")
	dsl.Attribute("id", dsl.Int, "TodoタスクのユニークID")
	dsl.Attribute("title", dsl.String, "Todoタスクのタイトル")
	dsl.Attribute("status", dsl.String, "Todoタスクの進捗状況")
	dsl.Attribute("label", dsl.String, "Todoタスクのラベル")
	dsl.Required("id", "title", "status")
})
