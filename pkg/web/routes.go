package web

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, app.session.Enable)

	mux := pat.New()
	mux.Get("/", standardMiddleware.ThenFunc(app.healthtapiForm))
	mux.Get("/envapi", standardMiddleware.ThenFunc(app.envapiForm))
    mux.Get("/envapi-get", standardMiddleware.ThenFunc(app.envapiGet))
	mux.Get("/envapi-airlevel", standardMiddleware.ThenFunc(app.envapiGetAirLevel))
	mux.Get("/healthapi", standardMiddleware.ThenFunc(app.healthtapiForm))
	mux.Get("/healthapi-get", standardMiddleware.ThenFunc(app.healthapiGet))
	mux.Get("/healthapi-typeimpact", standardMiddleware.ThenFunc(app.healthapiGetByImpact))
	
	return standardMiddleware.Then(mux)
}