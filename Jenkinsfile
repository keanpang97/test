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
                sh 'docker run go:latest'
            }
        }
    }
}