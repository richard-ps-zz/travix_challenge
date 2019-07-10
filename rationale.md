## Design Rationale

This document describes the design choices behind the automation built to solve the challenge. It's important to say that the implementation was done with simplicity in mind and with little
to no concern for security, as the objetive of this project was
to show ability to solve infrastructure problems using Kubernetes and Google Cloud Platform and not to produce a production ready cluster.

### Ingress 

As a way of creating an entrypoint to the requests reaching our 
cluster we needed to have an Ingress that could direct requests to the Stateless App Load Balancer. 

The choice was to use the Helm chart stable/nginx containing an NGINX Ingress Controller that is ready to run all the Kubernetes resources required to run NGINX as an ingress inside Kubernetes.

To effectively route the requests to the Stateless App, I also
needed to create an [ingress resource](https://github.com/richard-ps/travix_challenge/blob/master/ingress/ingress-resource.yml) that added this configuration to our Ingress Controller.

[Here](https://cloud.google.com/community/tutorials/nginx-ingress-gke) you can find the reference used to create the Ingress.    

### Stateless App

The stateless app had to be composed of two components: one reverse proxy and one API. Both components ended up being placed
into a Pod resource that was managed by the same [Deployment](https://github.com/richard-ps/travix_challenge/blob/master/apps/stateless-app/stateless-app-deploy.yml).

For the side car component I used NGINX again. No fancy [configuration](https://github.com/richard-ps/travix_challenge/blob/master/apps/stateless-app/nginx-configmap.yml) was needed. The web server simply executes a proxy pass to his back-end service, the API component.

The API part was developed using the Golang programming language 
for no particular reason. The API's Docker image used in the Deployment was built using the Docker Hub automated build capability and stored as a public image into my personal repository.

### Stateful App

In the stateful part of our application we needed a table of 
articles. As the type of articles was not specified, I used a
public dataset of news paper articles available [here](https://www.kaggle.com/asad1m9a9h6mood/news-articles/version/1).

As the word `table` was used, I presumed the creation of a relational database might be required so I chose to run a
PostgreSQL as a StatefulSet resource. To accomplish that I
used a Helm chart called stable/postgresql.

The StatefulSet give us some garantees that Deployments do not have, for example, guarantees about the ordering and uniqueness of Pods managed by it. For the purposes of this project this difference doesn't affect much of the functionaty of our cluster.
Nevertheless, StatefulSets are the recommended way of running databases into Kubernetes.

Also a pod was created to load the data into the correct table
inside PostgreSQL as soon as the database is ready.

### Autoscaling Capability

As the metric used to autoscale the number of replicas running 
the API handler was not mentioned, I have used the average cpu 
utilization as the metric in which the scaling was predicated to be triggered.

To implement this feature into our Kubernetes I simply created 
a Horizontal Pod Autoscaling resource and configured it to trigger autoscaling of the API Pod as soon as average CPU utilization reached 30%.

### Possible improvements

This implementation is far from perfect and many adjustments could be made to make it better.

The first thing that was left out of the scope was the creation
of a docker image containing all the dependencies to run the automation, this way anyone who's trying to reproduce it would
have their work cut out.

The next enhancement would be to run all the Kubernetes configurations with Ansible. Using Ansible would enable us to apply parts of the infrastructure separetely with tags to different tasks, making our automation much more flexible. 