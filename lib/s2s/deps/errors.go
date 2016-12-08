// written by London Trust Media
// released under the MIT license
package deps

import "errors"

var (
	// ErrorNoProtocol is what it says on the tin.
	ErrorNoProtocol = errors.New("Protocol not found")
	// ErrorSIDIncorrect means that the SID wasn't defined or was incorrect.
	ErrorSIDIncorrect = errors.New("ServerID is either incorrect or not defined")
)
