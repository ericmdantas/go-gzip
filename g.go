package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	am = 100000
)

type mrWriter interface {
	Write([]byte) (int, error)
	Close() error
	Flush() error
}

type newWriter interface {
	NewWrite(*[]byte)
}

type todo struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

type info struct {
	k string
	b []byte
	p os.FileMode
}

func cmp(mw mrWriter, b []byte) {
	mw.Write(b)
	mw.Close()
	mw.Flush()
}

func writeFiles(infos [4]info) {
	for _, v := range infos {
		ioutil.WriteFile(v.k, v.b, v.p)
	}
}

func main() {
	var ts []todo

	for i := 0; i < am; i++ {
		ts = append(ts, todo{
			ID:  i,
			Msg: "message " + strconv.Itoa(i),
		})
	}

	bTs, _ := json.Marshal(ts)

	var b bytes.Buffer
	var z bytes.Buffer

	gz := gzip.NewWriter(&b)
	gzip.NewWriterLevel(gz, gzip.BestCompression)
	cmp(gz, bTs)

	zl := zlib.NewWriter(&z)
	zlib.NewWriterLevel(zl, zlib.BestCompression)
	cmp(zl, bTs)

	var infos [4]info

	infos[0] = info{k: "g.gz", b: b.Bytes(), p: 0644}
	infos[1] = info{k: "g.zlib", b: z.Bytes(), p: 0644}
	infos[2] = info{k: "g.json", b: bTs, p: 0644}
	infos[3] = info{k: "g.b64.txt", b: []byte(base64.StdEncoding.EncodeToString(bTs)), p: 0644}

	writeFiles(infos)
}
