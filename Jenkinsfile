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
                    def map = [:]
                    map.put("BUILD_BRANCH", BUILD_BRANCH)
                    stage_check(map)
                }
            }
        }
        stage("Build") {
            parallel {
                script {
                    stage_build_core()
                }
                script {
                    stage_build_web()
                }
            }
        }
        stage("Test") {
            parallel {
                script {
                    stage_test_core()
                }
                script {
                    stage_test_web()
                }
            }
        }
        stage("Deploy") {
            parallel {
                script {
                    stage_deploy_core()
                }
                script {
                    stage_deploy_web()
                }
            }
        }
    }
}

def stage_check(Map params) {
    stage('Check Parameters') {
        echo "build branch: ${params.BUILD_BRANCH}"
    }
    stage('Check others') {
        echo 'OK'
    }
}

def stage_build_core() {
    stage("Build core") {
        echo "build core"
    }
}
def stage_build_web() {
    stage("Build web") {
        echo "build web"
    }
}

def stage_test_core() {
    stage("Test core") {
        echo "test core"
    }
}

def stage_test_web() {
    stage("Test web") {
        echo "test web"
    }
}

def stage_deploy_core() {
    stage("Deploy core") {
        echo "deploy core"
    }
}

def stage_deploy_web() {
    stage("Deploy web") {
        echo "deploy web"
    }
}