package config

type Config struct {
	Version bool     `docopt:"--version"`
	Right   bool     `docopt:"--right"`
	Icon    string   `docopt:"--icon"`
	Width   int      `docopt:"--width"`
	Pad     string   `docopt:"--pad"`
	Words   []string `docopt:"<word>"`
	Name    string   `docopt:"--name"`
}
