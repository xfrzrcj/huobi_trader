package untils

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/xfrzrcj/huobi_trader/conf"
)

func generateParameters() map[string]string {
	params := make(map[string]string)
	params["AccessKeyId"] = conf.ACCESS_KEY
	params["SignatureMethod"] = "HmacSHA256"
	params["SignatureVersion"] = "2"
	params["Timestamp"] = "2018-07-16T06:06:42"

	return params
}

func generateSignature() string {

	params := generateParameters()

	hostName := conf.HOST_NAME
	strRequestPath := "/v1/order/orders/place"
	secretKey := conf.SECRET_KEY

	signature := CreateSign(params, "POST", hostName, strRequestPath, secretKey)
	return signature
}

func generateSignature_to_channel(ch chan string) {
	ch <- generateSignature()
}

func Test_digest_hash(t *testing.T) {
	expectSignature := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	signature := generateSignature()
	assert.Equal(t, signature, expectSignature, "The signatures should be the same")
}

func Test_digest_in_parallel(t *testing.T) {
	len := 20
	channel := make(chan string, len)
	for i := 0; i < len; i++ {
		go generateSignature_to_channel(channel)
	}

	//todo: replace with actual signature
	expectSignature := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < cap(channel); i++ {
		signature, _ := <-channel
		assert.Equal(t, signature, expectSignature)
	}
}

func Benchmark_private_signature_performance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		params := generateParameters()

		hostName := conf.HOST_NAME
		strRequestPath := "/v1/order/orders/place"
		secretKey := conf.SECRET_KEY

		signature := CreateSign(params, "POST", hostName, strRequestPath, secretKey)
		CreatePrivateSignByJWT(signature)
	}
}
