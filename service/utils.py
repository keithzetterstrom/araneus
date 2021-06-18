import string
from datasketch import MinHash
import items_pb2


def shingle(text, k):
    text = text.translate(str.maketrans('', '', string.punctuation))

    text_for_shingle = text.split(' ')

    shingle_set = []

    for i in range(len(text_for_shingle) - k):
        shingle_set.append("".join(text_for_shingle[i:i + k]))

    return shingle_set


def get_minhash_signature(shingle_set):
    mh = MinHash()
    for el in shingle_set:
        mh.update(el.encode('utf8'))
    return mh.hashvalues


def get_signature(item: items_pb2.ItemGRPC):
    s = shingle(item.Text, 3)
    signature = get_minhash_signature(s)
    return signature


def compare_with_minhash(set1, set2):
    mh1, mh2 = MinHash(), MinHash()
    for el in set1:
        mh1.update(el.encode('utf8'))
    for el in set2:
        mh2.update(el.encode('utf8'))

    return mh1.jaccard(mh2)
