## Travix Challenge

![alt text](https://github.com/richard-ps/travix_interview/raw/master/src/k8s.png)

To reproduce the infrastructure defined on this repository you
will need to set up an environment with the following:

### Dependencies

terraform (version 0.12.3)README.md

kubectl   (version 1.15)

helm      (version 2.13.0)

gcloud    (version 244.0.0)

Luckily, if you are using a Debian-like distro, you can install
these dependencies using the command below in the root folder:

`$ sudo ./install_dependencies`

Also, you will have to have a GCP project and a service account
in this project with the roles:

- Kubernetes Engine Admin
- Owner 

### Configuring your variables

In order to create the infrastructure in your own GCP project it's important to configure your service account and project as 
inputs to the automation. 

To do it, just download a key to your service account in JSON format and place it in the root folder of the project, renaming it to sa_key.json.

There are other two variables that have to be configured inside the script 'run' that lies into the root folder of this repo.

`export PROJECT_ID={YOUR-PROJECT-ID}`

### Running the automation

Now that you have your service account and environment variables set up, it's time to run the automation inside the 'run' script.

`$ ./run`

The first thing this script will do, after exporting the variables mentioned in the previous section, is launch a new tab in your web browser in order to log into your google account.

After that, 'run' script will call other scripts that are responsible for creating diferent components and resources of the infrastructure (eg: run_db and run-stateless).

Wait for a few minutes and the API will be ready to receive requests.

### Testing the API

To check if the automation has worked out well, we need to find out the external IP of the ingress controller created:

`kubectl get svc | grep nginx-ingress-controller`

You should see two different IP addresses, the first one is the cluster IP, used internally, and the second is the external IP.

Copy the external IP value and call this URL on your browser:

`http://{INGRESS-EXTERNAL-IP}/articles`

or, if you prefer, on the command line:

`curl http://{INGRESS-EXTERNAL-IP}/articles`

You should see values of news paper articles returning from the API.

### Scaling the API

If you want to test the stateless app autoscaling capability you
can do it by running a pod that generates some load to the API:

```
kubectl run -i --tty load-generator --image=busybox /bin/sh 
Hit enter for command prompt
while true; do wget -q -O- http://{INGRESS-EXTERNAL-IP}/articles; done
```

And then you can watch the number of replicas grow with:

`kubectl get hpa -w` 

### Cleaning up

To destroy the cluster created previously and all its resources,
just run the following command from the root folder:

`terraform destroy cluster`