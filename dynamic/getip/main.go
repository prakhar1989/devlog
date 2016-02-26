package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "unicode/utf8"
    "net/url"
    "strconv"
    "os"
    "io/ioutil"
    "errors"
    "net"
)
//import "bytes"

const (
    GET    = "GET"
    POST   = "POST"
    PUT    = "PUT"
    DELETE = "DELETE"
)

type Resource interface {
    Get(values url.Values) (int, interface{})
    Post(values url.Values) (int, interface{})
    Put(values url.Values) (int, interface{})
    Delete(values url.Values) (int, interface{})
}

type ResourceBase struct{}

func (ResourceBase) Get(values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Post(values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Put(values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, ""
}

func (ResourceBase) Delete(values url.Values) (int, interface{}) {
    return http.StatusMethodNotAllowed, ""
}

func requestHandler(resource Resource) http.HandlerFunc {
    return func(rw http.ResponseWriter, request *http.Request) {

        var data interface{}
        var code int

        request.ParseForm()
        method := request.Method
        values := request.Form

        switch method {
        case GET:
            code, data = resource.Get(values)
        case POST:
            code, data = resource.Post(values)
        case PUT:
            code, data = resource.Put(values)
        case DELETE:
            code, data = resource.Delete(values)
        default:
            rw.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        content, err := json.Marshal(data)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            return
        }

        rw.Header().Set("Access-Control-Allow-Origin", "*")
        rw.Header().Set("Access-Control-Allow-Methods", "GET")
        rw.Header().Set("Access-Control-Allow-Header", "Content-Type")
        rw.WriteHeader(code)
        rw.Write(content)
    }
}

func AddResource(resource Resource, path string) {
    http.HandleFunc(path, requestHandler(resource))
}

type GetIp struct {
    // Default implementation of all Resource methods
    ResourceBase
}

func verifyMid(mid_str string) (mid_int int, err error) {
    // Verify that a machine id was provided
    id_len := utf8.RuneCountInString(mid_str)
    if id_len <= 0 {
        return 0, errors.New("Please provide a machine id\n")
    }

    // Verify the machine name is numeric
    mid_int, err = strconv.Atoi(mid_str)
    if err != nil {
        return 0, errors.New(fmt.Sprintf("%d does not look like a number.\n", mid_str))
    }

    // Verify that the file nae actually exists
    fname := fmt.Sprintf("%d_latest.txt", mid_int)
    _, err = os.Stat(fname)
    if os.IsNotExist(err) {
        return 0, errors.New(fmt.Sprintf("%s is an invalid machine id.\n", mid_str))
    }

    // Simply return the contents of the file
    f, err := os.Open(fname)
    if err != nil {
        return 0, errors.New(fmt.Sprintf("There was an error reading: %s\n%s\n", fname, err))
    }

    f.Close()

    return mid_int, nil
}

func verifyMip(ip_str string) (net.IP, error) {
    ip := net.ParseIP(ip_str)
    if ip == nil {
        return net.IPv4(255,255,255,255), errors.New(fmt.Sprintf("The IP address is invalid: %s", ip_str))
    }

    return ip, nil
}

// This method stores the provided IP and timestamp
func (t GetIp) Put(values url.Values) (int, interface{}) {
    machine_id := values.Get("mid")
    machine_ip := values.Get("mip")
    
    mid, err := verifyMid(machine_id)
    if err != nil {
        fmt.Printf("%s\n", err)
        return http.StatusNotFound, ""
    }

    mip, err := verifyMip(machine_ip)
    if err != nil {
        fmt.Printf("%s\n", err)
        return http.StatusNotFound, ""
    }

    fmt.Printf("%s\n", mip)
    fname := fmt.Sprintf("%d_latest.txt", mid)
    f, err := os.Create(fname)
    if err != nil {
        fmt.Printf("%s\n", err)
        return http.StatusNotFound, ""
    }

    f.WriteString(machine_ip)
    defer f.Close()

    fname = fmt.Sprintf("%d_all.txt", mid)
    f, err = os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0600)
    if os.IsNotExist(err) {
        f, err = os.Create(fname)
    } else if err != nil {
        panic(err)
    }
    defer f.Close()
    f.WriteString(fmt.Sprintf("%s\n", machine_ip))

    return http.StatusOK, mid
}

// This method retrieves the latest IP and timestamp fo the given
func (t GetIp) Get(values url.Values) (int, interface{}) {
    machine_id := values.Get("mid")

    mid, err := verifyMid(machine_id)
    if err!= nil {
        fmt.Printf("%s\n", err);
        return http.StatusNotFound, ""
    }

    fname := fmt.Sprintf("%d_latest.txt", mid)

    // Try reading the data
    ipaddr, err := ioutil.ReadFile(fname)
    if err != nil {
        fmt.Printf("Error reading the ip address file:%s\n%s\n", fname, err)
        return http.StatusNotFound, ""
    }

    // Try opening the latest file for this machine
    return http.StatusOK, ipaddr
}

func main() {
    var getip GetIp
    AddResource(getip, "/getip")
    fmt.Println("Starting server on 8089")
    err := http.ListenAndServe(":8089", nil)
    if err != nil {
        fmt.Printf("There was an error: %s\n", err)
    }
}
