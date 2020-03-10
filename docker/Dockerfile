FROM golang:alpine as builder
RUN mkdir /app
WORKDIR /app
RUN apk add git
RUN go get github.com/guaychou/BotSemhas
RUN cp $GOPATH/bin/BotSemhas /app
RUN adduser -S -D -H -h /app appuser


FROM alpine:3.11
LABEL maintainer="Kevinchou kevin.harnanta@gmail.com"
ENV TZ="Asia/Jakarta"
# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
# Add timezone
RUN apk add tzdata
# Copy our static executable
COPY --from=builder /app/BotSemhas /app/BotSemhas
# Use an unprivileged user.
USER appuser
COPY entrypoint.sh /entrypoint.sh
# Run the binary.
ENTRYPOINT ["/entrypoint.sh"]
