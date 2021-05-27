package schedule

import (
	"crypto/tls"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"proxy_pool/app/global"
	"proxy_pool/app/repositories"
	"proxy_pool/config"
	"reflect"
	"strconv"
	"time"
)

type CronLog struct {
	CronLog *logrus.Logger
}

type Job struct {
	Shut chan int `json:"shut"`
}

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
	global.Logger.Info("Start job to fetch proxy.")

	ch := make(chan map[string]interface{}, 10)
	
	finish := saveToDb(ch)

	for _, fetch := range config.FetcherList {
		ref := reflect.ValueOf(fetch).Type()
		elem := reflect.New(ref).Elem()
		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(10)
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

	global.Logger.Info("Exit Fetch Proxies Job.")
}

// 验证IP是否可用
func validate(proxy map[string]interface{}, ch chan map[string]interface{})  {
	// ch <- proxy
	request, _ := http.NewRequest("HEAD", "https://www.qq.com", nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	proxyUrlStr := proxy["protocol"].(string) + "://" + proxy["proxy"].(string)
	proxyURL, err := url.Parse(proxyUrlStr)
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10, //超时时间
	}

	resp, err := client.Do(request)
	if err != nil {
		global.Logger.Errorln(err)
		return
	}
	if status, err := strconv.Atoi(resp.Status); err != nil && status==http.StatusOK {
		ch <- proxy
	}
}


// 保存验证可用的ip到db.
func saveToDb(ch <-chan map[string]interface{}) <-chan struct{} {

	finish := make(chan struct{})

	go func() {
		defer func() {
			global.Logger.Info("Exit save proxies worker.")
			finish <- struct{}{}
			close(finish)
		}()

		for  {
			select {
			case m := <-ch:
				isExists := repositories.ProxyRepository{}.IsExists(m["proxy"].(string))
				if !isExists {
					repositories.ProxyRepository{}.Create(m)
				}
			case <-time.After(300 * time.Second):
				global.Logger.Info("Finished to save proxies.")
				return
			}
		}
	}()

	return finish
}



func StartServer()  {
	job := Job{
		Shut: make(chan int, 1),
	}
	// 每分钟执行一次
	go StartJob("* */1 * * *", job)
}

