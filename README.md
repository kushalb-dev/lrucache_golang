# <u>LRU Cache in Golang</u>

The code above is a simple implementation of LRU Cache in Golang. The above caching mechanisms represent the caching observed in a true LRU cache that follows the following rules:

* If an item already exists in the cache, we remove the item and add it to the beginning of cache
* Order of items is maintained
* Deletion of items happens at the tail, addition happens at the beginning

## <u>Simple Workflow</u>

The code utilizes structs in golang to build a cache that primarily contains the following things:

* LinkedList: The elements in the cache are presented as nodes in a linkedlist. A hashMap is used to preserve the uniqueness of the node. Addition and Deletion of items in the linkedlist is so preserved that order of items remains preserved. 

* Hashmap: A hashmap is used to maintain the uniqueness of items in the cache. The hashmap hashes the data in the cache, which is a string, to form a key value pair of string and a linkedlist node

![GO_Image](https://img-c.udemycdn.com/course/750x422/4689134_cb97.jpg)

Thanks for Checking Out!