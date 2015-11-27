package main
import "fmt"
import "net/http"
import "encoding/json"
import "net/url"
import "bytes"
import "io/ioutil"

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

type Ambience struct {
    // Default implementation of all Resource methods
    ResourceBase
}

// Override the Get method
func (t Ambience) Put(values url.Values) (int, interface{}) {
    r_light:= values["l"]
    r_hum:= values["h"]
    r_temp:= values["t"]

    client := &http.Client{}

    payload_light := []byte(fmt.Sprintf(`r_light value=%s`, r_light[0]))
    payload_temp  := []byte(fmt.Sprintf(`r_temp value=%s`, r_temp[0]))
    payload_hum   := []byte(fmt.Sprintf(`r_hum value=%s`, r_hum[0]))

    req, err := http.NewRequest("POST", "http://localhost:8086/write?db=ambience", bytes.NewBuffer(payload_light))
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("There was an error storing the lighting data")
    }
    defer resp.Body.Close()

    req, err = http.NewRequest("POST", "http://localhost:8086/write?db=ambience", bytes.NewBuffer(payload_temp))
    resp, err = client.Do(req)
    if err != nil {
        fmt.Println("There was an error storing the temperature data")
    }
    defer resp.Body.Close()

    req, err = http.NewRequest("POST", "http://localhost:8086/write?db=ambience", bytes.NewBuffer(payload_hum))
    resp, err = client.Do(req)
    if err != nil {
        fmt.Println("There was an error storing the humidity data")
    }
    defer resp.Body.Close()

    fmt.Println("Added 3 new measurements to the database")
    return http.StatusOK, ""
}

func (t Ambience) Get(values url.Values) (int, interface{}) {
    client := &http.Client{}
    url := "http://localhost:8086/query?db=ambience&q=SELECT%20value%20from%20r_light,r_temp,r_hum"
    req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("There was an error in fetching data")
    }
    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("There was an error fetching data")
    }

    defer resp.Body.Close()
    s := string(b)
    return http.StatusOK, s
}

func main() {
    var ambience Ambience
    AddResource(ambience, "/ambience")
    fmt.Println("Starting server on 8085")
    http.ListenAndServe(":8085", nil)
}
