pipeline {
	environment {
		registryCredential = 'dockerHub'
		tag = VersionNumber (
			versionNumberString: '${BUILD_DATE_FORMATTED, "yyMMdd"}.${BUILDS_TODAY}',
			versionPrefix: 'v'
		)
	}
  
    agent any
	
    stages {
		stage('Checkout') {
			steps {
				sh 'docker system prune --volumes -f'
				sh 'docker rmi keanpang97/jenkins_cicd_go:1.3.20'
				sh 'docker rmi keanpang97/jenkins_cicd_go:1.3.22'
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