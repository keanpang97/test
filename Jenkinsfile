pipeline {
    agent none
	
    stages {
        stage('build') {
            steps {
				sh 'docker build -t go - f dockerfile'
                sh 'docker run go:latest'
            }
        }
    }
}