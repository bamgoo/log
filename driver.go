package log

type (
	Driver interface {
		Connect(*Instance) (Connect, error)
	}

	Connect interface {
		Open() error
		Close() error
		Write(logs ...Log) error
	}
)
