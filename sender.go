package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Vehicle struct {
	Id   string
	Name string
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2 * time.Second)
		payLoad := Vehicle{Id: randomString(5), Name: randomString(20)}
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(payLoad)
		url := "http://localhost:8080/handleme"
		res, _ := http.Post(url, "application/json;charset=utf-8", b)
		io.Copy(os.Stdout, res.Body)
	}
}
