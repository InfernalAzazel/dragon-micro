package jd

import (
	"context"
)

type GetFormWidgetsArgs struct {
	A int
	B int
}

type GetFormWidgetsReply struct {
	C int
}

func (t *GetFormWidgets) Mul(ctx context.Context, args *GetFormWidgetsArgs, reply *GetFormWidgetsReply) error {
	reply.C = args.A * args.B
	return nil
}