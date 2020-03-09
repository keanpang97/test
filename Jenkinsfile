pipeline {
    agent {
		dockerfile {
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