FROM golang:1.21.2 as build-stage
# WORKDIR /app
# COPY go.mod .
# RUN go mod download
# COPY main.go .
# RUN go get
# RUN go build -o bin .

# #Run stage
# FROM golang:1.21.2 AS run-stage
# WORKDIR /app
# COPY --from=build-stage /app/bin .
# CMD [ "/app/bin"]


