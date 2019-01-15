pipeline {
    environment {
        registry = "andres2508/docker-helloworld"
        repository = 'https://jaime2508@bitbucket.org/jaime2508/jenkins-helloworld.git'
        registryCredential = 'dockerhub'
        dockerImage = ''

        // SSH Credentials
        host = '192.168.0.23'
        user = 'sdelosrios'
        password = '4kidz'
    }
    agent any
    stages {
        stage('Cloning Git') {
            // Clona el repositorio -- En el que se encuentre el SCM
            steps {
                checkout scm
            }
        }
        stage('Building image') {
            steps{
                script {
                    dockerImage = docker.build registry + ":latest"
                }
            }
        }
        stage('Push Docker Image') {
            steps{
                script {
                    docker.withRegistry('', registryCredential ) {
                        dockerImage.push("${env.BUILD_NUMBER}")
                        dockerImage.push("latest")
                    }
                }
            }
        }
    
        stage('Deploy Production') {
            when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            steps{
                script {
                    def remote = [:]
                    remote.name = 'jenkins'
                    remote.host = host
                    remote.user = user
                    remote.password = password
                    remote.allowAnyHosts = true
                    sshCommand remote: remote, command: "./deploy_script.sh deploy-master-jenkins 3535"
                }
            }
        }

        stage('Deploy Development') {
            when {
                expression {
                    return env.GIT_BRANCH == "origin/dev"
                }
            }
            steps{
                script {
                    def remote = [:]
                    remote.name = 'jenkins'
                    remote.host = host
                    remote.user = user
                    remote.password = password
                    remote.allowAnyHosts = true
                    sshCommand remote: remote, command: "./deploy_script.sh deploy-dev-jenkins 3536"
                }
            }
        }
    }
}