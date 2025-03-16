# Configure the AWS Provider (commented out for local development)
# provider "aws" {
#   region = var.aws_region
# }

# EKS Cluster (commented out for local development)
# module "eks" {
#   source  = "terraform-aws-modules/eks/aws"
#   version = "~> 19.0"
#
#   cluster_name    = var.cluster_name
#   cluster_version = "1.27"
#
#   vpc_id     = module.vpc.vpc_id
#   subnet_ids = module.vpc.private_subnets
#
#   eks_managed_node_groups = {
#     default = {
#       min_size     = 1
#       max_size     = 3
#       desired_size = 2
#
#       instance_types = ["t3.medium"]
#     }
#   }
# }

# Output for local development
output "kubernetes_cluster_info" {
  value = "Using local minikube cluster for development"
} 