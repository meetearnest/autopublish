# autopublish

## What
autopublish is a simple command-line tool to automagically publish an npm module to an npm registry when its version has changed. Autopublish is designed to be plugged into a CI/CD pipeline. It supports custom registries and private modules (e.g. as defined in an `.npmrc` file).

## Why
It's usually a good idea for humans to decide when a change to a module's implementation warrants a version change, as well as how big that change should be. However it's not a good idea to rely on the human to remember to do the toil of testing and publishing the module. Humans tend to forget things and to not always do tasks the same way. That's why we use CI/CD tools. Autopublish allows humans to remain in control of module version changes while removing the boring parts.

## How
autopublish detects whether the current module should be published by checking to see whether the registry already contains a logically equivalent version, as defined by the `semver` module's `eq` function. If an equivalent version has not already been published then the local version is published.

# Contributing

Pull requests welcome!

Make sure tests pass before submitting your PR. `npm test` will run them. Note that there is a test for private module support which will only be run if you pass the name for a private package that the currently logged-in user has access to. You can run that test by specifying the module name via a `EXAMPLE_PRIVATE_PACKAGE` environment variable when running the tests, e.g. `EXAMPLE_PRIVATE_PACKAGE=our_private_package npm test`. If you don't specify an example private package then that test will be skipped.

# Prior art
the ['publish'](https://www.npmjs.com/package/publish) module *almost* does what we want, but doesn't seem to support custom registries. It also has rather clunky support for npm auth (triggering an internal re-auth using credentials passed via environment variables, and only if a `TRAVIS` env var is also set).

# License
MIT

# TODO
- [ ] setup CI/CD pipeline (oh, how meta!)
