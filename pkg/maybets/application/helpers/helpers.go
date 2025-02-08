package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/enums"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/domain"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"golang.org/x/exp/rand"
)

func CheckIfCurrentDBIsLocal() bool {
	environment := os.Getenv("ENVIRONMENT")

	return environment == enums.Local.String()
}

// ConvertPortToInt a helper function to conver string port number to int
func ConvertPortToInt() (int, error) {
	result, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return 0, err
	}

	return result, nil
}

// GenerateTestData is a helper method used to generate test betting data
func GenerateTestData(filename string, numRecords int) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	fixedUserID := uuid.NewString()

	for i := 0; i < numRecords; i++ {
		var (
			userID  string
			outcome enums.Outcome
		)

		if i < 4000 {
			// fix 4000 users having won for testing purposes
			userID = fixedUserID
			outcome = enums.Win
		} else {
			userID = uuid.NewString()
			outcome = []enums.Outcome{enums.Win, enums.Lose}[rand.Intn(2)]
		}

		bet := &domain.Bet{
			BetID:     uuid.NewString(),
			UserID:    userID,
			Amount:    rand.Float64() * 100,
			Odds:      rand.Float64() * 10,
			Outcome:   outcome,
			Timestamp: time.Now(),
		}

		if err := encoder.Encode(bet); err != nil {
			return fmt.Errorf("failed to encode JSON: %w", err)
		}
	}

	return nil
}

// LoadBetsFromFile is a helper function to load a set of bet records from a file into a given data class
func LoadBetsFromFile(filename string) ([]*domain.Bet, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	var bets []*domain.Bet

	decoder := json.NewDecoder(file)

	for {
		var bet domain.Bet
		if err := decoder.Decode(&bet); err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %w", err)
		}

		bets = append(bets, &bet)
	}

	return bets, nil
}

// SetupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func SetupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}

		shutdownFuncs = nil

		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up propagator.
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)

	// Set up trace provider.
	tracerProvider, err := newTraceProvider()
	if err != nil {
		handleErr(err)
		return
	}

	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)

	return
}

func newTraceProvider() (*trace.TracerProvider, error) {
	serviceName := fmt.Sprintf("maybets-%v", os.Getenv("ENVIRONMENT"))

	traceExporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(os.Getenv("JAEGER_URL")),
	)
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			traceExporter,
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
			trace.WithBatchTimeout(trace.DefaultScheduleDelay*time.Millisecond),
			trace.WithMaxExportBatchSize(trace.DefaultMaxExportBatchSize),
		),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(serviceName),
			),
		),
	)

	_ = traceProvider.Tracer("maybets-analytics-svc")

	return traceProvider, nil
}
