# test-sigma-catalog-services-api
Tool to investigate Sigma Catalog Services API

## API client
API client code was generated using https://swagger.io/tools/swagger-codegen/

[API documentation](https://localhost/CS3/ApiDocs/)
[Swagger URL](https://localhost/CS3/ApiDocs/v1/CatalogServices3.json)

Command to generate swagger go client using swagger generator.

```bash
curl -X POST -H "content-type:application/json" -d '{"swaggerUrl":"https://localhost/CS3/ApiDocs/v1/CatalogServices3.json"}' https://generator.swagger.io/api/gen/clients/go
```
### A note on the generated client stub code
All read methods seem to return empty with a 200 http on success. So the generated code does not expose the actual data..
This looks like a bug in the spec. The client code needs to be patched to prevent this.

To fix this remove this part of the client code after the client http call.

```golang
localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
localVarHttpResponse.Body.Close()
if err != nil {
  return localVarHttpResponse, err
}

if localVarHttpResponse.StatusCode >= 300 {
  newErr := GenericSwaggerError{
    body: localVarBody,
    error: localVarHttpResponse.Status,
  }

  return localVarHttpResponse, newErr
}
```

## Usage

To get the categories defined in the catalog a call using the tool could look like below, which will put the data into the categories.xml.

```bash
./test-sigma-catalog-services-api --config .test-sigma-catalog-services-api.yaml catalog categories > categories.xml
```

## Building and Running the tool using Docker

You must have [Docker](https://www.docker.com/) installed for the commands to work.

### Build it
```bash
mkdir techsimplifier
cd techsimplifier
git clone https://github.com/techsimplifier/test-sigma-catalog-services-api.git
cd test-sigma-catalog-services-api
docker build -t test-sigma-catalog-services-api:0.0.0 .
```

## Run it
```bash
docker run test-sigma-catalog-services-api:0.0.0 catalog categories --endpoint https://localhost/CS3/
```