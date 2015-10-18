package main

import (
	"net/http"
	// "fmt"
	// "bytes"
	// "io/ioutil"
	// "encoding/json"
	// "io"
	// "log"
	"strconv"
	"log"
	"github.com/gameit/lol"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	)

func main() {
	// str := lol.GetTournaments()
	// fmt.Println(str)

	m := martini.New()
    route := martini.NewRouter()

    m.Use(func(c martini.Context, w http.ResponseWriter, r *http.Request) {
        // Use indentations. &pretty=1
        pretty, _ := strconv.ParseBool(r.FormValue("pretty"))
        // Use null instead of empty object for json &null=1
        null, _ := strconv.ParseBool(r.FormValue("null"))
        // Some content negotiation
        switch r.Header.Get("Content-Type") {
        case "application/xml":
            c.MapTo(encoder.XmlEncoder{PrettyPrint: pretty}, (*encoder.Encoder)(nil))
            w.Header().Set("Content-Type", "application/xml; charset=utf-8")
        default:
            c.MapTo(encoder.JsonEncoder{PrettyPrint: pretty, PrintNull: null}, (*encoder.Encoder)(nil))
            w.Header().Set("Content-Type", "application/json; charset=utf-8")
        }
    })

    route.Get("/tournament", func(enc encoder.Encoder) (int, []byte) {
        result := lol.GetTournaments()
        return http.StatusOK, encoder.Must(enc.Encode(result))
    })

    m.Action(route.Handle)

    log.Println("Waiting for connections...")

    if err := http.ListenAndServe(":8000", m); err != nil {
        log.Fatal(err)
    }

 //  m := martini.Classic()
 //  m.Get("/news", func() string {
 //    return getNews()
 //  })

 //  m.Get("/tournament", func() string {
	// tournaments = lol.GetTournaments()
	// return json.NewEncoder(w).Encode(todos)
 //  })
  // m.Get("/tournament", func(params martini.Params) string {
  //   return getTournaments(params["status"])
  // })
  m.Run()
}