package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/thegrandpackard/go-testing/cases"
)

type RESTService struct {
	r   *mux.Router
	srv *http.Server
	c   *cases.Cases
}

type RESTEndpoint interface {
	StartServer(RESTAddress string) (err error)
	HandleFunc(route string, f func(w http.ResponseWriter, r *http.Request, c *cases.Cases))
}

func Init(cases *cases.Cases) (restService *RESTService, err error) {
	restService = &RESTService{
		c: cases,
		r: mux.NewRouter(),
	}

	return
}

func (r *RESTService) HandleFunc(route string, f func(w http.ResponseWriter, r *http.Request, c *cases.Cases)) {
	r.r.HandleFunc(route, r.CaseHandler(f))
}

func (r *RESTService) CaseHandler(f func(w http.ResponseWriter, req *http.Request, c *cases.Cases)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		f(w, req, r.c)
	}
}

func (r *RESTService) StartServer(RESTAddress string) (err error) {

	// myH is your app's http handler, perhaps a http.ServeMux or similar.
	var myH = context.ClearHandler(r.r)
	// wrappedH wraps myH in order to log every request.
	wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(myH, w, r)
		// user, _ := getRequestUser(r)
		log.Printf("%s %s \"%s %s\" %d %d %s \"%s\"", r.RemoteAddr, "-", r.Method, r.URL.Path, m.Code, m.Written, m.Duration, r.UserAgent())
	})

	r.srv = &http.Server{
		Handler: wrappedH,
		Addr:    RESTAddress,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	return r.srv.ListenAndServe()
}
