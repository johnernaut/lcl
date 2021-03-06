# lcl

[![CircleCI](https://circleci.com/gh/johnernaut/lcl.svg?style=svg)](https://circleci.com/gh/johnernaut/lcl)
[![Maintainability](https://api.codeclimate.com/v1/badges/1e40352165fd9b281485/maintainability)](https://codeclimate.com/github/johnernaut/lcl/maintainability)

## Usage
* `go install https://github.com/johnernaut/lcl`
* `lcl -mlf=/path/to/master.strings -tlf=/path/to/target.strings -lk=es`

### How it works
**lcl** takes a "master" file, for example, your english `.strings` file, and compares its keys *and* values against a "lock file" called `last_translated_version.strings`.  Any differences between your "master" file and that file will be updated in all of the other locale files that you have, and the `last_translated_version.strings` files content will be automatically replaced.

**lcl** uses a file called the `last_translated_version.strings` file as a "lock file" to use as a diff against the master file.  The `last_translated_version.strings` should live in within a `config` folder in the same directory as your executable and should never be directly edited as it is populated with the content of all of the strings that were just translated from your master file.

Once a diff is determined between your "master" file and your `last_translated_version.strings` file, **lcl** will pull the current strings from your existing locale files and update them to match your "master" file.  Only the diff'd strings from your existing locale files and your "master" file will be replaced.  Translations are generated on these diff'd strings by hitting the Google Translate API.

### Available Options
```console
-mlf string
    Path to the master locale file used to generate locales
-tlf string
    Path to the target locale .strings file you want to translate INTO
-lk string
    Comma separated string with locale values to use.  E.x. -lk=es,gb
```

#### TODO
- [ ] dynamically pull in files from build settings / env vars 
- [ ] add a lock file for the last translated version to diff off of
- [ ] set up google translate API code to run against diffs
