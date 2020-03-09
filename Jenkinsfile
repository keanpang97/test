pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile'
			label 'go'
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