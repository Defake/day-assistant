package env

import (
    "os"
    str "github.com/Defake/day-assistant/util/strings"
    cf "github.com/Defake/day-assistant/util/controlflow"
)

func ReadDevEnvs() {
    dat, err := os.ReadFile("./.env")
    cf.PanicOnErr(err)

    content := string(dat)
    for _, s := range str.SplitLines(content) {
        env := str.SplitBy(s, '=')
        os.Setenv(env[0], env[1])
    }

    // fmt.Print(os.Getenv("token"))
}
