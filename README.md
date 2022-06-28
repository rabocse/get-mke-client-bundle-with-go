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

This is just an initial version, refactoring and additional capabilities will be added later on.

## How To Use the Scripts (As in main branch from this repo)

The first versions of the scripts work like this:

```
❯ ./GetAuthToken
{"auth_token":"896386ba-15f2-4bc6-9be4-99eaf10ca345"}


❯ ./GetClientBundle
Done.


❯ ls -al
total 32
drwxr-xr-x   3 aescobar  staff     96 Apr  1 14:25 .
drwxr-xr-x  16 aescobar  staff    512 Apr  1 14:22 ..
-rw-r--r--   1 aescobar  staff  15680 Apr  1 14:25 bundle.zip


❯ unzip bundle.zip
Archive:  bundle.zip
 extracting: clientbundle

❯ ls -al
total 64
drwxr-xr-x   4 aescobar  staff    128 Apr  1 14:25 .
drwxr-xr-x  16 aescobar  staff    512 Apr  1 14:22 ..
-rw-r--r--   1 aescobar  staff  15680 Apr  1 14:25 bundle.zip
-rw-r--r--   1 aescobar  staff  15524 Apr  1 14:25 clientbundle
```

## Improvements for GetAuthToken (Work In Progress)

Below is the current status for GetAuthToken script:

```
❯ ./GetAuthToken -clus=xxxxxxxxxxxxxx -user=xyxyxyxyxyx -pass=kkkkkkkkkkkkkkk
########### INPUT: Server ##########################
Cluster:  https://xxxxxxxxxxxxxx/auth/login
 
########### INPUT: Credentials #####################
Username:  xyxyxyxyxyx
Password:  kkkkkkkkkkkkkkk
 
########### PARSED INPUT: Credentials in JSON ######
{"username":"xyxyxyxyxyx","password":"kkkkkkkkkkkkkkk"}
 
########### OUTPUT: AUTH TOKEN #####################
{"auth_token":"a4f4682c-1a61-41be-aa27-176b741420de"}
```

- GetAuthToken accepts flags that allow it to pass the url, username and password.
- The printing of "INPUT" is only temporary and for verification purposes.
  

## Improvements for GetClientBundle.go (Work In Progress)

Below is the current status for GetAuthToken script:

```
❯ ./GetClientBundle -clus=xxxxxxxxxxx -tokn=a4f4682c-1a61-41be-aa27-176b741420de
########### INPUT: Server ##########################
Cluster:  https://xxxxxxxxxxxx/api/clientbundle

########### INPUT: Token #####################
Authentication Token: a4f4682c-1a61-41be-aa27-176b741420de
Done.
```

- Similar as GetAuthToken, now GetClientBundle accepts flags.
- The printing is only for verification purposes.


## Refactoring

- Group code in fuctions.
- Create a single script.
- Pass parameter to the script and avoid hardcoded values.
- Mask the password entered by user.




