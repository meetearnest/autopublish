#!/usr/bin/env groovy
@Library('jenkins-pipeline-library') _

pipeline {
  agent {
    label "generic"    // Our preferred agent (generic, platform, etc.)
  }
  options {
    timeout 60 // minutes
    ansiColor colorMapName: 'XTerm'
  }
  stages {
    stage("Display ENV data") {
      steps {
        printEnvSorted ()
      }
    }

    stage("Run all unit tests") {
      steps {
        prepareNpmEnv()
        prepareDockerEnv()
        sh "./go test"
      }
    }

    stage("Publish to NPM") {
      steps {
        prepareNpmEnv()
        prepareDockerEnv()
        sh "./go publish"
      }
    }
  }
}
