pipeline {
  agent any

  environment {
    GHCR_IMAGE = "ghcr.io/abrahimzaman360/go-core:latest"
    GHCR_CREDENTIALS = "154a5619-ecf6-49fc-91cc-9bac632775bf" // Jenkins credentials (PAT)
  }

  stages {
    stage('Build Docker Image') {
      steps {
        sh "docker build -t $GHCR_IMAGE ."
      }
    }

    stage('Push to GHCR') {
      steps {
        withCredentials([usernamePassword(credentialsId: GHCR_CREDENTIALS, usernameVariable: 'GH_USER', passwordVariable: 'GH_PAT')]) {
          sh """
            echo "$GH_PAT" | docker login ghcr.io -u "$GH_USER" --password-stdin
            docker push $GHCR_IMAGE
            docker logout ghcr.io
          """
        }
      }
    }
  }

  post {
    success { echo "✅ Build and push successful." }
    failure { echo "❌ Build failed." }
  }
}
