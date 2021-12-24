
# DEV
```
docker build --target dev --file Dockerfile.dev -t go .
docker run -it -v $(PWD):/work go sh
```

# BUILD
```
docker build . -t videos
docker run -it videos get --help
docker run -it -v "${PWD}/videos/videos.json":/videos.json videos add --help
docker run -it -v "${PWD}/videos/videos.json":/videos.json videos get --all
```