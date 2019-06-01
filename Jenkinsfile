#!/usr/bin/env groovy

pipeline {
    agent any
    environment {
        SERVICE_NAME = 'go-http-sample'
    }
    stages {
        stage('Build image') {
            steps {
                script {
                    def dockerImage = docker.build("${SERVICE_NAME}:latest")
                }
            }
        }
    }
}
