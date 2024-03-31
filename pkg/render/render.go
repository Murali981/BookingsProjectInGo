package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Murali981/bookings/pkg/config"
	"github.com/Murali981/bookings/pkg/models"
)

var app *config.AppConfig


// NewTemplates function sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter , tmpl string , td *models.TemplateData) { // td stands for template data


	var tc map[string]*template.Template

	if app.UseCache {
      
	// get the template cache from the app config....


      tc  = app.TemplateCache
	} else {
		tc , _  = CreateTemplateCache()
	}
    //   // Create a template cache
	//   tc , err := CreateTemplateCache()
	//   if err != nil {
	// 	log.Fatal(err)
	//   }


       // Get the requested template from the cache.....
	   t , ok := tc[tmpl]
	   if !ok {
		log.Fatal("Could not get the template from the template cache")
	   }

	   buf := new(bytes.Buffer)
   
	   td = AddDefaultData(td)
	   err := t.Execute(buf,td)

	   if err != nil {
		  log.Println(err)
	   }



	   // Render the template

	   _ , err = buf.WriteTo(w)
	   if (err != nil) {
		log.Println(err)
	   }

	// parsedTemplate , _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.html")
	// err := parsedTemplate.Execute(w,nil)
	// if err != nil {
	// 	fmt.Println("Error parsing the template" , err)
	// 	return
	// }
 }


   func CreateTemplateCache() (map[string]*template.Template , error) {
	//   myCache := make(map[string]*template.Template)

	    myCache := map[string]*template.Template{}



		// Get  all of the files named *page.html from the "./templates" folder....
          pages , err := filepath.Glob("./templates/*page.html")

		  if err != nil {
			return myCache , err
		  }

		  // range through all the files ending with *page.html....
		  for _,page := range pages {
			name := filepath.Base(page)
			ts,err := template.New(name).ParseFiles(page)
			if err != nil {
				return myCache , err
			  }
	   // ts stands for template set...

        matches , err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache , err
		  }

		  if len(matches) > 0 {
			ts , err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache , err
			}
		  }

		  myCache[name] = ts
       }
         
	    return myCache , nil
   }

//    var tc = make(map[string]*template.Template)

//  func RenderTemplate(w http.ResponseWriter , t string) {
// 	var tmpl *template.Template
// 	var err error


// 	// Check to see if we already have the template in our cache
// 	_ , inMap := tc[t]
// 	if !inMap {
// 		// Need to create a fresh template
// 		log.Println("Creating template and adding it to the cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// We have already a template existing in our cache
// 		log.Println("Using the cached template")
// 	}


// 	    tmpl = tc[t]

// 		err = tmpl.Execute(w,nil)
// 		if err != nil {
// 			log.Println(err)
// 		}
//  }


//  func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s" , t) , "./templates/base.layout.html",
// 	}

//      // Parse the template.......

// 	 tmpl , err := template.ParseFiles(templates...)
// 	 if err != nil {
// 		return err
// 	 }

// 	 // Add the template to the cache  (map)
// 	 tc[t] = tmpl

// 	 return nil


//  }