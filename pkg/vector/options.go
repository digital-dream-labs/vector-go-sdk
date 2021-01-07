package vector

// Option provides a function definition to set options
type Option func(*options)

type options struct {
	target string
	token  string
}

// WithTarget sets the ip of the vector robot.
func WithTarget(s string) Option {
	return func(o *options) {
		o.target = s
	}
}

// WithToken set the token for the vector robot.
func WithToken(s string) Option {
	return func(o *options) {
		o.token = s
	}
}
