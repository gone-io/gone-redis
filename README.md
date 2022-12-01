# gone-redis
gone adapter for redis

【notice】
1. some command contains so many params, so split it as two or there func.
2. all command here just supported redis before v5.0.0, and some command rarely used not supported here. You can use it by "conn.do()" with redigo.