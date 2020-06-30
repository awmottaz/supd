package update

// An Error describes known or expected errors that may arise while using
// this package.
type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// NotFound error occurs when an update cannot be found.
	NotFound = Error("Update not found")
)
