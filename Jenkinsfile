pipeline {
  agent {
    label "generic"
  }    
  options {
    ansiColor colorMapName: 'XTerm'
    timestamps()
  }
  stages {
    stage("test") {
      steps {
        sh "echo hello"
      }
    }
  }
}
