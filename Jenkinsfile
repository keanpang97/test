pipeline {
    agent none
	
    stages {
        stage('build') {
			def image = docker.build("hello", "-f dockerfile")
            steps {
				
                sh 'docker run hello:latest'
            }
        }
    }
}