docker rm -vf dashboard
docker run --name dashboard -v ~/.kube/config:/config -e CONFIG_LOCATION=/config -p 8060:80 patnaikshekhar/kubedashboard:0.0.1