FROM golang:latest


RUN mkdir /delete_all
ADD . /delete_all/
WORKDIR /delete_all
RUN go build -o main .
CMD ["/delete_all/main"]