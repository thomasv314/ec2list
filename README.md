## wtf is this?
i dunno it fetches ec2 instance lists, caches them locally, refreshes/filters by vpc/tag/role/etc, works w/ multiple profiles, and does helpful things that make life easier

## why?
i had a ruby task to do this but each execution took 2.26 seconds to run.it felt like forever and i couldnt stand to sit around waiting for it. i was upset. so then i wrote this and now it only takes 0.01 seconds to run and its fast so im happy (boom)

```
$ time bundle exec thor aws:list_instances # 2.264 seconds * 100 = 3.6 minutes
$ time ec2list # 0.011 seconds * 100 = 1.1 seconds
```

## install it
```
git clone git@github.com:thomasv314/ec2list.git
cd ec2list
go build -o ec2list *.go
mv ec2list /somewhere/in/your/path/bin
ec2list --setup
```

## use it
```
ec2list --setup         # create a cache directory
ec2list                 # list instances in a table w/ role/tags/name/vpc/ips/id/etc
ec2list -f              # filter table by key:val where key = (vpc|role|name)
ec2list -s              # instead of a table output ips like: "<ip1> <ip2> ..."
```

## examples
```
# show instances in `my-vpc`
ec2list -f vpc:my-vpc

# show instances w/ role of `app`
ec2list -f role:app

# show instances w/ role of `app` in `my-vpc` vpc
ec2list -f role:app vpc:my-vpc

# show only private ips of `apps` in `my-vpc`
ec2list -f role:app vpc:my-vpc -s

# list all elasticsearch instances in an ec2 role
ec2list -f vpc:my-vpc-name role:elasticsearch

# start a cluster SSH session w/ your ec2list query
csshX --login ec2-user $(eplist -s -f vpc:my-vpc role:app)
```

## multiple profiles

Use [direnv](http://direnv.net/) to set the following
ENV vars in different directories:

```
# in ~/my-profile/.envrc
export ACTIVE_EP="my-profile-name"
export AWS_ACCESS_KEY_ID="profile-key-id"
export AWS_SECRET_ACCESS_KEY="kldjafs"


# in ~/other-client/.envrc:
export ACTIVE_EP="other-profile-name"
export AWS_ACCESS_KEY_ID="other-profile-key-id"
export AWS_SECRET_ACCESS_KEY="342938"
```

Switching into either of those directories will enable ec2list to
use the proper instance cache as long as direnv is installed.

## todo:

- easy ssh sockets to boxes
- bust cache every N minutes
- ssh gateway config stuffs
