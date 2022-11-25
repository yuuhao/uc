package api

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"uc/app/models"
	"uc/utils"

	"github.com/go-playground/validator/v10"

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
	//s := `{"code":100,"data":{"message":"\u7b7e\u5230\u6210\u529f,\u83b7\u5f97\u4e86 267MB \u6d41\u91cf."}}`

	//sText := "hello 你好"
	//textQuoted := strconv.QuoteToASCII(sText)
	//textUnquoted := textQuoted[1 : len(textQuoted)-1]
	//fmt.Println(textUnquoted)
	//v, _ := zhToUnicode([]byte(s))
	//v := []byte([1,2,3])
	fmt.Println(string([]byte(`[1111]`)))
}

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func Ping(c *gin.Context) {
	var user models.User
	utils.DB.Find(&user, "id = 1")

	err := utils.Redis.Set(c, "user", "11", time.Minute).Err()
	if err != nil {
		FailResponse(c, "fail", nil)
	}
	_, err = utils.Redis.Get(c, "user").Result()
	if err != nil {
		FailResponse(c, "fail", nil)
	}

	//utils.InitElastic()
	FailResponse(c, "ok", nil)
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
