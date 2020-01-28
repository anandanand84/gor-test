FROM scratch
# Copy our static executable.
COPY ./gor /go/bin/gor
WORKDIR /go/bin/
# Run the hello binary.
ENTRYPOINT ["/go/bin/gor"]