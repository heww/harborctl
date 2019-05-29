package version

var (
	// Package is filled at linking time
	Package = "github.com/heww/harborctl"

	// Version holds the complete version number. Filled in at linking time.
	Version = "0.0.1+unknown"

	// Revision is filled with the VCS (e.g. git) revision being used to build
	// the program at linking time.
	Revision = ""
)
