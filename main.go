package main

import (
  "fmt"
  "net/http"
  "html/template"
  "log"
  "os/exec"
)


func openBrowser(url string) {
	var err error

	err = exec.Command("xdg-open", url).Start()
	if err != nil {
		log.Println("Error opening browser:", err)
	}
}


type itemList struct {
  checked bool
  item string
}


func newItemList(checked bool, item string) *itemList {
  it := itemList{checked: checked, item: item}
  return &it
}


func main() {
    
    
    thingsToPack := newItemList(false, "Passport")

    fmt.Println("Hello! Going to start the web server!")
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    //now this is the function keep in mind that h1 means handler #1
    h1 := func (w http.ResponseWriter, r *http.Request) {
      //io.WriteString(w, "Hello World!")
      //io.WriteString(w, r.Method)
      tmp1 := template.Must(template.ParseFiles("index.html"))
      tmp1.Execute(w, nil)
      fmt.Println(thingsToPack.item)
    }


    //This will handle the function.
    //since this is "/" basicly when you go to the homepage it tells the person Hey! go run h1 when visitng the homepage
    http.HandleFunc("/", h1)


    http.HandleFunc("/createCheckList", func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/html")
      w.Write([]byte(`
          <input type="checkbox" id="1" name="1" value="Passport">
          <label for="1"> Passport</label><br>
        `)) 
    })


    openBrowser("http://localhost:8000")
    //log.fatal will recond something if it failed to make a webserver
    //http listen and serve creates that server.
    log.Fatal(http.ListenAndServe(":8000", nil)) 
  }
