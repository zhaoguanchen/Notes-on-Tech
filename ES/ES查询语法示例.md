GET /history_index.2019-02-28/dm_autoreport_app_business/_search


GET /bitauto_dm__dm_autoreport_app_business/2019-02-28/_search


GET /da_index__app_index_heat_serial_day/2019-02-21/_search
{
  "query": { 
    "term": { 
      "/da_index/app_index_heat_serial_day/serial_id": "carmodel_4707" 
    } 
  } 
}



GET /bitauto_dm__dm_autoreport_app_business/2019-02-28/_search
{
  "query": { 
    "term": { 
      "/bitauto_dm/dm_autoreport_app_business/product_type": "productType_10318" 
    } 
  } 
}

GET /da_index__app_index_heat_serial_day/2019-02-27/_search
{
  "query": { 
    "term": { 
      "/da_index/app_index_heat_serial_day/serial_id": "carmodel_4707" 
    } 
  } 
}



POST /history_index.2019-02-19/dm_autoreport_app_traffic_channel/_delete_by_query
{
  "query":{
    "match_all":{}
  }
}

POST /history_index.2019-02-28/dm_autoreport_app_business/_delete_by_query
{
  "query":{
    "match_all":{
      
    }
  }
}


