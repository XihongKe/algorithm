package rpc_objects

import "fmt"

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	fmt.Println(args)
	return nil
}
