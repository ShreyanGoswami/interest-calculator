package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ShreyanGoswami/interest-calculator/adapter/validator"
	"github.com/ShreyanGoswami/interest-calculator/api/action"
	"github.com/ShreyanGoswami/interest-calculator/infra/logger"
	"github.com/ShreyanGoswami/interest-calculator/usecase"
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

func (a GorillaMux) Listen() {
	a.setAppHandlers(a.router)
	a.middleware.UseHandler(a.router)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", a.port),
		Handler:      a.middleware,
	}
	if err := server.ListenAndServe(); err != nil {
		a.log.WithError(err).Fatalln("Error starting REST endpoint server")
	}
}

func (a GorillaMux) setAppHandlers(router *mux.Router) {
	api := router.PathPrefix("/v1").Subrouter()
	api.Handle("/investment/{investment_id}", a.buildGetTodaysInterestAmountAction()).Methods(http.MethodGet)
}

func (a GorillaMux) buildGetTodaysInterestAmountAction() http.Handler {
	var handler http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		var (
			uc  = usecase.NewGetTodaysAmountInteractor({}, )
			act = action.NewCalculateTodaysAmountAction(uc, a.log)
		)
		act.Execute(res, req)
	}

	return negroni.New(
		// TODO add logging middleware
		negroni.NewRecovery(),
		negroni.Wrap(handler),
	)
}
