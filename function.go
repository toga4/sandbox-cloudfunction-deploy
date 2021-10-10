// https://github.com/GoogleCloudPlatform/functions-framework-go#quickstart-hello-world-on-your-local-machine
package function

import (
	"fmt"
	"net/http"
)

func SandboxCFDeploy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")
}
