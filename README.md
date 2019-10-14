# secret-sidecar
The secret-side car project is an example of how you can retrieve a secret from AWS Secrets Manager using an init container and mount it as a RAM disk that is shared with an application container. The init container is written in Go and uses IAM Roles for Service Accounts (IRSA) to assume an identity that has permission to read the secret. 

## Using the init container
The init container looks for 2 environment variables: AWS_REGION and SECRET_NAME. The values of these variables should be included in your pod manifest. The AWS_REGION designates the region where the secret is stored and the SECRET_NAME refers to the name of the secret in AWS Secrets Manager.  For an example, see hello.deployment.yaml.

The serviceAccountName references the Kubernetes service account that allows the init container to assume a IAM role that allows it to read secrets from AWS Secrets Manager.  When running in production, this service account and IAM role should be scoped to read a specific secret or set of secrets.  

## ECS
This sidecar pattern also works with ECS using [Enhanced Container Dependencies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_dependson).  Instead of a RAM disk, the secret is mounted as a Docker volume on the local host. See [using data volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html) for additional information.  An sample task definition is included in the ecs-task-def directory. 

## TODO
Create a mutating webhook that automatically adds the init container to a pod when a specific annotation is added to the pod. 
