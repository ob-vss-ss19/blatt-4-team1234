pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'go install -v github.com/gogo/protobuf/protoc-gen-gogoslick'
                sh 'go get github.com/micro/protoc-gen-micro'
                sh 'cd hallservice && make build'
                sh 'cd movieservice && make build'
                sh 'cd reservationservice && make build'
                sh 'cd showservice && make build'
                sh 'cd userservice && make build'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
                }
                steps {
                    sh 'cd hallservice && golangci-lint run --deadline 20m --enable-all'
                    sh 'cd movieservice && golangci-lint run --deadline 20m --enable-all'
                    sh 'cd reservationservice && golangci-lint run --deadline 20m --enable-all'
                    sh 'cd showservice && golangci-lint run --deadline 20m --enable-all'
                    sh 'cd userservice && golangci-lint run --deadline 20m --enable-all'
                }
        }
        stage('Test') {
            agent{
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd hallservice && make test'
                sh 'cd movieservice && make test'
                sh 'cd reservationservice && make test'
                sh 'cd showservice && make test'
                sh 'cd userservice && make test'
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
