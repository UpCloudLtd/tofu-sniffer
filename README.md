# Tofu Sniffer

Sniff which tool was used to start a OpenTofu/Terraform Provider.

The OpenTofu/Terraform provider process is a child process of `tofu` or `terraform`. This allows using the parent processes PID to find out the name of the process (and thus tool) that started the provider server.

## Usage

The `tf.Sniff()` function can be used when configuring a provider to, for example, include the tool and its version in the user-agent value.

```go
func (p *SomeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	version := tf.Sniff(req.TerraformVersion)
	userAgent := fmt.Sprintf("terraform-provider-some/%s%s", p.version, version.UserAgentPostfix())
	// ...
}
```
