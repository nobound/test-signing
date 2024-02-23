""" 
Test RSA Signing
"""
import base64
from OpenSSL.crypto import FILETYPE_PEM, load_privatekey, sign

def get_signing_key(key_file: str) -> str:
    """
    Get the key from the PEM file
    """
    key_file = "./rsa_private_key.pem"
    with open(key_file, "r", encoding="utf8") as file:
        contents = file.read()
    return contents

def test_rsa_signing() -> None:
    """
    Test signing with RAS key
    """
    key_file = "./rsa_private_key.pem"
    payload = bytes("message to be signed", 'utf-8')
    pkey = load_privatekey(FILETYPE_PEM, get_signing_key(key_file=key_file))
    signature = sign(pkey, payload, "sha256")
    s = str(base64.b64encode(signature), encoding='utf-8')
    print(s)

test_rsa_signing()
