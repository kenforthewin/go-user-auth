FROM golang

RUN go get github.com/elwinar/rambler
RUN go install github.com/elwinar/rambler

COPY ./db ./

ENTRYPOINT [ "rambler" ]
CMD [ "apply" ]
