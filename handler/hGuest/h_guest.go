package hGuest

import (
	"be04gomy/handler"
	"be04gomy/model/mStudent"
	"encoding/json"
	"net/http"
)

func StudentList(ctx *handler.Ctx) {
	students, err := mStudent.SelectAll(ctx.Db)
	if ctx.IsError(err) {
		return
	}
	ctx.End(students)
}

func StudentCreate(ctx *handler.Ctx) {
	if ctx.Request.Method == `GET` {
		http.ServeFile(ctx,ctx.Request,ctx.ViewsDir+`guest/student_create.html`)
		return
	}
	m := mStudent.Student{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&m)
	if ctx.IsError(err) {
		return
	}
	err = mStudent.Insert(ctx.Db,&m)
	if ctx.IsError(err) {
		return
	}
	ctx.End(m)
}

func StudentUpdate(ctx *handler.Ctx) {
	
}

func StudentDelete(ctx *handler.Ctx) {
	
}
