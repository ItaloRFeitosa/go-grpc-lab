package main

import (
	"context"
	"sync"

	pb "github.com/italorfeitosa/go-grpc-lab/proto"
)

type Handler struct {
	pb.UnimplementedQuotesServer
	quotes []*pb.QuoteData
	mu     sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		quotes: []*pb.QuoteData{},
	}
}

func (h *Handler) ListQuotes(ctx context.Context, req *pb.ListQuotesRequest) (*pb.ListQuotesReply, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	quotes := &pb.ListQuotesReply{
		Quotes: h.quotes,
	}

	return quotes, nil
}

func (h *Handler) AddQuote(ctx context.Context, req *pb.QuoteData) (*pb.AddQuoteReply, error) {
	h.mu.Lock()
	h.quotes = append(h.quotes, req)
	h.mu.Unlock()

	reply := &pb.AddQuoteReply{
		Quote: req,
	}

	return reply, nil
}
