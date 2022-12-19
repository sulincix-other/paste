package main

import (
    b64 "encoding/base64"
)


func b64_encode(data string) (string){
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    return sEnc
}

func b64_decode(data string) (string){
    sDec, _ := b64.StdEncoding.DecodeString(data)
    return string(sDec)
}
