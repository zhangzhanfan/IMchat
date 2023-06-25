package utils

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//封装一个日志中间件
func Logger() gin.HandlerFunc {
	file, err := os.OpenFile("./logger.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777) // 打开日志文件
	if err != nil {
		println("打开日志文件失败!", err)
	}
	log := logrus.New() //实例化logrus对象
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		DisableLevelTruncation:    true,
	})
	log.SetLevel(logrus.TraceLevel) //设置日志级别
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)         //设置终端和文件一起标准输出
	log.SetReportCaller(true) //打印文件路径和行号
	// log.WithFields(logrus.Fields{

	// })
	// log.Println("log test")
	// log.Error("Error")
	// log.Info("Info")
	// log.Warn("Warn")

	// //设置日志切割器
	// logWriter, _ := rotatelogs.New(
	// 	//分割后的日志文件名
	// 	// fileName+"_%Y%m%d.log",
	// 	"./logger.log",
	// 	//生成软链，指向最新日志文件
	// 	rotatelogs.WithLinkName("./logger.log"),
	// 	// 设置最大保存时间(7天)
	// 	rotatelogs.WithMaxAge(7*24*time.Hour),
	// 	// 设置日志切割时间间隔(1天)
	// 	rotatelogs.WithRotationTime(24*time.Hour),
	// )

	// writerMap := lfshook.WriterMap{
	// 	logrus.InfoLevel:  logWriter,
	// 	logrus.DebugLevel: logWriter,
	// 	logrus.WarnLevel:  logWriter,
	// 	logrus.FatalLevel: logWriter,
	// 	logrus.ErrorLevel: logWriter,
	// 	logrus.PanicLevel: logWriter,
	// }
	// //设置格式
	// lfshook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
	// 	TimestampFormat: "2023-04-28 22:37:08",
	// })
	// log.AddHook(lfshook)
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//执行操作
		c.Next()
		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		// clientIP := c.ClientIP()
		// 日志格式
		log.WithFields(logrus.Fields{
			"status_code": statusCode,
			"执行时间":        latencyTime,
			// "client_ip":   clientIP,
			"req_method": reqMethod,
			"req_uri":    reqUri,
		}).Println()
	}

}
