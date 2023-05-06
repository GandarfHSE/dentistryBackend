package img

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/GandarfHSE/dentistryBackend/utils/algo"
	"github.com/GandarfHSE/dentistryBackend/utils/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

// return URL
func UploadImage(image *Image) (string, error) {
	if image == nil {
		return "", errors.New("Image is nil!")
	}

	s3Config := config.GetS3Config()

	session, err := session.NewSession(&aws.Config{
		Region:   &s3Config.Region,
		Endpoint: &s3Config.Endpoint,
		Credentials: credentials.NewStaticCredentials(
			s3Config.AccessKey,
			s3Config.SecretKey,
			"",
		),
	})
	if err != nil {
		return "", err
	}

	imgKey, err := uploadImpl(session, image)
	if err != nil {
		return "", err
	}

	imgURL := s3Config.Endpoint + "/" + s3Config.Bucket + "/" + imgKey
	log.Info().Msg(fmt.Sprintf("Image was successfully uploaded to %s!", imgURL))
	return imgURL, nil
}

func uploadImpl(s *session.Session, img *Image) (string, error) {
	s3Config := config.GetS3Config()

	uploader := s3manager.NewUploader(s)
	imgKey := "images/" + algo.GenerateRandomString(10) + "." + img.Ext

	log.Info().Msg(fmt.Sprintf("Uploading image to s3 with key %s...", imgKey))
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &s3Config.Bucket,
		Key:    aws.String(imgKey),
		Body:   bytes.NewReader(img.Data),
	})
	if err != nil {
		return "", err
	}

	return imgKey, nil
}
