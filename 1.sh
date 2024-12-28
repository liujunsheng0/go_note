[liujunsheng@1305.worker.livepush.tt.bjs.p1staff.com live_idlepush]$ cat data_cp_2_fallback.sh
#!/bin/bash

root_dir="/data/live_idlepush"
today_time=$(date "+%Y-%m-%d %H:%M:%S")
today_date=$(date "+%Y-%m-%d")
before_date=$(date -d "yesterday" +%Y-%m-%d)
echo "current date is $today_date, before date is $before_date"

fun_curl_cmd(){
  curl --location --request POST 'http://tsp-notification.vip.p1staff.com/notifications' \
  --header 'x-access-username: worker.livepush.tt' \
  --header 'x-access-token: zRFUCbouGB89Ba' \
  --header 'Content-Type: application/json' \
  -d '{
      "custom_key":"test_idle_push_dingding1",
      "messages":[
          {
              "type":"dingding",
              "body":{
                  "ding_text_type":"text",
                  "data":{
                      "content":"'$1'"
                  }
              }
          }
      ],
      "user_ldap_list":["liujunsheng","jiaxintian"]
  }'
}

function send_dingding(){
    time=$(date "+%Y-%m-%d_%H:%M:%S")
    fun_curl_cmd $1$time/ip10.100.54.93
    echo $1 , $time
}

en_size=`ls -l $root_dir/$before_date/en.online | awk '{print $5}'`
echo $en_size
en_file=$root_dir/$before_date/en.online
zh_file=$root_dir/$before_date/zh.online
msg_file=$root_dir/$before_date/msg.list
pk_file=$root_dir/$before_date/second_pk_result.online
user_file=$root_dir/$before_date/userInfo.csv
user_all_file=$root_dir/$before_date/userInfo_all.csv
fallback_file=$root_dir/fallback


test -s
if test -s $en_file && test -s $zh_file && test -s $msg_file && test -s $pk_file && test -s $user_file ; then
       echo "$before_date/files not empty , cp to fallback"

       echo $en_file
       cp -f $en_file $fallback_file || send_dingding "$user_file/cp_failed"
       echo "$en_file cp done"

       cp -f $zh_file $fallback_file || send_dingding "$user_file/cp_failed"
       echo "$zh_file cp done"

       cp -f $msg_file $fallback_file || send_dingding "$user_file/cp_failed"
       echo "$msg_file cp done"

       cp -f $pk_file $fallback_file || send_dingding "$user_file/cp_failed"
       echo "$pk_file cp done"

       cp -f $user_file $fallback_file || send_dingding "$user_file/cp_failed"
       echo "$user_file cp done"

       cp -f $user_all_file $fallback_file || send_dingding "$user_all_file/cp_failed"
       echo "$user_all_file cp done"
else
       echo "$root_dir/$before_date/files has empty"
       send_dingding "$root_dir/$before_date/files_has_empty_cp_fail"
fi

echo "date cp to fallback done"

curl --location --request POST '10.100.65.226:20792/v2/debug/idlepush/1/userId/397951647' || send_dingding "397951647|curl_failed"
curl --location --request POST '10.100.65.226:20792/v2/debug/idlepush/2/userId/397951647' || send_dingding "397951647|curl_failed"
curl --location --request POST '10.100.65.226:20792/v2/debug/idlepush/3/userId/397951647' || send_dingding "397951647|curl_failed"
curl --location --request POST '10.100.65.226:20792/v2/debug/idlepush/4/userId/397951647' || send_dingding "397951647|curl_failed"
curl --location --request POST '10.100.65.226:20792/v2/debug/idlepush/5/userId/397951647' || send_dingding "397951647|curl_failed"


[liujunsheng@1305.worker.livepush.tt.bjs.p1staff.com live_idlepush]$
[liujunsheng@1305.worker.livepush.tt.bjs.p1staff.com live_idlepush]$
[liujunsheng@1305.worker.livepush.tt.bjs.p1staff.com live_idlepush]$