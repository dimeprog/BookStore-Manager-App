package  utils


import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func parseBody(r *http.Request, x interface{}){
if body,err:= ioutil.ReadAll(r.Body); err==nil{
	if err := json.Unmarshal([]byte(body),x);err!=nil{
		return
	}
}
}