pipeline {
    agent none
    stages {
        stage('Protoc') {
            agent{
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps{
                sh 'go install -v github.com/gogo/protobuf/protoc-gen-gogoslick'
                sh 'go get github.com/micro/protoc-gen-micro'
                sh 'cd hallservice && make proto'
                sh 'cd movieservice && make proto'
                sh 'cd reservationservice && make proto'
                sh 'cd showservice && make proto'
                sh 'cd userservice && make proto'
            }
        }
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd hallservice && make build'
                sh 'cd movieservice && make build'
                sh 'cd reservationservice && make build'
                sh 'cd showservice && make build'
                sh 'cd userservice && make build'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...'
                sh 'cd hallservice && make test'
                sh 'cd movieservice && make test'
                sh 'cd reservationservice && make test'
                sh 'cd showservice && make test'
                sh 'cd userservice && make test'
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
                sh 'docker-build-and-push -b ${BRANCH_NAME} -s hallservice -f hallservice/Dockerfile'
                sh 'docker-build-and-push -b ${BRANCH_NAME} -s treecli -f movieservice/Dockerfile'
                sh 'docker-build-and-push -b ${BRANCH_NAME} -s reservationservice -f reservationservice/Dockerfile'
                sh 'docker-build-and-push -b ${BRANCH_NAME} -s showservice -f showservice/Dockerfile'
                sh 'docker-build-and-push -b ${BRANCH_NAME} -s userservice -f userservice/Dockerfile'
            }
        }
    }
}
