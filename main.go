package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/johnernaut/lcl/hashdiff"
	"github.com/johnernaut/lcl/parser"
)

var config = new(configuration)

type configuration struct {
	googleAPIKey       string
	lastTranslatedFile string
	masterLocaleFile   string
	targetLocaleFile   string
	targetLocaleKey    string
}

func main() {
	loadConfig()
	diffedValues := getChangesFromMaster(config.lastTranslatedFile, config.masterLocaleFile)

	log.Printf("%+v\n", diffedValues)
}

func loadConfig() {
	lastTranslatedPath, err := filepath.Abs("config/last_translated_version.strings")
	if err != nil {
		log.Fatalf("Unable to find the last_translated_verstion.strings file.")
	}

	masterLocaleFile := flag.String("mlf", "", "Path to the master locale file used to generate locales")
	targetLocaleFile := flag.String("tlf", "", "Path to the target locale .strings file you want to translate INTO")
	localeKey := flag.String("lk", "", "Comma separated string with locale values to use.  E.x. -lk=es,gb")

	flag.Parse()

	if *masterLocaleFile == "" {
		log.Fatalln("You must provide a -mlf option with a path to the master .strings locale file.")
	}

	if *targetLocaleFile == "" {
		log.Fatalln("You must provide a -tlf option with a path to the target .strings file.")
	}

	if *localeKey == "" {
		log.Fatalln("You must provide a -lk option with a comma separated list of locales you'd like to support.")
	}

	config = &configuration{
		lastTranslatedFile: lastTranslatedPath,
		masterLocaleFile:   *masterLocaleFile,
		targetLocaleFile:   *targetLocaleFile,
		targetLocaleKey:    *localeKey,
	}
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
