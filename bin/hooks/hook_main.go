package hook_main
import (
    "fmt"
    "hook"
)

struct HookMain {
    hook.Hook
    hook.Execer
    data map
}

func (h HookMain) Exec {
    fmt.Println("main hook goes here")
    fmt.Println("It is for %s", h.Name)
    fmt.Println("data will be %s", h.data)
}


