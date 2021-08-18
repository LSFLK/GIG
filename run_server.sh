#!/bin/bash
sudo docker build -f Dockerfile -t ldflk/gig .
sudo docker push ldflk/gig:latest
sudo docker pull ldflk/gig
sudo docker kill local_gig
sudo docker run --rm -p 9000:9000 -d --name local_gig ldflk/gig
sudo docker logs local_gig