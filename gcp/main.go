package instance

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Instance(emailId, machineType, zone, image, sshKey string) {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewInstance(ctx, "ubuntu-instance", &compute.InstanceArgs{
			MachineType: pulumi.String(machineType),
			Zone:        pulumi.String(zone),
			Tags: pulumi.StringArray{
				pulumi.String("name"),
				pulumi.String("type"),
			},
			BootDisk: &compute.InstanceBootDiskArgs{
				InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String(image),
					// Type:  pulumi.String("SSD persistance disk"),
					// Size:  pulumi.Int(30),
				},
			},
			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
				&compute.InstanceNetworkInterfaceArgs{
					Network: pulumi.String("default"),
					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
						compute.InstanceNetworkInterfaceAccessConfigArgs{
							NetworkTier: pulumi.String("STANDARD"),
						},
					},
				},
			},
			Metadata: pulumi.StringMap{
				"name":     pulumi.String("Debian"),
				"type":     pulumi.String("Dev"),
				"ssh-keys": pulumi.String(sshKey),
			},
			ServiceAccount: &compute.InstanceServiceAccountArgs{
				Email: pulumi.String(emailId),
				Scopes: pulumi.StringArray{
					pulumi.String("cloud-platform"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
