FROM ubuntu:latest AS build

RUN groupadd -g 65536 appgroup
RUN useradd -M -r -u 65536 -g 65536 appuser

FROM scratch

COPY --from=build /etc/passwd /etc/passwd

COPY password-api /

USER 65536

ENTRYPOINT [ "/password-api" ]
