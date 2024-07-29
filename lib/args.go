package lib

import (
	"flag"
	"os"
)

type Args struct {
	TargetHost string
	TargetPort int
}

func ArgParser() Args {
	targetHost := flag.String("t", "", "Target host")
	targetPort := flag.Int("p", 0, "Target port")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(-1)
	}
	args := Args{
		TargetHost: *targetHost,
		TargetPort: *targetPort,
	}
	return args
}
