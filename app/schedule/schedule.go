package schedule

import (
	"fmt"
	"github.com/asmcos/requests"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"proxy_pool/app/fetcher"
	"proxy_pool/app/global"
	"reflect"
	"time"
)

type CronLog struct {
	CronLog *logrus.Logger
}

type Job struct {
	Shut chan int `json:"shut"`
}

var FetcherList map[string]interface{}

// 普通日志记录
func (L *CronLog) Info(msg string, keysAndValues ...interface{}) {
	L.CronLog.WithFields(logrus.Fields{
		"data": keysAndValues,
	}).Info(msg)
}

// 错误日志记录
func (L *CronLog) Error(err error, msg string, keysAndValues ...interface{}) {
	L.CronLog.WithFields(logrus.Fields{
		"msg":  msg,
		"data": keysAndValues,
	}).Warn(msg)
}


// 开始任务
func StartJob(spec string, job Job) {

	logger := &CronLog{CronLog: global.Logger}

	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(logger)))

	c.AddJob(spec, &job)

	// 启动执行任务
	c.Start()

	// 退出时关闭计划任务
	defer c.Stop()

	select {
		case <-job.Shut:
			return
	}
}

// 停止任务
func StopJob(shut chan int) {
	shut <- 0
}


// Job执行爬取ip
func (j *Job) Run() {
	ch := make(chan map[string]interface{}, 10)
	
	finish := saveToDb(ch)

	for _, fetch := range FetcherList {
		ref := reflect.ValueOf(fetch).Type()
		elem := reflect.New(ref).Elem()
		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(20)
		data := elem.Method(0).Call(params)

		for i := 0; i < data[0].Len(); i++ {
			v := data[0].Index(i)
			 m := make(map[string]interface{}, 10)
			for _, k := range v.MapKeys() {
				m[k.String()] = v.MapIndex(k).Interface()
			}
			go validate(m, ch)
		}
	}

	<-finish
	fmt.Println("Job Exit...")
}

// 验证IP是否可用
func validate(proxy map[string]interface{}, ch chan map[string]interface{})  {
	request := requests.Requests()
	request.SetTimeout(time.Second * 10)
	proxyUrl := proxy["protocol"].(string) + "://" + proxy["proxy"].(string)
	fmt.Println(proxyUrl)
	request.Proxy(proxyUrl)
	resp, err := request.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.R.Status == "200" {
		ch <- proxy
	}

}


// 保存验证可用的ip到db.
func saveToDb(ch <-chan map[string]interface{}) <-chan struct{} {

	finish := make(chan struct{})

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		for  {
			select {
			case m := <-ch:
				fmt.Println(m)
			case <-time.After(30 * time.Second):
				fmt.Println("timed out")
				return
			}
		}
	}()

	return finish
}


func init()  {
	FetcherList = make(map[string]interface{})
	FetcherList["CloudFetcher"] = fetcher.CloudFetcher{}
}

func StartServer()  {
	job := Job{
		Shut: make(chan int, 1),
	}
	// 每分钟执行一次
	go StartJob("*/1 * * * *", job)
}

