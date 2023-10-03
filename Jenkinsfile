pipeline {
    agent {
        node {
            label 'debian'
        }
    }

    stages {
        stage('Pull Request') {
          steps {
          git branch: 'master', url: "https://github.com/harezmii/test-api/"


          }  
        }
        stage('Image Build') {
            steps {
                sh 'docker build -t harezmi06/api:latest .'
            }
        }
        stage('Docker Login') {
            steps {
            withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId:'dockerhub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
                     sh 'echo $PASSWORD | docker login -u $USERNAME --password-stdin'
            }
                
            }
        }
        stage('Image Push') {
            steps {
              sh 'docker push harezmi06/api:latest'
            }
        }
        stage('Stack Deploy') {
            steps {
                sh 'docker stack deploy --compose-file docker-compose.yml test --with-registry-auth'
            }
        }
       
    }
    post {
        always {
            sh 'docker logout'
        }
        success {
        slackSend channel: '#jenkins-build',
                  color: 'good',
                  message: "The pipeline ${currentBuild.fullDisplayName} completed successfully."
    }
    }
}
