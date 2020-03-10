pipeline {
	environment {
		registryCredential = 'dockerHub'
	}
  
    agent any
	
    stages {
		stage('Checkout') {
			steps {
				sh 'docker rm -vf $(docker ps -a -q)'
				sh 'docker rmi -f $(docker images -a -q)'
			}
		}
		
        stage('build') {
            steps {
                sh 'docker build -f abc -t keanpang97/jenkins_cicd_go:testing .'
				sh 'docker run --name testing keanpang97/jenkins_cicd_go:testing'
				script {
					docker.withRegistry('',registryCredential) {
						sh 'docker push keanpang97/jenkins_cicd_go'
					}
				}
				sh 'docker rm testing'
				sh 'docker rmi keanpang97/jenkins_cicd_go:testing'
				sh 'docker run keanpang97/jenkins_cicd_go:testing'
            }
        }
    }
}