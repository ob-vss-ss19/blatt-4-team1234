pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd halladministration && go build main.go'
                sh 'cd movieadministration && go build main.go'
                sh 'cd showadministration && go build main.go'
                sh 'cd reservationadministration && go build main.go'
                sh 'cd useradministration && go build main.go'
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
                sh "docker-build-and-push -b ${BRANCH_NAME} -s halladministration -f halladministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movieadministration -f movieadministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s showadministration -f showadministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservationadministration -f reservationadministration.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s useradministration -f useradministration.dockerfile"
            }
        }
    }
}
