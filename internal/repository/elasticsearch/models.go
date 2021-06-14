package elasticsearch

type AggregationOutput struct {
	Key   interface{}
	Count int64
}

const PutMapping = `
	{
	  "properties":{
			"title":{
				"type": "text",
				"analyzer":"russian",
				"fields": {
          			"keyword": { 
            			"type": "keyword"
					}
        		}
			},
			"author": {
				"type":"keyword"
			},
			"link":{
				"type":"keyword"
			},
			"description":{
				"type":"text",
				"fields": {
          			"keyword": { 
            			"type": "keyword"
					}
        	},
				"analyzer":"russian"
			},
			"pub_date":{
				"type":"date"
			},
			"text":{
				"type":"text",
				"analyzer":"russian"
			}
	 }
}`
