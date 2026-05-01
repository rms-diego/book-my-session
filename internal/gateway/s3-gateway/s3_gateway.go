package s3gateway

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appconfig "github.com/rms-diego/book-my-session/pkg/config"
)

type s3Gateway struct {
	client *s3.Client
}

type S3GatewayInterface interface {
	Upload(ctx context.Context, file io.Reader, filename string) (*string, error)
	Delete(ctx context.Context, filename string) error
	buildObjectUrl(filename string) string
}

var S3Gateway S3GatewayInterface

func newS3Gateway() (S3GatewayInterface, error) {
	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(appconfig.Env.AWS_REGION),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				appconfig.Env.AWS_ACCESS_KEY_ID,
				appconfig.Env.AWS_SECRET_ACCESS_KEY,
				"",
			),
		),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &s3Gateway{client}, nil
}

func Init() error {
	gateway, err := newS3Gateway()
	if err != nil {
		return err
	}

	S3Gateway = gateway
	return nil
}

func (g *s3Gateway) Upload(ctx context.Context, file io.Reader, filename string) (*string, error) {
	tm := transfermanager.New(g.client)

	_, err := tm.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket: aws.String(appconfig.Env.AWS_S3_BUCKET),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return nil, err
	}

	url := g.buildObjectUrl(filename)
	return &url, nil
}

func (g *s3Gateway) buildObjectUrl(filename string) string {
	segments := strings.Split(filename, "/")
	for i := range segments {
		segments[i] = url.PathEscape(segments[i])
	}

	staticURL := fmt.Sprintf(
		"https://%v.s3.%v.amazonaws.com/%v",
		appconfig.Env.AWS_S3_BUCKET,
		appconfig.Env.AWS_REGION,
		strings.Join(segments, "/"),
	)

	return staticURL
}

func (g *s3Gateway) Delete(ctx context.Context, filename string) error {
	_, err := g.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(appconfig.Env.AWS_S3_BUCKET),
		Key:    aws.String(filename),
	})

	return err
}
