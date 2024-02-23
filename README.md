
Simple test to compare Go RSA signing function vs the Python ones (OpenSSL crypto function)

```sh
$ python msg_signing.py 
UcjiIQ3q+Ysp8GhjRV0oxFyqmdLAXdSbQSJa1fTgJQo/tSrukC9oXyjvHyCxm5+1aV0aJaoiR79ubwnzUxj+XxOGP2JTxQPDDZ6sHARx2cFzdT+q34BrlnKGKO/j74pG6d61pw6yMekvCI72XliJQR4VliFmDGBgfbbw9N8fUY4h0hJ5UIB1S2Ddz/0au3cY+Xg2tp15pLnz9or5LbenNCoDrw41yuUMjNOAiBKKp7BrYIx9JtI3hSJpU4RcT45IC8D0qUKTqp9uyEXTPmp9twVPFu21UwJo9xj75CtCsksNP3CT/liQjs96GG9WOHFBuq3nUFbWrCzureymkK2V3Q==
```

```sh
$ go run msg_signing.go 
UcjiIQ3q+Ysp8GhjRV0oxFyqmdLAXdSbQSJa1fTgJQo/tSrukC9oXyjvHyCxm5+1aV0aJaoiR79ubwnzUxj+XxOGP2JTxQPDDZ6sHARx2cFzdT+q34BrlnKGKO/j74pG6d61pw6yMekvCI72XliJQR4VliFmDGBgfbbw9N8fUY4h0hJ5UIB1S2Ddz/0au3cY+Xg2tp15pLnz9or5LbenNCoDrw41yuUMjNOAiBKKp7BrYIx9JtI3hSJpU4RcT45IC8D0qUKTqp9uyEXTPmp9twVPFu21UwJo9xj75CtCsksNP3CT/liQjs96GG9WOHFBuq3nUFbWrCzureymkK2V3Q==%      
```                                    
