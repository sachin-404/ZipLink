# Stage 1: Build Stage
FROM golang:alpine as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build -o main .

# Stage 2: Production Stage
FROM alpine

# Create and Switch to a non-root user for running the application
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY . /app

# Copy the built executable from the build stage to the production directory
COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]
