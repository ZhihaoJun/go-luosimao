package luosimao

import (
    "strings"
    "net/url"
    "net/http"
)

const (
    SEND_SMS_JSON = "http://sms-api.luosimao.com/v1/send.json"
    SEND_SMS_XML = "http://sms-api.luosimao.com/v1/send.xml"
    SEND_BATCH_JSON = "http://sms-api.luosimao.com/v1/send_batch.json"
    SEND_BATCH_XML = "http://sms-api.luosimao.com/v1/send_batch.xml"
    STATUS_JSON = "http://sms-api.luosimao.com/v1/status.json"
    STATUS_XML = "http://sms-api.luosimao.com/v1/status.xml"
)

type Luosimao struct {
    ApiKey string
}

func NewLuosimao(apiKey string) *Luosimao {
    return &Luosimao{
        ApiKey: apiKey,
    }
}

func signatureFix(signature string) string {
    if !strings.HasPrefix(signature, "【") {
        signature = "【" + signature
    }
    if !strings.HasSuffix(signature, "】") {
        signature += "】"
    }
    return signature
}

func (this *Luosimao) SendSMSJSON(mobile string, content string, signature string) (*http.Response, error) {
    values := url.Values{}
    values.Add("mobile", mobile)
    values.Add("message", content + signatureFix(signature))
    query := values.Encode()
    req, err := http.NewRequest("POST", SEND_SMS_JSON, strings.NewReader(query))

    // add auth
    req.SetBasicAuth("api", "key-" + this.ApiKey)
    if err != nil {
        return nil, err
    }
    return http.DefaultClient.Do(req)
}
