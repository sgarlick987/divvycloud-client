# divvycloud-client
This is a go-swagger generated client for the divvycloud platform

Notable files
`divvyswagger-Swagger20.json` is the official divvy swagger doc but doesn't fully generate correctly with go-swagger    
`swagger.json` is a stripped down swagger doc with just the operations we require and few modifications for the terraform provider and we use it to generate the client  
`generate.sh` go-swagger command used to generate the client  
