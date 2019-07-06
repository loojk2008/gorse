===
API
===

Insert
======

Insert New Items
----------------

.. code-block:: bash

   curl -X PUT -H 'Content-Type: application/json' 127.0.0.1:8080/items -d '[2048,2049,2050]'


Insert three new items ``[2048,2049,2050]``.

.. code-block:: json

    {
        "Failed": false,
        "BeforeCount": 1687,
        "AfterCount": 1690
    }

Insert New Feedbacks
--------------------

.. code-block:: bash

    curl -X PUT -H 'Content-Type: application/json' 127.0.0.1:8080/ratings \
        -d '[{"UserId":2048,"ItemId":1000,"Rating":3},{"UserId":2049,"ItemId":1000,"Rating":3},{"UserId":2050,"ItemId":1000,"Rating":3}]'

Insert three new ratings: ``<2048,1000,3>``, ``<2049,1000,3>`` and ``<2050,1000,3>``.

.. code-block:: json

    {
        "Failed": false,
        "BeforeCount": 100000,
        "AfterCount": 100003
    }


Query
=====

Get Popular Items
-----------------

.. code-block:: bash

    curl 127.0.0.1:8080/popular?number=10

Get top 10 popular items.

.. code-block:: json

    {
        "Failed": false,
        "Items": [50, 258, 100, 181, 294, 286, 288, 1, 300, 121]
    }

Get Random Items
----------------

.. code-block:: bash

    curl 127.0.0.1:8080/random?number=10

Get 10 random items.

.. code-block:: json

    {
        "Failed": false,
        "Items": [454, 1584, 1124, 1165, 1617, 945, 1268, 610, 783, 1091]
    }


Get Recommended Items
---------------------

.. code-block:: bash

    curl 127.0.0.1:8080/recommends/100/?number=10


Recommend 10 items for the 100th user.

.. code-block:: json

    {
        "Failed": false,
        "Items": [948, 748, 906, 984, 352, 327, 242, 890, 1294, 680]
    }

Get Similar Items
-----------------

.. code-block:: bash

    curl 127.0.0.1:8080/neighbors/100/?number=10


Get top 10 similar items for the 100th item.

.. code-block:: json

    {
        "Failed": false,
        "Items": [247, 897, 442, 643, 757, 788, 838, 957, 907, 891]
    }
