docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i "/local/api/stakes.yml" \
    -g go-server \
    -o /local/build/go