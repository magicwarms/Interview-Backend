package infrastructures

import (
	"basic-api/controllers"

	"github.com/gorilla/mux"
)

func Dispatch() *mux.Router {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", controllers.InitialRoute)
	myRouter.HandleFunc("/question/one", controllers.BackendQuestionOne)
	myRouter.HandleFunc("/question/two", controllers.BackendQuestionTwo)
	myRouter.HandleFunc("/question/three", controllers.BackendQuestionThree)

	return myRouter
}
