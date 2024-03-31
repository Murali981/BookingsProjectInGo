package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Murali981/bookings/pkg/config"
	"github.com/Murali981/bookings/pkg/handlers"
	"github.com/Murali981/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"


var app config.AppConfig

var session *scs.SessionManager


// main is the main application function
func main() {

	


  /// Change this to true when in production....
       app.InProduction = false  

	session = scs.New()

	session.Lifetime = 24 * time.Hour // The session we have created will last for 24 hours and time is a built-in package of GO..
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session



	 tc , err := render.CreateTemplateCache()

	  if(err != nil) {
		log.Fatal("Cannot create the template cache")
	  }

	  app.TemplateCache = tc

	  app.UseCache = false

	  repo := handlers.NewRepo(&app)

	  handlers.NewHandlers(repo)

	  render.NewTemplates(&app)
	// http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request) {
    //        n , err := fmt.Fprintf(w , "Hello world")
	// 	   if err != nil {
	// 		 fmt.Println(err)
	// 	   }
	// 	   fmt.Println(fmt.Sprintf("Number of bytes written: %d" , n))
	// })

	//   http.HandleFunc("/" , handlers.Repo.Home)

	//   http.HandleFunc("/about" , handlers.Repo.About)

	  

	fmt.Println(fmt.Sprintf("Starting the application on port no %s" , portNumber))  

	// _ = http.ListenAndServe(portNumber , nil)


     srv := &http.Server {
		Addr : portNumber,
		Handler : routes(&app),
	 }

	 err = srv.ListenAndServe()
	 log.Fatal(err)

}