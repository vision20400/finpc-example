################################################################
##
##  AWS VPC
##

locals {
  availability_zones = ["a", "b"]
}

resource aws_vpc this {
  # 10.0.0.0/16
  cidr_block = var.cidr_block

  tags = {
    Name = var.project
  }
}

resource aws_subnet private {
  # 10.0.0.0/18
  # 10.0.64.0/18
  for_each = { for i, v in local.availability_zones : v => i }

  availability_zone = "${var.aws_region}${each.key}"
  cidr_block        = cidrsubnet(var.cidr_block, 2, each.value)
  vpc_id            = aws_vpc.this.id

  tags = {
    Name = "${var.project}-private-${each.key}"
  }
}

resource aws_subnet public {
  # 10.0.128.0/18
  # 10.0.192.0/18
  for_each = { for i, v in local.availability_zones : v => i }

  availability_zone       = "${var.aws_region}${each.key}"
  cidr_block              = cidrsubnet(var.cidr_block, 2, each.value + 2)
  map_public_ip_on_launch = true
  vpc_id                  = aws_vpc.this.id

  tags = {
    Name = "${var.project}-public-${each.key}"
  }
}

##--------------------------------------------------------------
##  gateways

resource aws_internet_gateway this {
  vpc_id = aws_vpc.this.id

  tags = {
    Name = var.project
  }
}

resource aws_nat_gateway this {
  allocation_id = aws_eip.this.id
  subnet_id     = aws_subnet.public["a"].id

  tags = {
    Name = "${var.project}"
  }
}

resource aws_eip this {
  tags = {
    Name = "${var.project}"
  }
}

##--------------------------------------------------------------
##  route tables

resource aws_route_table_association private {
  for_each = aws_subnet.private

  route_table_id = aws_route_table.private.id
  subnet_id      = each.value.id
}

resource aws_route_table private {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.this.id
  }

  tags = {
    Name = "${var.project}-private"
  }
}

resource aws_route_table_association public {
  for_each = aws_subnet.public

  route_table_id = aws_route_table.public.id
  subnet_id      = each.value.id
}

resource aws_route_table public {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }

  tags = {
    Name = "${var.project}-public"
  }
}

################################################################
##
##  AWS CloudWatch
##

resource aws_cloudwatch_log_group this {
  name              = var.project
  retention_in_days = 1
}
