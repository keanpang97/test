pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile'
			additionalBuildArgs '-t go'
		}
	}
	
    stages {
        stage('build') {
            steps {
                sh 'docker run go:latest .'
            }
        }
    }
}