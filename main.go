package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/05blue04/Blog-Aggregator/internal/config"
	"github.com/05blue04/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)

	if err != nil {
		log.Fatalf("error connecting to database %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollows))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", handlerUnfollow)

	if len(os.Args) < 2 {
		log.Fatal("Usage: <command> [args..]")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
