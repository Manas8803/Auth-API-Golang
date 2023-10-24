package configs

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

const connectMsg string = "---------------------------------------------------------------------------------------------\nConnected to DB\n---------------------------------------------------------------------------------------------"

func ConnectDB() error {
	ctx := context.Background()
	uri := SQLURI()

	//* Parsing the uri
	config, err := pgx.ParseConfig(uri)
	if err != nil {
		return err
	}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	fmt.Println(connectMsg)
	return nil

}
