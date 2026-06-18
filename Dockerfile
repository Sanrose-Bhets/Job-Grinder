From golang:1.26-alpine

WORKDIR /home/app

COPY ./app /home/app

RUN go build -o task-scheduler .

CMD ["./task-scheduler"]