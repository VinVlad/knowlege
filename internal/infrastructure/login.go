package infrastructure

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Login struct {
	app.Compo

	Pass string
}

func (l *Login) Render() app.UI {
	return app.Section().
		Class("consultation").
		Body(
			app.Div().
				Class("container").
				Body(
					app.Form().
						ID("consultation-form").
						Class("feed-form").
						Body(
							app.Div().
								Class("input-effect").
								Body(
									app.Input().
										Class("effect-18").
										Name("email").
										Required(true).
										Placeholder("E-mail").
										Type("text").
										OnChange(l.onClick),
									app.Span().
										Class("focus-border"),
								),
							app.Div().
								Class("input-effect").
								Body(
									app.Input().
										Class("effect-18").
										Name("password").
										Required(true).
										Placeholder("Пароль").
										Type("text").
										OnChange(l.onClick),
									app.Span().
										Class("focus-border"),
								),
							app.Button().
								Class("button").
								Class("button_submit").
								Text("Войти").
								OnClick(l.onClick).
								OnSubmit(l.onClick).
								OnChange(l.onClick),
						).
						OnSubmit(l.onClick),
				),
		)
}

func (l *Login) onClick(ctx app.Context, e app.Event) {
	l.Pass = ctx.JSSrc().Get("value").String()
	fmt.Println(l.Pass)
}
