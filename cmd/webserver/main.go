package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/matheuss0xf/gobid/internal/api"
	"github.com/matheuss0xf/gobid/internal/services"
	"net/http"
	"os"
	"time"
)

func main() {
	gob.Register(uuid.UUID{})
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBID_DB_USER"),
		os.Getenv("GOBID_DB_PASSWORD"),
		os.Getenv("GOBID_DB_HOST"),
		os.Getenv("GOBID_DB_PORT"),
		os.Getenv("GOBID_DB_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	s := scs.New()
	s.Store = pgxstore.New(pool)
	s.Lifetime = 24 * time.Hour
	s.Cookie.HttpOnly = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	api := api.Api{
		Router:         chi.NewMux(),
		UserService:    services.NewUserService(pool),
		ProductService: services.NewProductService(pool),
		Sessions:       s,
		BidsService:    services.NewBidsService(pool),
		WsUpgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		AuctionLobby: services.AuctionLobby{
			Rooms: make(map[string]*services.AuctionRoom),
		},
	}

	api.BindRoutes()

	fmt.Println("Staring Server on port :8080")
	if err := http.ListenAndServe("localhost:8080", api.Router); err != nil {
		panic(err)
	}
}
