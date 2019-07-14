package hook_main
import "fmt"
import "hook"

struct HookMain {
    Hook
    Execer
}

func (h HookMain) Execer {
    fmt.Println("main hook goes here")
}

func Exec:
