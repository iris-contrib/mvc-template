package main

import (
	"github.com/iris-contrib/mvc-template/controllers"
	"github.com/iris-contrib/mvc-template/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	// Serve our front-end and its assets.
	app.HandleDir("/", iris.Dir("./app/public"))

	// Note, it's buffered, so make sure it's closed so it can flush any buffered contents.
	ac := accesslog.File("./access.log")
	defer ac.Close()

	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())
	// Group routes and mvc apps based on /api path prefix.
	api := app.Party("/api")
	{
		// Group based on /api/counter path prefix.
		counterAPI := api.Party("/counter")
		// Optionally, a <trick> to keep the `m` local variable
		// unaccessible outside of this block's scope. That
		// way you can register many mvc apps for different Parties
		// with a "m" variable.
		// Alternatively you can use the mvc.Configure function :)

		// Register a new MVC Application to the counterAPI Party.
		m := mvc.New(counterAPI)
		m.Register(
			// Register a static dependency (static because it doesn't accept an iris.Context,
			// only one instance of that it's used). Helps us to keep a global counter across
			// clients requests.
			services.NewGlobalCounter(),
			// Register a dynamic dependency (GetFields accepts an iris.Context,
			// it binds a new instance on every request). Helps us to
			// set custom fields based on the request handler.
			accesslog.GetFields,
		)
		// Register our controller.
		m.Handle(new(controllers.Counter))
	}

	// GET http://localhost:8080/api/counter
	// POST http://localhost:8080/api/counter/increment
	app.Listen(":8080")
}
