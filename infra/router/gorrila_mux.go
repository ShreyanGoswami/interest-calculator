package router

import (
	"log"
	"net/http"
	"time"

	"github.com/ShreyanGoswami/interest-calculator/adapter/validator"
	"github.com/ShreyanGoswami/interest-calculator/infra/logger"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// TODO inject repository here
type GorillaMux struct {
	router     *mux.Router
	middleware *negroni.Negroni
	log        logger.Logger
	validator  validator.Validator
	port       int64
	ctxTimeout time.Duration
}

func NewGorillaMux(
	log logger.Logger,
	validator validator.Validator,
	port int64,
	timeout time.Duration,
) *GorillaMux {
	return &GorillaMux{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		log:        log,
		validator:  validator,
		port:       port,
		ctxTimeout: timeout,
	}
}

func Start() {
	router := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8080", router)) // TODO inject port here
}

func (a GorillaMux) setAppHandlers(router *mux.Router) {
	api := router.PathPrefix("/v1").Subrouter()
	api.Handle("/investment/{investment_id}", a.buildGetTodaysInterestAmountAction()).Methods(http.MethodGet)

}

func (a GorillaMux) buildGetTodaysInterestAmountAction() http.Handler {
	// Effectively a closure
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {

	}

	return negroni.New(
		// TODO add logging middleware
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
