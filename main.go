package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"temp/consumer"
	"temp/producer"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/patrickmn/go-cache"
	"github.com/urfave/cli"
	"github.com/widuu/goini"
	"golang.org/x/sync/errgroup"
	"gopkg.in/ini.v1"
)

// type a struct {
// 	b string
// 	c string
// }

// type b struct {
// 	b string
// }

func main() {
	// app := gin.Default()
	// app.GET("/index/:ID", func(c *gin.Context) {
	// 	id := c.Param("ID")

	// 	if id == "" {
	// 		c.JSON(200, gin.H{
	// 			"message": "err",
	// 		})
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"message": "Hi" + id,
	// 	})
	// })
	// app.Run(":8080")
	// a := new(a)
	// a.c = "2"
	// b := new(b)
	// b.b = "3"
	// s := reflect.ValueOf(a)
	// d := reflect.ValueOf(b)
	// fmt.Println(d)
	// v := s.Elem().Field(0)
	// fieldName := s.Elem().Type().Field(0).Name
	// skip := s.Elem().Type().Field(0).Tag.Get("structs")
	// fmt.Println(v.Kind() > reflect.Float64)
	// fmt.Println(s.Elem().NumField(), "--", v, "--", fieldName, "--", skip, "--", v.Kind())
	// Merge(a, b)
	// fmt.Print(a)
	// err := isPassIP("123.207.254.129")
	// log.Println("----", err)
	// switch ft(); {
	// case true:
	// 	fmt.Println(true)
	// case false:
	// 	fmt.Println(false)
	// }
	//PEDU:精准教育,TUTOR:辅导机构,FLAGSHIP:旗舰店,STC:大学云

	// get1 := make(map[string]string)
	// get1["PEDU"] = "精准教育"
	// get1["TUTOR"] = "辅导机构"
	// get1["FLAGSHIP"] = "旗舰店"
	// for k, v := range get1 {
	// 	log.Println(k, "--", v)
	// }
	// str := "2021-06-24 10:56:14"
	// payTime := strings.Split(str, " ")
	// str1 := strings.Split(payTime[0], "-")
	// date := time.Now().Format("20060102")
	// payTimeTemp := str1[0] + str1[1] + str1[2]

	// log.Println(str1[0], "--", str1[1], "--", str1[2], "--", str, "--", str1, "--", date, "--", payTimeTemp)

	// var x float64 = 398.00 * 100
	// var y = int(x)
	// var j int = 1
	// c := gron.New()
	// c.AddFunc(gron.Every(5*time.Second), func() {
	// 	log.Println("三点钟了喂！,够钟饮茶啦！")
	// })
	// c.Start()
	// log.Println("你好")
	// for {
	// 	log.Println("----")
	// 	time.Sleep(time.Second * 4)
	// }

	// cron2 := cron.New() //创建一个cron实例
	// var oldDate string
	// date := time.Now().Format("20060102")
	// oldDate = date
	// log.Println("old", oldDate)
	// //执行定时任务（每5秒执行一次）
	// cron2.AddFunc("*/5 * * * * *", func() {
	// 	date := time.Now().Format("20060102")
	// 	log.Println("data", date)
	// 	if oldDate != date {
	// 		log.Println("执行定期任务")
	// 		oldDate = date
	// 	}
	// })
	// cron2.Start()
	// c := cron.New()
	// spec := "0 54 16 * * ?"
	// c.AddFunc(spec, func() {
	// 	log.Println("每天执行一次定时任务")
	// })
	// hello()
	// spec1 := "0 59 09 * * *"
	// spec2 := "0 00 10 * * *"

	// c := cron.New()
	// c.AddFunc(spec1, func() {
	// 	log.Println("每天执行一次定时任9务1")
	// })
	// c.AddFunc(spec2, func() {
	// 	log.Println("每天执行一次定时任务2")
	// })
	// c.Start()

	// is := []int{}
	// var tempstr int
	// tempstr++
	// is1 := []int{1, 2, 1, 5, 6, 1}
	// is2 := make(map[int]int)
	// for _, v := range is1 {
	// 	is2[v] = 2
	// }
	// _, ok := is2[1]
	// if ok {
	// 	is2[1+tempstr] = 100
	// }
	// is1 := ""
	// is2 := strings.Split(is1, ",")
	// log.Println(is2, "--")
	// log.Println(is2)
	// for k, v := range is1 {
	// 	if !is2[k] {
	// 		log.Println("--")
	// 	}
	// }
	// lock := make(map[int]int)
	// for _, v := range is {
	// 	lock[v] = v
	// }
	// log.Println(lock)
	// time_ticker.RegisterTimeTicker()
	// s := GetUUID()
	// log.Println(s)
	// http.HandleFunc("/", IndexHandler)
	// http.ListenAndServe("127.0.0.1:8081", nil)
	// crateday := GetBetweenDates("2020-07-01", "2021-10-01")
	// s := RemoveRepeatedElement(crateday)
	// log.Println(s)
	// now := time.Now()
	// today, _ := time.Parse("2006-01-02", "2021-09-08")
	// s := today.After(now)
	// log.Println(s)
	// teamid := ""
	// _, err := strconv.Atoi(teamid)
	// if err != nil {
	// 	log.Println("---")
	// }
	// defulttime := "2021-08-25"
	// GdTime, _ := time.ParseInLocation("2006-01-02 15:04:05", defulttime+" 00:00:00", time.Local)
	// defulttime1 := "2021-08-66"
	// _, err := time.Parse("2006-01-02", defulttime1)
	// if err != nil {
	// 	log.Println(err)
	// }
	// GdTime1, _ := time.ParseInLocation("2006-01-02 15:04:05", defulttime1+" 00:00:00", time.Local)
	// log.Println(GdTime)
	// s := GdTime.Before(GdTime1)
	// log.Println(s)
	// _, err := strconv.Atoi("")
	// if err != nil {
	// 	log.Println(`----`)
	// config.Redis.Set("token", "1234567", 0)
	// config.Redis.Set("cookie", "123456a", 100)
	// cookieisNull, _ := config.Redis.Exists("cookie")
	// log.Println(cookieisNull)
	// redisstr, err := config.Redis.GetString("token")
	// if err != nil {
	// 	log.Println(err)
	// }
	// for {
	// 	_, err := config.Redis.GetString("token")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }
	// time.Sleep(time.Second * 5)
	// config.Redis.Del("cookie")
	// log.Println(redisstr)
	// config.Redis.Del("token")
	// err = config.Redis.Del("token")
	// if err != nil {
	// 	log.Println(err)
	// }

	// }
	// // log.Println(crateday)
	// log.Println(time.Now().AddDate(0, 0, -1))
	// var get = make([]string, 0)
	// for _, v := range crateday {
	// 	get = append(get, "4"+"_"+v)
	// }
	// wuhu := strings.Join(get, ",")
	// log.Println(get, wuhu)
	// createdateIsnull, _ := time.s("2006-01-02 15:04:05", "0001-01-01 00:00:00 +0000 UTC")
	// log.Println(createdateIsnull)
	// result := GetToken()
	// log.Println("code", MD5("123456a")
	// timeStamp, err := time.Parse("2006-01-02", "2021-08-16")

	// log.Println(timeStamp.Unix())
	// DATE_FORMAT := "2006-01-02"
	// dateTime, _ := time.Parse("2006-01", "2021-08")
	// year, month, _ := dateTime.AddDate(0, 1, 0).Date()
	// thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	// start := thisMonth.AddDate(0, -1, 0).Format(DATE_FORMAT)
	// end := thisMonth.AddDate(0, 0, -1).Format(DATE_FORMAT)
	// timeRange := fmt.Sprintf("%s~%s", start, end)
	// fmt.Println(timeRange)
	// var i = []int{-1, 0, 1, -2, 3, 1, -2, 3, 4, -3, 2, 0}
	// s := maxArea(i)
	// s := BubbleAsort(i)
	// q := threeSum(i)
	// go go_test()
	// for {
	// 	log.Println("main to sleep")
	// 	time.Sleep(time.Second * 2)
	// }
	//匿名子协程
	// go func() {
	// 	defer fmt.Println("A.defer")
	// 	//匿名函数
	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		//此时只有defer执行
	// 		runtime.Goexit()
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()
	// limiter := NewSliding(100*time.Millisecond, time.Second, 10)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(limiter.LimitTest())
	// }
	// time.Sleep(100 * time.Millisecond)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(limiter.LimitTest())
	// }
	// fmt.Println(limiter.LimitTest())
	// for _, v := range limiter.windows {
	// 	fmt.Println(v.timestamp, v.count)
	// }

	// fmt.Println("moments later...")
	// time.Sleep(time.Second)
	// for i := 0; i < 7; i++ {
	// 	fmt.Println(limiter.LimitTest())
	// }
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	// wg.Add(1)
	// go cat()
	// go dog()
	// cc <- struct{}{}
	// wg.Wait()
	// ch := make(chan int)
	// wg := sync.WaitGroup{}
	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 5; i++ {
	// 		ch <- i
	// 	}
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	for c := range ch {
	// 		log.Println(c)
	// 	}
	// }()
	// wg.Wait()
	// for {
	// 	time.Sleep(time.Minute * 10)
	// }
	// for _, v := range limiter.windows {
	// 	fmt.Println(v.timestamp, v.count)
	// }
	// for {
	// 	time.Sleep(time.Second)
	// }
	// limiter := rate.NewLimiter(10, 1000)
	// for {
	// 	log.Println(limiter.WaitN(context.Background(), 100))
	// 	log.Println(limiter.AllowN(time.Now(), 100))
	// 	time.Sleep(time.Second * 3)
	// }
	// ci := make(chan int, 1)
	// go write(ci, 4)
	// go write(ci, 5)
	// go write(ci, 6)
	// ch := make(chan struct{}, 0)
	// mLock.Add(1)
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// go test_chan1(ch, ctx)
	// ch <- struct{}{}
	// time.Sleep(1)
	// cancel()
	// mLock.Wait()
	// value := <-ci
	// value1 := <-ci
	// value2 := <-ci
	// fmt.Println(value, value1, value2)
	// i := []int{1}
	// n2 := []int{3, 4}
	// var n1 = make([][]int, 3)
	// n1[0] = []int{1}
	// n1[0] = []int{1, 3, 5, 7}
	// n1[1] = []int{10, 11, 16, 20}
	// n1[2] = []int{23, 30, 34, 60}
	// q := threeSumClosest(i, 2)
	// sort.Ints(i)
	// log.Println(i)
	// s := threeSumClosest(i, 1)
	// s := generateParenthesis(3)
	// s := fourSum(i, 0)
	// s := findMedianSortedArrays(n1, n2)
	// s := searchRanges(i, 8)
	// s := mySqrt(4)
	// s := searchMatrix(n1, 0)
	// s := jump(i)
	// s := canJump(i)
	// // s := maxProfit(i)
	// s := largestNumber(i)
	// i := "   fly me   to   the moon  "
	// s := lengthOfLastWord(i)
	// s := lengthOfLIS(i)
	// s := twoSum(i, 7)
	// s := isPowerOfThree(0)
	// s := compareVersion("1.0.0", "1.0.0.1")
	// s := smallestK(i, 2)
	// s := fib(66)
	// s := test1("asaxx")
	// s := strings.Count("asaxx", "a")
	// t := []rune("asds")
	// s := unicode.IsLetter(t[0])
	// pase_student()
	// t := time.NewTicker(time.Second * 1)
	// var i = []int{10, 9, 2, 5, 3, 7, 101, 18}
	// s := search(i, 3)
	// s := 3
	// s |= s >> 2
	// log.Println("|", s, "|")
	// ip := "192.168.3.97"
	// intrnalIp := ip[:strings.LastIndex(ip, ".")+1]
	// log.Println(intrnalIp)
	// var ipList []string
	// ipList = append(ipList, "192.168.1.1", "192.168.1.2", "192.168.1.3")
	// var ch = make(chan string)
	// for range ipList {
	// 	wg.Add(2)
	// 	go a(ch, &wg)
	// 	go b(ch, &wg)
	// }
	// wg.Wait()
	// log.Println("Demo Finish")
	// var os = "Microsoft Windows 10 1709 - 1909"
	// var OsType string
	// if strings.Contains(os, "Linux") {
	// 	OsType = "Linux"
	// } else if strings.Contains(os, "Windows") {
	// 	OsType = "Windows"
	// } else {
	// 	OsType = "UnKnow"
	// }
	// log.Println(OsType)
	// file := excelize.NewFile()
	// streamWriter, err := file.NewStreamWriter("Sheet1")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := streamWriter.SetRow("A1", []interface{}{
	// 	excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
	// 	fmt.Println(err)
	// }
	// for rowID := 2; rowID <= 10; rowID++ {
	// 	row := make([]interface{}, 2)
	// 	for colID := 0; colID < 2; colID++ {
	// 		row[colID] = "test"
	// 	}
	// 	cell, _ := excelize.CoordinatesToCellName(1, rowID)
	// 	if err := streamWriter.SetRow(cell, row); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// if err := streamWriter.Flush(); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := file.SaveAs("Book1.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	// type t interface{}
	// var x, y t
	// x = []X{{S: "xxx"}, {S: "zzz"}}
	// y = []Y{{N: 4}, {N: 5}, {N: 3}}
	// Process(x, y)
	// var col string
	// var num = 3
	// for num > 0 {
	// 	col = string(rune((num-1)%26+65)) + col
	// 	num = (num - 1) / 26
	// }
	// log.Println(col)
	// group := ColorGroup{
	// 	ID:     1,
	// 	Name:   "Reds",
	// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	// }
	// b, err := json.Marshal(group)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// err := errors.New("数据错无")
	// err.Error()
	//生成json文件
	// err = ioutil.WriteFile("test.json", b, os.ModeAppend)
	// if err != nil {
	// 	return
	// }
	// var data interface{}
	// err = json.Unmarshal(b, &data)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// dataJson := data.(map[string]interface{})["Colors"]
	// b11, err := json.Marshal(dataJson)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// err = ioutil.WriteFile("test11.json", b11, os.ModeAppend)
	// if err != nil {
	// 	return
	// }
	// WriteConfig("./host.json", b)
	// var req = make([]map[string]string, 0)

	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		go task1(req[0])
	// 		go task1(req[1])

	// 	}
	// }()
	// time.Sleep(5 * time.Second)
	// req[1]["jug"] = "大司马"
	// time.Sleep(10 * time.Second)
	// m := map[string]interface{}{"test": Test}
	// ret, err := Call(m, "test", "hello", "world")
	// if err != nil {
	// 	fmt.Println("method invoke error:", err)
	// }
	// fmt.Println(ret)
	// s := "aaa23fd54fds57gdsf89,7,9"
	// r, _ := regexp.Compile("[1-9]")
	// fmt.Println(r.FindString(s))
	// timer1 := time.NewTimer(2 * time.Second)
	// <-timer1.C
	// fmt.Println("Timer 1 fired")
	// timer2 := time.Until(timer)
	// log.Println(timer, timer2)
	// go func() {
	// 	<-timer2.C
	// 	fmt.Println("Timer 2 fired")
	// }()
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
	// infos, _ := cpu.Info()
	// for _, info := range infos {
	// 	data, _ := json.MarshalIndent(info, "", " ")
	// 	fmt.Print(string(data))
	// }
	// timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C
	// fmt.Println("Timer 1 fired")
	// fmt.Printf("time.Until(timer): %v\n", time.Until(timer))
	// start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	// oneDayLater := start.AddDate(0, 0, 1)
	// timenow := time.Now()
	// timer := time.Second * 1
	// timer += time.Second * 1
	// timer2 := time.NewTimer(timer)
	// <-timer2.C
	// timer, _ := time.Parse("2006-01-02 15:04:05", "2021-11-23 18:43:31")
	// timer, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-11-23 15:50:31", time.Local)
	// log.Println(timer)
	// times := time.Until(timer.Local()).String()
	// hour := strings.Index(times, "h")
	// minute := strings.Index(times, "m")
	// second := strings.Index(times, ".")
	// log.Println(times[:hour] + times[hour+1:minute] + times[minute+1:second])
	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer 2 stopped")
	// }
	// log.Println(times)
	// s := Student{Id: 1, Name: "咖啡色的羊驼"}

	// // 获取目标对象
	// t := reflect.TypeOf(s)
	// // .Name()可以获取去这个类型的名称
	// fmt.Println("这个类型的名称是:", t.Name())

	// // 获取目标对象的值类型
	// v := reflect.ValueOf(s)
	// // .NumField()来获取其包含的字段的总数
	// for i := 0; i < t.NumField(); i++ {
	// 	// 从0开始获取Student所包含的key
	// 	key := t.Field(i)

	// 	// 通过interface方法来获取key所对应的值
	// 	value := v.Field(i).Interface()

	// 	fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	// }

	// // 通过.NumMethod()来获取Student里头的方法
	// for i := 0; i < t.NumMethod(); i++ {
	// 	m := t.Method(i)
	// 	fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	// }

	// logfile.Log.Info("wowo")
	// logfile.log.Printlnf("key", "wocao")
	// logfile.Loginit.Warn("wuhuhu")
	// body := getToken()
	// fmt.Println(body)
	//设置配置
	// config := sarama.NewConfig()
	// //等待服务器所有副本都保存成功后的响应
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// //随机的分区类型
	// config.Producer.Partitioner = sarama.NewRandomPartitioner
	// //是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	// config.Producer.Return.Successes = true
	// config.Producer.Return.Errors = true
	// //设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	// config.Version = sarama.V0_11_0_0

	// //使用配置,新建一个异步生产者
	// producer, e := sarama.NewAsyncProducer([]string{"IP:9092", "IP:9092", "IP:9092"}, config)
	// if e != nil {
	// 	panic(e)
	// }
	// defer producer.AsyncClose()

	// //发送的消息,主题,key
	// msg := &sarama.ProducerMessage{
	// 	Topic: "logstash_test",
	// 	Key:   sarama.StringEncoder("test"),
	// }

	// var value string
	// for {
	// 	value = "this is a message"
	// 	//设置发送的真正内容
	// 	fmt.Scanln(&value)
	// 	//将字符串转化为字节数组
	// 	msg.Value = sarama.ByteEncoder(value)
	// 	fmt.Println(value)

	// 	//使用通道发送
	// 	producer.Input() <- msg

	// 	//循环判断哪个通道发送过来数据.
	// 	select {
	// 	case suc := <-producer.Successes():
	// 		fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
	// 	case fail := <-producer.Errors():
	// 		fmt.Println("err: ", fail.Err)
	// 	}
	// }
	// size, err := CheckLocalSecurePolicies()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(size)

	// time.Sleep(time.Hour * 10)
	// var f http.File
	// var err error
	// f, err = os.Open("./json")
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// defer f.Close()
	// fileList, err := f.Readdir(2000)
	// if err != nil {
	// 	log.Println(err.Error(), "json")
	// 	return
	// }
	// var jj []JsonText
	// for _, v11 := range fileList {
	// 	if strings.ToLower(path.Ext("json/"+v11.Name())) != ".json" {
	// 		continue
	// 	}
	// 	f, err := os.Open("json/" + v11.Name())
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	r := io.Reader(f)
	// 	var jsons JSONPlugin
	// 	json.NewDecoder(r).Decode(&jsons)
	// 	bytes, _ := json.Marshal(jsons)
	// 	k, v := ExampleNewGCMEncrypter(bytes)
	// 	jj = append(jj, JsonText{
	// 		Seed:  fmt.Sprintf("%x", k),
	// 		Value: fmt.Sprintf("%x", v),
	// 	})
	// }
	// bytesssss, _ := json.Marshal(jj)
	// if err = ioutil.WriteFile("996.json", bytesssss, 0666); err != nil {
	// 	log.Println("WriteFile Error =", err)
	// }
	// f1, err := os.Open("996.json")
	// if err != nil {
	// 	log.Println(err)
	// }
	// r1 := io.Reader(f1)
	// var jsonssss []JsonText
	// json.NewDecoder(r1).Decode(&jsonssss)
	// for _, val := range jsonssss {
	// 	jsonss, _ := ExampleNewGCMDecrypter(val.Seed, val.Value)
	// 	var jsonsssss JSONPlugin
	// 	json.Unmarshal(jsonss, &jsonsssss)
	// 	log.Println(jsonsssss)
	// }
	// var stopLock sync.Mutex
	// stopChan := make(chan struct{}, 1)
	// signalChan := make(chan os.Signal, 1)
	// go func() {
	// 	//阻塞程序运行，直到收到终止的信号
	// 	<-signalChan
	// 	stopLock.Lock()
	// 	stopLock.Unlock()
	// 	log.Println("Cleaning before stop...")
	// 	stopChan <- struct{}{}
	// 	os.Exit(0)
	// }()
	// signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	// //模拟一个持续运行的进程
	// time.Sleep(10 * time.Second)
	// 构建 生产者
	// 生成 生产者配置文件
	// groupID := "group-1"
	// topicList := "topic_1"
	// config := cluster.NewConfig()
	// config.Consumer.Return.Errors = true
	// config.Group.Return.Notifications = true
	// config.Consumer.Offsets.CommitInterval = 1 * time.Second
	// config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始
	// c, err := cluster.NewConsumer(strings.Split("192.168.3.128:9093", ","), groupID, strings.Split(topicList, ","), config)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer c.Close()
	// go func() {
	// 	for err := range c.Errors() {
	// 		log.Println(err.Error())
	// 	}
	// }()

	// go func() {
	// 	for note := range c.Notifications() {
	// 		log.Println(note)
	// 	}
	// }()

	// for msg := range c.Messages() {
	// 	fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
	// 	c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	// }
	// go startProduce()
	// var err error
	// if err != nil {
	// 	fmt.Printf("fail to start consumer,err:%v\n", err)
	// 	return
	// }
	// topics := []string{"Inspection", "mysql", "nsq"}
	// //topics := []string{"nsq"}
	// //topic := "nsq"
	// for _, topic := range topics {
	// 	ctx, _ := context.WithCancel(context.Background())
	// 	go testTopic(topic, ctx)
	// }
	// select {}
	// // url := os.Args[1]
	// url := "https://www.zhihu.com/search?q=golang&type=column"
	// // 根据URL获取资源
	// res, err := http.Get(url)
	// res.Header.Add("referrer-policy", "no-referrer-when-downgrade")
	// res.Header.Add("set-cookie", "KLBRSID=4843ceb2c0de43091e0ff7c22eadca8c|1643184663|1643175688; Path=/")

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	// 	os.Exit(1)
	// }

	//读取资源数据 body: []byte
	// body, err := ioutil.ReadAll(res.Body)

	// 关闭资源流
	// res.Body.Close()

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	// 	os.Exit(1)
	// }

	// 控制台打印内容 以下两种方法等同
	// fmt.Printf("%s", body)
	// fmt.Printf(string(body))

	// // 写入文件
	// ioutil.WriteFile("site.txt", body, 0644)
	// res, err := http.Get("https://zhuanlan.zhihu.com/p/97113566")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }
	// log.Println(res)
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	log.Println(err)
	// }
	// doc.Find(".Card SearchResult-Card .List-item .ContentItem .ContentItem-main .ContentItem-head .ContentItem-title").Each(func(i int, contentSelection *goquery.Selection) {
	// 	title := contentSelection.Find("a").Text()
	// 	log.Println("第", i+1, "个帖子的标题：", title)
	// })
	// const columnName = "c_1170012477445263360"
	// pinnedArticlePidAndAuthor, err := GetPinnedArticlePidAndAuthor(columnName)
	// if err != nil {
	// 	log.Println(err)
	// }
	// pinnedArticle, err := GetSingleArticle(pinnedArticlePidAndAuthor.ID)
	// if err != nil {
	// 	log.Println(err)
	// }
	// pids, err := GetArticlesListPids(columnName)
	// if err != nil {
	// 	log.Println(err)
	// }
	// for _, pid := range pids {
	// 	if pid == pinnedArticle.ID {
	// 		continue
	// 	}
	// 	article, err := GetSingleArticle(pid)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	s1 := strings.Replace(article.URL, "api", "zhuanlan", -1)
	// 	s1 = strings.Replace(s1, "articles", "p", -1)
	// 	Islike(article.Title, s1)
	// 	// log.Println(article.Title + ":" + s1)
	// }
	// log.Println(IsLikereq)
	// q := qqwry.NewQQwry("qqwry.dat")
	// ip := "31.173.137.47"
	// q.Find(ip)
	// // result := GetEvilIp(ip)
	// CACHE_IPS = cache.New(cache.NoExpiration, cache.DefaultExpiration)
	// gob.Register(models.IpList{})
	// err := CACHE_IPS.LoadFile("ips")
	// if err != nil {
	// 	log.Println(err)
	// }
	// v, has := CACHE_IPS.Get(ip)
	// count, _ := CacheStatus(CACHE_IPS)
	// data, ok := v.(models.IpList)

	// if !ok {
	// 	log.Printf("ip: %v, ret: %v, count:%v", ip, data, count)
	// }
	// log.Println(fmt.Sprintf("IP:%s,恶意IP:%v,IP地区:%s%s", ip, has, q.Country, q.City))
	// RegisterService()
	// GetService("Order")
	// str := `{
	// 	"service_name": "pokemon",
	// 	"service_desc": "pokemonnnnnn",
	// 	"rule_type": 0,
	// 	"rule": "/pokemon",
	// 	"need_https": 0,
	// 	"need_websocket": 0,
	// 	"need_strip_uri": 1,
	// 	"url_rewrite": "",
	// 	"header_transfor": "",
	// 	"round_type": 1,
	// 	"ip_list": "127.0.0.1:2010,127.0.0.1:2011,127.0.0.1:2012",
	// 	"weight_list": "50,50,50",
	// 	"upstream_connect_timeout": 0,
	// 	"upstream_header_timeout": 0,
	// 	"upstream_idle_timeout": 0,
	// 	"upstream_max_idle": 0,
	// 	"open_auth": 0,
	// 	"black_list": "",
	// 	"white_list": "",
	// 	"clientip_flow_limit": 0,
	// 	"service_flow_limit": 0
	// }`
	// RegisterService(str)
	// server01 := &http.Server{
	// 	Addr:         ":2010",
	// 	Handler:      router01(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// server02 := &http.Server{
	// 	Addr:         ":2011",
	// 	Handler:      router02(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// server03 := &http.Server{
	// 	Addr:         ":2012",
	// 	Handler:      router03(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// g.Go(func() error {
	// 	return server01.ListenAndServe()
	// })
	// g.Go(func() error {
	// 	return server02.ListenAndServe()
	// })
	// g.Go(func() error {
	// 	return server03.ListenAndServe()
	// })
	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
	// iot, err := ioutil.ReadFile("temp.txt") //读取文件
	// if err != nil {                         //做错误判断
	// 	fmt.Printf("读取文件错误,错误为:%v\n", err)
	// 	return
	// }
	// var ss Soul
	// ss.Type = "md5"
	// file, err := os.OpenFile("temp1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Printf("文件错误,错误为:%v\n", err)
	// 	return
	// }
	// defer file.Close()
	// for _, val := range strings.Split(string(iot), "\n") {
	// 	if !strings.Contains(val, "header") && !strings.Contains(val, "title") {
	// 		continue
	// 	}
	// 	file.Write([]byte(strings.Split(val, " ")[1]))
	// 	// var s Trait
	// 	// sl := strings.Split(val, "|")
	// 	// if len(sl) < 3 {
	// 	// 	continue
	// 	// }
	// 	// s.Path = sl[0]
	// 	// s.Name = sl[1]
	// 	// s.Md5 = strings.ReplaceAll(sl[2], "\r", "")
	// 	// ss.Traits = append(ss.Traits, s)
	// }
	// ms, _ := json.Marshal(ss)
	//将str字符串的内容写到文件中，强制转换为byte，因为Write接收的是byte。
	// log.Println(string([]byte{103, 111, 118, 46, 99, 110}))
	// type pos [2]int
	// a := pos{4, 5}
	// b := pos{4, 5}
	// fmt.Println(a == b)
	// var a string = "192.168.3.97/24"
	// log.Println(ScanIpList(a))

	// telnetClientObj := new(TelnetClient)
	// telnetClientObj.IP = "192.168.3.128"
	// telnetClientObj.Port = "23"
	// telnetClientObj.IsAuthentication = true
	// telnetClientObj.UserName = "root"
	// telnetClientObj.Password = ""
	// //	fmt.Println(telnetClientObj.PortIsOpen(5))
	// log.Println(telnetClientObj.Telnet(20))
	// conn, err := net.Dial("tcp", "192.168.3.128:23")
	// if err != nil {
	// 	fmt.Sprint(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// var buf [4096]byte

	// //	for {
	// n, err := conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// fmt.Println((buf[0:n]))
	// log.Println("---------")
	// buf[1] = 252
	// buf[4] = 252
	// buf[7] = 252
	// buf[10] = 252
	// fmt.Println((buf[0:n]))
	// n, err = conn.Write(buf[0:n])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// fmt.Println((buf[0:n]))
	// log.Println("---------")

	// buf[1] = 252
	// buf[4] = 251
	// buf[7] = 252
	// buf[10] = 254
	// buf[13] = 252
	// fmt.Println((buf[0:n]))
	// n, err = conn.Write(buf[0:n])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// fmt.Println((buf[0:n]))
	// log.Println("---------")

	// buf[1] = 252
	// buf[4] = 252
	// fmt.Println((buf[0:n]))
	// n, err = conn.Write(buf[0:n])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// fmt.Println((buf[0:n]))
	// log.Println("---------")

	// /*
	// 	buf[0] = 255
	// 	buf[1] = 252
	// 	buf[2] = 1
	// 	buf[3] = 255
	// 	buf[4] = 253
	// 	buf[5] = 1
	// 	buf[6] = 255
	// 	buf[7] = 252
	// 	buf[8] = 1
	// 	buf[9] = 255
	// 	buf[10] = 253
	// 	buf[11] = 1
	// 	fmt.Println((buf[0:12]))
	// 	n, err = conn.Write(buf[0:12])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 		return
	// 	}

	// 	n, err = conn.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 		return
	// 	}
	// */

	// n, err = conn.Write([]byte("root\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// fmt.Println(string(buf[0:n]))
	// log.Println("---------")
	// n, err = conn.Write([]byte("admin\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// log.Println("---------")
	// for {
	// 	n, err = conn.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(buf[0:n]))
	// 	if strings.HasSuffix(string(buf[0:n]), "> ") {
	// 		break
	// 	}
	// }

	// n, err = conn.Write([]byte("enable\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// log.Println("---------")
	// n, err = conn.Write([]byte("terminal length 0\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// n, err = conn.Read(buf[0:])
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }
	// fmt.Println(string(buf[0:n]))
	// log.Println("---------")
	// n, err = conn.Write([]byte("root\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// for {
	// 	n, err = conn.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(buf[0:n]))
	// 	if strings.HasSuffix(string(buf[0:n]), "# ") {
	// 		break
	// 	}
	// }

	// n, err = conn.Write([]byte("admin\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// for {
	// 	n, err = conn.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(buf[0:n]))
	// 	if strings.HasSuffix(string(buf[0:n]), "# ") {
	// 		break
	// 	}
	// }

	// n, err = conn.Write([]byte("show running-config\n"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 	return
	// }

	// reader := bufio.NewReader(conn)
	// if reader == nil {
	// 	fmt.Fprintf(os.Stderr, "Create reader failed.")
	// }

	// for {
	// 	n, err := reader.Read(buf[0:])
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
	// 		return
	// 	}
	// 	fmt.Println(string(buf[0:n]))
	// 	if strings.HasSuffix(string(buf[0:n]), "# ") {
	// 		break
	// 	}
	// }
	/*
		for {
				n, err = conn.Read(buf[0:])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
					return
				}
				fmt.Println(string(buf[0:n]))
				fmt.Println(n)
				if strings.HasSuffix(string(buf[0:n]), "# ") {
					break
				}

		}
	*/

	//	}
	// type taskRed struct {
	// 	M         Taskw `json:"task"`
	// 	startTime string
	// 	status    int
	// }
	// var t = Taskw{
	// 	Name: "wuhuhu",
	// }
	// var s = taskRed{
	// 	M:         t,
	// 	startTime: "22",
	// 	status:    0,
	// }
	// b, err := json.Marshal(s)
	// log.Println(string(b), err)
	// var Ss = []string{"a", "b", "b", "a", "a", "c", "c", "c", "a"}
	// log.Println(TopKFrequent(Ss, 0))
	// str := "2022-06-15 15:54:30"
	// 设置时区
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// date, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	// if err != nil || !date.After(time.Now()) {
	// 	return
	// }
	// log.Println(time.Until(date))
	// var rs Timetab
	// rs.ids = make(map[string]*time.Timer)
	// rs.AddByID("1", time.Until(date), func() {
	// 	log.Println("----------")
	// })
	// time.Sleep(10 * time.Second)
	// rs.DelByID("1")
	// time.Sleep(1 * time.Minute)
	// config := sarama.NewConfig()
	// config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	// config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	// config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	/*-------------------------*/
	// consumer, err := sarama.NewConsumer([]string{"192.168.31.159:9092"}, nil)
	// if err != nil {
	// 	fmt.Printf("fail to start consumer, err:%v\n", err)
	// 	return
	// }
	// partitionList, err := consumer.Partitions("command") // 根据topic取到所有的分区
	// if err != nil {
	// 	fmt.Printf("fail to get list of partition:err%v\n", err)
	// 	return
	// }
	// fmt.Println(partitionList)
	// for partition := range partitionList { // 遍历所有的分区
	// 	// 针对每个分区创建一个对应的分区消费者
	// 	pc, err := consumer.ConsumePartition("command", int32(partition), sarama.OffsetNewest)
	// 	if err != nil {
	// 		fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
	// 		return
	// 	}
	// 	defer pc.AsyncClose()
	// 	// 异步从每个分区消费信息
	// 	go func(sarama.PartitionConsumer) {
	// 		for msg := range pc.Messages() {
	// 			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
	// 		}
	// 	}(pc)
	// }
	// time.Sleep(5 * time.Second)
	// // 构造一个消息
	// msg := &sarama.ProducerMessage{}
	// msg.Topic = "command"
	// msg.Value = sarama.StringEncoder("this is 3 test command")
	// // 连接kafka
	// client, err := sarama.NewSyncProducer([]string{"192.168.31.159:9092"}, config)
	// if err != nil {
	// 	fmt.Println("producer closed, err:", err)
	// 	return
	// }
	// defer client.Close()
	// // 发送消息
	// pid, offset, err := client.SendMessage(msg)
	// if err != nil {
	// 	fmt.Println("send msg failed, err:", err)
	// 	return
	// }
	// fmt.Printf("pid:%v offset:%v\n", pid, offset)
	// signal.Ignore(syscall.SIGHUP)
	// runtime.Goexit()
	// 初始化生产生

	// 监听
	go func() {
		err := consumer.InitConsumer("192.168.31.159:9092")
		if err != nil {
			panic(err)
		}
		err = consumer.LoopConsumer("command", TopicCallBack)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := consumer.InitConsumer("192.168.31.159:9092")
		if err != nil {
			panic(err)
		}
		err = consumer.LoopConsumer("command", TopicCallBack)
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(5 * time.Second)
	var addr = "192.168.31.159:9092"
	err := producer.InitProducer(addr)
	if err != nil {
		panic(err)
	}
	// 关闭
	defer producer.Close()

	// 发送测试消息
	producer.Send("command", "This is Test command")
	producer.Send("command", "Hello command")

	signal.Ignore(syscall.SIGHUP)
	runtime.Goexit()
}

func TopicCallBack(data []byte) {
	log.Println("kafka", "Test:"+string(data))
}

type Timetab struct {
	ids map[string]*time.Timer
}

func (c *Timetab) AddByID(id string, spec time.Duration, sf func()) error {
	if _, ok := c.ids[id]; ok {
		return errors.New("timetab id exists")
	}
	timers := time.NewTimer(spec)
	go func() {
		curTime := <-timers.C
		fmt.Println("task start time: ", curTime.Format("2006-01-02 15:04:05"))
		sf()
	}()
	c.ids[id] = timers
	return nil
}

func (c *Timetab) DelByID(id string) {
	eid, ok := c.ids[id]
	if !ok {
		return
	}
	eid.Stop()
	delete(c.ids, id)
}

//TopKFrequent 排序数组
func TopKFrequent(nums []string, k int) (res []map[string]int) {
	ans := []string{}
	map_num := map[string]int{}
	for _, item := range nums {
		map_num[item]++
	}
	for key, _ := range map_num {
		ans = append(ans, key)
	}
	//核心思想：排序
	//可以不用包函数，自己实现快排
	sort.Slice(ans, func(a, b int) bool {
		return map_num[ans[a]] > map_num[ans[b]]
	})
	for _, val := range ans {
		res = append(res, map[string]int{val: map_num[val]})
	}
	return
}

type Taskw struct {
	Name string
}

func (this *TelnetClient) PortIsOpen(timeout int) bool {
	raddr := this.IP + ":" + this.Port
	conn, err := net.DialTimeout("tcp", raddr, time.Duration(timeout)*time.Second)
	if nil != err {
		log.Println("pkg: model, func: PortIsOpen, method: net.DialTimeout, errInfo:", err)
		return false
	}
	defer conn.Close()
	return true
}

type TelnetClient struct {
	IP               string
	Port             string
	IsAuthentication bool
	UserName         string
	Password         string
}

const (
	//经过测试，linux下，延时需要大于100ms
	TIME_DELAY_AFTER_WRITE = 500 //500ms
)

func (t *TelnetClient) Telnet(timeout int) bool {
	var buf [4096]byte
	raddr := t.IP + ":" + t.Port
	conn, err := net.DialTimeout("tcp", raddr, time.Duration(timeout)*time.Second)
	if nil != err || conn == nil {
		return false
	}
	n, err := conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	buf[1] = 252
	buf[4] = 251
	buf[7] = 252
	buf[10] = 254
	buf[13] = 252
	fmt.Println(string((buf[0:n])))
	defer conn.Close()
	return true
}

func (this *TelnetClient) telnetProtocolHandshake(conn net.Conn) bool {
	var buf [4096]byte
	n, err := conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	fmt.Println((buf[0:n]))

	buf[1] = 252
	buf[4] = 252
	buf[7] = 252
	buf[10] = 252
	fmt.Println((buf[0:n]))
	n, err = conn.Write(buf[0:n])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Write, errInfo:", err)
		return false
	}

	n, err = conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	fmt.Println((buf[0:n]))

	buf[1] = 252
	buf[4] = 251
	buf[7] = 252
	buf[10] = 254
	buf[13] = 252
	fmt.Println((buf[0:n]))
	n, err = conn.Write(buf[0:n])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Write, errInfo:", err)
		return false
	}

	n, err = conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	fmt.Println((buf[0:n]))

	buf[1] = 252
	buf[4] = 252
	fmt.Println((buf[0:n]))
	n, err = conn.Write(buf[0:n])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Write, errInfo:", err)
		return false
	}

	n, err = conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	fmt.Println((buf[0:n]))

	if false == this.IsAuthentication {
		return true
	}

	n, err = conn.Write([]byte(this.UserName + "\n"))
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Write, errInfo:", err)
		return false
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))

	n, err = conn.Write([]byte(this.Password + "\n"))
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Write, errInfo:", err)
		return false
	}
	time.Sleep(time.Millisecond * TIME_DELAY_AFTER_WRITE)

	n, err = conn.Read(buf[0:])
	if nil != err {
		log.Println("pkg: model, func: telnetProtocolHandshake, method: conn.Read, errInfo:", err)
		return false
	}
	fmt.Println(string(buf[0:n]))
	return true
}
func ScanIpList(str string) (result []string) {
	if strings.Contains(str, ",") {
		for _, v := range strings.Split(str, ",") {
			if ip := net.ParseIP(v); ip == nil {
				continue
			}
			result = append(result, v)
		}
	}
	if strings.Contains(str, "/") {
		_, _, err := net.ParseCIDR(str)
		if err != nil {
			return nil
		}
		result = append(result, "++")
	}
	if ip := net.ParseIP(str); ip != nil {
		result = append(result, str)
		return
	}
	return
}

type Soul struct {
	Type   string
	Traits []Trait
}

type Trait struct {
	Name string
	Path string
	Md5  string
}

var (
	g errgroup.Group
)

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/pokemon", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusOK,
				"GO":   "Welcome server 2010",
			},
		)
	})
	return e
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/pokemon", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusOK,
				"GO":   "Welcome server 2011",
			},
		)
	})

	return e
}

func router03() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/pokemon", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code": http.StatusOK,
				"GO":   "Welcome server 2012",
			},
		)
	})
	return e
}

func GetEvilIp(ip string) EvilIp {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:8000/api/ip/%s", ip), nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest) //提交
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var result EvilIp
	json.Unmarshal(body, &result)
	return result
}

type EvilIp struct {
	Evil bool `json:"evil"`
	Data struct {
		ID   int    `json:"Id"`
		IP   string `json:"Ip"`
		Info []struct {
			Desc   string `json:"Desc"`
			Source string `json:"Source"`
		} `json:"Info"`
	} `json:"data"`
}

// 通过传递模块名获得网关上存储的服务信息
func GetService(model string) {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	req := GetServerRes{model}
	var res GetServerponse
	err = conn.Call("Register.GetServer", req, &res)
	if err != nil {
		log.Fatalln("Register error: ", err)
	}
	fmt.Printf("%v\n", res.RemainingTimeList)
}

// 获取服务信息结构体
type GetServerRes struct {
	Model string
}

// 获取服务信息返回结构体
type GetServerponse struct {
	RemainingTimeList map[string]time.Time
}

// 将服务注册到网关上
func RegisterService(str string) {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	req := Register{"192.168.3.220", "1459", "Order", str}
	var res RegisterResponse
	err = conn.Call("Register.Server", req, &res)
	if err != nil {
		log.Fatalln("Register error: ", err)
	}
	fmt.Printf("%v\n", res.RemainingTime)
}

//读取configIni文件内容
func ReadConfigIni(key, value string) (result string) {
	filePath := "config/config.ini"
	prefix := goini.SetConfig(filePath)
	result = prefix.GetValue(key, value)
	if result == "no value" {
		log.Println(key, "In config.ini not exist")
		return ""
	}
	return
}

// 监听结构体
type Register struct {
	Addr  string
	Port  string
	Model string
	Str   string
}

// 监听结构体
type RegisterResponse struct {
	RemainingTime float64
}

func CacheStatus(cache *cache.Cache) (count int, items map[string]cache.Item) {
	count = cache.ItemCount()
	items = cache.Items()
	return count, items
}

var CACHE_IPS *cache.Cache

func SaveToFile(ctx *cli.Context) (err error) {
	CACHE_IPS.SaveFile("ips")
	return err
}

func MakeSign(t string, key string) (sign string) {
	sign = MD5(fmt.Sprintf("%s%s", t, key))
	return sign
}

func tcpGather(ip string, ports []string) map[string]string {
	// 检查 emqx 1883, 8083, 8080, 18083 端口

	results := make(map[string]string)
	wg.Add(len(ports))
	for _, port := range ports {
		go func(ports string) {

			address := net.JoinHostPort(ip, ports)
			// 3 秒超时
			conn, err := net.DialTimeout("tcp", address, 3*time.Second)
			if err != nil {
				results[ports] = "failed"
				// todo log handler
			} else {
				if conn != nil {
					results[ports] = "success"
					_ = conn.Close()
				} else {
					results[ports] = "failed"
				}
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	return results
}

var IsLikereq = make(map[string]string)

func Islike(c, q string) {
	var isLike = []string{"学好", "学习", "路线", "初入", "自学", "笔试", "程序员", "常用,", "小白", "编程", "面试", "笔记", "基础"}
	for _, val := range isLike {
		if strings.Contains(c, val) {
			IsLikereq[c] = q
			break
		}
	}
}

func GetPinnedArticlePidAndAuthor(columnName string) (*PinnedArticleAndAuthor, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}
	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/pinned-article", columnName)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	pinnedArticleAndAuthor := PinnedArticleAndAuthor{}
	err = res.ToJSON(&pinnedArticleAndAuthor)
	if err != nil {
		return nil, err
	}

	return &pinnedArticleAndAuthor, nil
}

// 获取单个文章
func GetSingleArticle(pid int) (*Article, error) {
	if pid == 0 {
		return nil, PidCanNotBeEmpty
	}
	u := fmt.Sprintf("https://api.zhihu.com/articles/%d", pid)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	article := Article{}
	err = res.ToJSON(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

var (
	ColumnNameCanNotBeEmpty = errors.New("专栏名不能为空")
	PidCanNotBeEmpty        = errors.New("pid 不能为空")
)

func sendNewZhihuRequest(u string) (*req.Resp, error) {
	return req.Get(u, req.Header{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36",
		"Referer":    "https://zhuanlan.zhihu.com/",
	}, nil)
}

// 获取文章pid列表
func GetArticlesListPids(columnName string) ([]int, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}

	var limit = 20
	var offset = 0

	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	articleList := ArticleList{}
	err = res.ToJSON(&articleList)
	if err != nil {
		return nil, err
	}

	var articleIds = []int{}

	for _, entry := range articleList.Data {
		articleIds = append(articleIds, entry.ID)
	}

	for offset = offset + limit; offset < articleList.Paging.Totals; offset = offset + limit {
		u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
		res, err := sendNewZhihuRequest(u)
		if err != nil {
			return nil, err
		}

		articleList := ArticleList{}
		err = res.ToJSON(&articleList)
		if err != nil {
			return nil, err
		}
		for _, entry := range articleList.Data {
			articleIds = append(articleIds, entry.ID)
		}
	}

	return articleIds, nil
}

type Zhuanlan struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// 作者
type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      int    `json:"gender"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	AvatarUrl   string `json:"avatar_url"`
	Type        string `json:"type"`
	UID         string `json:"uid"`
	URL         string `json:"url"`
	URLToken    string `json:"url_token"`
	UserType    string `json:"user_type"`
}

// 置顶文章及作者信息
type PinnedArticleAndAuthor struct {
	Type     string `json:"type"`
	ID       int    `json:"id"`
	Updated  int64  `json:"updated"`
	Created  int64  `json:"created"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	URL      string `json:"url"`
	Excerpt  string `json:"excerpt"`
	Author   Author
}

// 主题
type Topic struct {
	Url  string `json:"url"`
	Type string `json:"type"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

// 文章
type Article struct {
	ID       int     `json:"id"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	Updated  int64   `json:"updated"`
	Created  int64   `json:"created"`
	Excerpt  string  `json:"excerpt"`
	Content  string  `json:"content"`
	ImageURL string  `json:"image_url"`
	Topics   []Topic `json:"topics"`
}

// 文章列表
type ArticleList struct {
	Paging struct {
		IsEnd   bool `json:"is_end"`
		Totals  int  `json:"totals"`
		IsStart bool `json:"is_start"`
	} `json:"paging"`
	Data []struct {
		ID int `json:"id"`
	} `json:"data"`
}

func testTopic(topic string, ctx context.Context) {
	var Consumer, err = sarama.NewConsumer([]string{"192.168.3.128:9093"}, nil)
	if err != nil {
		log.Println(err, "---------------")
	}
	log.Println("i am here")
	partitionList, err := Consumer.Partitions(topic)
	fmt.Println(partitionList)
	if err != nil {
		fmt.Printf("fail to start consumer partition,err:%v\n", err)
		return
	}
	for partition := range partitionList {
		//  遍历所有的分区，并且针对每一个分区建立对应的消费者
		pc, err := Consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go testGetMsg(pc, ctx)
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}
func testGetMsg(partitionConsumer sarama.PartitionConsumer, ctx context.Context) {
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Partition:%d Offset:%v Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		select {
		case <-ctx.Done():
			return
		default:
			continue
		}
	}
}

// var (
// 	producer sarama.SyncProducer
// 	brokers  = []string{"192.168.3.128:9093"}
// 	topic    = "mysql"
// )

// func init() {
// 	config := sarama.NewConfig()
// 	config.Producer.RequiredAcks = sarama.WaitForLocal
// 	config.Producer.Retry.Max = 5
// 	config.Producer.Return.Successes = true
// 	brokers := brokers
// 	var err error
// 	producer, err = sarama.NewSyncProducer(brokers, config)
// 	if err != nil {
// 		fmt.Printf("init producer failed -> %v \n", err)
// 		panic(err)
// 	}
// 	fmt.Println("producer init success")
// }

// func produceMsg(msg string) {
// 	msgX := &sarama.ProducerMessage{
// 		Topic: topic,
// 		Value: sarama.StringEncoder(msg),
// 	}
// 	fmt.Printf("SendMsg -> %v\n", dumpString(msgX))

// 	partition, offset, err := producer.SendMessage(msgX)
// 	if err != nil {
// 		fmt.Printf("send msg error:%s \n", err)
// 	}
// 	fmt.Printf("msg send success, message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
// }

// func startProduce() {
// 	time.Sleep(5 * time.Second)
// 	tick := time.Tick(2 * time.Second)
// 	for {
// 		select {
// 		case <-tick:
// 			t := time.Now().Unix() * 1000
// 			msg := fmt.Sprintf("{\"timestamp\":%d}", t)
// 			produceMsg(msg)
// 		}
// 	}
// }

//解析为json字符串
func dumpString(v interface{}) (str string) {

	bs, err := json.Marshal(v)
	b := bytes.Buffer{}
	if err != nil {
		b.WriteString("{err:\"json format error.")
		b.WriteString(err.Error())
		b.WriteString("\"}")
	} else {
		b.Write(bs)
	}
	str = b.String()
	return str
}

type JsonText struct {
	Seed  string
	Value string
}

func ExampleNewGCMEncrypter(plaintext []byte) ([]byte, []byte) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return nonce, ciphertext
}

func ExampleNewGCMDecrypter(seed, value string) ([]byte, error) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("AES256Key-32Characters1234567890")
	ciphertext, _ := hex.DecodeString(value)
	nonce, _ := hex.DecodeString(seed)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	if len(nonce) != 12 {
		return nil, errors.New("Json文件非法")
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return plaintext, err
}

// JSONPlugin JSON插件
type JSONPlugin struct {
	Target  string `json:"target" hashids:"true"`
	Meta    Plugin `json:"meta"`
	Request struct {
		Path     string `json:"path"`
		PostData string `json:"postdata"`
	} `json:"request"`
	Verify struct {
		Type  string `json:"type"`
		Match string `json:"match"`
	} `json:"verify"`
	Extra bool
}

//References 插件附加信息
type References struct {
	URL  string `json:"url"`
	CVE  string `json:"cve"`
	KPID string `json:"kpid"`
}

// Plugin 漏洞插件信息
type Plugin struct {
	Name       string     `json:"name"`
	Remarks    string     `json:"remarks"`
	Level      int        `json:"level" hashids:"true"`
	Type       string     `json:"type"`
	Author     string     `json:"author"`
	References References `json:"references"`
	Request    string
	Response   string
}

func CheckLocalSecurePolicies() ([]string, error) {

	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	path1 := fmt.Sprintf("%s\\setup.ini", path)
	path2 := fmt.Sprintf("%s\\dengBaoLV2WinCheckValue.ini", path)

	c := exec.Command("secedit", "/export", "/cfg", path1)
	err := c.Run()
	if err != nil {
		return nil, err
	}
	setup, err := ini.Load(path1)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("!!!!!!read secedit setup failed: %s", err))
	}
	checkValue, err := ini.Load(path2)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("!!!!!!read secedit setup failed: %s", err))
	}

	var failPoints []string
	v1, _ := strconv.Atoi(setup.Section("System Access").Key("MinimumPasswordAge").Value())
	v2, _ := strconv.Atoi(checkValue.Section("System Access").Key("MinimumPasswordAge").Value())
	fmt.Println(v1, v2)
	if v1 > v2 {
		fmt.Println(v1, ">", v2)
		failPoints = append(failPoints, "MinimumPasswordAge")
	}
	return failPoints, nil
}

var (
	accessKey = "a97c367ee15228f6cd11d9e936813a3e89345572f6e0d40e99d42ec8a62d2414"
	secretKey = "d2ac61500f2fc8e49ec2f098ec44c2ed89e9207033c790dc98a9054a2eb55c15"
	ip        = "192.168.3.9"
	port      = "8834"
	username  = "admin"
	password  = "admin"
)

func getToken() string {
	nessusUrl := fmt.Sprintf("https://%s:%s/session", ip, port)
	fmt.Println(nessusUrl)
	urlValues := url.Values{}
	urlValues.Add("username", "admin")
	urlValues.Add("password", "admin")
	fmt.Println("urlValues:", urlValues)
	body := strings.NewReader(urlValues.Encode())
	//body := ioutil.NopCloser(strings.NewReader(urlValues.Encode()))
	fmt.Printf("%+v\n", body)
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://192.168.3.9:8834/session", body)
	if err != nil {
		fmt.Println("err:", err)
	}
	req.Header.Set("Content-Type", "application/json; param=value")

	fmt.Printf("req: %+v\n", req)
	resp, err := client.Do(req)
	fmt.Println("resp:", resp)
	//defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	//body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read err:", err)
		return ""
	} else {
		return string(data)

	}
	//if response.StatusCode == 200{
	// body,_ := ioutil.ReadAll(response.Body)
	// return body
	//}else {
	// return nil
	//}
}

type Student struct {
	Id   int
	Name string
}

func (s Student) Hello() {
	fmt.Println("我是一个学生")
}

// func getNil(req ...interface{}) bool {
// 	res
// 	return true
// }

func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}
func Test(a, b string) (string, error) {
	return a + " " + b, nil
}

func AddTask(req map[string]string) {
	// req = append(req, req)
}

func task1(req map[string]string) {
	log.Println(req)
}

func WriteConfig(cfg string, jsonByte []byte) { //这里的cfg就是我要写到的目标文件 ./host.json
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	_, err := WriteBytes(cfg, jsonByte)
	if err != nil {
		log.Fatalln("write config file:", cfg, "fail:", err)
	}

	log.Println("write config file:", cfg, "successfully")

}

func WriteBytes(filePath string, b []byte) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	fw, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fw.Close()
	return fw.Write(b)
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type X struct {
	S string
}

type Y struct {
	N int
}

func Process(i ...interface{}) {
	for _, v := range i {
		s := reflect.ValueOf(v)
		if s.Kind() != reflect.Slice {
			continue
		}
		fmt.Println(s.Len(), s)
	}
}

var wg sync.WaitGroup
var countdone int

func a(ch chan string, wg *sync.WaitGroup) {
	select {
	case aget := <-ch:
		log.Println(aget)
	case <-time.After(2 * time.Second):
		countdone++
		<-ch
		log.Println("a.time out", countdone)
		wg.Done()
		return
	}
	countdone++
	log.Println("a.done", countdone)
	wg.Done()
}

func b(ch chan string, wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	ch <- "a.string"
	countdone++
	log.Println("b.done", countdone)
	wg.Done()
}

func test_chan1(ch chan struct{}, ctx context.Context) {
	for {
		select {
		case <-ch:
			fmt.Println("ch1")
			mLock.Done()
		case <-ctx.Done():
			fmt.Println("timeout...")
		}
	}
}

type Sing struct {
}

var insertSing *Sing
var mu sync.Mutex
var mLock sync.WaitGroup

func getSing() *Sing {
	if insertSing == nil {
		mu.Lock()
		defer mu.Unlock()
		insertSing = &Sing{}
	}
	return insertSing
}

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	temp := len(nums) / 2
	sort.Ints(nums)
	for l <= r {
		if target == nums[temp] {
			return target
		}
		if target > nums[temp] {
			l = temp + 1
		} else {
			r = temp - 1
		}
		temp = (l + r) / 2
	}
	return -1
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
}

func test1(s string) string {
	str := []rune(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str)
}

var ac = make(chan struct{}, 0)
var cc = make(chan struct{}, 0)

func dog() {
	var i int
	for {
		select {
		case <-cc:
			i++
			if i > 3 {
				wg.Done()
				return
			}
			fmt.Println("dog")
			ac <- struct{}{}
		}
	}
}

func cat() {
	for {
		select {
		case <-ac:
			fmt.Println("cat")
			cc <- struct{}{}
		}
	}
}

func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (q + p) % mod
	}
	return r
}

func smallestK(arr []int, k int) (ans []int) {
	sort.Ints(arr)
	return arr[0:k]
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func compareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for i < n || j < m {
		x := 0
		for ; i < n && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++ // 跳过点号
		y := 0
		for ; j < m && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++ // 跳过点号
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

func isPowerOfThree(n int) bool {
	return isPowerOfThrees(n)
}

func isPowerOfThrees(n int) bool {
	if n == 3 {
		return true
	}
	if n > 0 && n%3 == 0 {
		return isPowerOfThrees(n / 3)
	}
	return false
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for k, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, k}
		}
		hashTable[v] = k
	}
	return nil
}

func lengthOfLIS(nums []int) (res int) {
	l, r := 0, len(nums)-1
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		if mid <= ans {
			ans = mid
		}
		r = mid - 1
		for _, v := range nums[:r+1] {
			if v <= ans {
				ans = v
			}
		}

	}
	for _, v := range nums {
		if v <= ans {
			res++
		}
	}
	return
}

func lengthOfLastWord(s string) int {
	ans := 0
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	quotient, remainder := n/k, n%k
	parts := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := quotient
		if i < remainder {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}

func canWinNim(n int) bool {
	if n == 1 {
		return true
	}
	return n%4 == 4
}

type timeSlot struct {
	timestamp time.Time // 这个timeSlot的时间起点
	count     int       // 落在这个timeSlot内的请求数
}

// 统计整个时间窗口中已经发生的请求次数
func countReq(win []*timeSlot) int {
	var count int
	for _, ts := range win {
		count += ts.count
	}
	return count
}

type SlidingWindowLimiter struct {
	mu           sync.Mutex    // 互斥锁保护其他字段
	SlotDuration time.Duration // time slot的长度
	WinDuration  time.Duration // sliding window的长度
	numSlots     int           // window内最多有多少个slot
	windows      []*timeSlot
	maxReq       int // 大窗口时间内允许的最大请求数
}

func NewSliding(slotDuration time.Duration, winDuration time.Duration, maxReq int) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		SlotDuration: slotDuration,
		WinDuration:  winDuration,
		numSlots:     int(winDuration / slotDuration),
		maxReq:       maxReq,
	}
}

func (l *SlidingWindowLimiter) validate() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	// 已经过期的time slot移出时间窗
	timeoutOffset := -1
	for i, ts := range l.windows {
		if ts.timestamp.Add(l.WinDuration).After(now) {
			break
		}
		timeoutOffset = i
	}
	if timeoutOffset > -1 {
		l.windows = l.windows[timeoutOffset+1:]
	}

	// 判断请求是否超限
	var result bool
	if countReq(l.windows) < l.maxReq {
		result = true
	}

	// 记录这次的请求数
	var lastSlot *timeSlot
	if len(l.windows) > 0 {
		lastSlot = l.windows[len(l.windows)-1]
		if lastSlot.timestamp.Add(l.SlotDuration).Before(now) {
			// 如果当前时间已经超过这个时间插槽的跨度，那么新建一个时间插槽
			lastSlot = &timeSlot{timestamp: now, count: 1}
			l.windows = append(l.windows, lastSlot)
		} else {
			lastSlot.count++
		}
	} else {
		lastSlot = &timeSlot{timestamp: now, count: 1}
		l.windows = append(l.windows, lastSlot)
	}

	return result
}

func (l *SlidingWindowLimiter) LimitTest() string {
	if l.validate() {
		return "Accepted"
	} else {
		return "Ignored"
	}
}

func largestNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i, _ := range nums {
		strs[i] = fmt.Sprintf("%d", nums[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if strs[0][0] == '0' {
		return "0"
	}

	return strings.Join(strs, "")
}

func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}

func canJump(nums []int) bool {
	n := len(nums) - 1
	if n == 0 {
		return true
	}
	maxPosition := 0
	for i := 0; i < n; i++ {
		if i <= maxPosition {
			maxPosition = max(maxPosition, i+nums[i])
		}
		if maxPosition >= n {
			return true
		}
	}
	return false
}

func jump(nums []int) int {

	position := len(nums) - 1
	steps := 0
	r := 0
	maxPosition := 0
	for i := 0; i < position; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == r {
			r = maxPosition
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func write(c chan int, n int) {
	c <- n
}

func go_test() {
	for {
		log.Println("time to sleep")
		time.Sleep(time.Second * 5)
	}
}

func searchs(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}
	sort.Ints(nums)
	res := sort.SearchInts(nums, target)
	if res == n {
		res--
	}
	return target == nums[res]
}

func searchMatrix(matrix [][]int, target int) bool {
	var nums = make([]int, 0)
	for k, _ := range matrix {
		nums = append(nums, matrix[k]...)
	}

	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func mySqrt(x int) int {
	l, r := 0, x
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func searchRanges(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	start := 0
	end := 0
	targetMid := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			targetMid = mid
			break
		}
		if nums[mid] < target && target <= nums[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if targetMid == -1 {
		return []int{targetMid, targetMid}
	}
	for i := targetMid; i >= 0; i-- {
		if nums[i] == target {
			start = i
		} else {
			break
		}
	}
	for i := targetMid; i < len(nums); i++ {
		if nums[i] == target {
			end = i
		} else {
			break
		}
	}
	return []int{start, end}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l, r := 0, len(nums1)-1
	if len(nums1)%2 != 0 {
		return float64(nums1[(l+r)/2])
	}
	return float64((nums1[(l+r)/2] + nums1[((l+r)/2)+1])) / 2
}

func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if i > 0 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				if sum := nums[i] + nums[j] + nums[l] + nums[r]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[l], nums[r]})
					for l++; l < r && nums[l] == nums[l-1]; l++ {
					}
					for r--; l < r && nums[r] == nums[r+1]; r-- {
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}

	}
	return
}

func generateParenthesis(n int) []string {
	result := make([][]string, n+1)
	result[0] = []string{""}
	for i := 1; i <= n; i++ {
		result[i] = make([]string, 0)
		for j := 0; j < i; j++ {
			for _, intside := range result[j] {
				for _, outside := range result[i-j-1] {
					result[i] = append(result[i], "("+intside+")"+outside)
				}
			}
		}
	}
	return result[n]
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	n, SumClosest := len(nums), nums[0]+nums[1]+nums[2] // 初始化为前三元素的值，避免任何个

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum-target) < abs(SumClosest-target) {
				SumClosest = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return SumClosest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if nums[i] > 0 {
			break
		}
		L, R := i+1, len(nums)-1
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for L < R {
			n2, n3 := nums[L], nums[R]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for L < R && nums[L] == n2 {
					L++
				}
				for L < R && nums[R] == n3 {
					R--
				}
			} else if n1+n2+n3 < 0 {
				L++
			} else {
				R--
			}
		}
	}
	return res
}

func BubbleAsort(values []int) []int {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var roman []byte
	for _, v := range valueSymbols {
		for num >= v.value {
			num -= v.value
			roman = append(roman, v.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var m = 0
	for i < j {
		if temp := (j - i) * min(height[i], height[j]); temp > m {
			m = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return m
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func removeDuplicateLetters(s string) string {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	st := []byte{}
	me := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		m[s[i]]--
		if _, ok := me[s[i]]; ok {
			continue
		}
		for len(st) > 0 && st[len(st)-1] > s[i] {
			if m[st[len(st)-1]] == 0 {
				break
			}
			delete(me, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, s[i])
		me[s[i]] = struct{}{}
	}
	return string(st)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

type Uer struct {
	Name     string
	Password string
}

func GetToken() string {
	//"name":"write@yondor.com","password":md5pwd}
	headerMap := make(map[string]string)
	var user = Uer{}
	user.Name = "write@yondor.com"
	user.Password = "9cbf8a4dcb8e30682b927f352d6559a0"
	headerMap["Content-Type"] = "application/json"
	httpUser, _ := json.Marshal(user)
	result, _ := HttpHandleToData("POST", "http://192.168.6.30:30671/crm/dashboard/login", httpUser)
	return result
}

var (
	httpClient *http.Client
	mutex      sync.Mutex
)

func GetHttpClient() *http.Client {
	if httpClient == nil {
		mutex.Lock()
		if httpClient == nil {
			tr := &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConnsPerHost:   1000, //连接池对所有host的最大链接数量
				MaxIdleConns:          1000, //连接池对每个host的最大链接数量
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			}
			httpClient = &http.Client{Transport: tr}
		}
		mutex.Unlock()
	}
	return httpClient
}

//http请求
func HttpHandleToData(method, urlValstring string, data []byte) (string, error) {
	client := GetHttpClient()
	var req *http.Request
	req, _ = http.NewRequest(method, urlValstring, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "nil", err
	}
	defer resp.Body.Close()
	return resp.Header.Get("X-Auth-Token"), nil
}

func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

//HasIndex 切片是否存在索引
func HasIndex(index int, data []interface{}, dataLength int) bool {
	if dataLength == 0 {
		dataLength = len(data)
	}
	return dataLength > index
}

func GetUUID() string {
	ID, err := uuid.NewV4()
	log.Println(err)
	UUID := strings.Replace(ID.String(), "-", "", -1)
	return UUID
}

func init() {

	// // date := time.Now().Format("20060102")
	// // oldDate := date

	// spec1 := "0 59 09 * * *"
	// spec2 := "0 00 10 * * *"

	// c := cron.New()
	// c.AddFunc(spec1, func() {
	// 	log.Println("每天执行一次定时任务1")
	// })
	// c.AddFunc(spec2, func() {
	// 	log.Println("每天执行一次定时任务2")
	// })
	// c.Start()
	// defer c.Stop()
	// select {}

	// for {
	// 	time.Sleep(time.Second * 5)
	// 	log.Println("--", time.Now())

	// }
	// var printFoo = PrintJob{"每天执行一次定时任务1"}
	// c := gron.New()
	// c.Add(gron.Every(1*xtime.Day).At("09:31"), printFoo)
	// log.Println("--")

	// c.Start()
	// defer c.Stop()

}

type PrintJob struct{ Msg string }

func (p PrintJob) Run() {
	fmt.Println(p.Msg)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func hello() {
	nulldata, _ := time.Parse("20060102150405", "00010101000000")

	dateIncome := time.Now().Format("2006-01-02")
	log.Println("定时执行任务", nulldata.Format("20060102030405"), "--", nulldata, "--", dateIncome+"%")
}

var gogogo = strings.Repeat("Go", 1024)

func f() {
	for range []byte(gogogo) {
	}
}

// func g() {
// 	bs := []byte(gogogo)
// 	for range bs {
// 	}
// }

func ft() bool { return false }

func isPassIP(ipnr string) bool {
	// bits := strings.Split(ipnr.String(), ".")
	isPassIPMap := make(map[string]bool)
	ipSlice := []string{"123.207.254.129", "111.230.141.226", "139.199.57.197", "139.199.59.73", "123.207.109.51", "111.231.244.234", "123.207.93.104", "118.89.59.88"}
	for _, val := range ipSlice {
		isPassIPMap[val] = true
	}
	a := []string{"A", "B", "C"}
	// b := make([]string,len(a))
	b := append(a[:0:0], a...)
	// copy(b, a)
	var c strings.Builder //效率高的字符拼接
	for _, v := range a {
		//将v写入c
		fmt.Fprint(&c, v)
	}

	var sql strings.Builder //+>strings.join>bytes.buffer
	fmt.Fprint(&sql, `select t1.u_user_id i_d, t1.realname realname,t2.mobile mobile, t2.createdate createdate, t1.grade, t1.name username, t2.is_initpwd is_init_pwd, t1.nickname, t3.head_url avatar_u_r_l, t1.es_school_id school_i_d, t1.es_school_name school_name,
	t1.es_class_id class_i_d, t1.es_class_name class_name, t1.student_no,t1.classnum
	from base.base_user t1 inner join base.base_userauth t2 on t1.u_user_id=t2.u_user_id inner join base.base_userext t3 on t1.u_user_id=t3.u_user_id where 1=1 `)
	fmt.Fprint(&sql, " and t1.u_user_id in ("+c.String()+")")
	fmt.Fprint(&sql, " and t1.realname='"+c.String()+"'")
	fmt.Fprint(&sql, " and t1.realname is not null ")
	log.Println(b, "--", c.String())
	sql1 := []string{"select", "and t1.realname='" + c.String() + "'"}
	// log.Println("FinallSQL:", sql.String())
	strings.Join(sql1, " ")
	log.Println("sql1", sql1)
	// if _, ok := isPassIPMap[ipnr]; ok {
	// 	return true
	// }
	return isPassIPMap[ipnr]
}

// Merge receives two structs, and merges them excluding fields with tag name: `structs`, value "-"
func Merge(dst, src interface{}) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dst)
	if s.Kind() != reflect.Ptr || d.Kind() != reflect.Ptr {
		return
	}
	for i := 0; i < s.Elem().NumField(); i++ {
		v := s.Elem().Field(i)
		fieldName := s.Elem().Type().Field(i).Name
		skip := s.Elem().Type().Field(i).Tag.Get("structs")
		if skip == "-" {
			continue
		}
		if v.Kind() > reflect.Float64 &&
			v.Kind() != reflect.String &&
			v.Kind() != reflect.Struct &&
			v.Kind() != reflect.Ptr &&
			v.Kind() != reflect.Slice {
			continue
		}
		if v.Kind() == reflect.Ptr {
			// Field is pointer check if it's nil or set
			if !v.IsNil() {
				// Field is set assign it to dest

				if d.Elem().FieldByName(fieldName).Kind() == reflect.Ptr {
					d.Elem().FieldByName(fieldName).Set(v)
					continue
				}
				f := d.Elem().FieldByName(fieldName)
				if f.IsValid() {
					f.Set(v.Elem())
				}
			}
			continue
		}
		d.Elem().FieldByName(fieldName).Set(v)
	}
}
