# AWSAdminAccess

In a lot of organizations there is a master AWS account and then
other accounts are added via consolidated billing. To better manage
the assets it is usually good to add a AdministratorAccess role to
the sub-account to allow the master account to monitor and control
costs on the sub-accounts. AWSAdminAccess provides a quick and easy
way to setup a trust policy for the AdministratorAccess.

## Running AWSAdminAccess

Download a binary for your system from the Releases page. You will
need to know the role name you want to create and the account number
of the master account.

```
AWSAdminAccess -r MasterAccountAccess -a 123456789012
```

## Building

To build the binaries it is preferable to use a docker build environment for
consistency. First build the docker buildn environment:

```
docker build -f build/Dockerfile-buildenv -t cloudtools:AWSAdminAccess-buildenv .
```

Next install the vendor package using glide:
```
glide install
```

And then build the binaries:
```
docker run -v `pwd`:/go/src/github.com/cloudtools/AWSAdminAccess -t cloudtools:AWSAdminAccess-buildenv bash -x build/build.sh
```
