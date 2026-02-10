package args

type ProgramArgs struct {
	Upstream string `arg:"-u,--upstream" help:"The host to proxy the request to"`
	Port     string `arg:"-p,--port" help:"Where this server should listen for requests" default:":8080"`
}
