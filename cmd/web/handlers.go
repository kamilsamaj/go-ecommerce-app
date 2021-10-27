package main

import "net/http"

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	td := templateData{
		String: map[string]string{
			"stripe_key": app.config.stripe.key,
		},
	}
	if err := app.renderTemplate(w, r, "terminal", &td); err != nil {
		app.errorLog.Println(err)
	}
}
