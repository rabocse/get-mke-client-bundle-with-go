# get-mke-client-bundle

Go scripts to download the MKE (Mirantis Kubernetes Engine) client bundle.


## What is MKE? 

Mirantis Kubernetes Engine (MKE, formerly Universal Control Plane or UCP) is a container orchestration platform for developing and running modern applications at scale, on private clouds, public clouds, and on bare metal.

For more information, visit the documentation: https://docs.mirantis.com/mke/3.5/overview.html

## What is the Client Bundle?

The MKE client bundle allow us to use MKE with Docker CLI and kubectl. The bundle includes:

- A private and public key pair for authorizing your requests using MKE

- Utility scripts for configuring Docker CLI and kubectl with your MKE deployment

For additional information, visit the official documentation:

https://docs.mirantis.com/mke/3.5/ops/access-cluster/download-configure-client-bundle.html



## About The Scripts

At this moment, the Go scripts succesfully download the client bundle. 

To acomplish this we must get and authentication token before donwloading the file.

NOTE: Refactoring and additional capabilities will be added later on.


## How to Use the Scripts

Here is how the current execution of the scripts looks like:

```
❯ ./GetAuthToken -clus=x.x.x.x -user=admin-cluster -pass=MySuperSecretPassword
{"auth_token":"50217f00-b934-5306-a1b6-d5c25e643fa0"}
```

<br/>

```
❯ ./GetClientBundle -clus=x.x.x.x -tokn=50217f00-b934-5306-a1b6-d5c25e643fa0
Done.
```

## Improvements

- Both scripts accept flags for each needed parameter. This removes the need of hardcoding values in the source code whenever executing.
- The source code was group in functions.
  

## Refactoring

- Group code in fuctions. [DONE]
- Pass parameter to the script and avoid hardcoded values. [DONE]
- Create a single script.
- Mask the password entered by user.




