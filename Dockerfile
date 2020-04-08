# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -ldflags '-w -s' -a -o hello-api

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/hello-api /app/
EXPOSE 8080
ENTRYPOINT ./hello-api