package cmd

// Command is an executable sub-command.
type Command interface {
	Summary() string
	Usage()
	Run(args []string) int
}
