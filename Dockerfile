FROM docker:latest
COPY statist /.
ENTRYPOINT ["/statist"]
CMD ["proxy"]