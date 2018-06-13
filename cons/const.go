package cons

const (
	RootTemp = "./temp"

	GoDocker = `
FROM golang:alpine

ENV WDIR $GOPATH/src/temp_folder_id

WORKDIR $WDIR

RUN mkdir -p $WDIR 

ADD ./ $WDIR

RUN go build -o temp_folder_id ./

CMD ["./temp_folder_id"]

`
)
