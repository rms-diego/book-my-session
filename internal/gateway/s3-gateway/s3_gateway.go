package s3gateway

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	appconfig "github.com/rms-diego/book-my-session/pkg/config"
)

type s3Gateway struct {
	client *s3.Client
	bucket string
}

type S3GatewayInterface interface {
	Upload(ctx context.Context, file io.Reader, filename string) (*string, error)
	Delete(ctx context.Context, filename string) error
}

var S3Gateway S3GatewayInterface

func newS3Gateway() (S3GatewayInterface, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &s3Gateway{client: client, bucket: appconfig.Env.S3_BUCKET}, nil
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

	result, err := tm.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket: aws.String(g.bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}

	return result.Location, nil
}

func (g *s3Gateway) Delete(ctx context.Context, filename string) error {
	_, err := g.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(g.bucket),
		Key:    aws.String(filename),
	})

	return err
}
