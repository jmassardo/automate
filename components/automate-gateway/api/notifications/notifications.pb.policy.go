// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/notifications/notifications.proto

package notifications

import policy "github.com/chef/automate/components/automate-gateway/api/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/AddRule", "notifications:rules", "notifications:notifyRules:create", "POST", "/notifications/rules", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/DeleteRule", "notifications:rules:{id}", "notifications:notifyRules:delete", "DELETE", "/notifications/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*RuleIdentifier); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/UpdateRule", "notifications:rules:{id}", "notifications:notifyRules:update", "PUT", "/notifications/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*RuleUpdateRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/GetRule", "notifications:rules:{id}", "notifications:notifyRules:get", "GET", "/notifications/rules/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*RuleIdentifier); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/ListRules", "notifications:rules", "notifications:notifyRules:list", "GET", "/notifications/rules", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/ValidateWebhook", "notifications:rules", "notifications:notifyRules:validate", "POST", "/notifications/webhook", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*URLValidationRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "url":
					return m.Url
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.notifications.Notifications/Version", "system:service:version", "system:serviceVersion:get", "GET", "/notifications/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}