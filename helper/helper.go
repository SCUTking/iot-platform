package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io/ioutil"
	"iot-platform/define"
	"net/http"
	"strings"
	"time"
)

// Md5 MD5加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func If(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// RFC3339ToNormalTime RFC3339 日期格式标准化
func RFC3339ToNormalTime(rfc3339 string) string {
	if len(rfc3339) < 19 || rfc3339 == "" || !strings.Contains(rfc3339, "T") {
		return rfc3339
	}
	return strings.Split(rfc3339, "T")[0] + " " + strings.Split(rfc3339, "T")[1][:8]
}

func GenerateToken(id uint, identity, name string, second int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(second))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyzeToken 解析token
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// httpRequest .  封装http请求  用于在代码里面调试

func httpRequest(url, method string, data, header []byte) ([]byte, error) {
	var err error
	reader := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 处理 header
	if len(header) > 0 {
		headerMap := new(map[string]interface{})
		err = json.Unmarshal(header, headerMap)
		if err != nil {
			return nil, err
		}
		for k, v := range *headerMap {
			if k == "" || v == "" {
				continue
			}
			request.Header.Set(k, v.(string))
		}
	}
	//request.SetBasicAuth(define.EmqxKey, define.EmqxSecret)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func HttpDelete(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "DELETE", data, header)
}

func HttpPut(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "PUT", data, header)
}

func HttpPost(url string, data []byte, header ...byte) ([]byte, error) {
	return httpRequest(url, "POST", data, header)
}

func HttpGet(url string, header ...byte) ([]byte, error) {
	return httpRequest(url, "GET", []byte{}, header)
}
