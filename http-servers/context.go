/*
CONTEXT: https://pkg.go.dev/context
The context package in Go's standard library is used to pass request-scoped information through your program.

In HTTP servers, the most important parts are cancellation and timeouts.
	When a client disconnects, a request times out, or the server shuts down,
	the request's context can tell the rest of your code to stop working on that request.

Every http.Request has a context:

	ctx := r.Context()

	That context belongs to the current HTTP request. If the request is canceled, ctx is canceled too.
*/

/*
EXAMPLE USE IS IN DATABASE CALLS.
	Many database APIs accept a context.Context as their first argument. SQLC-generated methods are no exception:

	user, err := cfg.db.CreateUser(ctx, params.Email)

	By passing the request context to the database call, the database work is tied to the lifetime of the HTTP request.
	If the client gives up before the query finishes, the query can be canceled instead of wasting server resources.

	In a handler, this usually means passing r.Context() directly:

	user, err := cfg.db.CreateUser(r.Context(), params.Email)
*/

/*
BACKGROUND CONTEXT: https://pkg.go.dev/context#Background
You'll also see context.Background() in Go code.
It's useful when a Context is expected but there's no incoming request or parent operation to start from
	 – like in startup code or a background job.

For web handlers, prefer r.Context(). It carries the cancellation signal for the specific request you're handling.
*/

package main
