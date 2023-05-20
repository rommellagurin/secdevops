pipeline {
  agent {label 'clustermgr'} 
  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub')
  }
  stages {
    stage('Build') {
      steps {
        sh 'docker build -t rommellagurin1022/jenkins-docker-hub:nginx-devops-vlatest$BUILD_NUMBER'
      }
    }
    stage('Login') {
      steps {
        sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
      }
    }
    stage('Push') {
      steps {
        sh 'docker push rommellagurin1022/jenkins-docker-hub:nginx-devops-vlatest$BUILD_NUMBER'
      }
    }
    
