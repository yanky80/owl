# Bug: REGISTER 200 响应缺少 X-GB-Ver

## 现象

GB28181 下级盒端连接 GoWVP 上级后，注册和目录交互均能完成，但盒端状态显示：

- `registration=VERSION_MISSING`
- `version_missing=true`
- `version_mismatch=false`

## 现场证据

- GoWVP biz 实例：`v1.4.29`
- 盒端在 `2026-09-14 12:02:22` 记录：上级 REGISTER 响应缺少 `X-GB-Ver`，期望值为 `3.0`。
- 同一时间 GoWVP 记录设备注册成功，并收到/解析 Catalog 响应，目录返回 1 个通道。
- 未发现 `VERSION_MISMATCH`。

## 根因

REGISTER 的 `200 OK` 响应通过 `NewResponseFromRequest` 构造，但该函数只附加 `User-Agent`，没有附加 `X-GB-Ver: 3.0`。

## 预期

GoWVP 对支持 GB28181-2022 的下级返回 REGISTER 成功响应时，应携带：

```text
X-GB-Ver: 3.0
```

## 验收标准

- REGISTER `200 OK` 包含 `X-GB-Ver: 3.0`。
- 盒端 `version_missing=false`，且不影响 Catalog 交互。
- SIP 包相关测试通过。

## 修复结果

- `NewResponseFromRequest` 已补发 `X-GB-Ver: 3.0`。
- 新增回归测试 `TestNewResponseFromRequestIncludesGBVersion`。
- `go test ./pkg/gbs/...` 已通过。
- biz 测试环境已更新并重启验证：盒端 `registration=ONLINE`、`version_missing=false`、`version_mismatch=false`，Catalog 返回 1 个通道。
- 全量测试中的外部服务依赖失败已确认与本问题无关。
