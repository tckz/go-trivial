package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	godotenv.Load()

	type Config struct {
		UserPoolID string `required:"true" split_words:"true"`
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("*** envconfig.Process: %v", err)
	}

	optEmail := flag.String("email", "", "email")
	optPasswd := flag.String("passwd", "", "password")
	flag.Parse()

	if *optEmail == "" {
		log.Fatalf("--email must be specified")
	}

	if *optPasswd == "" {
		log.Fatalf("--passwd must be specified")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ctx := context.Background()
	cog := cognitoidentityprovider.New(sess)
	{
		out, err := cog.AdminCreateUserWithContext(ctx, &cognitoidentityprovider.AdminCreateUserInput{
			// 招待メールを送信しない
			MessageAction: aws.String("SUPPRESS"),
			UserAttributes: []*cognitoidentityprovider.AttributeType{
				{
					Name:  aws.String("email"),
					Value: optEmail,
				},
				{
					// email確認済状態に
					Name:  aws.String("email_verified"),
					Value: aws.String("true"),
				},
			},
			UserPoolId: aws.String(cfg.UserPoolID),
			Username:   optEmail,
		})
		if err != nil {
			log.Fatalf("*** AdminCreateUserWithContext: %v", err)
		}
		log.Printf("AdminCreateUserWithContext.out: %v", out)
	}

	// FORCE_CHANGE_PASSWORD 状態になっているので強制的に再設定しPermanent指定することで確認済にできる
	{
		out, err := cog.AdminSetUserPasswordWithContext(ctx, &cognitoidentityprovider.AdminSetUserPasswordInput{
			Password:   optPasswd,
			Permanent:  aws.Bool(true),
			UserPoolId: aws.String(cfg.UserPoolID),
			Username:   optEmail,
		})
		if err != nil {
			log.Fatalf("*** AdminSetUserPasswordWithContext: %v", err)
		}
		log.Printf("AdminSetUserPasswordWithContext.out: %v", out)
	}
}
