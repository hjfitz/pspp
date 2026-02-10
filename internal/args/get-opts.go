package args

import "github.com/alexflint/go-arg"

func GetOpts() ProgramArgs {

	var opts ProgramArgs

	arg.MustParse(&opts)

	return opts

}
