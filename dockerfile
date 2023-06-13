FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "DASHBOARDAPP"
RUN go build -o airbnbapp

# run executable file
CMD ["./airbnbapp"]