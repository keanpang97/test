pipeline {
    agent none
	
    stages {
		agent none
        stage('build') {
            steps {
                sh 'docker build -f dockerfile -t go'
				sh 'docker run go:latest'
            }
        }
    }
}