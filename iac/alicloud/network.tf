# This resource caused issue https://github.com/aliyun/terraform-provider-alicloud/issues/7967
resource "alicloud_vpc" "default" {
  is_default  = true
  description = "test"
  cidr_block  = "10.0.0.0/16"
  vpc_name    = "test-vpc"
  enable_ipv6 = false
}
