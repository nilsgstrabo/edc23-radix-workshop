package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	mssql "github.com/microsoft/go-mssqldb"
	"github.com/microsoft/go-mssqldb/msdsn"
	"github.com/rs/zerolog/log"
)

type moviesGetter func() ([]Movie, error)

func createMoviesGetter(sqlConnector driver.Connector) moviesGetter {
	return func() ([]Movie, error) {
		db := sql.OpenDB(sqlConnector)
		defer db.Close()

		rows, err := db.Query("select Title, Released, Rating from dbo.Movie")
		if err != nil {
			log.Error().Msg(err.Error())
			return nil, err
		}

		var movies []Movie
		for rows.Next() {
			var (
				title    string
				released time.Time
				rating   float32
			)
			if err := rows.Scan(&title, &released, &rating); err != nil {
				return nil, err
			}
			movies = append(movies, Movie{title, released, rating})
		}
		return movies, nil
	}
}

func createSqlConnector(dsn string) (driver.Connector, error) {
	cred, err := azidentity.NewWorkloadIdentityCredential(nil)
	if err != nil {
		return nil, err
	}
	cfg, err := msdsn.Parse(dsn)
	if err != nil {
		return nil, err
	}
	return mssql.NewSecurityTokenConnector(cfg, func(ctx context.Context) (string, error) {
		t, err := cred.GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{"https://database.windows.net/.default"}})
		if err != nil {
			return "", err
		}
		return t.Token, nil
	})
}
