pipeline {
    agent {
        node {
            label 'go'
        }
    }
    parameters {
        string(name: 'BUILD_BRANCH', defaultValue: "develop/${VERSION}", description: '请输入分支号')
        string(name: 'VERSION', defaultValue: '', description: '请输入版本号')
        booleanParam(name: 'RELEASE', defaultValue: false, description: '是否是release版本')
    }

    stages {
        stage('Check') {
            steps {
                script {
                    stage_check()
                }
            }
        }
        stage("Build") {
            steps {
                script {
                    stage_build()
                }
            }
        }
        stage("Test") {
            steps {
                script {
                    stage_test()
                }
            }
        }
        stage("Deploy") {
            steps {
                script {
                    stage_deploy()
                }
            }
        }
    }
}

def stage_check() {
    stage('Check Parameters') {
        echo "build branch: ${params.BUILD_BRANCH}"
    }
    stage('Check others') {
        echo 'OK'
    }
}

def stage_build() {
    parallel {
        stage("Build core") {
            steps {
                echo "build core"
            }
        }
        stage("Build web") {
            steps {
                echo "build web"
            }
        }
        stage("Build mobile") {
            steps {
                echo "build mobile"
            }
        }
    }
}

def stage_test() {
    parallel {
        stage("Test core") {
            steps {
                echo "test core"
            }
        }
        stage("Test web") {
            steps {
                echo "test web"
            }
        }
        stage("Test mobile") {
            steps {
                echo "test mobile"
            }
        }
    }
}

def stage_deploy() {
    parallel {
        stage("Deploy core") {
            steps {
                echo "deploy core"
            }
        }
        stage("Deploy web") {
            steps {
                echo "deploy web"
            }
        }
        stage("Deploy mobile") {
            steps {
                echo "deploy mobile"
            }
        }
    }
}