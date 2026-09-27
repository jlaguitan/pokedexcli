package main

var cfg = &config{
	commands: getCommands(),
}

func main() {
	startRepl(cfg)
}

