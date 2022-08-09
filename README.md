# https://www.youtube.com/watch?v=89PdRvRUcPU


export GOPATH=$(go env GOPATH)
export kubernetes_code_generator_path=~/TEST/mykluster/code-generator-master

git clone https://tfs.de.kworld.kpmg.com/tfs/ITS_Projects/Advanced%20Analytics%20Platform/_git/hossein2

"${kubernetes_code_generator_path}"/generate-groups.sh all github.com/viveksinghggits/kluster/pkg/client github.com/viveksinghggits/kluster/pkg/apis viveksingh.dev:v1alpha1 --go-header-file "${kubernetes_code_generator_path}"/hack/boilerplate.go.txt
