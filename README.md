# go-game-chatgpt

```markdown
# API REST para Consumir a API ChatGPT

Este repositório contém uma API REST escrita em Go que consome a API ChatGPT para responder a perguntas e retornar o texto fornecido como entrada.

## Pré-requisitos

- Go (versão 1.22.0 ou superior)
- Credenciais da API ChatGPT

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/seu-usuario/nome-do-repositorio.git
   ```

2. Acesse o diretório do projeto:

   ```bash
   cd nome-do-repositorio
   ```

3. Instale as dependências do projeto:

   ```bash
   go mod tidy
   ```

4. Configure suas credenciais da API ChatGPT no arquivo de configuração `config.json`.

5. Execute a aplicação:

   ```bash
   go run main.go
   ```

A API estará disponível em `http://localhost:8000`.

## Uso

### Endpoint

`POST /ask`

Este endpoint aceita uma solicitação JSON contendo a pergunta e retorna um JSON com a resposta da API ChatGPT.

#### Requisição

```json
{
  "question": "Qual é a capital do Brasil?"
}
```

#### Resposta

```json
{
  "answer": "A capital do Brasil é Brasília."
}
```

## Configuração

O arquivo `config.json` contém as configurações da aplicação, incluindo as credenciais da API ChatGPT.

```json
{
  "chatgpt_api_key": "SUA_CHAVE_DE_API_DO_CHATGPT"
}
```