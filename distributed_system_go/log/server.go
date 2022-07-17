package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

//声明log为 Logger指针类型的变量
var log *stlog.Logger

//声明 filelog 为 string 的别名
type filelog string

//Write方法属于之前定义的结构体，参数有一个数组切片，返回值有int和error
func (fl filelog) Write(data []byte) (int, error) {
	//fl这里指的是文件地址，三个os的标签为创建，只读，添加，编码方式为0600
	f, err := os.Open(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	//结尾关闭操作
	defer f.Close()
	//返回
	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(filelog(destination), "Go: ", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

//made write function to write
func write(message string) {
	log.Printf("%v\n", message)
}
