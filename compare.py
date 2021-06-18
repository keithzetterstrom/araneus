from elasticsearch import Elasticsearch
from datasketch import MinHash


def compare(signature1, signature2):
    mh1, mh2 = MinHash(hashvalues=signature1), MinHash(hashvalues=signature2)

    return mh1.jaccard(mh2)


res = Elasticsearch().search(index="items", body={'size': 2, 'query': {'match_all': {}}})

signature1 = res['hits']['hits'][0]['_source']['signature']
signature2 = res['hits']['hits'][1]['_source']['signature']

jaccard_k = compare(signature1, signature2)

print(jaccard_k)
