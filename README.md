## About

This project is a web api. It has 2 endpoints.

## Before running

You should set GO_ENV of your OS. Default GO_ENV is "prod". GO_ENV value reads app settings from settings.json

## Requirements

 - go 1.17

## Endpoints

**[GET]**

    /api/in-memory

You must specify query param key. Example /api/in-memory?key=john

**Possible Responses**:

*Successful response (status code : 200)*
  

   

    {
	    "key":"value"
    }

*Error response (status code: 400 || 404 || 500)*

    {
    	"message":"string"
    }



**[POST]**

    /api/in-memory

**Request**

    {
	   "key":"string",
	   "value":"string"
    }



**Possible Responses**:

*Successful response (status code : 201)*
 

    No content

*Error response (status code: 400)*

    {
	    "message":"string"
    }

**[POST]**

    /api/records

**Request**

"startDate" and "endDate" must be short date format. Format: yyyy-mm-dd

    {
	    "minCount":0,
	    "maxCount":0,
	    "endDate":"string",
	    "startDate": "string"
    }

**Response**

If the status code is 200, the response code is equal to zero and the response msg equal to "Success". If the status code is not 200, the response code is equal to the server's status code and the response msg is equal to the server's message

*Example response (status code : 200 || 400 || 500)*

    {
	    "code":0,
	    "msg":"string",
	    "records":[{
		    "key":"string",
		    "createdAt": "date-time",
		    "count": 0
	    }]
    }

