# Twits

[![Build Status](https://travis-ci.org/ThrashingCode/Twits.svg?branch=master)](https://travis-ci.org/ThrashingCode/Twits)

This application is about all of us lovely twits on Twitter! A CLI tool to help introspect acounts to follow and make Twitter more useful to one's interests.

## Installation

No official install process yet besides the ole' `go get` or `go install` method. I'll write more docs in the future when I get an official version released.

## Dev Setup

`dep ensure` to get the dependencies.

`go build` to build and use the executable.

etc.

## Post Installation Prereqs

Once the command is built and ready for use, you'll need to get your .twits.yaml, twitterers.txt, and export files configured.

First open up the .twits.yaml file and set the values per your preferred file to parse and what file you want to export out to. An example .twits.yaml file has the following values and some standard settings.

```
file: twitterers.txt
export: tweeters.txt
```

Next is the twitterers.txt file. Which, depending on what you've set in the .twits.yaml file might be named whatever you've designated it. A twitterers.txt file can have a pretty wide array of text in the file, but specifically it needs at least one Twitter account somewhere in the deluge of text. An example is included below.

```
This is a sample twitterers.txt file created by @Adron.

You can add a list like this; @Adron, @angryseattle, and @jessefraz or you could go horizontal!

@Adron
@angryseattle
@pdxtst
```

The result of that file would actually spit out something just like this when issuing the `twits parse` command.

```
Using config file:  .twits.yaml
[@Adron @Adron @angryseattle @jessefraz @Adron @angryseattle @pdxtst]
```

There is also a more elaborate example [here](twitterers.txt).

## Commands & Usage

First and foremost you'll need a text file of Twitter accounts listed in a file called `twitterers.txt`. This file will be parsed and pull out the accounts within the file. For more information on the file and the formatting check out [twitterers file](twitterers-file.md).

`twits` will just list out some basic documentation, commands, and other information related to the CLI itself.

`twits config` is a command that will show what is set in the .twits.yaml configuration file.

`twits parse` this is the main command that'll parse out the *twitterers.txt* file and provide a list of any Twitter accounts in the file to the console.

`twits parse`

`twits xyz` other commands will be documented shortly...  
