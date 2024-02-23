
Simple test to compare Go RSA signing function vs the Python ones (OpenSSL crypto function)

```bash
❯ python msg_signing.py
ZKTAxJ49JfbuHJCBJXZmQ23YlloPdhuiQqB5KUsLxi+4e1I0N0+TnJWnem0+p2jvzC/V/89AW/TGGH+radHCvA/OVdWW9x4sBVB+l1VlmVsiZJER66sj+ndereGLbU5HyL4gZSOTff6SoTcKL38iI65PrOIE/ChNiwF+1HIpCOyRTles6+LlkNKVnLEYwWJf0MpUm66lesAwGTqBVAuujvAhOVL98GdWgt58YXaYEdU7Fk6VW+YrWID+NOs8fewqN9AR2R+oCbnB+hir9RKgOXu9Zz42OYIXBZSMSV4oMjqCM532cdT5nq1hlOPGZqoUPSDx6UILaVxL/QucmwTVKA==

❯ go run msg_signing.go
ZKTAxJ49JfbuHJCBJXZmQ23YlloPdhuiQqB5KUsLxi+4e1I0N0+TnJWnem0+p2jvzC/V/89AW/TGGH+radHCvA/OVdWW9x4sBVB+l1VlmVsiZJER66sj+ndereGLbU5HyL4gZSOTff6SoTcKL38iI65PrOIE/ChNiwF+1HIpCOyRTles6+LlkNKVnLEYwWJf0MpUm66lesAwGTqBVAuujvAhOVL98GdWgt58YXaYEdU7Fk6VW+YrWID+NOs8fewqN9AR2R+oCbnB+hir9RKgOXu9Zz42OYIXBZSMSV4oMjqCM532cdT5nq1hlOPGZqoUPSDx6UILaVxL/QucmwTVKA==%                                          

```


Note: Use openssl command to generate RSA key pair

```bash
# Generate a private key
openssl genpkey -algorithm RSA -out private_key.pem

# Extract the public key from the private key
openssl rsa -pubout -in private_key.pem -out public_key.pem
```
