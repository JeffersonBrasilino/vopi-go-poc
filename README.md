# vopi-go-poc

## Versão do Go

Este projeto utiliza Go 1.24.4

## Descrição do Projeto

Este projeto é uma prova de conceito (POC) de backend para chat, bots e gerenciamento de pessoas, utilizando Go, Clean Architecture e Package by Feature. O objetivo é demonstrar uma arquitetura escalável, testável e de fácil manutenção, com separação clara de responsabilidades.

## Abordagens Utilizadas

- **Clean Architecture**: Separação de camadas (entidades, casos de uso, interfaces, infraestrutura).
- **Package by Feature**: Cada módulo (ex: chat, person, bot) possui sua própria estrutura interna.
- **Go Proverbs & Boas Práticas**: Interfaces pequenas, nomes claros, baixo acoplamento.
- **SOLID**: Princípios aplicados na modelagem dos casos de uso e dependências.

## Estrutura de Pastas

```
internal/
  chat/
    entity/        # Entidades de domínio do chat
    usecase/       # Casos de uso (ex: criar chat)
    infra/         # Infraestrutura (repositórios, handlers HTTP)
    chat.go        # Orquestrador do módulo
  person/
    ...
cmd/
  api/
    main.go        # Ponto de entrada da aplicação
```

## Fluxo de Processamento das Requisições

1. **Entrada HTTP**: Recebida por um handler (ex: CreateChatHandler).
2. **DTO**: Payload convertido para DTO de entrada.
3. **Usecase**: Handler chama o caso de uso, que executa regras de negócio.
4. **Entidades**: Validação e criação de entidades de domínio.
5. **Repositório**: Persistência via interface de repositório (injeção de dependência).
6. **Resposta**: Usecase retorna resultado ou erro, handler formata resposta HTTP.

## Como Manter e Gerar Novos Pacotes

- Para criar um novo módulo, siga o padrão de package by feature:
  - Crie uma pasta em `internal/` com subpastas `entity`, `usecase`, `infra`.
  - Implemente entidades, casos de uso e handlers/repositórios específicos.
- Use interfaces para dependências e injete via construtores.
- Siga o padrão de nomes idiomáticos do Go.

## Fluxos de Requests (Exemplo: Criar Chat)

1. POST `/chat` com JSON de participantes e mensagens.
2. Handler converte para DTO e chama o usecase.
3. Usecase valida, cria entidades e persiste via repositório.
4. Resposta HTTP com sucesso ou erro de validação.

## Fluxograma do Processamento das Requisições

```mermaid
flowchart TD
  A[API HTTP] --> B[Handler HTTP]
  B --> C[DTO]
  C --> D[Usecase]
  D --> E[Entidades]
  D --> F[Repositório Interface]
  F --> G[Repositório Implementação]
  G --> H[Banco de Dados]
  D -- Erro --> I[Resposta de Erro]
  D -- Sucesso --> J[Resposta de Sucesso]
```

**Legenda:**

- API HTTP: Recebe requisições externas
- Handler HTTP: Converte payloads, chama casos de uso
- DTO: Objeto de transferência de dados
- Usecase: Orquestra regras de negócio
- Entidades: Modelos de domínio
- Repositório: Interface e implementação para persistência
- Banco de Dados: Persistência final
- Resposta: HTTP de sucesso ou erro

## Diagrama de Sequência da Interação dos Componentes

```mermaid
sequenceDiagram
  participant Cliente as Cliente/API
  participant Handler as Handler HTTP
  participant Usecase as Usecase
  participant Entidade as Entidade
  participant Repo as Repositório
  participant DB as Banco de Dados

  Cliente->>Handler: Envia requisição HTTP
  Handler->>Usecase: Converte e repassa DTO
  Usecase->>Entidade: Valida/cria entidades
  Usecase->>Repo: Persiste/consulta dados
  Repo->>DB: Executa operação no banco
  Repo-->>Usecase: Retorna resultado
  Usecase-->>Handler: Retorna resposta ou erro
  Handler-->>Cliente: Resposta HTTP
```

## Como Rodar

```sh
make start-dev
```

## Observabilidade e OpenTelemetry (Otel)

O projeto utiliza uma camada de abstração para OpenTelemetry, localizada em `internal/core/otel`, que facilita a instrumentação de rastreamento distribuído (tracing) de forma desacoplada e idiomática.

### Objetivo

- Permitir instrumentação de spans, eventos e status em handlers, usecases e repositórios, sem acoplamento direto à implementação do OpenTelemetry.
- Facilitar a troca, extensão ou desativação da instrumentação sem impacto no domínio.

### Como Utilizar

#### 1. Inicialização do Tracer

No início do módulo (exemplo: `internal/chat/chat.go`):

```go
var tracer = otel.InitTrace("chat-module")
```

#### 2. Uso em Handlers ou Usecases

Ao criar handlers HTTP, injete o tracer e utilize-o para criar spans:

```go
router.Group("/chat").
  POST("", func(ctx *gin.Context) { http.CreateChatHandler(ctx, m.createUseCase, tracer) })
```

No handler, crie spans para monitorar operações relevantes:

```go
ctx, span := tracer.Start(ctx, "CreateChatHandler")
defer span.End()
```

Adicione eventos, status ou erros conforme necessário:

```go
span.AddEvent("validando payload")
span.Success("chat criado com sucesso")
// ou
span.Error(err, "erro ao criar chat")
```

#### 3. Interface e Extensibilidade

A abstração define interfaces para `OtelTracer` e `OtelSpan`, permitindo testes e substituição fácil da implementação.

#### 4. Exemplo Completo

Veja o módulo `internal/chat` para um exemplo de uso integrado do tracer em rotas, handlers e casos de uso.

---

### Benefícios

- **Desacoplamento**: O domínio não depende diretamente do pacote OpenTelemetry.
- **Testabilidade**: Possível mockar spans em testes.
- **Padronização**: Uso consistente de tracing em todos os módulos.

---

## Observações

- O projeto não versiona o vendor/ por padrão.
- Para builds reprodutíveis/offline, adicione vendor/ ao repositório.
- Siga as convenções do Go para novos módulos e contribuições.
