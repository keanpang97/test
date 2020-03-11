pipeline {
	environment {
		registryCredential = 'dockerHub'
	}
  
    agent any
	
    stages {
		stage('Checkout') {
			steps {
				sh 'docker system prune --volumes -f'
			}
		}
		
        stage('build') {
            steps {
                sh 'docker build -f abc -t keanpang97/jenkins_cicd_go:VersionNumber (versionNumberString: "${YEARS_SINCE_PROJECT_START}.${BUILD_MONTH}.${BUILDS_THIS_MONTH}") .'
				sh 'docker run --name testing keanpang97/jenkins_cicd_go:VersionNumber (versionNumberString: "${YEARS_SINCE_PROJECT_START}.${BUILD_MONTH}.${BUILDS_THIS_MONTH}")'
				script {
					docker.withRegistry('',registryCredential) {
						sh 'docker push keanpang97/jenkins_cicd_go'
					}
				}
				sh 'docker rm testing'
				sh 'docker rmi keanpang97/jenkins_cicd_go:VersionNumber (versionNumberString: "${YEARS_SINCE_PROJECT_START}.${BUILD_MONTH}.${BUILDS_THIS_MONTH}")'
				sh 'docker run keanpang97/jenkins_cicd_go:latest'
            }
        }
    }
}