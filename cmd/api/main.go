package main

import (
	"time"

	"github.com/sz-whereable/pants/internal/auth"
	"github.com/sz-whereable/pants/internal/db"
	"github.com/sz-whereable/pants/internal/env"
	"github.com/sz-whereable/pants/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config  config
	store   store.Storage
	jwtAuth auth.Authenticator
	logger  *zap.SugaredLogger
}

type config struct {
	addr     string
	dbConfig dbConfig
	env      string
	apiURL   string
	auth     authConfig
}

type dbConfig struct {
	addr        string
	maxOpenConn int
	maxIdleConn int
	maxIdleTime int
}

type authConfig struct {
	basic basicAuthConfig
	jwt   jwtAuthConfig
}

type basicAuthConfig struct {
	username string
	password string
}

type jwtAuthConfig struct {
	secret string
	exp    time.Duration
	issuer string
}

func main() {
	config := config{
		addr: env.GetKeyString("ADDR", ":8000"),
		dbConfig: dbConfig{
			addr:        env.GetKeyString("DB_ADDR", "postgres://admin:freak@localhost:5432/pathfinder?sslmode=disable"),
			maxOpenConn: env.GetKeyInt("DB_MAX_OPEN_CONN", 25),
			maxIdleConn: env.GetKeyInt("DB_MAX_IDLE_CONN", 25),
			maxIdleTime: env.GetKeyInt("DB_MAX_IDLE_TIME", 15),
		},
		env:    env.GetKeyString("ENV", "development"),
		apiURL: env.GetKeyString("API_URL", "http://localhost:8000"),
		auth: authConfig{
			basic: basicAuthConfig{
				username: env.GetKeyString("BASIC_AUTH_USERNAME", "admin"),
				password: env.GetKeyString("BASIC_AUTH_PASSWORD", "admin"),
			},
			jwt: jwtAuthConfig{
				secret: env.GetKeyString("JWT_SECRET", "secret"),
				exp:    time.Hour * 24 * 7,
				issuer: env.GetKeyString("JWT_ISSUER", "pants"),
			},
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Connect to the database.
	db, err := db.InitDB(
		config.dbConfig.addr,
		config.dbConfig.maxOpenConn,
		config.dbConfig.maxIdleConn,
		config.dbConfig.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	logger.Info("Connected to database")

	jwtAuth := auth.NewJWTAuth(
		config.auth.jwt.secret,
		config.auth.jwt.issuer,
		config.auth.jwt.issuer,
	)

	store := store.NewStorage(db)

	app := &application{
		config:  config,
		store:   store,
		jwtAuth: jwtAuth,
		logger:  logger,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
