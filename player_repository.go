package charlakata

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type PlayerRepository interface {
	LoadAllPlayers(ctx context.Context) ([]DBPlayer, error)
}

type postgresRepository struct {
	conn *pgx.Conn
}

type DBPlayer struct {
	ID     int
	Health int
	Level  int
}

func (p postgresRepository) LoadAllPlayers(ctx context.Context) ([]DBPlayer, error) {
	rows, err := p.conn.Query(ctx, "SELECT id, health, level FROM player ORDER BY id")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return processPlayers(rows)
}

func processPlayers(rows pgx.Rows) ([]DBPlayer, error) {
	var players []DBPlayer

	for rows.Next() {
		var p DBPlayer

		err := rows.Scan(&p.ID, &p.Health, &p.Level)
		if err != nil {
			return nil, fmt.Errorf("failed to scan player row: %w", err)
		}

		players = append(players, p)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error during row iteration: %w", rows.Err())
	}

	return players, nil
}

func NewPostgresPlayerRepository(ctx context.Context, connStr string) (PlayerRepository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		slog.Error("unable to connect to database", slog.String("error", err.Error()))
		return nil, err
	}
	return &postgresRepository{
		conn: conn,
	}, nil
}
