# seninhedefin

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

## API setup with Docker
You need the Docker infrastructure to run the API part of the project.

### Build and create docker image
```
docker build -t senin_hedefin_api:latest .
```

### Run docker image and go on
```
docker run --rm -it -p 8080:8080 -v $(pwd):/go/src/api senin_hedefin_api
```
