## Tools to validate and convert Kubernetes manifest into JSON

### Getting started
1. ensure that you have installed go with version 1.24.3. use `go version` to check if you already have go installed.
2. run `go mod tidy` to download depedency of the project.
3. ensure that you already put the kubernetes manifest that you want to validate and convert in the `input` folder.
4. you can use `go run main.go` to run in development mode.
5. or you can build using `go build` and run using your shell `sh ./seek-assessment.exe`.
6. check log in your shell when running this solution to check the status of the kubernetes manifest. any manifest that doesnt pass the validation will throw an error together with the line that does not pass validation.


### Limitation
- currently this script can validate required property on the manifest but some property validation might be missing. we can only check for required/standard property for now.
- this script couldnt validate the manifest based on apiVersion and Kind yet. further modification can be made to make this script have the capability to do that.

### Future Development Idea
- turn this script into CLI command that can enable selection of file instead of processing every file in input folder
- add method to validate the manifest based on kind, Deployment, pods, Service volume and etc.

