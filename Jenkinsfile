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

def stage_check(Map params) {
    stage('Check Parameters') {
        echo("build branch: ${params.BUILD_BRANCH}")
    }
    stage('Check others') {
        echo('OK')
    }
}

def stage_build() {
    parallel(
        'Build core': {
            echo("build core")
            sleep(0.5)
            echo("core built")
        },
        'Build web': {
            echo("build web")
            sleep(0.5)
            echo("web built")
        }
    )
}

def stage_test() {
    parallel(
        'Test core': {
            echo("test core")
            sleep(0.5)
            echo("core tested")
        },
        'Test web': {
            echo("test web")
            sleep(0.5)
            echo("web tested")
        }
    )
}

def stage_deploy() {
    parallel(
        'Deploy core': {
            echo("deploy core")
            sleep(0.5)
            echo("core deployed")
        },
        'Deploy web': {
            echo("deploy web")
            sleep(0.5)
            echo("web deployed")
        }
    )
}