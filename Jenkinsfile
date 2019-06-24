pipeline {
    agent none
    stages {
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
