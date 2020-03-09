pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile'
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