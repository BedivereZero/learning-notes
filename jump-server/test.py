# Python 示例
# pip install requests drf-httpsig
import requests, datetime, json
from httpsig.requests_auth import HTTPSignatureAuth

def get_auth(KeyID, SecretID):
    signature_headers = ['(request-target)', 'accept', 'date']
    auth = HTTPSignatureAuth(key_id=KeyID, secret=SecretID, algorithm='hmac-sha256', headers=signature_headers)
    return auth

def get_user_info(jms_url, auth):
    url = jms_url + '/api/v1/users/users/'
    gmt_form = '%a, %d %b %Y %H:%M:%S GMT'
    headers = {
        'Accept': 'application/json',
        'X-JMS-ORG': '00000000-0000-0000-0000-000000000002',
        'Date': datetime.datetime.utcnow().strftime(gmt_form)
    }

    response = requests.get(url, auth=auth, headers=headers)
    print(json.loads(response.text))

def get_assets(jms_url, auth):
    url = jms_url + '/api/v1/assets/assets/'
    gmt_form = '%a, %d %b %Y %H:%M:%S GMT'
    headers = {
        'Accept': 'application/json',
        'X-JMS-ORG': '00000000-0000-0000-0000-000000000002',
        'Date': datetime.datetime.utcnow().strftime(gmt_form)
    }

    response = requests.get(url, auth=auth, headers=headers)
    print(response.text)
    # print(json.loads(response.text))x

if __name__ == '__main__':
    jms_url = 'http://localhost'
    KeyID = '3E7076FB-D5A4-4EC3-B34A-B50DFA546EF9'
    SecretID = '7F6644D2-F6BE-447D-95A1-DEC718DEEC26'
    auth = get_auth(KeyID, SecretID)
    get_assets(jms_url, auth)
