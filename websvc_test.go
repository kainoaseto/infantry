package infantry

import (
	"fmt"
	"github.com/kainoaseto/infantry/web/clients"
	"github.com/kainoaseto/infantry/web/server"
	"testing"
)

// Test out setting up a basic web service with the tank framework
func TestWebService(t *testing.T) {

	// Setup our Application context which contains our Logger initialized from ENV Vars
	// and tracks service metrics and other core settings to the server
	// Overrides will be possible by using kernel seetings passed through here
	// or by setting stuff up further down before the kernel is passed to the websvc
	app_context, err := NewAppContext()
	if err != nil {
		t.Errorf("Failed to create server Kernel!")
	}

	// Tools to allow this application to use
	app_context.AddTool(EnvManager)
	app_context.AddTool(Logrus)

	// Not allowing new clients initialized after this
	// make external communication  decisions here
	app_context.AddClient(Mongo)
	app_context.AddClient(Redis)

	// Create our new web service
	websvc := NewService(app_context)

	// Export our "Context" out which contains references to the clients, env variables,
	// logger, and any other utilties added into the kernel on init.
	// The seed and kernel are immutable after the server is created.
	websvcKrnl := websvc.Ctx.Get()

	// Pass through the context and the Handler will expose
	// struct level variables to access everything to reduce breadcrumbs
	testHandler := test.NewTestHandler(websvcKrnl)
	websvc.Register("/api/v2/test", testHandler.RouteGroup())

	// Start the web service
	websvc.Start()
}
