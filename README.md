# webdavserver

[![Container Repository on Quay](https://quay.io/repository/dominikholler/webdavserver/status "Container Repository on Quay")](https://quay.io/repository/dominikholler/webdavserver)

webdavserver is a example application utilizing the webdav server of the golang standard library.

```
export WEBDAVDIR=/tmp/webdav
mkdir $WEBDAVDIR
chmod 777 $WEBDAVDIR

podman  run --rm -p 8080:8080 -v $WEBDAVDIR:/srv/webdav:z -it quay.io/dominikholler/webdavserver
```
