pipeline {
	environment {
		registryCredential = 'dockerHub'
		tag = VersionNumber (
			projectStartDate: '2020-01-01',
			versionNumberString: '${YEARS_SINCE_PROJECT_START}.${BUILD_MONTH}.${BUILDS_ALL_MONTH}'
		)
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
                sh 'docker build -f abc -t keanpang97/jenkins_cicd_go:$tag -t keanpang97/jenkins_cicd_go:latest .'
				sh 'docker run --name testing keanpang97/jenkins_cicd_go:$tag'
				script {
					docker.withRegistry('',registryCredential) {
						sh 'docker push keanpang97/jenkins_cicd_go'
					}
				}
				sh 'docker rm testing'
				sh 'docker rmi keanpang97/jenkins_cicd_go:$tag'
				sh 'docker run keanpang97/jenkins_cicd_go:latest'
            }
        }
    }
}