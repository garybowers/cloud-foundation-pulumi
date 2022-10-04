FROM ubuntu:latest

ENV GOVER=1.19.1
ENV PULUMIVER=3.40.2
ENV GCLOUDVER=404.0.0
ENV CODEVER=4.7.0

RUN apt-get update -y && apt-get install -y wget curl ca-certificates sudo vim git tig make

RUN wget https://go.dev/dl/go${GOVER}.linux-amd64.tar.gz && \
	tar -xvzf go${GOVER}.linux-amd64.tar.gz -C /opt && \
	rm go${GOVER}.linux-amd64.tar.gz

RUN wget https://github.com/pulumi/pulumi/releases/download/v${PULUMIVER}/pulumi-v${PULUMIVER}-linux-x64.tar.gz && \
	tar -xvzf pulumi-v${PULUMIVER}-linux-x64.tar.gz -C /opt && \
	rm pulumi-v${PULUMIVER}-linux-x64.tar.gz 

RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-${GCLOUDVER}-linux-x86_64.tar.gz && \
 	tar -xvzf google-cloud-cli-${GCLOUDVER}-linux-x86_64.tar.gz -C /opt && \
 	rm google-cloud-cli-${GCLOUDVER}-linux-x86_64.tar.gz

RUN wget https://github.com/coder/code-server/releases/download/v${CODEVER}/code-server_${CODEVER}_amd64.deb && \
	dpkg -i code-server_${CODEVER}_amd64.deb && \
	rm code-server_${CODEVER}_amd64.deb

RUN echo 'PATH=$PATH:/opt/go/bin:/opt/pulumi:/opt/google-cloud-sdk/bin' >> /etc/skel/.bashrc && \
    mkdir -p /etc/skel/go/src /etc/skel/.config/code-server /etc/skel/.local/share/code-server && touch /etc/skel/go/src/README.txt


RUN useradd -m developer -s /bin/bash
USER developer
WORKDIR /home/developer

EXPOSE 8080
VOLUME /home/developer/go/src /home/developer/.config/gcloud
ENTRYPOINT ["/usr/bin/code-server", "--disable-telemetry", "--auth", "none", "--disable-update-check", "--bind-addr", "0.0.0.0:8080"]
