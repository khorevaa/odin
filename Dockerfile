FROM scratch
COPY odin /

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

ENTRYPOINT ["/odin"]