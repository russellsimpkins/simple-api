# Simple API for a code test.

Getting back into the flow writing APIs in golang. A special shout out to https://github.com/MartinHeinz/go-project-blueprint for
a very nice project that helped move things along much faster.

### Adding New Libraries/Dependencies
Don't forget to run this if you add new imports:
```bash
go mod vendor
```

### Building/testing/running
If you're on a mac, test and then build the container with the following:
```
GOOS=linux GOARCH=amd64 make test
GOOS=linux GOARCH=amd64 make container
```

Then you can run with
```
docker run -p 1234:1234 docker.pkg.github.com/rsimpkins/simple-api/codetest:latest
```

### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/russellsimpkins/simple-api/codetest:latest .
# make container
docker push docker.pkg.github.com/russellsimpkins/simple-api/codtest:latest
# make push
```

Pull and Run:

```bash
docker pull docker.pkg.github.com/russellsimpkins/simple-api/codtest:latest
docker run docker.pkg.github.com/russellsimpkins/simple-api/codtest:latest
```
