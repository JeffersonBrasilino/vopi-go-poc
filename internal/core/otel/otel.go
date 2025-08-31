
// Pacote otel provê utilitários para integração com OpenTelemetry, facilitando a instrumentação,
// criação e gerenciamento de spans, eventos e status para rastreamento distribuído em aplicações Go.

// Exemplo de uso:
//     exporter, err := otel.InitOtelTracerProvider("meuServico")
package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	traceTypes "go.opentelemetry.io/otel/trace"
)

// OtelTracer define interface para criação de spans customizados.
type OtelTracer interface {

	// Start inicia um novo span.
	// Parâmetros:
	//   ctx context.Context - contexto de propagação.
	//   name string - nome do span.
	//   attributes ...otelAttribute - atributos opcionais.
	// Retorno:
	//   context.Context - contexto atualizado.
	//   OtelSpan - span criado.
	Start(ctx context.Context, name string, attributes ...otelAttribute) (context.Context, OtelSpan)
}

// OtelSpan define interface para manipulação de spans.
type OtelSpan interface {
	// End finaliza o span.
	//
	// Exemplo de uso:
	//   span.End()
	End()

	// AddEvent adiciona um evento ao span, com atributos opcionais.
	//
	// Parâmetros:
	//   eventMessage string - mensagem do evento.
	//   attributes ...otelAttribute - atributos opcionais.
	//
	// Exemplo de uso:
	//   span.AddEvent("evento", attr1)
	AddEvent(eventMessage string, attributes ...otelAttribute)

	// SetStatus define o status do span (sucesso ou erro) e uma descrição.
	//
	// Parâmetros:
	//   status SpanStatus - status do span.
	//   description string - descrição do status.
	//
	// Exemplo de uso:
	//   span.SetStatus(SpanStatusOK, "ok")
	SetStatus(status SpanStatus, description string)

	// Success marca o span como sucesso, com mensagem descritiva.
	//
	// Parâmetros:
	//   message string - mensagem de sucesso.
	//
	// Exemplo de uso:
	//   span.Success("operação concluída")
	Success(message string)

	// Error marca o span como erro, com mensagem descritiva.
	//
	// Parâmetros:
	//   err error - erro ocorrido.
	//   message string - mensagem descritiva do erro.
	//
	// Exemplo de uso:
	//   span.Error(err, "falha na operação")
	Error(err error, message string)
}

// otelAttribute representa um par chave-valor para atributos de span/evento.
type otelAttribute struct {
	key   string
	value string
}

// InitOtelTracerProvider inicializa e configura o provider do tracer do OpenTelemetry.
//
// Parâmetros:
//   serviceName string - nome do serviço para rastreamento.
//
// Retorno:
//   *otlptrace.Exporter - exportador OTLP configurado.
//   error - erro ocorrido, se houver.
//
// Exemplo de uso:
//   exporter, err := otel.InitOtelTracerProvider("meuServico")
func InitOtelTracerProvider(serviceName string) (*otlptrace.Exporter, error) {
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		return nil, fmt.Errorf("falha ao criar OTLP exporter: %w", err)
	}

	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(exporter)
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batchSpanProcessor),
	)
	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return exporter, nil
}

// makeAttributes converte uma lista de otelAttribute em atributos para o span/evento.
//
// Parâmetros:
//   attributes []otelAttribute - lista de atributos.
//
// Retorno:
//   traceTypes.SpanStartEventOption - opção com os atributos para criação do span/evento.
//
// Exemplo de uso:
//   opts := makeAttributes([]otelAttribute{core.NewOtelAttr("chave", "valor")})
func makeAttributes(attributes []otelAttribute) traceTypes.SpanStartEventOption {
	var attrs []attribute.KeyValue
	if len(attributes) > 0 {
		for _, attr := range attributes {
			attrs = append(attrs, attribute.String(attr.key, attr.value))
		}
	}
	return traceTypes.WithAttributes(attrs...)
}

// NewOtelAttr cria um novo otelAttribute com a chave e valor fornecidos.
//
// Parâmetros:
//   key string - chave do atributo.
//   value string - valor do atributo.
//
// Retorno:
//   otelAttribute - atributo criado.
//
// Exemplo de uso:
//   attr := otel.NewOtelAttr("chave", "valor")
func NewOtelAttr(key string, value string) otelAttribute {
	return otelAttribute{
		key:   key,
		value: value,
	}
}

