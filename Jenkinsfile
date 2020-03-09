pipeline {
    agent none
	
    stages {
        stage('build') {
            steps {
				def image = docker.build("hello", "-f dockerfile")
                sh 'docker run hello:latest'
            }
        }
    }
}