package infantry

import (
	"fmt"
	"github.com/kainoaseto/infantry/web/clients"
	"github.com/kainoaseto/infantry/web/server"
	"testing"
)

// Test out setting up a basic web service with the tank framework
func TestWebService(t *testing.T) {

	// Setup Our Kernel which contains our Logger initialized from ENV Vars
	// and tracks service metrics and other core settings to the server
	// Overrides will be possible by using kernel seetings passed through here
	// or by setting stuff up further down before the kernel is passed to the websvc
	kernel, err := NewKernel()
	if err != nil {
		t.Errorf("Failed to create server Kernel!")
	}

	// Get an environement variable from our ENV, this will be accssed via server
	// after creation
	some_test_var := kernel.Env.Get("TEST_ENV_VAR", "DEFAULT_VALUE_IF_NOT_FOUND")
	port := kernel.Env.Get("PORT", "4001")
	kernel.Server.Port = port

	// Not allowing new clients initialized after this
	// make external communication  decisions here
	kernel.Clients = clients.NewClients(
		clients.Redis,
		clients.Mongo,
		clients.Sql,
		clients.Cassandra,
	)

	// Create our new web service
	websvc := server.New(kernel)

	// Export our "Context" out which contains references to the clients, env variables,
	// logger, and any other utilties added into the kernel on init.
	// The seed and kernel are immutable after the server is created.
	websvcKrnl := websvc.Kernel.Get()

	// Pass through the context and the Handler will expose
	// struct level variables to access everything to reduce breadcrumbs
	testHandler := test.NewTestHandler(websvcKrnl)
	websvc.Register("/api/v2/test", testHandler.RouteGroup())

	// Start the web service
	websvc.Start()
}
