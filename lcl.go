package main

import (
	"flag"
	"log"

	"github.com/johnernaut/lcl/hashdiff"
	"github.com/johnernaut/lcl/parser"
)

func main() {
	lastTranslatedPath := flag.String("lt", "", "Path to the last_translated_version.strings file")
	masterVersionPath := flag.String("master", "", "Path to the master file used to generate locales")
	locales := flag.String("locales", "", "Comma separated string with locale values to use.  E.x. -locales=es,gb")

	flag.Parse()

	if *lastTranslatedPath == "" {
		log.Fatalln("You must provide a -lt option with a path to the last_translated_version.strings file.")
	}

	if *masterVersionPath == "" {
		log.Fatalln("You must provide a -master option with a path to the master .strings locale file.")
	}

	if *locales == "" {
		log.Fatalln("You must provide a -locales option with a comma separated list of locales you'd like to support.")
	}

	diffedValues := getChangesFromMaster(*lastTranslatedPath, *masterVersionPath)

	log.Printf("%+v\n", diffedValues)
}

func getChangesFromMaster(lastTranslated, original string) map[string]string {
	lastTranslatedVersion, err := parser.ParseLocaleFile(lastTranslated)
	if err != nil {
		log.Fatalf("Error parsing last translated file: %s", err.Error())
	}

	newVersion, err := parser.ParseLocaleFile(original)
	if err != nil {
		log.Fatalf("Error parsing original file: %s", err.Error())
	}

	return hashdiff.DiffMap(newVersion, lastTranslatedVersion)
}
