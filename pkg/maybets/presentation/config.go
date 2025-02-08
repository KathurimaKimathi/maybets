package presentation

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/KathurimaKimathi/maybets/pkg/maybets/application/helpers"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/cache"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/infrastructure/database/postgres/gorm"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/presentation/rest"
	"github.com/KathurimaKimathi/maybets/pkg/maybets/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var allowedOriginPatterns = []string{
	`^(https?://)?(.+)-?.ingeniumct\.com$`,
}

// Compile the regex patterns into a slice of *regexp.Regexp
func compilePatterns(patterns []string) []*regexp.Regexp {
	var compiledPatterns []*regexp.Regexp

	for _, pattern := range patterns {
		compiledPattern := regexp.MustCompile(pattern)
		compiledPatterns = append(compiledPatterns, compiledPattern)
	}

	return compiledPatterns
}

// Check if the origin is allowed by matching it against the compiled regex patterns
func isAllowedOrigin(origin string, compiledPatterns []*regexp.Regexp) bool {
	for _, pattern := range compiledPatterns {
		if pattern.MatchString(origin) {
			return true
		}
	}

	return false
}

// StartServer sets up gin
func StartServer(ctx context.Context, port int) error {
	otelShutdown, err := helpers.SetupOTelSDK(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = errors.Join(err, otelShutdown(ctx))
	}()

	maybetUsecases, err := ConfigureStartUpDependencies()
	if err != nil {
		return err
	}

	r := gin.Default()

	SetupRoutes(r, *maybetUsecases)

	addr := fmt.Sprintf(":%d", port)

	if err := r.Run(addr); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// ConfigureStartUpDependencies is used to initialize all the constructors required for the application to start
func ConfigureStartUpDependencies() (*usecases.UsecaseMayBets, error) {
	db, err := gorm.NewDBInstance()
	if err != nil {
		return nil, err
	}

	redisURL := os.Getenv("REDIS_URL")

	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	c := redis.NewClient(opt)
	cacheSvc := cache.NewStoreCache(c)

	_, err = c.Ping().Result()
	if err != nil {
		return nil, err
	}

	err = postgres.RunMigrations()
	if err != nil {
		return nil, err
	}

	database := postgres.NewMaybetsDB(cacheSvc, db, db)

	infra := infrastructure.NewInfrastructureInteractor(cacheSvc, database)

	maybetUsecases, err := usecases.NewUsecaseMayBetsImpl(*infra)
	if err != nil {
		return nil, fmt.Errorf("can't instantiate service : %w", err)
	}

	return maybetUsecases, nil
}

func SetupRoutes(r *gin.Engine, usecases usecases.UsecaseMayBets) {
	compiledPatterns := compilePatterns(allowedOriginPatterns)

	r.Use(cors.New(cors.Config{
		AllowWildcard: true,
		AllowMethods:  []string{http.MethodPut, http.MethodPatch, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{
			"Accept",
			"Accept-Charset",
			"Accept-Language",
			"Accept-Encoding",
			"Origin",
			"Host",
			"User-Agent",
			"Content-Length",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length", "Link"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Specific localhost origins
			if origin == "http://localhost:5000" || origin == "http://localhost:4200" ||
				origin == "http://localhost:8878" || origin == "http://localhost:8080" {
				return true
			}

			allowed := isAllowedOrigin(origin, compiledPatterns)
			return allowed
		},
		MaxAge:          12 * time.Hour,
		AllowWebSockets: true,
	}))

	r.Use(otelgin.Middleware(fmt.Sprintf("maybets-%v", os.Getenv("ENVIRONMENT"))))

	handlers := rest.NewHandlersInterfaces(&usecases)

	// version our APIS
	apiV1RoutesGroup := r.Group("/api/v1")

	// group analytics apis
	analytics := apiV1RoutesGroup.Group("/analytics")
	analytics.GET("/total_bets", handlers.GetUserTotalBets)
	analytics.GET("/total_winnings", handlers.GetUserTotalWinnings)
	analytics.GET("/top_users", handlers.GetTopFiveUsers)
	analytics.GET("/anomalies", handlers.GetAllAnomalousUsers)
}
