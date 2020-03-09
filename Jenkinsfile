pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile'
			args '-t go'
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