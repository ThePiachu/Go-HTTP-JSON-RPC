package gaehttpjsonrpc

// Copyright 2011-2014 ThePiachu. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/ThePiachu/Go/Log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var TimeoutDuration time.Duration

func init() {
	TimeoutDuration, _ = time.ParseDuration("60s")
}

func Call(c appengine.Context, address string, allowInvalidServerCertificate bool, method string, id interface{}, params []interface{}) (map[string]interface{}, error) {
	Log.Debugf(c, "gaehttpjsonrpc Call - %v", method)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		Log.Infof(c, "Marshal: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(string(data)))
	if err != nil {
		Log.Infof(c, "Request: %v", err)
		return nil, err
	}

	tr := &urlfetch.Transport{Context: c, Deadline: TimeoutDuration, AllowInvalidServerCertificate: allowInvalidServerCertificate}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		Log.Infof(c, "Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Infof(c, "ReadAll: %v", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		Log.Infof(c, "Unmarshal: %v", err)
		Log.Infof(c, "%s", body)
		return nil, err
	}
	//log.Println(result)
	return result, nil
}

func CallWithBasicAuth(c appengine.Context, address string, username string, password string, allowInvalidServerCertificate bool, method string, id interface{}, params []interface{}) (map[string]interface{}, error) {
	Log.Debugf(c, "gaehttpjsonrpc CallWithBasicAuth - %v", method)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		Log.Infof(c, "Marshal: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(string(data)))
	if err != nil {
		Log.Infof(c, "Request: %v", err)
		return nil, err
	}
	req.SetBasicAuth(username, password)

	tr := &urlfetch.Transport{Context: c, Deadline: TimeoutDuration, AllowInvalidServerCertificate: allowInvalidServerCertificate}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		Log.Infof(c, "Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Infof(c, "ReadAll: %v", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		Log.Infof(c, "Unmarshal: %v", err)
		Log.Infof(c, "%s", body)
		return nil, err
	}
	//log.Println(result)
	return result, nil
}

func CallWithBasicAuthSingleParam(c appengine.Context, address string, username string, password string, allowInvalidServerCertificate bool, method string, id interface{}, params interface{}) (map[string]interface{}, error) {
	Log.Debugf(c, "gaehttpjsonrpc CallWithBasicAuthSingleParam - %v", method)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
	if err != nil {
		Log.Infof(c, "Marshal: %v", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", address, strings.NewReader(string(data)))
	if err != nil {
		Log.Infof(c, "Request: %v", err)
		return nil, err
	}
	req.SetBasicAuth(username, password)

	tr := &urlfetch.Transport{Context: c, Deadline: TimeoutDuration, AllowInvalidServerCertificate: allowInvalidServerCertificate}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		Log.Infof(c, "Post: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Log.Infof(c, "ReadAll: %v", err)
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		Log.Infof(c, "Unmarshal: %v", err)
		Log.Infof(c, "%s", body)
		return nil, err
	}
	//log.Println(result)
	return result, nil
}
