package xmlrpc

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//Request : Is used to send the HTTP Request to the XMLRPC WP server
func Request(url string, method string, params ...interface{}) ([]interface{}, error) {
	unserializedResponse := make([]interface{}, 0)
	var err error
	var response *http.Response

	request := Serialize(method, params)
	buffer := bytes.NewBuffer([]byte(request))

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   0,
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	timeout := time.Duration(time.Duration(30) * time.Second)
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	req, _ := http.NewRequest("POST", url, buffer)
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("Connection", "close")

	response, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return unserializedResponse, err
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	unserializedResponse, err = Unserialize(response.Body)

	if err != nil {
		log.Println(err)
		return unserializedResponse, err
	}

	return unserializedResponse, nil
}

type MethodResponse struct {
	Params []Param `xml:"params>param"`
}

type Param struct {
	Value Value `xml:"value"`
}

type Value struct {
	List     []Value  `xml:"array>data>value"`
	Object   []Member `xml:"struct>member"`
	String   string   `xml:"string"`
	Int      string   `xml:"int"`
	Boolean  string   `xml:"boolean"`
	DateTime string   `xml:"dateTime.iso8601"`
}

type Member struct {
	Name  string `xml:"name"`
	Value Value  `xml:"value"`
}

//unserialize : Is used to unserialize the response
func unserialize(value Value) interface{} {
	if value.List != nil {
		result := make([]interface{}, len(value.List))
		for i, v := range value.List {
			result[i] = unserialize(v)
		}
		return result

	} else if value.Object != nil {
		result := make(map[string]interface{}, len(value.Object))
		for _, member := range value.Object {
			result[member.Name] = unserialize(member.Value)
		}
		return result

	} else if value.String != "" {
		return fmt.Sprintf("%s", value.String)

	} else if value.Int != "" {
		result, _ := strconv.Atoi(value.Int)
		return result

	} else if value.Boolean != "" {
		return value.Boolean == "1"

	} else if value.DateTime != "" {
		var format = "20060102T15:04:05"
		result, _ := time.Parse(format, value.DateTime)
		return result
	}

	return nil
}

func Unserialize(buffer io.ReadCloser) ([]interface{}, error) {
	var body []byte
	var err error
	result := make([]interface{}, 0)
	body, err = ioutil.ReadAll(buffer)
	if err != nil {
		log.Println(err)
		return result, err
	}

	var response MethodResponse
	xml.Unmarshal(body, &response)

	result = make([]interface{}, len(response.Params))
	for i, param := range response.Params {
		result[i] = unserialize(param.Value)
	}

	return result, nil
}

func Serialize(method string, params []interface{}) string {
	request := "<methodCall>"
	request += fmt.Sprintf("<methodName>%s</methodName>", method)
	request += "<params>"

	for _, value := range params {
		request += "<param>"
		request += serialize(value)
		request += "</param>"
	}

	request += "</params></methodCall>"

	return request
}

func serialize(value interface{}) string {
	result := "<value>"
	switch value.(type) {
	case string:
		result += fmt.Sprintf("<string>%s</string>", value.(string))
		break
	case int:
		result += fmt.Sprintf("<int>%d</int>", value)
		break
	default:
		if reflect.ValueOf(value).Kind() == reflect.Map {
			result += "<struct>"
			for k, v := range value.(map[string]interface{}) {
				result += "<member>"
				result += fmt.Sprintf("<name>%s</name>", k)
				if strings.EqualFold(k, "bits") {
					result += fmt.Sprintf("<value><base64><![CDATA[---%s---]]></base64></value>", v)
				} else {
					result += serialize(v)
				}
				result += "</member>"
			}
			result += "</struct>"
		}

	}
	result += "</value>"
	return result
}
