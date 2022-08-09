# https://www.youtube.com/watch?v=89PdRvRUcPU


export GOPATH=$(go env GOPATH)
export kubernetes_code_generator_path=~/TEST/mykluster/code-generator-master

cd $GOPATH/src/Kluster

"${kubernetes_code_generator_path}"/generate-groups.sh all Kluster/pkg/client Kluster/pkg/apis mygroupname.dev:v1myopr -h "${kubernetes_code_generator_path}"/hack/boilerplate.go.txt

