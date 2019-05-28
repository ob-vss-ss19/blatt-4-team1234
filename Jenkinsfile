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
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'echo run tests...NO TESTS EXISTING'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'skipping lint to test docker config'
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
