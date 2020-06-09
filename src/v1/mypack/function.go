package mypack

import (
	"fmt"
	"math/rand"
	"time"
)

func CheckError(err error) {

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
