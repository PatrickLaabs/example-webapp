FROM scratch
EXPOSE 8090
COPY example-webapp example-webapp
COPY . .
ENTRYPOINT ["/example-webapp"]

