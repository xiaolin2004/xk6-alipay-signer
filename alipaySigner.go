package xk6_alipay_signer

import (
	"fmt"
	"github.com/alipay/global-open-sdk-go/com/alipay/api/tools"
	"go.k6.io/k6/js/modules"
	"strconv"
	"time"
)

func init() {
	modules.Register("k6/x/alipaySigner", new(AlipaySigner))
}

type AlipaySigner struct {
}

func (signer AlipaySigner) GenSignatureHeader(clientId, path, httpMethod, privateKey, jsonBody string) (headers map[string]string, err error) {
	reqTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	sign, err := tools.GenSign(fmt.Sprintf("%s", httpMethod), path, clientId, reqTime, jsonBody, privateKey)
	if err != nil {
		return nil, err
	}
	headers = BuildBaseHeader(reqTime, clientId, "1", sign)
	return headers, nil
}

func BuildBaseHeader(reqTime string, clientId string, keyVersion string, signatureValue string) map[string]string {
	if keyVersion == "" {
		keyVersion = "1"
	}
	signatureValue = "algorithm=RSA256,keyVersion=" + keyVersion + ",signature=" + signatureValue
	return map[string]string{
		"Content-Type": "application/json; charset=UTF-8",
		"Request-Time": reqTime,
		"Client-Id":    clientId,
		"Key-Version":  keyVersion,
		"Signature":    signatureValue,
	}
}
