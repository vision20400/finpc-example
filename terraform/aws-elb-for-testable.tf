################################################################
##
##  AWS ELB
##

##--------------------------------------------------------------
##  external network load balancer

resource aws_lb_listener nlb_grpc {
  load_balancer_arn = aws_lb.nlb.arn
  port              = aws_lb_target_group.nlb_grpc.port
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.nlb_grpc.arn
  }

  tags = {
    Name = "${var.project}-nlb-grpc"
  }
}

resource aws_lb_target_group_attachment nlb_grpc {
  port             = aws_lb_listener.alb_grpc.port
  target_group_arn = aws_lb_target_group.nlb_grpc.arn
  target_id        = aws_lb.alb.id
}

resource aws_lb_target_group nlb_grpc {
  name        = "nlb-${var.project}-grpc"
  port        = aws_lb_listener.alb_grpc.port
  protocol    = "TCP"
  target_type = "alb"
  vpc_id      = aws_vpc.this.id

  health_check {
    matcher  = "200,404"
    path     = "/"
    protocol = "HTTPS"
  }
}

resource aws_lb_listener nlb_proxy_postgres {
  load_balancer_arn = aws_lb.nlb.arn
  port              = aws_lb_target_group.nlb_proxy_postgres.port
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.nlb_proxy_postgres.arn
  }

  tags = {
    Name = "${var.project}-nlb-proxy-postgres"
  }
}

resource aws_lb_target_group nlb_proxy_postgres {
  name             = "nlb-${var.project}-proxy-postgres"
  port             = local.postgres_port
  protocol         = "TCP"
  target_type      = "ip"
  vpc_id           = aws_vpc.this.id

  deregistration_delay = 10

  health_check {
    interval = 10
    matcher  = "200"
    path     = "/"
    port     = 80
    protocol = "HTTP"
  }
}
