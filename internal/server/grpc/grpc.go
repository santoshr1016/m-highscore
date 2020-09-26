package grpc
import (
	"context"
	"github.com/rs/zerolog/log"
	pbhighscore "github.com/santoshr1016/squash_go/m-apis/m-highscore/v1"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string;
	srv *grpc.Server;
}

var HighScore = 99999999.0;

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
