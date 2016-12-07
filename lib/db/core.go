// written by London Trust Media
// released under the MIT license
package db

// DB is the core database interface. Our DBs are based on key-value storage.
type DB interface {
	// use same interface as buntdb -- should work nicely across different providers
	// Get(key string) (string, error)
	// Set(key string, value string) error
	Upgrade() error
	Save(filename string) error // ignored for ones that don't use files
}
