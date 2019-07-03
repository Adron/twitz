# Twitz

[![CircleCI](https://circleci.com/gh/Adron/twitz.svg?style=svg)](https://circleci.com/gh/Adron/twitz) [![Trello Board](https://img.shields.io/badge/trello-board-purple.svg)](https://trello.com/b/1vxsOsUa/thrashing-code-projects)

[![Twitter @Adron](https://img.shields.io/twitter/follow/adron.svg?style=social&logo=twitter)](https://twitter.com/intent/follow?screen_name=adron) [![Twitter @ThrashingCode](https://img.shields.io/twitter/follow/ThrashingCode.svg?style=social&logo=twitter)](https://twitter.com/intent/follow?screen_name=ThrashingCode)

This application is about all of us lovely twitz on Twitter! A CLI tool to help introspect acounts to follow and make Twitter more useful to one's interests.

## Installation

No official install process yet besides the ole' `go get` or `go install` method. I'll write more docs in the future when I get an official version released.

## Dev Setup

1. Make sure you have Go modules enabled.
2. `go build` to build and use the executable.
3. Profit. i.e. use the CLI tool.

## Post Installation Prerequisites

Once the command is built and ready for use, you'll need to get your .twitz.yaml, twitterers.txt, and export files configured.

First open up the .twitz.yaml file and set the values per your preferred file to parse and what file you want to export out to. An example .twitz.yaml file has the following values and some standard settings.

```
file: twitterers.txt
fileExport: tweeters
fileFormat: txt
```

Next is the twitterers.txt file. Which, depending on what you've set in the .twitz.yaml file might be named whatever you've designated it. A twitterers.txt file can have a pretty wide array of text in the file, but specifically it needs at least one Twitter account somewhere in the deluge of text. An example is included below.

```
This is a sample twitterers.txt file created by @Adron.

You can add a list like this; @Adron, @angryseattle, and @jessefraz or you could go horizontal!

@Adron
@angryseattle
@pdxtst
```

The result of that file would actually spit out something just like this when issuing the `twitz parse` command.

```
Using config file:  .twitz.yaml
[@Adron @Adron @angryseattle @jessefraz @Adron @angryseattle @pdxtst]
```

There is also a more elaborate example [here](twitterers.txt).

## Commands & Usage

First and foremost you'll need a text file of Twitter accounts listed in a file called `twitterers.txt`. This file will be parsed and pull out the accounts within the file. For more information on the file and the formatting check out [twitterers file](twitterers-file.md).

`twitz` will just list out some basic documentation, commands, and other information related to the CLI itself.

`twitz config` is a command that will show what is set in the .twitz.yaml configuration file.

`twitz parse` this is the main command that'll parse out the *twitterers.txt* file and provide a list of any Twitter accounts in the file to the console.

`twitz findem` this command retrieves Twitter account information from the Twitter API.

`twitz cassie` this command will retrieve Twitter information for the accounts and insert the data into an Apache Cassandra Database.
