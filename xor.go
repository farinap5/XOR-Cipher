package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// A simple XOR cipher function required arguments: A key
// as a string and a data as a string of bytes. A string of
// bytes is returned.
func XORCipher(k string,m []byte) []byte {
	var t []byte
	for len(k) < len(m) {
		k = k + k
	}
	for i := 0;i < len(m);i++ {
		z := m[i]^k[i]
		t = append(t,z)
	}
	return t
}

func encode_file(k string,p string)  {
	f,_ := ioutil.ReadFile(p)
	c := XORCipher(k,f)
	fa,_ := os.OpenFile(p, os.O_RDWR, 0644)
	fa.Write(c)
	println("[OK]-",len(c),"bytes converted!")
}

func encry(k string,m []byte) {
	data := XORCipher(k,m)
	sEnc := base64.StdEncoding.EncodeToString(data)
	fmt.Println(sEnc)
}

func decry(k string,m []byte)  {
	dec,_ := base64.StdEncoding.DecodeString(string(m))
	data := XORCipher(k,dec)
	fmt.Println(string(data))
}

func start_file(key string,path string,e bool,d bool) {
	if (e == true) {
		encode_file(key,path)
	}
}

func start(key string,data string,e bool,d bool) {
	if (e == true) {
		encry(key,[]byte(data))
	}
	if (d == true) {
		decry(key,[]byte(data))
	}
}

func main() {
	var p = flag.String("p","","Path.")
	var file = flag.Bool("f",false,"")

	var h = flag.Bool("h",false,"")
	var e = flag.Bool("e",false,"")
	var d = flag.Bool("d",false,"")

	var key = flag.String("k","","Password.")
	var data = flag.String("m","","Data.")
	flag.Parse()

	if (*h == true) {
		println(`
COMMAND MENU
------- ----
-h      Help
-e      Encode
-d      Decode
-k      key
-m      Message 
        encoded/decoded

FILE ENCODING
---- --------
-e      Encode
-k      Key
-f      File
-p      Path
        encoded/decoded
`)
	} else {
		if (*file == true) {
			start_file(*key,*p,*e,*d)
		} else {
			start(*key,*data,*e,*d)
		}
	}
}
