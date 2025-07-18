# FedMCP CLI

Command-line tool for creating, signing, and managing FedMCP artifacts.

## Installation

```bash
go install github.com/FedMCP/cli/cmd/fedmcp@latest
```

Or download pre-built binaries from the [releases page](https://github.com/FedMCP/cli/releases).

## Quick Start

```bash
# Initialize configuration
fedmcp config init

# Create an artifact
fedmcp create ssp-fragment --name "ML Pipeline Security" --file pipeline.yaml

# Sign the artifact
fedmcp sign artifact.json --key-file ~/.fedmcp/keys/signing.key

# Verify signature
fedmcp verify artifact.json

# Push to server
fedmcp push artifact.json
```

## Commands

### `fedmcp create`
Create new FedMCP artifacts of various types:
- `ssp-fragment` - System Security Plan fragments
- `poam-template` - Plan of Action & Milestones templates
- `agent-recipe` - Agent configuration recipes
- `baseline-module` - Security baseline modules
- `audit-script` - Audit automation scripts

### `fedmcp sign`
Sign artifacts using ECDSA P-256 with local keys or AWS KMS.

### `fedmcp verify`
Verify artifact signatures and integrity.

### `fedmcp push`
Upload signed artifacts to a FedMCP server.

### `fedmcp config`
Manage CLI configuration, workspaces, and defaults.

## Configuration

The CLI looks for configuration in `~/.fedmcp/config.yaml`:

```yaml
default_server: https://fedmcp.agency.gov
default_workspace: development

workspaces:
  development:
    server: https://dev.fedmcp.agency.gov
    key_file: ~/.fedmcp/keys/dev-signing.key
    
  production:
    server: https://fedmcp.agency.gov
    kms_key_id: arn:aws:kms:us-gov-west-1:123:key/prod-key
```

## Environment Variables

- `FEDMCP_CONFIG` - Path to config file
- `FEDMCP_DEFAULT_SERVER` - Default server URL
- `FEDMCP_DEFAULT_WORKSPACE` - Default workspace
- `FEDMCP_KEYS_DIRECTORY` - Directory for storing keys

## Examples

### Create and Sign SSP Fragment

```bash
# Create from YAML file
fedmcp create ssp-fragment \
  --name "ML Pipeline Security Controls" \
  --file controls.yaml

# Sign with local key
fedmcp sign ml-pipeline-ssp.json \
  --key-file ~/.fedmcp/keys/signing.key

# Push to server
fedmcp push ml-pipeline-ssp.json \
  --tag ml-pipeline \
  --tag production
```

### Verify Downloaded Artifact

```bash
# Download and verify
fedmcp verify artifact-abc123.json --verbose

# Verify from server
fedmcp verify \
  --artifact-id abc123 \
  --server https://fedmcp.agency.gov
```

### Workspace Management

```bash
# Add new workspace
fedmcp config add-workspace staging \
  --server https://staging.fedmcp.gov \
  --kms-key-id arn:aws:kms:us-gov-west-1:123:key/staging

# Switch workspace
fedmcp config set default-workspace staging

# List workspaces
fedmcp config list-workspaces
```

## License

Apache License 2.0 - See [LICENSE](https://github.com/FedMCP/cli/blob/main/LICENSE) for details.
