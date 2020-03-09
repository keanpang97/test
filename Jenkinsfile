pipeline {
    agent {
		docker {
		    filename 'dockerfile'
			additionalBuildArgs '-t go'
			args '--entrypoint='
		}
	}
	
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}