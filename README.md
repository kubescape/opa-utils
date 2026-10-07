# OPA Utilities package

## Framework-scoped result views

`exceptions.FilterExceptionsByFrameworks` filters posture exceptions with the
same case-insensitive and anchored-regexp matching rules used by the exception
processor. Returned exceptions are copies whose `PosturePolicies` contain only
the matching `(framework, control, rule)` tuples, so callers can derive an
effective framework status without mutating the shared evaluation result. An
empty framework-name slice selects framework-agnostic policies.

Result and summary `GetStatus` methods carry the effective substatus through
rule, control, resource, framework, and top-level views. Summary substatuses
are derived from serialized control status information, so this does not add a
new report field.

### Scan request control exclusions

`httpserver/meta/v1.PostScanRequest` accepts optional `excludeControls`, a list of
control IDs (or legacy CIS section numbers), for example `["C-0069", "C-0070"]`.
Scanners supporting this field omit matching controls before evaluation and
compliance scoring, using the CLI's `--exclude-controls` behavior. Omitting the
field or using an empty list preserves the selected assessment scope. Exclusions
narrow that scope; they do not establish that the excluded checks would pass.
