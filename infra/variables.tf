variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_region" {}

variable "instance_ami" {}
variable "instance_type_site" {}
variable "instance_key_name" {}

variable "ssh_user" {}
variable "ssh_private_key" {}

variable "dns_zone_id" {}
variable "dns_site_domain" {}

variable "script_setup_site" {
  type = "list"
}

