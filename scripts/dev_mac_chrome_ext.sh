#!/bin/sh
set -x
set -e

export CLIENT_ID="" 
# retrieve the correct key from Chrome Developer Dashboard (https://chrome.google.com) -> open the extension page -> Package tab -> click View public key.
export KEY='"key": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApy7O4tuda2CWVXMgu3EFrrs4qv4IXnt2KQJG3KC404aYG1Hb8nvjwVpIqIdHG0QGQUH3nxIIuGnQtbFJxXzEF9bgWsWTNFe5jzrC9wKxVEjA/W50jMKIcqmypZb4Yi4D9oTHxhKvIyaNnxnljXR99f7Y8kZ3aDq5pduUPsubsWS76IvTX/7wog8KEjB2esF0+GmhRElA1e1MEZIn0iujhd1TlPvkFtouLg3PgC7mOV7Sgzb2SfbD36N+G/ImO6f0mn3b85GQXZ+mMcrbhPG9Q9feHgBpZVdkcwmmWfx4DxU/CoXyLu/989VIf79n7GgoeoL03RtIGK0TCWK8HXbM9wIDAQAB",'
export CHROME_EXT_VERSION=$1

envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
zip chrome-exten.zip chrome-exten/*
