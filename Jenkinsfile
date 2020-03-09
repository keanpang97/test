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
                sh 'docker exec go:latest'
            }
        }
    }
}