# lcl

[![CircleCI](https://circleci.com/gh/johnernaut/lcl.svg?style=svg)](https://circleci.com/gh/johnernaut/lcl)

## Usage
* `go install https://github.com/johnernaut/lcl`
* `lcl -lt=/path/to/last_translated_version.txt -master=/path/to/test_en.txt -locales=es`

### Available Options
```console
  -locales string
        Comma separated string with locale values to use.  E.x. -locales=es,gb
  -lt string
        Path to the last_translated_version.strings file
  -master string
        Path to the master file used to generate locales
```

#### TODO
- [ ] dynamically pull in files from build settings / env vars 
- [ ] add a lock file for the last translated version to diff off of
- [ ] set up google translate API code to run against diffs
