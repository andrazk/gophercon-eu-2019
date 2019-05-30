package diagnostics

import "fmt"

var (
	revision  string
	buildTime string
	branch    string
)

// Version returns build info
func Version() string {
	return fmt.Sprintf("revision: %s, build_time: %s, branch: %s", revision, buildTime, branch)
}
