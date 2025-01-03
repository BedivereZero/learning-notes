# GORM

涉及多张表的增删差改。

1. 一个应用案例包含多个应用场景
2. 一个应用场景引用多个资源
3. 一个资源可以被多个应用场景引用

数据库表 `examples` 记录应用案例。

| Field | Description    |
| :---- | :------------- |
| id    | 应用案例的 ID  |
| name  | 应用案例的名称 |

数据库表 `scenes` 记录应用场景。

| Field | Description    |
| :---- | :------------- |
| id    | 应用场景的 ID  |
| name  | 应用场景的名称 |

数据哭表 `scene_resource_bindings` 记录应用场景与资源的绑定关系。
