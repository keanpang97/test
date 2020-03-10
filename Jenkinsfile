pipeline {
    agent any
	
    stages {
        stage('build') {
            steps {
                sh 'docker build -f abc -t go .'
				sh 'docker run go:latest'
            }
        }
    }
}