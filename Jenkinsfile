pipeline {
    agent none
	
    stages {
        stage('build') {
            steps {
                sh 'docker build -f dockerfile -t go'
				sh 'docker run go:latest'
            }
        }
    }
}