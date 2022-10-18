package main
import(
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/dimeprog/BookStore-Manager-App/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:= mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}