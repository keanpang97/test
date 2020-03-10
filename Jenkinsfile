pipeline {
	environment {
		registryCredential = 'dockerhub'
	}
  
    agent any
	
    stages {
        stage('build') {
            steps {
                sh 'docker build -f abc -t keanpang97/jenkins_cicd_go:testing .'
				sh 'docker run --name abc keanpang97/jenkins_cicd_go:testing'
				script {
					docker.withRegistry('',registryCredential) {
						sh 'docker push keanpang97/jenkins_cicd_go'
					}
				}
				sh 'docker rm abc'
				sh 'docker rmi keanpang97/jenkins_cicd_go:testing'
				sh 'docker run keanpang97/jenkins_cicd_go:testing'
            }
        }
    }
}