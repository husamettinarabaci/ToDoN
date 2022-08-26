# ToDoN
Web Based ToDo List Application for Modanisa

## Deciriptions

## Installation

## Usage

## For Developers
 - installation golang 
 - installation protoc
 - github secrets
 - folder structere


Namespace
- kubectl create -f Configs/K8S/Namespaces/todon-prod-namespace.yaml
- kubectl create -f Configs/K8S/Namespaces/todon-test-namespace.yaml
Config
- kubectl create -f Configs/K8S/Configs/prod/memcacheserver-prod-config.yaml
- kubectl create -f Configs/K8S/Configs/prod/webserver-prod-config.yaml 
- kubectl create -f Configs/K8S/Configs/test/memcacheserver-test-config.yaml 
- kubectl create -f Configs/K8S/Configs/test/webserver-test-config.yaml 
Cluster IP
- kubectl create -f Configs/K8S/ClusterIPs/prod/memcacheserver-prod-clusterip.yaml 
- kubectl create -f Configs/K8S/ClusterIPs/test/memcacheserver-test-clusterip.yaml 
LoadBalancer
- kubectl create -f Configs/K8S/LoadBalancers/prod/memcacheserver-prod-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/prod/webserver-prod-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/test/memcacheserver-test-loadbalancer.yaml 
- kubectl create -f Configs/K8S/LoadBalancers/test/webserver-test-loadbalancer.yaml 
Deployment
- kubectl create -f Configs/K8S/Deployments/prod/memcacheserver-prod-deployment.yaml
- kubectl create -f Configs/K8S/Deployments/prod/webserver-prod-deployment.yaml 
- kubectl create -f Configs/K8S/Deployments/test/memcacheserver-test-deployment.yaml 
- kubectl create -f Configs/K8S/Deployments/test/webserver-test-deployment.yaml 



 namespace todon-prod,todon-test

 memcahceserver Prod LoadBalancer IP : ac871331e10a24d3fb1775da164f9258-2020354426.eu-central-1.elb.amazonaws.com
 memcahceserver Test LoadBalancer IP : a06cfcc7dfbca4c829963fcf485472bb-53359176.eu-central-1.elb.amazonaws.com
 webserver Prod LoadBalancer IP : aab6a856cd8984348b4e34a58610665f-1150680336.eu-central-1.elb.amazonaws.com
 webserver Test LoadBalancer IP : a970bf7ee7cce4ea8be39748807dcc44-429638464.eu-central-1.elb.amazonaws.com