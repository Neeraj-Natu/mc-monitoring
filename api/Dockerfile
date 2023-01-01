FROM golang:1.20-rc-alpine3.17

# create a working directory inside the image
RUN mkdir /app

# copy Go modules and dependencies to image
COPY go.mod /app


# copy directory files i.e all files ending with .go
COPY main.go /app


WORKDIR /app

# compile application
RUN go build -o main .
 
# tells Docker that the container listens on specified network ports at runtime
EXPOSE 9090

# command to be used to execute when the image is used to start a container
ENTRYPOINT [ "/app/main" ]
CMD ["cluster-0", "service-0"]