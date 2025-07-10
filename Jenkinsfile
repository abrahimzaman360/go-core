pipeline {
  agent {
    kubernetes {
      label 'kaniko'
      defaultContainer 'kaniko'
    }
  }

  environment {
    GHCR_IMAGE = "ghcr.io/abrahimzaman360/go-core:latest"
    GHCR_CREDENTIALS = "154a5619-ecf6-49fc-91cc-9bac632775bf" // Jenkins secret with GitHub PAT
  }

  stages {
    stage('Checkout') {
      steps {
        git credentialsId: '154a5619-ecf6-49fc-91cc-9bac632775bf', url: 'https://github.com/abrahimzaman360/go-core.git'
      }
    }

    stage('Build & Push using Kaniko') {
      steps {
        withCredentials([usernamePassword(credentialsId: GHCR_CREDENTIALS, usernameVariable: 'GH_USER', passwordVariable: 'GH_PAT')]) {
          sh '''
            mkdir -p /kaniko/.docker
            cat <<EOF > /kaniko/.docker/config.json
            {
              "auths": {
                "ghcr.io": {
                  "username": "${GH_USER}",
                  "password": "${GH_PAT}"
                }
              }
            }
            EOF

            /kaniko/executor \
              --context `pwd` \
              --dockerfile `pwd`/Dockerfile \
              --destination=${GHCR_IMAGE} \
              --skip-tls-verify
          '''
        }
      }
    }
  }

  post {
    success { echo '✅ Image built and pushed with Kaniko!' }
    failure { echo '❌ Failed during Kaniko build or push.' }
  }
}
