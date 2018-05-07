node {
    try{        
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/") {
            withEnv(["GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"]) {
                env.PATH="${GOPATH}/bin:$PATH"
                environment {
                  SRC_PATH = 'src/github.com/Arvinderpal/jenkins-test-1'
                }
                stage('Checkout'){
                    echo 'Checking out SCM'
                    dir('src/github.com/Arvinderpal/jenkins-test-1') {
                      checkout scm
                    }
                    //checkout scm
                }
                
                stage('Pre Test'){
                    echo 'Pulling Dependencies'
            
                    sh 'go version'
                    //sh 'go get -u github.com/golang/lint/golint'
                    //sh 'go get github.com/tebeka/go2xunit'
                }
        
                stage('Test'){
                    dir('src/github.com/Arvinderpal/jenkins-test-1') {
                      
                      // List all our project files
                      // Push our project files relative to ./src
                      // FIXME (awander): this does not work with jenkins. The
                      // output format is not same as if run on laptop. 
                      sh 'go list ./... | grep -v /vendor/
                      sh 'go list ./... | grep -v /vendor/ > projectPaths'
                      
                      //Print them with 'awk '$0="./src/"$0' projectPaths' in 
                      // order to get full relative path to $GOPATH
                      def paths = sh returnStdout: true, script: """awk '\$0="./src/"\$0' projectPaths"""
                    
                      //echo 'Vetting'
                      sh """go tool vet ${paths}"""

                      //echo 'Linting'
                      //sh """golint ${paths}"""
                    
                      echo 'Testing'
                      //sh """go test -race -cover ${paths}"""
                      sh """go test ./..."""
                    }
                }
            
                stage('Build'){
                  dir('src/github.com/Arvinderpal/jenkins-test-1') {
                    echo 'Building Executable'
                
                    //Produced binary is $GOPATH/<name>
                    sh """go build -ldflags '-s'"""
                  }
                }
            }
        }
    }catch (e) {
        // If there was an exception thrown, the build failed
        currentBuild.result = "FAILED"
        echo 'FAILED'
    } finally {
       
        def bs = currentBuild.result ?: 'SUCCESSFUL'
        if(bs == 'SUCCESSFUL'){
            echo 'SUCCESSFUL'
        }
    }
}

