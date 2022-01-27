package impl

type ActionOpts struct {
	RemoveArtifact      bool   `env:"REMOVE_ARTIFACT, default=false"`
	CancelWhenJobPassed bool   `env:"CANCEL_WHEN_JOB_SUCCESSFULLY_PASSED, default=false"`
	Verbose             bool   `env:"VERBOSE, default=false"`
	GitHub              GitHub `env:"GITHUB"`
}

type GitHub struct {
	Token      string `env:"TOKEN, require=true"`
	Repository string `env:"REPOSITORY, require=true"`
	RunID      string `env:"RUN_ID, require=true"`
}
