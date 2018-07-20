package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/always-sunny/gRPC-t/proto"
	"golang.org/x/net/context"
	"github.com/always-sunny/gRPC-t/server/blockchain"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen on poart 8080: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	srv.Serve(listener)
}


type Server struct{
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

//Why do the proto parameters need pointers vs the return? What's 'new' on a struct vs not?
func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash: b.Hash,
			Data: b.Data,
		})
	}
	return resp, nil
}