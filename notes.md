task list

- when the keepalive fails only get a single notification, checks should not alerts
- expire checks that are not longer needed
- remove stashes that are no longer needed
- execute a check upon call

* when a check definition is removed from the client, the client should also remove any stashes or results from redis.

-- handler-check-cleaner --
* runs every 5 minutes
1. get a list of all checks currently active on the host (sensu-client -d /etc/sensu/conf.d/ -P)
2. get a list of all checks currently in redis for the client (curl -s -k -u admin:password https://localhost:4567/results/CLIENT)
3. compare them and obtain a list of checks in redis that are not configured locally
4. delete the check results from redis for checks that are not configured locally 
