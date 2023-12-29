INFILE=$(pwd)/docs/api/openapi.yml
OAPI_GEN_IMAGE=tomoyamachi/oapi-codegen:v0.0.1
OUTFILE=$(pwd)/gen/api/api.go

docker run --platform linux/amd64 --rm -v ${INFILE}:/target.yml $OAPI_GEN_IMAGE -package openapi -generate "types,server" /target.yml > ${OUTFILE}
echo "success gen"