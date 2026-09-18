package main

import (
    "log"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/plugins/migratecmd"
    "github.com/pocketbase/pocketbase/tools/osutils"


    _ "github.com/Damillora/griseo/migrations"
)

func main() {
    app := pocketbase.New()

    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        // enable auto creation of migration files when making collection changes in the Dashboard
        // (the IsProbablyGoRun check is to enable it only during development)
        Automigrate: osutils.IsProbablyGoRun(),
    })

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
