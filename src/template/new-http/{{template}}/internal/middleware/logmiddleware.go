package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogMiddleware struct {
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		startTime := time.Now()

		// 读取请求主体
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("Failed to read request body: %v", err)
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// 打印请求参数日志
		logx.WithContext(r.Context()).Infof(`[HTTP Request]: %s %s [body: %s]`, r.Method, r.RequestURI, body)

		// 创建一个自定义的 ResponseWriter，用于记录响应
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}

		// 调用下一个处理器，捕获响应
		next(recorder, r)

		// 打印响应日志
		responseBoy := string(recorder.body)

		logx.WithContext(r.Context()).Infof(`[HTTP Response]: %s %s [http status code: %d] [rt: %v] [body: %s]`, r.Method, r.RequestURI, recorder.statusCode, time.Since(startTime), responseBoy)
	}
}

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}
