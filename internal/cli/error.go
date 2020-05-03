package cli

import (
	"github.com/fatih/color"
)

var (
	redFprint = color.New(color.Bold, color.FgRed).FprintfFunc()
	redSprint = color.New(color.FgRed).SprintFunc()

	magentaFprint = color.New(color.FgMagenta).FprintfFunc()
)

// Error must implemented by all cli errors in this package
type Error interface {
	printAndExit()
}

// HandleErrorAndExit can wrap all top level command functions.
// Note: HandleErrorAndExit does not print 'Usage'
// Example: HandleErrorAndExit(logOut())
func HandleErrorAndExit(err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case Error:
		e.printAndExit()
		break
	default:
		// Unspecified error: will be cast as an internal error
		err := &InternalError{err, "refer debug-logs"}
		err.printAndExit()
	}
}
