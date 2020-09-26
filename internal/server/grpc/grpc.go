package grpc
import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	pbhighscore "github.com/santoshr1016/squash_go/m-apis/m-highscore/v1"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string;
	srv *grpc.Server;
}

var HighScore = 99999999.0;

func NewServer(address string) *Grpc{
	return &Grpc{
		address: address,
	}

}

func(g *Grpc)SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error){
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore

	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

func(g *Grpc)GetHighScore(ctx context.Context, out *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error){
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listner, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to open TCP port")
	}
	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)
	log.Info().Str("Address", g.address).Msg("Starting GRPC Server for HighScore MS")
	err = g.srv.Serve(listner)

	if err != nil{
		return errors.Wrap(err, "Failed to Listen")
	}
	return nil
}