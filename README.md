# https://www.youtube.com/watch?v=89PdRvRUcPU


export GOPATH=$(go env GOPATH)
export PATH=$PATH:~/controller_gen:~/kuber_code_generator/code-generator

git clone https://github.com/hosseinjdkde/kluster.git

generate-groups.sh all github.com/viveksinghggits/kluster/pkg/client github.com/viveksinghggits/kluster/pkg/apis viveksingh.dev:v1alpha1 --go-header-file "${kubernetes_code_generator_path}"/hack/boilerplate.go.txt

controller-gen paths=github.com/viveksinghggits/kluster/pkg/apis/viveksingh.dev/v1alpha1  crd:crdVersions=v1 output:crd:artifacts:config=manifests


