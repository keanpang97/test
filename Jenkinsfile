pipeline {
    agent {
		dockerfile {
		    filename 'dockerfile.build'
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