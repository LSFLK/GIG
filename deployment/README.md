### Server Setup using Kubernetes (Optional):
If you want to create a kubernetes node with all dependencies follow the steps given below. Otherwise you can modify the project configuration to match your server environment.

Create Persistent Directory

    sudo mkdir /home/data/db -p
    sudo mkdir /home/data/minio -p
    sudo chmod -R 777 /home/data/
    
Create separate node for GIG Server configurations:
    
    kubectl create namespace gig-api-node
    kubens gig-api-node
    
Initiate MongoDB and Minio Servers using following commands

For MongoDB:

    kubectl apply -f deployment/mongodb/persistent-volume.yaml
    kubectl apply -f deployment/mongodb/secrets.yaml
    kubectl apply -f deployment/mongodb/configmap.yaml
    kubectl apply -f deployment/mongodb/statefulsets.yaml
    kubectl apply -f deployment/mongodb/service.yaml
    kubectl apply -f deployment/mongodb/ingress.yaml
    
For Minio: For more details check [MinIO Kubernetes YAML Files](https://github.com/minio/minio/blob/master/docs/orchestration/kubernetes/k8s-yaml.md)

    kubectl apply -f deployment/minio/minio-standalone-pv.yaml
    kubectl apply -f deployment/minio/minio-standalone-deployment.yaml
    kubectl apply -f deployment/minio/minio-standalone-service.yaml
    kubectl apply -f deployment/minio/ingress.yaml
    
Use the following command to get the mongodb Server IP

    kubectl get svc |grep database
    
Use the following command to get the minio Server IP

    kubectl get svc |grep minio-service
