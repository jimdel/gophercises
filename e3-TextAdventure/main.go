package main

import (
	"cyoa/jsonParser"
	"cyoa/web"
	"flag"
)

func main() {
	const Intro = "intro"

	storyPath := flag.String("story", "./gopher.json", "path to story file JSON")
	startWith := flag.String("startwith", Intro, "name of story to start the adventure with")
	cyoa := jsonParser.ReadAndParse(*storyPath)

	web.StartServer(cyoa, *startWith)

}
