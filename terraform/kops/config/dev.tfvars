environment   = "dev"
cluster_name  = "goreact-webapp-cluster"
region        = "us-east-1"

### VPC MODULE
vpc= {
    cidr          = "10.2.0.0/16",
    dns_hostnames = true,
    dns_support   = true,
    tenancy       = "default",
  }
public_subnets  = ["10.2.0.0/24","10.2.1.0/24","10.2.5.0/24"]
private_subnets = ["10.2.2.0/24","10.2.3.0/24","10.2.4.0/24"]

### KUBERNETES MODULE
kops_state_bucket = "kops-mytest-store/kops"
worker_node_type = "t3.micro"
min_worker_nodes = "1"
max_worker_nodes = "2"
master_node_type = "t3.micro"