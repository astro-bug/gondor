@ECHO OFF

REM ��װ������ssl֤��
mkcert.exe -install
mkcert.exe -key-file certs\key.pem -cert-file certs\cert.pem localhost 127.0.0.1

REM ��װ������Windows����
winsw.exe install
winsw.exe start

PAUSE