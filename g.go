package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

const (
	am = 100000
)

type todo struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

func main() {
	var ts []todo

	for i := 0; i < am; i++ {
		ts = append(ts, todo{
			ID:  i,
			Msg: "message " + strconv.Itoa(i),
		})
	}

	var b bytes.Buffer

	bTs, _ := json.Marshal(ts)

	gz := gzip.NewWriter(&b)
	gz.Write(bTs)
	gz.Close()
	gz.Flush()
	ioutil.WriteFile("g.gz", b.Bytes(), 0644)
	ioutil.WriteFile("g.json", bTs, 0644)
	ioutil.WriteFile("g.b64.txt", []byte(base64.StdEncoding.EncodeToString(bTs)), 0644)

}
