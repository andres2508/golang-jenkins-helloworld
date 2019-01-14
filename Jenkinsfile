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
            when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            // Clona el repositorio -- Solo el Master
            steps {
                git repository
            }
        }
        stage('Building image') {
            when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            steps{
                script {
                    dockerImage = docker.build registry + ":latest"
                }
            }
        }
        stage('Deploy Image') {
            when {
                expression {
                    return env.GIT_BRANCH == "origin/master"
                }
            }
            steps{
                script {
                    docker.withRegistry('', registryCredential ) {
                        dockerImage.push("${env.BUILD_NUMBER}")
                        dockerImage.push("latest")
                    }
                }
            }
        }
    
        stage('Deploy SSH') {
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
                    sshCommand remote: remote, command: "./deploy_script.sh"
                }
            }
        }
    }
}