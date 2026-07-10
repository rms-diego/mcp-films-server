# MCP Movies server

## O que o projeto faz

Servidor [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) escrito em Go que expõe ferramentas de busca de filmes. Ele consome a [API do TMDB](https://developer.themoviedb.org/) e disponibiliza os resultados para clientes MCP (como assistentes de IA) através de HTTP.

O servidor registra ferramentas (tools) MCP que um cliente — como um assistente de IA — pode descobrir e invocar. Cada tool encapsula uma consulta a uma fonte de dados externa (ex: catálogo de filmes e séries do TMDB) e devolve o resultado ao cliente. A comunicação MCP acontece via HTTP no endpoint `/mcp`, usando o transporte _Streamable HTTP_ do SDK oficial de Go. Novas capacidades são adicionadas como novos módulos, sem alterar o núcleo do servidor.

## Stack

| Camada           | Tecnologia                                                                    |
| ---------------- | ----------------------------------------------------------------------------- |
| Linguagem        | Go 1.25                                                                       |
| HTTP framework   | [gin](https://github.com/gin-gonic/gin)                                       |
| Hot reload (dev) | [air](https://github.com/air-verse/air)                                       |
| Protocolo MCP    | [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk) |
| Fonte de dados   | [TMDB API](https://developer.themoviedb.org/)                                 |

## Arquitetura

O projeto segue uma arquitetura em camadas, organizada por módulos (features) e com injeção de dependência via interfaces, separando entrada (handler), regra de negócio (service) e acesso a dados externos (gateway).

```
cmd/main.go                     Ponto de entrada: carrega config, cria o MCP server e sobe o Gin
│
├── internal/config             Carrega e valida variáveis de ambiente (.env)
├── internal/consts             Constantes (ex: URLs base de APIs externas)
├── internal/routes
│   ├── routes.go               Rotas HTTP
│   └── mcp-routes.go           Grupo /mcp (GET, POST, DELETE) + registro dos módulos MCP
│
├── internal/modules/<module>   Cada módulo é uma feature (ex: movie)
│   ├── <module>.go             Init: inicializa as dependências do módulo (registra rotas/tools, monta handlers, services e gateways)
│   ├── handler/                Adapta a chamada MCP ou HTTP para o service
│   ├── service/                Regra de negócio
│   └── dto/                    Input/Output da tool/endpoint
│
├── internal/gateway/<source>   Cliente de uma API externa (ex: tmdb), atrás de uma interface
├── internal/model              Structs de domínio
└── internal/utils              Helpers genéricos reutilizáveis (ex: cliente HTTP Fetch[T])
```

### Fluxo de uma chamada

```
Cliente MCP → POST /mcp → Handler → Service → Gateway → API externa
```

O `main.go` inicializa a config, cria o `mcp.Server` e registra as rotas HTTP e MCP. Cada módulo, ao ser registrado, monta sua própria cadeia de dependências: o gateway é injetado no service, que é injetado no handler, e a tool correspondente é adicionada ao servidor MCP. Novos módulos e gateways seguem exatamente esse mesmo padrão.

## Pré-requisitos

- [Go 1.25+](https://go.dev/dl/)
- Conta e credenciais na [API do TMDB](https://developer.themoviedb.org/reference/intro/getting-started)
- (Opcional) [air](https://github.com/air-verse/air) para hot reload no ambiente de dev

## Configuração

1. Copie o arquivo de exemplo de variáveis de ambiente:

   ```bash
   cp .env.example .env
   ```

2. Preencha o `.env` com suas credenciais do TMDB:

   ```env
   TMDB_API_KEY="sua_api_key"
   TMDB_READ_ACCESS_TOKEN="seu_read_access_token"
   ```

   | Variável                 | Obrigatória | Descrição                                               |
   | ------------------------ | :---------: | ------------------------------------------------------- |
   | `TMDB_API_KEY`           |     Sim     | API key do TMDB                                         |
   | `TMDB_READ_ACCESS_TOKEN` |     Sim     | Token de leitura (Bearer) usado nas requisições ao TMDB |
   | `PORT`                   |     Não     | Porta do servidor (padrão: `8080`)                      |

   > A aplicação falha na inicialização se `TMDB_API_KEY` ou `TMDB_READ_ACCESS_TOKEN` não estiverem definidos.

## Como rodar

Instale as dependências:

```bash
go mod download
```

### Desenvolvimento

```bash
make dev          # go run cmd/main.go
```

Com hot reload (requer `air` instalado):

```bash
make watch-dev    # air
```

### Build

```bash
make build        # gera o binário em bin/app
```

O servidor sobe em `http://localhost:8080` (ou na porta definida em `PORT`).

## Endpoints

| Método                | Rota           | Descrição                                              |
| --------------------- | -------------- | ------------------------------------------------------ |
| `GET`                 | `/heath-check` | Verifica se o servidor está no ar                      |
| `GET` `POST` `DELETE` | `/mcp`         | Endpoint do protocolo MCP (transporte Streamable HTTP) |

Clientes MCP devem se conectar em `/mcp` para descobrir e invocar as ferramentas disponíveis.
