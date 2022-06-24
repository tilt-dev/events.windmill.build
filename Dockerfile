FROM alpine
WORKDIR /app
ADD build build
ENTRYPOINT ./build/events-windmill-build
