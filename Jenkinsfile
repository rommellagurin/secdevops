pipeline {
  agent any
  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub')
  }
  stages {
    stage('Build') {
      steps {
        sshagent(credentials: ['your_ssh_credentials_id']) {
            sh 'ssh rommellagurin@192.168.64.2'
            sh 'ssh rommellagurin@192.168.64.2'
        }
      }
    }
  }
  post {
    always {
      sh 'docker logout'
    }
  }
}
