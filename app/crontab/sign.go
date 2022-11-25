package crontab

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"uc/utils/mylog"

	"github.com/robfig/cron/v3"
)

func CronTab() {
	//sign()

	c := cron.New(cron.WithSeconds())
	//c.AddFunc("0 6 * * ?", sign)
	_, err := c.AddFunc("0 0 6 * * *", sign)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println("crontab is running...")
	}
	c.Start()
}

func sign() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://renzhe.cloud/user/checkin", &strings.Reader{})
	if err != nil {
		panic(err)
	}
	req.Header.Set(
		"cookie",
		"_ga=GA1.1.604092515.1665367854; crisp-client%2Fsession%2F50829b14-5009-4006-b9ec-984351fa9529%2F91d2d0f3-fbd3-3b1c-bfdf-80bb2595613a=session_82f43fd0-f997-4fd3-8628-03c15794a194; crisp-client%2Fsession%2F50829b14-5009-4006-b9ec-984351fa9529%2Feabd61aa-ad69-39a9-b3d9-24bbad788cd5=session_80cc7eb0-23a0-41a0-8107-465a0caecb18; crisp-client%2Fsession%2F50829b14-5009-4006-b9ec-984351fa9529%2F43250a77-60d6-39a2-a3a2-7ef1c992e3cd=session_f6b543c1-b259-4b05-8bb1-4ce9a569aecb; _ga_86TCXRZQN5=GS1.1.1665655914.3.1.1665656821.0.0.0; _ga4=eb46c191-cb5c-4f74-b0b9-56288a0ee584; crisp-client%2Fsession%2F50829b14-5009-4006-b9ec-984351fa9529=session_864d19e7-0748-4f5f-8cbc-079c05a8584d; crisp-client%2Fsession%2F50829b14-5009-4006-b9ec-984351fa9529%2Ff188c164-15f3-35f4-b03f-e12ef8abfdc7=session_864d19e7-0748-4f5f-8cbc-079c05a8584d; uid=224462; email=bigyuhao%40163.com; key=9d14f09ba268ffcc27035727f31c4340906651a93a659; ip=8538a74d23fa9bbbfceb495e03bfee9a; expire_in=1669969134; crisp-client%2Fsocket%2F50829b14-5009-4006-b9ec-984351fa9529=1",
	)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("签到错误,状态码:%v\n", resp.StatusCode)
	}
	fmt.Println(body)
	v, _ := zhToUnicode(body)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + string(v))
	mylog.Info().Msg(string(v))
}

// 转中文字符串
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

type Handler struct {
	// ...
}

// 用于触发编译期的接口的合理性检查机制
// 如果 Handler 没有实现 http.Handler，会在编译期报错
var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	// ...
}
