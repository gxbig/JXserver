package util

import (
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type Results struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetResults(data interface{}, code string, message string) ([]byte, error) {
	result := &Results{
		Message: message,
		Code:    code,
		Data:    data,
	}
	res, err := json.Marshal(result)
	log.Debug(string(res))
	return res, err
}
func GetSuccess(data interface{}) []byte {
	result := &Results{
		Message: "请求成功！",
		Code:    "200",
		Data:    data,
	}
	res, _ := json.Marshal(result)
	log.Debug(string(res))
	return res
}
func GetError(data interface{}) []byte {
	result := &Results{
		Message: "系统错误",
		Code:    "999",
		Data:    data,
	}
	res, _ := json.Marshal(result)
	log.Error(string(res))
	return res
}

//请求之前

type HandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "*") //header的类型
	w.Header().Set("content-type", "application/json")
	log.Debug(fmt.Sprintf("【%s】%s %s %s", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, r.RequestURI))
	f(w, r)
}

// 设置cookie
func SetCookie(w http.ResponseWriter, sessionId string) {
	//设置cookie
	cookies := &http.Cookie{
		Name:     "sessionId",
		Value:    sessionId,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookies)
}

// Unpack 从 HTTP 请求 req 的参数中提取数据填充到 ptr 指向结构体的各个字段
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 创建字段映射表，键为有效名称
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("json")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// 对请求中的每个参数更新结构体中对应的字段
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 忽略不能识别的 HTTP 参数
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

// 生成验证码
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskSrcUnsafe(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
