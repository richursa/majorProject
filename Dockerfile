FROM ubuntu
RUN apt update
RUN apt install gnupg -y
RUN apt install wget -y 
RUN mkdir /data
RUN mkdir /data/db
RUN wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | apt-key add -
RUN echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.2.list
RUN apt update
RUN apt install -y mongodb-org
RUN apt install golang -y
RUN apt install git -y
RUN apt install nano -y
RUN apt install curl -y
RUN go get -u github.com/gorilla/mux
RUN go get -u go.mongodb.org/mongo-driver/bson
RUN mkdir /app
ADD . /app
ENTRYPOINT mongod
