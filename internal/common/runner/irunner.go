package runner

// Runner is running commandline commands
type Runner interface {
	Run() error
	String() string
}
