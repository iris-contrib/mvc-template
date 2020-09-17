package controllers

import (
	"github.com/iris-contrib/mvc-template/models"
	"github.com/iris-contrib/mvc-template/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/mvc"
)

// The Counter Controller.
// The controller instance is always the same and shared across requests, why?
// Because it contains a single field which binds to a static dependency.
// Service is a static dependency, why? Because it does not rely on iris.Context
// or any other dynamic dependency.
//
// The `GetIncrement` method relies on a dynamic dependency,
// because it depends on the incoming request (access log).
// That input argument will be a new instance of accesslog.Fields on every single new request.
type Counter struct {
	Service services.Counter
}

// HandleError catches controller's methods and servetime dependency-injection errors.
func (c *Counter) HandleError(ctx iris.Context, err error) {
	ctx.StopWithError(iris.StatusBadRequest, err)
	// Note that,
	// you can ignore this error and continue by not stopping the execution.
}

// HandleHTTPError catches HTTP Errors under the controller's party (prefix path).
//
// Responds the http error with JSON.
func (c *Counter) HandleHTTPError(err mvc.Err, statusCode mvc.Code) models.ErrorResponse {
	/* OR
	err := ctx.GetErr()
	code := ctx.GetStatusCode()
	*/
	code := int(statusCode)
	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = iris.StatusText(code)
	}

	return models.ErrorResponse{
		Code:    code,
		Message: msg,
	}
}

// PostIncrement handles the POST */increment of this controller's Party.
//
// Responds the new counter value with JSON.
// Access a registered request-time dependency of *accesslog.Fields
// in order to set custom log fields (e.g. current 'counter') in the log file.
func (c *Counter) PostIncrement(fields *accesslog.Fields) models.CounterResponse {
	value := c.Service.Increment()
	fields.Set("counter", value)

	return models.CounterResponse{
		Value: value,
	}
}

// Get handles the GET */ path.
// Returns the current counter.
func (c *Counter) Get() models.CounterResponse {
	value := c.Service.Get()

	return models.CounterResponse{
		Value: value,
	}
}
