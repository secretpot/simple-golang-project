pipeline {
    agent {
        node {
            label 'go'
        }
    }
    parameters {
        string(name: 'VERSION', defaultValue: '', description: '请输入版本号')
        booleanParam(name: 'RELEASE', defaultValue: false, description: '是否是release版本')
    }

    stages {
        stage('Check') {
            stages {
                stage('Check GO VERSION') {
                    steps {
                        script {
                            sh 'go version'
                        }
                    }
                }
                stage('Check Project Version') {
                    steps {
                        script {
                            sh 'git describe --tags'
                        }
                    }
                }
            }
        }
        stage("Build") {
            steps {
                script {
                    sh('./scripts/build.sh')
                }
            }
        }
        stage("Test") {
            steps {
                script {
                    sh('./scripts/test.sh')
                }
            }
        }
    }
}
