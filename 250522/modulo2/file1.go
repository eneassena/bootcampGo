package tarefas

import (
	"context"
	"fmt"
)

type contextId int

const contextKeyId contextId = iota

func saldacoes(ctx context.Context) {
	fmt.Println(ctx.Value(contextKeyId))
}

func PlayPratica() {
	ctx := context.Background()

	saldacoes(ctx)
}