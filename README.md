# GIG
General Information Graph

## Project Structure

The directory structure of GIG:

    conf/                       Configuration directory
        app.conf                GIG server configuration file
        routes                  Routes definition file

    app/                        App sources
        init.go                 Interceptor registration
        cache/                  Cache Directory
        controllers/            App controllers
            api/                Inbound/Outbound API controllers
        data/                   Data Files for importing             
        models/                 Model classes
        repositories            Model Repositories
        storages/               Storage Handlers
        utilities/
            config/             Configuration Handler Class
            crawlers/           Data Crawler Classes
            entity_handlers/    Entity Management Classes
            importers/          Data Importer Classes
            normalizers/        Normalizer Classes
            parsers/            Source Parser Classes
            request_handlers/   Request Handler Classes
        views/                  Templates directory            

    messages/                   Message files

    public/                     Public static assets
        css/                    CSS files
        js/                     Javascript files
        images/                 Image files

    tests/                      Test suites
    
GIG Eco-System:

![GIG High Level Architecture](docs/images/gig_dataflow_diagram.png)

## Get Started

### Deployment Requirements
* Golang
* Revel
* Mongo Server for Database Hosting
* Docker for Deploying Minio Server
* Minio Server for File Hosting
* Python for NER Recognition Server
* Google Custom Search API

### Server Setup using Kubernetes (Optional):
If you want to create a kubernetes node with all dependencies follow the steps given below. Otherwise you can modify the project configuration to match your server environment.

Create Persistent Directory

    sudo mkdir /home/data/db -p
    sudo chmod -R 777 /home/data/
    
Install Kubernetes: use the following commands inside the project directory to create a namespace.

If you have not configured kubernetes already:

    sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=127.0.0.1
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml
    kubectl taint nodes --all node-role.kubernetes.io/master-
    
Create separate node for GIG Server configurations:
    
    kubectl create namespace gig-api-node
    kubens gig-api-node
    
Initiate MongoDB and Minio Servers using following commands

    kubectl apply -f deployment/mongodb/persistent-volume.yaml
    kubectl apply -f deployment/mongodb/persistent-volume-claim.yaml
    kubectl apply -f deployment/mongodb/secrets.yaml
    kubectl apply -f deployment/mongodb/configmap.yaml
    kubectl apply -f deployment/mongodb/statefulsets.yaml
    kubectl apply -f deployment/mongodb/service.yaml
    kubectl apply -f deployment/mongodb/ingress.yaml
    
Use the following command to get the mongodb Server IP

    kubectl get svc |grep database| cut -d' ' -f7

### First time run:

Create cache directory:

    mkdir app/cache
    
    
Configure mongo.path at conf/app.conf using the IP. Refer [How to Configure the Server](conf/README.md)

    [dev]
    ...
    mongo.path = mongodb://gig:GIGLSF2019@18.221.69.238:27017/gig
    ...
    
### Run Server:

### `revel run`
    
### Build Command:

    revel build -m prod -t build
    ./build/run.sh

## Help
* [API Documentation](https://app.swaggerhub.com/apis-docs/LSFGIG/GIG_API/1.0.0)
* [Utility Documentation](commons/README.md)
* [Crawlers Documentation](scripts/crawlers/README.md)

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [Revel API documentation](https://godoc.org/github.com/revel/revel).