FROM scratch
COPY bin/fulfillment-tester-linux /
COPY ui /ui
VOLUME ["/tmp"]
ENTRYPOINT ["/fulfillment-tester-linux"]