package main

import (
	"flag"
	grpcSetup "github.com/santoshr1016/squash_go/m-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main(){
	var addrPtr = flag.String("address", ":50051", "Address to connect")
	flag.Parse()
	s := grpcSetup.NewServer(*addrPtr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to star GRPC Server")
	}
}
