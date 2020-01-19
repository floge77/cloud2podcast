FROM golang:alpine as builder

RUN adduser -D -g '' cloud2podcast
RUN apk update && apk add --update nodejs npm git

WORKDIR /cloud2podcast
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /go/bin/cloud2podcast
RUN cd frontend && npm install

# Build small sratch image
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/cloud2podcast /cloud2podcast
COPY --from=builder /cloud2podcast/frontend/dist/ /frontend/dist/
USER cloud2podcast
VOLUME /downloads

EXPOSE 8080

ENTRYPOINT ["/cloud2podcast"]
