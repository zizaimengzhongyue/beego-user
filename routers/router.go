package routers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"github.com/zizaimengzhongyue/beego-user/controllers"
)

var (
	log = logs.NewLogger(10000)
)

func init() {
	log.SetLogger("file", `{"filename": "log/access.log"}`)
	logs.SetLogger("file", `{"filename": "log/beego_user.log"}`)

	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8", 30)

	beego.InsertFilter("*", beego.BeforeRouter, SetRequestTime, false)
	beego.InsertFilter("*", beego.FinishRouter, PrintAccessLog, false)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/getAllUser", &controllers.UserController{}, "*:GetAll")
	beego.Router("/user/delete/:uid", &controllers.UserController{}, "*:Delete")
	beego.Router("/user/add/:uid/:name", &controllers.UserController{}, "*:Add")
	beego.Router("/user/update/:id/:uid/:name", &controllers.UserController{}, "*:Update")
	beego.Router("/user/find/:uid", &controllers.UserController{}, "*:Find")
}

func SetRequestTime(ctx *context.Context) {
	ctx.Input.SetData("request_time", time.Now())
}

func PrintAccessLog(ctx *context.Context) {
	reqTime, _ := ctx.Input.GetData("request_time").(time.Time)
	log.Write([]byte(GeneralAccessLog(ctx, &reqTime, ctx.Output.Status)))
}

func GeneralAccessLog(ctx *context.Context, startTime *time.Time, statusCode int) string {
	var (
		requestTime time.Time
		elapsedTime time.Duration
		r           = ctx.Request
	)
	if startTime != nil {
		requestTime = *startTime
		elapsedTime = time.Since(*startTime)
	}
	record := &logs.AccessLogRecord{
		RemoteAddr:     ctx.Input.IP(),
		RequestTime:    requestTime,
		RequestMethod:  r.Method,
		Request:        fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto),
		ServerProtocol: r.Proto,
		Host:           r.Host,
		Status:         statusCode,
		ElapsedTime:    elapsedTime,
		HTTPReferrer:   r.Header.Get("Referer"),
		HTTPUserAgent:  r.Header.Get("User-Agent"),
		RemoteUser:     r.Header.Get("Remote-User"),
		BodyBytesSent:  0, //@todo this one is missing!
	}
	return fmt.Sprintf("%s - - \"%s %d %d\" %f %s %s", record.RemoteAddr, record.Request, record.Status, record.BodyBytesSent,
		record.ElapsedTime.Seconds(), record.HTTPReferrer, record.HTTPUserAgent)
}
