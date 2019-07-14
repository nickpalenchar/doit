package hook

type Hook struct {
    Name string
    Priority int
    WillAccept []string
    Required bool
}

type Execer interface {
    Main()
}
type PreHook interface {
    PreHook()
}
type PostHook interface {
    PostHook()
}



