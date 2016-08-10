# autopublish

[![Build Status](https://snap-ci.com/meetearnest/autopublish/branch/master/build_image)](https://snap-ci.com/meetearnest/autopublish/branch/master)
[![npm version](https://badge.fury.io/js/autopublish.svg)](https://badge.fury.io/js/autopublish)

## What
autopublish is a simple command-line tool to automagically publish an npm module to an npm registry when its version has changed. Autopublish is designed to be plugged into a CI/CD pipeline. It supports custom registries and private modules (e.g. as defined in an `.npmrc` file).

There are more details in [this blog post](https://medium.com/earnest-engineering/autopublish-your-npm-modules-1f0f5ebc64c5).

## Why
It's usually a good idea for humans to decide when a change to a module's implementation warrants a version change, as well as how big that change should be. However it's not a good idea to rely on the human to remember to do the toil of testing and publishing the module. Humans tend to forget things and to not always do tasks the same way. That's why we use CI/CD tools. Autopublish allows humans to remain in control of module version changes while removing the boring parts.

## How
autopublish detects whether the current module should be published by checking to see whether the registry already contains a logically equivalent version, as defined by the `semver` module's `eq` function. If an equivalent version has not already been published then the local version is published.

If the current module has never been published to the registry then autopublish will *not* publish a new module.

## Plugging autopublish in a CI/CD pipeline
The cleanest way to integrate autopublish into your pipeline is to create a "publish module" stage in your CI tool which installs the autopublish module on-demand and then runs it directly. Here's an example publish script which installs autopublish and then runs it for the module in the current working directory:
``` publish.sh
#!/bin/bash

npm install autopublish
./node_modules/.bin/autopublish .
```

You can also install autopublish as a dev dependency, and perhaps add an `autopublish` script to your package.json. However if your goal is to encourage all your module publishing through the same mechanism - your CI tool - then keeping it out of your package.json entirely might avoid the temptation to run it by hand.

# Contributing
Pull requests welcome!

Make sure tests pass before submitting your PR. `npm test` will run them. Note that there is a test for private module support which will only be run if you pass the name for a private package that the currently logged-in user has access to. You can run that test by specifying the module name via a `EXAMPLE_PRIVATE_PACKAGE` environment variable when running the tests, e.g. `EXAMPLE_PRIVATE_PACKAGE=our_private_package npm test`. If you don't specify an example private package then that test will be skipped.

# Prior art
the ['publish'](https://www.npmjs.com/package/publish) module *almost* does what we want, but doesn't seem to support custom registries. It also has rather clunky support for npm auth (triggering an internal re-auth using credentials passed via environment variables, and only if a `TRAVIS` env var is also set).

# License
MIT
