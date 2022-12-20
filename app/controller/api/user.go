package api

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"uc/app/helper"
	"uc/utils"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" format:"2006-01-02"`
	Checkout time.Time `form:"check_out" binding:"required,gtfield=CheckIn" format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
func UserList(c *gin.Context) {
	s := `{"code":100,"data":{"message":"\u7b7e\u5230\u6210\u529f,\u83b7\u5f97\u4e86 267MB \u6d41\u91cf."}}`

	sText := "hello 你好"
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)
	fmt.Println(helper.ZhToUnicode([]byte(s)))
}

func Ping(c *gin.Context) {
	//

	// ticker := time.NewTicker(time.Second * 1)
	// defer ticker.Stop()

	for {
		select {
		case <-time.Tick(time.Millisecond * 1):
			fmt.Println("ok")
		}
	}

}

func Cache(c *gin.Context) {
	rds := utils.RedisPool.Get()
	defer rds.Close()

	_, err := rds.Do("SET", "name", "zhangsan")
	if err != nil {
		fmt.Println(err)
	}

	name, _ := rds.Do("GET", "name")
	fmt.Println(cast.ToString(name))
}

func demo() {
	p := new(struct1)
	p.f1 = 2.1
	(*p).i1 = 11
}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}
type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func post(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// always an error is EOF cont be reused
	} else if errB := c.ShouldBind(&objB); errB != nil {
		c.String(http.StatusOK, "the body should be formB %v", errB)
	}
}

func FailResponse(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(200, gin.H{"code": 200, "msg": msg, "data": data})
}
