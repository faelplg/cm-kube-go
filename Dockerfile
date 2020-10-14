FROM scratch
COPY cm-kube-go .
ENTRYPOINT ["./cm-kube-go"]
EXPOSE 8000