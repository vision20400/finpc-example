output endpoint {
  value = aws_lb.nlb.dns_name
}

output grpc_cacert {
  value = base64encode(tls_self_signed_cert.ca.cert_pem)
  sensitive = true
}

output grpc_endpoint {
  value = aws_lb.alb.dns_name
}

output postgres_endpoint {
  value = aws_db_instance.this.endpoint
}

output postgres_password {
  value     = random_password.rds.result
  sensitive = true
}
