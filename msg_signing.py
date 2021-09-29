
import base64
from OpenSSL.crypto import FILETYPE_PEM, load_privatekey, sign

def get_signing_key():
    key_file = "./rsa_private_key.pem"
    key = open(key_file)
    return key.read()

def rsa_signing(payload):
    pkey = load_privatekey(FILETYPE_PEM, get_signing_key())
    signature = sign(pkey, payload, "sha256")
    s = str(base64.b64encode(signature), encoding='utf-8')
    print(s)

def test_rsa_signing():
    payload = bytes("message to be signed", 'utf-8')
    rsa_signing(payload)


test_rsa_signing()
