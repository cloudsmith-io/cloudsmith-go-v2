# Plugin

What a Claude Code plugin bundles, for policy purposes.

`hooks`, `mcp_servers`, `lsp_servers`, `monitors` and `executables` run code
on the user's machine once the plugin is enabled. They are detected on a
best-effort basis, so a policy that must not allow code should match
`subtype == "plugin"` rather than rely on them being false.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Skills`           | []`string`         | :heavy_check_mark: | N/A                |
| `Agents`           | []`string`         | :heavy_check_mark: | N/A                |
| `Commands`         | []`string`         | :heavy_check_mark: | N/A                |
| `OutputStyles`     | []`string`         | :heavy_check_mark: | N/A                |
| `Workflows`        | []`string`         | :heavy_check_mark: | N/A                |
| `Hooks`            | `bool`             | :heavy_check_mark: | N/A                |
| `McpServers`       | `bool`             | :heavy_check_mark: | N/A                |
| `LspServers`       | `bool`             | :heavy_check_mark: | N/A                |
| `Monitors`         | `bool`             | :heavy_check_mark: | N/A                |
| `Executables`      | `bool`             | :heavy_check_mark: | N/A                |