# webdavserver

[![Container Repository on Quay](https://quay.io/repository/dominikholler/webdavserver/status "Container Repository on Quay")](https://quay.io/repository/dominikholler/webdavserver)

webdavserver is a example application utilizing the webdav server of the golang standard library.

## Linux

```
export WEBDAVDIR=/tmp/webdav
mkdir $WEBDAVDIR
chmod 777 $WEBDAVDIR

podman  run --rm -p 8080:8080 -v $WEBDAVDIR:/srv/webdav:z -it quay.io/dominikholler/webdavserver
```

## Windows

```
$WEBDAVDIR = "$env:TEMP\\webdav"
New-Item -ItemType Directory -Path $WEBDAVDIR

wslc run --rm -p 8080:8080 -v ${WEBDAVDIR}:/srv/webdav -it quay.io/dominikholler/webdavserver
```
