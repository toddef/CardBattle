# AWS Region (for future use)
variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "us-west-2"
}

# Cluster Name (for future use)
variable "cluster_name" {
  description = "Name of the EKS cluster"
  type        = string
  default     = "card-battle-cluster"
} 