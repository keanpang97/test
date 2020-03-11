def Skip = readFile('file.txt')

pipeline {
	environment {
		registryCredential = 'dockerHub'
		tag = VersionNumber (
			versionNumberString: '${BUILD_DATE_FORMATTED, "yyMMdd"}.${BUILDS_TODAY}',
			versionPrefix: 'v'
		)
	}
  
    agent any
	
	/*parameters {
		booleanParam (
			name: 'Skip',
			defaultValue: false,
			description: 'Do you want to skip the Test stage?'
		)
	}*/
	
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
				sh 'docker rmi keanpang97/jenkins_cicd_go:latest'
			}
		}
		
		stage('read data') {
			steps {
				script {
					println(Skip)
				}
			}
		}
		
		stage('check condition') {
			when {
				/*expression { Skip == 'true'}*/
				
				expression { Skip == 'true'}
			}
			steps {
				sh 'docker run keanpang97/jenkins_cicd_go:latest'
            }
        }
    }
}