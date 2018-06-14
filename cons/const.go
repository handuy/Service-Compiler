package cons

const (
	RootTemp = "./temp"

	BaseBash =`
#!/usr/bin/env bash
docker build -t temp_folder_id ./temp_folder_id/
docker run temp_folder_id
`

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
