pipeline {
    agent any // Runs on the controller node
    environment {
        DOCKER_REGISTRY = 'ghcr.io'
        DOCKER_IMAGE = 'abrahimzaman360/go-core'
        DOCKER_TAG = "${env.BUILD_NUMBER}"
        KUBECONFIG = '/var/lib/jenkins/.kube/config'
    }
    stages {
        stage('Clone Private Repo') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    userRemoteConfigs: [[
                        url: 'https://github.com/abrahimzaman360/go-core',
                        credentialsId: 'ghcr-io'
                    ]]
                ])
                sh 'ls -la' // Verify cloned files
                sh 'ls -la k8s/' // Verify k8s directory
            }
        }
        stage('Build and Push Docker Image') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'ghcr-io', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                    echo "$DOCKER_PASS" | docker login $DOCKER_REGISTRY -u $DOCKER_USER --password-stdin
                    docker build -f Containerfile -t $DOCKER_REGISTRY/$DOCKER_IMAGE:$DOCKER_TAG .
                    docker push $DOCKER_REGISTRY/$DOCKER_IMAGE:$DOCKER_TAG
                    '''
                }
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'ghcr-io', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh '''
                    kubectl create secret docker-registry ghcr-secret \
                        --docker-server=$DOCKER_REGISTRY \
                        --docker-username=$DOCKER_USER \
                        --docker-password=$DOCKER_PASS \
                        --docker-email=jenkins@exoscale-vm.local \
                        --dry-run=client -o yaml | kubectl apply -f -
                    for file in k8s/*.yaml; do
                        sed "s|BUILD_NUMBER|$DOCKER_TAG|g" $file | kubectl apply -f -
                    done
                    '''
                }
            }
        }
        stage('Test Environment') {
            steps {
                sh 'git --version'
                sh 'java -version'
                sh 'docker --version'
                sh 'kubectl version --client'
            }
        }
    }
}