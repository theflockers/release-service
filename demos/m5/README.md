# Milestone 5 demo

This directory contains the resources needed for running the release demo for Milestone 5.

# Boostraping

After bootstrapping the cluster with App Studio, you can prepare the demo by running the [boostrap.sh](demos/m5/boostrap.sh) script.
```
$ ./demos/m5/boostrap.sh
```
This will create all the required resources to run the demo. Those are:
* Demo resources in the base directory.
* Quay secret to pull and push images (a password to the robot account will be requested).
* Cosign secret with the public key extracted from tekton-chains.

# Triggering the demo

To trigger the demo, the only thing needed after boostraping is to apply the [release.yaml](demos/m5/release.yaml) file, which will create a new Release in the cluster.
