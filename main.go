package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Barrioslopezfd/gator/internal/config"
	"github.com/Barrioslopezfd/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
    db	    *database.Queries
    conf    *config.Config
}

func main(){
    conf, err:=config.Read()
    if err != nil {
        log.Fatal(err)
    }
    db, err := sql.Open("postgres", conf.DbURL)
    if err != nil {
	log.Fatal(err)
    }
    defer db.Close()
    dbQueries := database.New(db)
    toolState := &state{
	conf:	&conf,
	db:	dbQueries,
    }
    if (os.Args[1] == "login" || os.Args[1] == "register") && len(os.Args) < 3 {
	log.Fatal("Not enough Arguments")
    }
    cmd := command {
	name:	    os.Args[1],
	arguments:  os.Args[2:],
    }

    cmds := commands {
	commands:   make(map[string]func(*state, command) error),
    }

    cmds.register("login", handleLogin)
    cmds.register("register", handleRegister)
    cmds.register("reset", resetUsers)
    cmds.register("users", getUsers)
    cmds.register("agg", agg)
    cmds.register("addfeed", handleAddFeed)
    err=cmds.run(toolState, cmd) 
    if err != nil {
	log.Fatal(err)
    }


}
