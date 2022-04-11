package utils

import (
	"github.com/jinzhu/copier"
	"myHome/gin/utils/logs"
	"time"
)

var nilTime = time.Time{}

func IsNilTime(t time.Time) bool {

	return t == nilTime
}

func Copy(toValue interface{}, fromValue interface{}) {

	err := copier.Copy(toValue, fromValue)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		//fmt.Println(err)
		logs.Error(err).Msg("")
	}
}

func FatalError(err error) {
	if err != nil {
		logs.Fatal(err).Msg("")
	}
}

func PanicError(err error) {
	if err != nil {
		//panic(err)
		logs.Panic(err).Msg("")
	}
}
