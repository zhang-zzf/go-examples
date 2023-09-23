# TODO

## go file servers


## UT

### mock

## RESTFul design

### categories

```http
GET     `/api/v1/categories`
```

### videos

```http
# 最热
GET     `/api/v1/videos?sort_by=-popularity&limit=24&offset=48`  
# 分类查询，按创建时间倒序排序
GET     `/api/v1/videos?tag=hot&sort_by=-created_at&limit=24&offset=48
# 分类查询，按热度倒序排序
GET     `/api/v1/videos?tag=hot&sort_by=-popularity&limit=24&offset=48
# 按 view_key 获取视频信息
GET     `/api/v1/videos/:view_key`  
# 按 title 模糊搜索
GET     `/api/v1/videos?title=abc`  

DELETE  `/api/v1/videos/:view_key`  
PUT     `/api/v1/videos/:view_key`  
POST    `/api/v1/videos`  
```

pagination resp

```json
{
    "data":[
        {
            "title": "titleOfTheVideo",
            "popularity": 35555,
        },
        {}
    ],
    "pagination":{
        "total": 72,
        "count": 24,
        "limit" 24,
        "offset": 24,
        "links": {
            "prev": "/api/v1/videos?tag=hot&sort_by=-created_at&limit=24&offset=0",
            "next": "/api/v1/videos?tag=hot&sort_by=-created_at&limit=24&offset=48", 
        }
    }
}
```

## ORM

### GORM

demo

### Pagination

### Cursor

```sql
select * from users where id > #{last_id} limit 24
```

### Cursor with Random Paging

not checked

`last_id = 254` 每页分页大小为 24，向后翻10页。

```sql
-- calc the new last_id
select max(id) from users where id > #{last_id} limit 240
select * from users where id > #{new_last_id} limit 24
```

## GIN
