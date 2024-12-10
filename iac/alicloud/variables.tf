variable "access_key" {
  description = "The Access Key ID for AliCloud"
  type        = string
}

variable "secret_key" {
  description = "The Secret Access Key for AliCloud"
  type        = string
  sensitive   = true
}

variable "region" {
  description = "The region to deploy resources in AliCloud"
  type        = string
  default     = "cn-hangzhou"
}
