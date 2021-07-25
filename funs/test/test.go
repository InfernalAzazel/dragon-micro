package test

import "context"

func (t *Arith) Test(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}