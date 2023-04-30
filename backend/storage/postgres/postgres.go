package postgres

import (
	"context"
	"fmt"

	"github.com/dilmurodov/xakaton_nt/config"
	"github.com/dilmurodov/xakaton_nt/storage"

	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db         *sql.DB
	userRepo   *userRepo
	smsRepo    *smsRepo
	emailRepo  *emailRepo
	leagueRepo *leagueRepo
	seasonRepo *seasonRepo
	teamRepo   *teamRepo
	matchRepo  *matchRepo
	newsRepo   *newsRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {

	pgUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)

	db, err := sql.Open("postgres", pgUrl)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return NewStore(db), nil
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:         db,
		userRepo:   &userRepo{db: db},
		smsRepo:    &smsRepo{db: db},
		emailRepo:  &emailRepo{db: db},
		leagueRepo: &leagueRepo{db: db},
		seasonRepo: &seasonRepo{db: db},
		teamRepo:   &teamRepo{db: db},
		matchRepo:  &matchRepo{db: db},
		newsRepo:   &newsRepo{db: db},
	}
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.userRepo != nil {
		return NewUserRepo(s.db)
	}
	return s.userRepo
}

func (s *Store) Sms() storage.SmsRepoI {
	if s.smsRepo != nil {
		return NewSmsRepo(s.db)
	}
	return s.smsRepo
}

func (s *Store) Email() storage.EmailRepoI {
	if s.emailRepo != nil {
		return NewEmailRepo(s.db)
	}
	return s.emailRepo
}

func (s *Store) League() storage.LeagueRepoI {
	if s.leagueRepo != nil {
		return NewLeagueRepo(s.db)
	}
	return s.leagueRepo
}

func (s *Store) Season() storage.SeasonRepoI {
	if s.seasonRepo != nil {
		return NewSeasonRepo(s.db)
	}
	return s.seasonRepo
}

func (s *Store) Team() storage.TeamRepoI {
	if s.teamRepo != nil {
		return NewTeamRepo(s.db)
	}
	return s.teamRepo
}

func (s *Store) Match() storage.MatchRepoI {
	if s.matchRepo != nil {
		return NewMatchRepo(s.db)
	}
	return s.matchRepo
}

func (s *Store) News() storage.NewsRepoI {
	if s.newsRepo != nil {
		return NewNewsRepo(s.db)
	}
	return s.newsRepo
}