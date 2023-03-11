package sum

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SumThevalue(ctx context.Context, v *Variabls) (*Result, error) {

	log.Printf("Recieved Message body %d and %d", v.Num, v.Num2)

	summme := v.Num + v.Num2
	return &Result{Thesum: summme}, nil
}

func (s *Server) mustEmbedUnimplementedSumServiceServer() {
	// This is a placeholder method to satisfy the SumServiceServer interface
}
