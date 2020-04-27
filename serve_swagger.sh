#!/bin/bash
# docker pull swaggerapi/swagger-ui
sudo docker run -p 8081:8080 -e SWAGGER_JSON=/swagger.yaml -v /<absolute path>/swagger.yaml:/swagger.yaml swaggerapi/swagger-ui
