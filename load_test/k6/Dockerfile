# We propose to use the image of the creator of K6, Load Impact AB.
FROM grafana/k6:0.44.1

COPY script script/
COPY script/k6run.sh /usr/bin/k6run.sh
USER root
RUN chmod +x /usr/bin/k6run.sh
USER k6

WORKDIR script

ENTRYPOINT ["sh","-c"]

CMD ["k6run.sh"]
