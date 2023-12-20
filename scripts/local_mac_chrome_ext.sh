#!/bin/sh
set -x
set -e

export CLIENT_ID='"client_id": "149500152182-1nlfm111isrv0c8828mg3u16q22l434q.apps.googleusercontent.com",' 
export KEY='"key": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxk66E7tDSLL2ZqxhrAHpJOXRxmwSxCSSyRSzYjgx6B150ZdyGe1lAQWOkMfN4h1aK679mJaNTXRWllvxRNaZ+qNagZoy+96/57H7Au+RnX7U7NPv79jVxNq1fDI/hFp5tVZ1cGS31gNeZewR5cPvsTkekHSNFZv8isUtyDomwsRJrxrvIj1a1sL2O9w70aWcfZ/m8eu2gOWWtzfpOuB6K79DAUHw0xBOJ4pykgfYOfrDeZqrZbEzT0Y5d1kACTdg/h7WrE1KPWe0lUJfs7yvf7wMdjOEORR5Dl9i2DCq4ugHsYWMeQ8KmMSBqrlCLnAWFJ9hxhYCa9UJC6i7PaYncQIDAQAB",'
export CHROME_EXT_VERSION=$1 

envsubst < chrome-exten/manifest.json.tmpl > chrome-exten/manifest.json
