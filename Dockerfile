FROM golang:1.25.5

WORKDIR /go_projects

COPY . .

RUN go build -o composite_types_bin ./composite_types
RUN go build -o maps_bin ./composite_types/maps

COPY run_all.sh .

RUN chmod +x run_all.sh 

CMD ["./run_all.sh"]