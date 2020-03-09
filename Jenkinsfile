pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile'
			tag 'go'
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