package database

type DbOptions struct {
	host string
	user string
	pass string
	db   string
}

// create default database connect options
func defaultOptions() DbOptions {
	return DbOptions{
		host: "localhost",
		user: "sa",
		pass: "sa123456789@",
		db:   "UniQUEUE",
	}
}

// func NewConn(opts ...DbOptions) {
// 	// opts := defaultOptions()
// 	for _, a := range opts {

// 	}

// }
