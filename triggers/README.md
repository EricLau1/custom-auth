# Custom Authetication com Aws Cognito e Lambda Triggers

https://github.com/aws/aws-lambda-go/blob/master/events/README_Cognito_UserPools_CustomAuthLambdaTriggers.md

## Compilar os Lambdas

```bash
cd triggers

GOOS=linux go build -o define_custom_auth define_custom_auth.go
GOOS=linux go build -o create_custom_auth create_custom_auth.go
GOOS=linux go build -o verify_custom_auth verify_custom_auth.go

```

## Zipar os Lambdas

```bash
cd triggers

zip define_custom_auth.zip define_custom_auth
zip create_custom_auth.zip create_custom_auth
zip verify_custom_auth.zip verify_custom_auth
```

> É necessário digitar o nome do arquivo binário no campo Handler do painel Lambda no AWS

> Para o trigger do `create_custom_auth` é necessário setar 3 env vars no painel do lambda: 
```bash
    # setar essas variáveis no painel do lambda create_custom_auth
    CONSOLE_AWS_ACCESS_KEY
    CONSOLE_AWS_SECRET_KEY
    CONSOLE_AWS_REGION
```