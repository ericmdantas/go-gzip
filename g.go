package main

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
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
}

func writeFiles(infos [6]info) {
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
	var l bytes.Buffer
	var f bytes.Buffer

	gz := gzip.NewWriter(&b)
	gzip.NewWriterLevel(gz, gzip.BestCompression)
	cmp(gz, bTs)

	zl := zlib.NewWriter(&z)
	zlib.NewWriterLevel(zl, zlib.BestCompression)
	cmp(zl, bTs)

	lz := lzw.NewWriter(&l, lzw.LSB, 8)
	cmp(lz, bTs)

	fl, _ := flate.NewWriter(&f, flate.BestCompression)
	cmp(fl, bTs)

	var infos [6]info

	infos[0] = info{k: "g.gz", b: b.Bytes(), p: 0644}
	infos[1] = info{k: "g.zlib", b: z.Bytes(), p: 0644}
	infos[2] = info{k: "g.lzw", b: l.Bytes(), p: 0644}
	infos[3] = info{k: "g.flate", b: f.Bytes(), p: 0644}
	infos[4] = info{k: "g.json", b: bTs, p: 0644}
	infos[5] = info{k: "g.b64.txt", b: []byte(base64.StdEncoding.EncodeToString(bTs)), p: 0644}

	writeFiles(infos)
}
