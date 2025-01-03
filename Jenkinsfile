pipeline {
    agent node {
        label 'go'
    }
    parameters {
        string(name: 'BUILD_BRANCH', defaultValue: "develop/${VERSION}", description: '请输入分支号')
        string(name: 'VERSION', defaultValue: '', description: '请输入版本号')
        booleanParam(name: 'RELEASE', defaultValue: false, description: '是否是release版本')
    }

    stages {
        stage('Check') {
            stage('Check Parameters') {
                steps {
                    echo "build branch: ${params.BUILD_BRANCH}"
                    echo "version: ${params.VERSION}"
                    echo "is release: ${params.RELEASE}"
                }
            }
            stage('Check others') {
                steps {
                    echo 'OK'
                }
            }
        }
        stage("Build") {
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
        stage("Test") {
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
        stage("Deploy") {
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
    }
}