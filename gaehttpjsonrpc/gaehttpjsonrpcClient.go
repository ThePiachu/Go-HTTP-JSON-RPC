package gaehttpjsonrpc

import (
    "json"
    "io/ioutil"
    //"log"
    "strings"
    "os"
    "appengine"
    "appengine/urlfetch"
    "http"
)

func Call(c appengine.Context, address string, allowInvalidServerCertificate bool, method string, id interface{}, params []interface{})(map[string]interface{}, os.Error){
    data, err := json.Marshal(map[string]interface{}{
        "method": method,
        "id":     id,
        "params": params,
    })
    if err != nil {
        c.Infof("Marshal: %v", err)
    	return nil, err
    }
    
    req, err:=http.NewRequest("POST", address, strings.NewReader(string(data)))
    if err!=nil{
        c.Infof("Request: %v", err)
    	return nil, err
    }
    
    tr := &urlfetch.Transport{Context: c}
    tr.AllowInvalidServerCertificate=allowInvalidServerCertificate
    
    resp, err:=tr.RoundTrip(req)
    if err != nil {
        c.Infof("Post: %v", err)
    	return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.Infof("ReadAll: %v", err)
    	return nil, err
    }
    result := make(map[string]interface{})
    err = json.Unmarshal(body, &result)
    if err != nil {
        c.Infof("Unmarshal: %v", err)
        c.Infof("%s", body)
    	return nil, err
    }
    //log.Println(result)
    return result, nil
}

func CallWithBasicAuth(c appengine.Context, address string, username string, password string, allowInvalidServerCertificate bool, method string, id interface{}, params []interface{})(map[string]interface{}, os.Error){
    data, err := json.Marshal(map[string]interface{}{
        "method": method,
        "id":     id,
        "params": params,
    })
    if err != nil {
        c.Infof("Marshal: %v", err)
    	return nil, err
    }
    
    req, err:=http.NewRequest("POST", address, strings.NewReader(string(data)))
    if err!=nil{
        c.Infof("Request: %v", err)
    	return nil, err
    }
    req.SetBasicAuth(username, password)
    
    tr := &urlfetch.Transport{Context: c}
    tr.AllowInvalidServerCertificate=allowInvalidServerCertificate
    
    resp, err:=tr.RoundTrip(req)
    if err != nil {
        c.Infof("Post: %v", err)
    	return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.Infof("ReadAll: %v", err)
    	return nil, err
    }
    result := make(map[string]interface{})
    err = json.Unmarshal(body, &result)
    if err != nil {
        c.Infof("Unmarshal: %v", err)
        c.Infof("%s", body)
    	return nil, err
    }
    //log.Println(result)
    return result, nil
}