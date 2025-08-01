package app

import "github.com/nachoconques0/diagnosis_svc/internal/db"

// Options holds the configuration of the instance
type Options struct {
	dbOptions []db.Option
	// HTTP server configuration
	httpPort string
	// JWT secret config key
	jwtSecret string
}

// Option type to add dependencies to the given Options
type Option func(*Options)

// WithHTTPPort http server port
func WithHTTPPort(p string) Option {
	return func(o *Options) {
		o.httpPort = p
	}
}

// WithJWTSecret secret key
func WithJWTSecret(p string) Option {
	return func(o *Options) {
		o.jwtSecret = p
	}
}

func (b *Options) appendDBOption(o db.Option) {
	if b.dbOptions == nil {
		b.dbOptions = []db.Option{}
	}

	b.dbOptions = append(b.dbOptions, o)
}

// WithDBHost sets the database host provided
func WithDBHost(h string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithHost(h))
	}
}

// WithDBPort sets the database port provided
func WithDBPort(p string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithPort(p))
	}
}

// WithDBName sets the database name provided
func WithDBName(d string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithDatabase(d))
	}
}

// WithDBUser sets the database user provided
func WithDBUser(u string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithUser(u))
	}
}

// WithDBPassword sets the database password provided
func WithDBPassword(p string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithPassword(p))
	}
}

// WithDBMaxConnections sets the database max connections
func WithDBMaxConnections(d int) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithMaxConnections(d))
	}
}

// WithSSLMode sets the database SSL mode
func WithSSLMode(s string) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithSSLMode(s))
	}
}

// WithDebug controls if queries should be logged
func WithDBDebug(d bool) Option {
	return func(o *Options) {
		o.appendDBOption(db.WithDBDebug(d))
	}
}
