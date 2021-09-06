package elasticsearch

type AggregationOutput struct {
	Key   interface{}
	Count int64
}

const PutSetting = `"settings": {
        "analysis": {
            "filter": {
                "delimiter": {
                    "type": "word_delimiter",
                    "preserve_original": "true"
                },
                "jmorphy2_russian": {
                    "type": "jmorphy2_stemmer",
                    "name": "ru"
                }
            },
            "analyzer": {
                "text_ru": {
                    "tokenizer": "standard",
                    "filter": [
                        "lowercase",
                        "delimiter",
                        "jmorphy2_russian"
                    ]
                }
            }
        }
    },`

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

const PutMap = `{
	  "properties":{
			"title":{
				"type": "text",
				"analyzer":"text_ru",
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
				"analyzer":"text_ru"
			},
			"pub_date":{
				"type":"date"
			},
			"text":{
				"type":"text",
				"analyzer":"text_ru"
			}
	 }
}`