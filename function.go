// https://github.com/GoogleCloudPlatform/functions-framework-go#quickstart-hello-world-on-your-local-machine
package function

import (
	"fmt"
	"net/http"
)

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}
