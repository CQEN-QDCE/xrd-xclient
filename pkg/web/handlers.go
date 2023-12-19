package web

import (
	"fmt"
	"github.com/brynjarh/xclient/pkg/forms"
	"net/http"
	"net/url"
	"strings"
)

func (app *application) envapiForm(w http.ResponseWriter, r *http.Request) {
	t:="url=http://10.99.123.78:80&client=CS/GOUV/003/CLIENT1&service=CS/GOUV/004/PROVIDER1/001"
	v, err := url.ParseQuery(t)
	if err != nil {
		panic(err)
	}
    fmt.Println("petapiForm1:")
	app.render(w, r, "envapi.page.gohtml", &templateData{
		Form: forms.New(v),
		Active: "envapi",
	})
} 

func (app *application) healthtapiForm(w http.ResponseWriter, r *http.Request) {
	t:="url=http://10.99.123.78:80&client=CS/GOUV/003/CLIENT1&service=CS/GOUV/002/PROVIDER1/001"
   v, err := url.ParseQuery(t)
   if err != nil {
	   panic(err)
   }
   app.render(w, r, "healthapi.page.gohtml", &templateData{
	   Form: forms.New(v),
	   Active: "healthapi",
   })
} 

func (app *application) envapiGet(w http.ResponseWriter, r *http.Request) {
    // Process the form submission
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.Form)
    form.Required("url", "client", "service")
    form.MatchesPattern("url", forms.UrlRX)
    form.ValidUrl("url")

    if !form.Valid() {
        app.render(w, r, "envapi.page.gohtml", &templateData{
            Form:   form,
            Active: "time",
        })
        return
    }

    baseURL := form.Get("url")
    service := form.Get("service")

    // Construct the URL without query parameters
    baseURL = strings.TrimRight(baseURL, "/")
    u, err := url.Parse(fmt.Sprintf("%s/r1/%s/v2/env/", baseURL, service))
    if err != nil {
        app.clientError(w, http.StatusInternalServerError)
        return
    }  

    // Now we know we have a valid form
    c := &Client{
        BaseURL:     u,
        XRoadClient: form.Get("client"),
        XRoadService: service,
        httpClient:   http.DefaultClient,
    }

	// Construct the URL with the encoded parameter
    // Perform the GET request
    result, req, rep, err := c.do("env/")
    
	// Handle the result and rendering as needed
    app.render(w, r, "envapi.page.gohtml", &templateData{
        Result:         result,
        Form:           form,
        Active:         "envapi",
        RequestHeaders: req,
        ReplyHeaders:   rep,
    })
}

func (app *application) envapiGetAirLevel(w http.ResponseWriter, r *http.Request) {
    // Process the form submission
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.Form)
    form.Required("url", "client", "service", "airlevel")
    form.MatchesPattern("url", forms.UrlRX)
    form.ValidUrl("url")

    if !form.Valid() {
        app.render(w, r, "envapi.page.gohtml", &templateData{
            Form:   form,
            Active: "time",
        })
        return
    }

    baseURL := form.Get("url")
    service := form.Get("service")
    airlevel := form.Get("airlevel")

    // Construct the URL without query parameters
    baseURL = strings.TrimRight(baseURL, "/")
    u, err := url.Parse(fmt.Sprintf("%s/r1/%s/v2/env/", baseURL, service))
    if err != nil {
        app.clientError(w, http.StatusInternalServerError)
        return
    }  

    // Now we know we have a valid form
    c := &Client{
        BaseURL:     u,
        XRoadClient: form.Get("client"),
        XRoadService: service,
        AirLevel:    airlevel,
        httpClient:   http.DefaultClient,
    }

	// Construct the URL with the encoded parameter
    // Perform the GET request
    result, req, rep, err := c.do("env/air/" + airlevel)
    
	// Handle the result and rendering as needed
    app.render(w, r, "envapi.page.gohtml", &templateData{
        Result:         result,
        Form:           form,
        Active:         "envapi",
        RequestHeaders: req,
        ReplyHeaders:   rep,
    })
}

func (app *application) healthapiGet(w http.ResponseWriter, r *http.Request) {
    // Process the form submission
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.Form)
    form.Required("url", "client", "service")
    form.MatchesPattern("url", forms.UrlRX)
    form.ValidUrl("url")

    if !form.Valid() {
        app.render(w, r, "healthapi.page.gohtml", &templateData{
            Form:   form,
            Active: "time",
        })
        return
    }

    baseURL := form.Get("url")
    service := form.Get("service")

    // Construct the URL without query parameters
    baseURL = strings.TrimRight(baseURL, "/")
    u, err := url.Parse(fmt.Sprintf("%s/r1/%s/v2/health/", baseURL, service))
    if err != nil {
        app.clientError(w, http.StatusInternalServerError)
        return
    }  

    // Now we know we have a valid form
    c := &Client{
        BaseURL:     u,
        XRoadClient: form.Get("client"),
        XRoadService: service,
        httpClient:   http.DefaultClient,
    }

	// Construct the URL with the encoded parameter
    // Perform the GET request
    result, req, rep, err := c.do("health/")
    
	// Handle the result and rendering as needed
    app.render(w, r, "healthapi.page.gohtml", &templateData{
        Result:         result,
        Form:           form,
        Active:         "healthapi",
        RequestHeaders: req,
        ReplyHeaders:   rep,
    })
}

func (app *application) healthapiGetByImpact(w http.ResponseWriter, r *http.Request) {
    // Process the form submission
    err := r.ParseForm()
    if err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    form := forms.New(r.Form)
    form.Required("url", "client", "service", "typeimpact")
    form.MatchesPattern("url", forms.UrlRX)
    form.ValidUrl("url")

    if !form.Valid() {
        app.render(w, r, "healthapi.page.gohtml", &templateData{
            Form:   form,
            Active: "time",
        })
        return
    }

    baseURL := form.Get("url")
    service := form.Get("service")
    typeimpact := form.Get("typeimpact")

    // Construct the URL without query parameters
    baseURL = strings.TrimRight(baseURL, "/")
    u, err := url.Parse(fmt.Sprintf("%s/r1/%s/v2/health/", baseURL, service))
    if err != nil {
        app.clientError(w, http.StatusInternalServerError)
        return
    }  

    // Now we know we have a valid form
    c := &Client{
        BaseURL:     u,
        XRoadClient: form.Get("client"),
        XRoadService: service,
        TypeImpact:   typeimpact,
        httpClient:   http.DefaultClient,
    }

	// Construct the URL with the encoded parameter
    // Perform the GET request
    result, req, rep, err := c.do("health/impact/" + typeimpact)
    
	// Handle the result and rendering as needed
    app.render(w, r, "healthapi.page.gohtml", &templateData{
        Result:         result,
        Form:           form,
        Active:         "healthapi",
        RequestHeaders: req,
        ReplyHeaders:   rep,
    })
}