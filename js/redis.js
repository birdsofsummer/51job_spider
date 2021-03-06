const { promisify } = require("util");
const bird=require("bluebird")
const redis=bird.promisifyAll(require("redis"));
//redis = require("redis");
const c= redis.createClient();

const k="black"
const add=(d=[])=>c.SADDAsync(k,d)
const del=(d=[])=>c.SREMAsync(k,d)
const list=()=>c.smembersAsync(k)

module.exports={
    add,
    del,
    list,
}

/*
APPEND
ASKING
AUTH
BATCH
BGREWRITEAOF
BGSAVE
BITCOUNT
BITFIELD
BITOP
BITPOS
BLPOP
BRPOP
BRPOPLPUSH
BZPOPMAX
BZPOPMIN
CLIENT
CLUSTER
COMMAND
CONFIG
DBSIZE
DEBUG
DECR
DECRBY
DEL
DISCARD
DUMP
ECHO
EVAL
EVALSHA
EXEC
EXISTS
EXPIRE
EXPIREAT
FLUSHALL
FLUSHDB
GEOADD
GEODIST
GEOHASH
GEOPOS
GEORADIUS
GEORADIUSBYMEMBER
GEORADIUSBYMEMBER_RO
GEORADIUS_RO
GET
GETBIT
GETRANGE
GETSET
HDEL
HEXISTS
HGET
HGETALL
HINCRBY
HINCRBYFLOAT
HKEYS
HLEN
HMGET
HMSET
HOST:
HOST_
HSCAN
HSET
HSETNX
HSTRLEN
HVALS
INCR
INCRBY
INCRBYFLOAT
INFO
KEYS
LASTSAVE
LATENCY
LINDEX
LINSERT
LLEN
LPOP
LPUSH
LPUSHX
LRANGE
LREM
LSET
LTRIM
MEMORY
MGET
MIGRATE
MODULE
MONITOR
MOVE
MSET
MSETNX
MULTI
OBJECT
PERSIST
PEXPIRE
PEXPIREAT
PFADD
PFCOUNT
PFDEBUG
PFMERGE
PFSELFTEST
PING
POST
PSETEX
PSUBSCRIBE
PSYNC
PTTL
PUBLISH
PUBSUB
PUNSUBSCRIBE
QUIT
RANDOMKEY
READONLY
READWRITE
RENAME
RENAMENX
REPLCONF
RESTORE
RESTORE-ASKING
RESTORE_ASKING
ROLE
RPOP
RPOPLPUSH
RPUSH
RPUSHX
SADD
SAVE
SCAN
SCARD
SCRIPT
SDIFF
SDIFFSTORE
SELECT
SET
SETBIT
SETEX
SETNX
SETRANGE
SHUTDOWN
SINTER
SINTERSTORE
SISMEMBER
SLAVEOF
SLOWLOG
SMEMBERS
SMOVE
SORT
SPOP
SRANDMEMBER
SREM
SSCAN
STRLEN
SUBSCRIBE
SUBSTR
SUNION
SUNIONSTORE
SWAPDB
SYNC
TIME
TOUCH
TTL
TYPE
UNLINK
UNSUBSCRIBE
UNWATCH
WAIT
WATCH
XACK
XADD
XCLAIM
XDEL
XGROUP
XINFO
XLEN
XPENDING
XRANGE
XREAD
XREADGROUP
XREVRANGE
XTRIM
ZADD
ZCARD
ZCOUNT
ZINCRBY
ZINTERSTORE
ZLEXCOUNT
ZPOPMAX
ZPOPMIN
ZRANGE
ZRANGEBYLEX
ZRANGEBYSCORE
ZRANK
ZREM
ZREMRANGEBYLEX
ZREMRANGEBYRANK
ZREMRANGEBYSCORE
ZREVRANGE
ZREVRANGEBYLEX
ZREVRANGEBYSCORE
ZREVRANK
ZSCAN
ZSCORE
ZUNIONSTORE
append
asking
auth
batch
bgrewriteaof
bgsave
bitcount
bitfield
bitop
bitpos
blpop
brpop
brpoplpush
bzpopmax
bzpopmin
client
cluster
command
config
connection_gone
cork
create_stream
dbsize
debug
decr
decrby
del
discard
drain
dump
duplicate
echo
emit_idle
end
eval
evalsha
exec
exists
expire
expireat
flush_and_error
flushall
flushdb
geoadd
geodist
geohash
geopos
georadius
georadius_ro
georadiusbymember
georadiusbymember_ro
get
getbit
getrange
getset
handle_reply
hdel
hexists
hget
hgetall
hincrby
hincrbyfloat
hkeys
hlen
hmget
hmset
host:
host_
hscan
hset
hsetnx
hstrlen
hvals
incr
incrby
incrbyfloat
info
initialize_retry_vars
internal_send_command
keys
lastsave
latency
lindex
linsert
llen
lpop
lpush
lpushx
lrange
lrem
lset
ltrim
memory
mget
migrate
module
monitor
move
mset
msetnx
multi
object
on_connect
on_error
on_info_cmd
on_ready
persist
pexpire
pexpireat
pfadd
pfcount
pfdebug
pfmerge
pfselftest
ping
post
psetex
psubscribe
psync
pttl
publish
pubsub
punsubscribe
quit
randomkey
readonly
readwrite
ready_check
rename
renamenx
replconf
restore
restore-asking
restore_asking
return_error
return_reply
role
rpop
rpoplpush
rpush
rpushx
sadd
save
scan
scard
script
sdiff
sdiffstore
select
sendCommand
send_command
send_offline_queue
set
setbit
setex
setnx
setrange
shutdown
sinter
sinterstore
sismember
slaveof
slowlog
smembers
smove
sort
spop
srandmember
srem
sscan
strlen
subscribe
substr
sunion
sunionstore
swapdb
sync
time
touch
ttl
type
uncork
unlink
unref
unsubscribe
unwatch
wait
warn
watch
write
write_buffers
write_strings
xack
xadd
xclaim
xdel
xgroup
xinfo
xlen
xpending
xrange
xread
xreadgroup
xrevrange
xtrim
zadd
zcard
zcount
zincrby
zinterstore
zlexcount
zpopmax
zpopmin
zrange
zrangebylex
zrangebyscore
zrank
zrem
zremrangebylex
zremrangebyrank
zremrangebyscore
zrevrange
zrevrangebylex
zrevrangebyscore
zrevrank
zscan
zscore
zunionstore
*/
