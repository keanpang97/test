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
                echo 'success'
            }
        }
    }
}