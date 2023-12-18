TAGS=$1
INFILE=$(pwd)/docs/swagger/openapi.yml
OAPI_GEN_IMAGE=tomoyamachi/oapi-codegen:v0.0.1
for tag in $TAGS
do
    mkdir -p gen/api/${tag}
    rm -f gen/api/${tag}/api.go
    OUTFILE=$(pwd)/gen/api/${tag}/api.go
    docker run --platform linux/amd64 --rm -v ${INFILE}:/target.yml $OAPI_GEN_IMAGE -package ${tag} -generate "types,server" /target.yml > ${OUTFILE}
    echo "success gen ${tag}"
done