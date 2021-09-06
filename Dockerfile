FROM scratch
COPY odin /

# Exposes port 3001 because our program listens on that port
EXPOSE 3001

ENTRYPOINT ["/odin"]