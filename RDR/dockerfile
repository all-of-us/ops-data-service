FROM ubuntu:latest
RUN apt-get update && apt-get -y update
RUN apt-get install -y build-essential python3.7 python3-pip python3-dev curl
RUN pip3 -q install pip --upgrade

RUN mkdir src
COPY ops_data_api.ipynb src/
COPY requirements.txt src/
ADD gcloud_key.json src/
WORKDIR src/

RUN pip3 install -r requirements.txt
RUN pip3 install jupyter

#RUN python3 module.py

#RUN rm /src/data/raw_data.csv

WORKDIR /src/

# Downloading gcloud package
RUN curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > /tmp/google-cloud-sdk.tar.gz

# Installing the package
RUN mkdir -p /usr/local/gcloud \
  && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
  && /usr/local/gcloud/google-cloud-sdk/install.sh

# Adding the package path to local
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin


# Add Tini. Tini operates as a process subreaper for jupyter. This prevents kernel crashes.
ENV TINI_VERSION v0.6.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /usr/bin/tini
RUN chmod +x /usr/bin/tini

ENTRYPOINT ["/usr/bin/tini", "--"]

CMD ["jupyter", "notebook", "ops_data_api.ipynb", "--port=8888", "--no-browser", "--ip=0.0.0.0", "--allow-root"]
