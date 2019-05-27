pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd hallAdministration && go build main.go'
                sh 'cd movieAdministration && go build main.go'
                sh 'cd showAdministration && go build main.go'
                sh 'cd reservationAdministration && go build main.go'
                sh 'cd userAdministration && go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...TODO'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'golangci-lint run --deadline 20m --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s hallAdministration -f hallAdministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movieAdministration -f movieAdministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s showAdministration -f showAdministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservationAdministration -f reservationAdministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s userAdministration -f userAdministration.dockerfile"
            }
        }
    }
}
