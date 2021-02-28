package fortune

import (
	"html/template"
	"os"

	"niffler/chat"
	"niffler/controller"

	"github.com/lunny/log"
	"github.com/lunny/tango"
	"github.com/tango-contrib/events"
	"github.com/tango-contrib/renders"
)

func webServ() {
	l := log.New(os.Stdout, "[tango] ", log.Ldefault())
	l.SetOutputLevel(log.Lfatal)
	t := tango.Classic(l)
	t.Use(
		events.Events(),
		tango.Static(tango.StaticOptions{
			RootPath: "./views/static",
			Prefix:   "static",
		}),
		renders.New(renders.Options{
			Reload:    true,
			Directory: "./views/templates",
			Funcs:     template.FuncMap{},
			Charset:   "UTF-8", // Appends the given charset to the Content-Type header. Default is UTF-8
			// Allows changing of output to XHTML instead of HTML. Default is "text/html"
			HTMLContentType: "text/html",
			DelimsLeft:      "<<<",
			DelimsRight:     ">>>", // default Delims is {{}}, if it conflicts with your javascript template such as angluar, you can change it.
		}),
	)

	t.Get("/", new(controller.Index))
	t.Get("/lucky", new(controller.Lucky))
	t.Get("/update", new(controller.Update))
	t.Run(8888)
}

func Serv() {
	go webServ()
	go chat.MessageLoop()
	sched()
	niffler := NewNiffler()
	niffler.Run()
}
