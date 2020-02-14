# Aws Autenticação Customizada com Cognito

## Pré requisitos

- Ter uma conta no AWS
- Criar um usuário com permissão de acesso ao COGNITO, o SES e o LAMBDA do AWS
- Criar um Pool de usuários no Cognito
- Setar as váriaveis de ambiente no arquivo `.env`:

```bash
    CONSOLE_AWS_ACCESS_KEY="COLOQUE A CHAVE DE ACESSO DO USUÁRIO COM PERMISSÂO AO COGNITO/SES"
    CONSOLE_AWS_SECRET_KEY="COLOQUE A CHAVE DE SECRETA DO USUÁRIO COM PERMISSÂO AO COGNITO/SES"
    CONSOLE_AWS_REGION="COLOQUE SUA REGIÃO SETADA NO AWS"
    AWS_USER_POOL_ID="COLOQUE O ID DO POOL DE USUÁRIOS CRIADO NO COGNITO"

    # para este exemplo eu crei um Client ID sem Secret Key
    AWS_POOL_CLIENT_ID="COLOQUE O CLIENTE ID DO POOL DO DE USUÁRIOS" 
```

- É necessário compilar e zipar os binários dos triggers, pra enviar para os lambdas
- No painel lambda crie uma função para cada trigger
- Para o trigger `create_custom_auth` é necessário setar as variáveis de ambiente: `CONSOLE_AWS_ACCESS_KEY`, `CONSOLE_AWS_SECRET_KEY` e `CONSOLE_AWS_REGION`, na aba lambda desse trigger.
- Adicionar cada trigger na aba Triggers do Pool de usuários no Cognito que foi criado.