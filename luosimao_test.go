package luosimao

import (
    "log"
    "testing"
    "io/ioutil"
)

func TestSendJSON(t *testing.T) {
    luosimao := NewLuosimao("")
    resp, err := luosimao.SendSMSJSON("18502710852", "hello,wolrd", "【luosimao】")
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(string(body))
}
