# cronoscopoe
This small utility allows to run other programs inside docker while sending their metrics to prometheus pushergateway. This is especially useful for running cron jobs on clusters such as Amazon ECS or Kubernetes. Suppose you have a program that runs like this:

```
python main.py --retries 3 --user-id 43
```

To be able to push metrics to prometheus gateway without touching the code, you can use invoke it like this:

```
$ cronoscope python main.py --retires 2 --user-id 23
```

You can also use it with cronolock with cronolock being the first argument and everything else will become the arguments to the cronoscope.

Please note that the `python` program must be in the serach `PATH` and if not, please specify the full path like this:

```
$ cronoscope /usr/bin/python main.py --retires 2 --user-id 23

```

From above, the main.py is the program to be executed whereas rest are the 

### Confguration
The cronoscope is totally configured via environment variables and there are no configuration files. Following are the environment variables:

* `CRONOSCOPE_PUSHERGATEWAY_HOST` requrid. Just the host name without `http` or `https` and currently only works with `http` 
* `CRONOSCOPE_PUSHERGATEWAY_PORT` optional, defaults to `9091`
* `CRONOSCOPE_LABEL_JOB` required 
* `CRONOSCOPE_POLLING_INTERVAL`  optional, defaults to 10 seconds
* `CRONOSCOPE_PUSH_RETRIES` optional, number of times cronoscope will try to push metrics to the pushergateway in case of failure. Defaults to 3 retries.
* `CRONOSCOPE_PUSHRETRIESINTERVAL` optional, intervals between retries. Defaults to 2 seconds.
* `CRONOSCOPE_DISABLED` just runs the given program and no monitoring agent is starated. Useful in local or test enivronments.
* `CRONOSCOPE_METRICS_PREFIX` the prefix to be added to the generated metrics. Defaults to `crons_`

### Passing additional labels
If you want to pass additional lables along with the metrics, youc set any number of them starting with `CRONOSCOPE_LABEL_.*` pattern and all of them would end up as labels. As 
an example, if you have following environment varibles:

```
export CRONOSCOPE_LABEL_ENVIRONMENT="stage"
export CRONOSCOPE_LABEL_JOB="my-job"
export CRONOSCOPE_LABEL_INSTANCE="my-instance"
export CRONOSCOPE_LABEL_CLUSTER="my-production-cluster"
```

Then the generated metrics would be like this:

```
# HELP crons_memory_usage_in_bytes Current usage for memory
# TYPE crons_memory_usage_in_bytes gauge
crons_memory_usage_in_bytes{cluster="my-production-cluster",environment="stage",instance="my-instance,job="my-job"} 7.040380928e+09
```