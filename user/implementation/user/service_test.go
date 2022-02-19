package user

import (
	"context"
	"testing"

	"github.com/civitops/Ecommercify/user/pkg/config"
	logger "github.com/civitops/Ecommercify/user/pkg/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func returnConf(path string, t *testing.T) *config.UserConfig {
	// fetchs configuration
	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Logf("failed to load config: %s", err.Error())
		t.Fail()
	}
	return cfg
}
func returnPgConn(t *testing.T) *gorm.DB {
	conn, err := gorm.Open(
		postgres.Open("postgres://postgres:mypass@localhost:5432/ecommercify"), &gorm.Config{
			SkipDefaultTransaction: true,
		})
	if err != nil {
		t.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}

func setTestBed(t *testing.T) Service {
	cfg := returnConf("../../", t)
	// setting up logger with the config
	zapLogger := logger.NewLogger(cfg)

	// tracer

	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		t.Fatalf("jaeger exported creation failed: %s", err.Error())
	}

	// Defining resource attributes
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("user-svc-test"),
		semconv.ServiceVersionKey.String("1.0.0"),
		attribute.Int64("ID", 1),
	)

	// Create the trace provider with the exporter and resources
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)

	// b3 propagator initilizes
	tracer := provider.Tracer("userSvc")

	pgConn := returnPgConn(t)
	pgConn.AutoMigrate(Entity{})

	// declare service here
	pgRepo := NewPostgresRepo(zapLogger, pgConn, tracer)

	userSvc := NewUserService(zapLogger, *cfg, pgRepo, tracer)
	t.Cleanup(func() {
		pgConn.Delete(&Entity{})
	})
	return userSvc
}

func TestUserCreate(t *testing.T) {
	testCases := []struct {
		desc string
		ctx  context.Context
		e    Entity
	}{
		{
			desc: "Creation should PASS ",
			ctx:  context.TODO(),
			e: Entity{
				ID:      5,
				Name:    "Suman Das",
				PhoneNo: "9856452136",
				Email:   "suman@gmail.com",
				Role:    4,
			},
		},
		{
			desc: "Creation should Fail ",
			ctx:  context.TODO(),
			e: Entity{
				ID:      5,
				Name:    "Suman Das",
				PhoneNo: "9856452136",
				Email:   "suman@gmail.com",
				Role:    3,
			},
		}, {
			desc: "Creation should Fail No Name Provided",
			ctx:  context.TODO(),
			e: Entity{
				ID:      5,
				Name:    "",
				PhoneNo: "9856452536",
				Email:   "suman@yahoo.com",
				Role:    2,
			},
		},
	}
	for i, tC := range testCases {
		user := setTestBed(t)
		id, err := user.Create(tC.ctx, tC.e)
		t.Run(tC.desc, func(t *testing.T) {
			if i == 1 {
				if id != tC.e.ID && err != nil {
					t.FailNow()
				}
			} else if i == 2 {
				if id == tC.e.ID && err == nil {
					t.FailNow()
				}
			} else if i == 3 {
				if id == tC.e.ID && err == nil {
					t.FailNow()
				}
			}

		})
	}
}
