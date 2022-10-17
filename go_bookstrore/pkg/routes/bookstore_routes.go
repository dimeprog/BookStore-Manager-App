package routes

import(
	"github.com/gorilla/mux"
	"github.com/dimeprog/BookStore-Manager-App/pkg/controllers"
)


var RegisterBookstoreRoutes= func (router *mux.Router)  {
	
	 router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	 router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	 router.HandleFunc("/book/", controllers.GetBookBYID).Methods("GET")
	 router.HandleFunc("/book/", controllers.DeleteBook).Methods("DELETE")
	 router.HandleFunc("/book/", controllers.UpdateBook).Methods("PUT")
}