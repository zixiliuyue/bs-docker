ARG BASEIMAGE=centos:centos7
FROM $BASEIMAGE

COPY hyper/bcs-ca-1.16 /bcs-cluster-autoscaler-1.16
COPY hyper/bcs-ca-1.22 /bcs-cluster-autoscaler-1.22
CMD ["/bcs-cluster-autoscaler-1.22"]