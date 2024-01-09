package sentry_helper

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

var (
	helper = New()
)

func New() *Helper {
	return &Helper{
		skippedCaller:       3,
		parentSkippedCaller: 3,
	}
}

type RouteContext interface {
	GetContextName(interface{}) (context.Context, string) //return context and name for parent span
	GetInfo(interface{}) map[string]interface{}
}

type NamingRules interface {
	ManageChildName(string) string
	ManageChildOperation(string) string
	ManageParentName(string) string
	ManageParentOperation(string) string
}

type Helper struct {
	dsn                 string
	env                 string
	name                string // app name
	flushTime           string
	skippedCaller       int
	parentSkippedCaller int

	showSentryLog bool

	ctx    RouteContext
	naming NamingRules
}

func (h *Helper) SetDSN(dsn string) *Helper {
	h.dsn = dsn
	return h
}

func (h *Helper) SetRouter(routeContext RouteContext) {
	h.ctx = routeContext
}

func (h *Helper) SetNamingRules(namingRules NamingRules) {
	h.naming = namingRules
}

func (h *Helper) SetShowSentryLog(isShow bool) {
	h.showSentryLog = isShow
}

func (h *Helper) PrintLog(msg string, values ...any) {
	if h.showSentryLog {
		if len(values) == 0 {
			fmt.Println(msg)
		} else {
			fmt.Printf(msg+"\n", values...)
		}
	}
}

func getCaller(skip ...int) (description, function string) {
	//default skip
	defaultSkip := 3

	if len(skip) == 1 {
		defaultSkip = skip[0]
	}

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(defaultSkip, pc)

	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	description = fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	function = frame.Function

	return
}

func getFunction(function string) string {
	temp := strings.Split(function, ".")

	if len(temp) != 0 {
		return temp[len(temp)-1]
	}

	return function
}
