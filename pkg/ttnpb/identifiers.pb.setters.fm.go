// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ApplicationIdentifiers) SetFields(src *ApplicationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "application_id":
			if len(subs) > 0 {
				return fmt.Errorf("'application_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationID = src.ApplicationID
			} else {
				var zero string
				dst.ApplicationID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ClientIdentifiers) SetFields(src *ClientIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "client_id":
			if len(subs) > 0 {
				return fmt.Errorf("'client_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ClientId = src.ClientId
			} else {
				var zero string
				dst.ClientId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceIdentifiers) SetFields(src *EndDeviceIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "device_id":
			if len(subs) > 0 {
				return fmt.Errorf("'device_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeviceID = src.DeviceID
			} else {
				var zero string
				dst.DeviceID = zero
			}
		case "application_ids":
			if len(subs) > 0 {
				var newDst, newSrc *ApplicationIdentifiers
				if src != nil {
					newSrc = &src.ApplicationIdentifiers
				}
				newDst = &dst.ApplicationIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ApplicationIdentifiers = src.ApplicationIdentifiers
				} else {
					var zero ApplicationIdentifiers
					dst.ApplicationIdentifiers = zero
				}
			}
		case "dev_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevEUI = src.DevEUI
			} else {
				dst.DevEUI = nil
			}
		case "join_eui":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUI = src.JoinEUI
			} else {
				dst.JoinEUI = nil
			}
		case "dev_addr":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddr = src.DevAddr
			} else {
				dst.DevAddr = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GatewayIdentifiers) SetFields(src *GatewayIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_id":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayId = src.GatewayId
			} else {
				var zero string
				dst.GatewayId = zero
			}
		case "eui":
			if len(subs) > 0 {
				return fmt.Errorf("'eui' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.EUI = src.EUI
			} else {
				dst.EUI = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationIdentifiers) SetFields(src *OrganizationIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "organization_id":
			if len(subs) > 0 {
				return fmt.Errorf("'organization_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.OrganizationId = src.OrganizationId
			} else {
				var zero string
				dst.OrganizationId = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *UserIdentifiers) SetFields(src *UserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_id":
			if len(subs) > 0 {
				return fmt.Errorf("'user_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UserId = src.UserId
			} else {
				var zero string
				dst.UserId = zero
			}
		case "email":
			if len(subs) > 0 {
				return fmt.Errorf("'email' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Email = src.Email
			} else {
				var zero string
				dst.Email = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *OrganizationOrUserIdentifiers) SetFields(src *OrganizationOrUserIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "organization_ids":
					_, srcOk := src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds).OrganizationIds
						}
						if dstOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_OrganizationIds).OrganizationIds
						} else {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_OrganizationIds{OrganizationIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcOk := src.Ids.(*OrganizationOrUserIdentifiers_UserIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*OrganizationOrUserIdentifiers_UserIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*OrganizationOrUserIdentifiers_UserIds).UserIds
						}
						if dstOk {
							newDst = dst.Ids.(*OrganizationOrUserIdentifiers_UserIds).UserIds
						} else {
							newDst = &UserIdentifiers{}
							dst.Ids = &OrganizationOrUserIdentifiers_UserIds{UserIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EntityIdentifiers) SetFields(src *EntityIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {

		case "ids":
			if len(subs) == 0 && src == nil {
				dst.Ids = nil
				continue
			} else if len(subs) == 0 {
				dst.Ids = src.Ids
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "application_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_ApplicationIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_ApplicationIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'application_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ApplicationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_ApplicationIDs).ApplicationIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_ApplicationIDs).ApplicationIDs
						} else {
							newDst = &ApplicationIdentifiers{}
							dst.Ids = &EntityIdentifiers_ApplicationIDs{ApplicationIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "client_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_ClientIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_ClientIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'client_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *ClientIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_ClientIds).ClientIds
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_ClientIds).ClientIds
						} else {
							newDst = &ClientIdentifiers{}
							dst.Ids = &EntityIdentifiers_ClientIds{ClientIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "device_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_DeviceIDs)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_DeviceIDs)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'device_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *EndDeviceIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_DeviceIDs).DeviceIDs
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_DeviceIDs).DeviceIDs
						} else {
							newDst = &EndDeviceIdentifiers{}
							dst.Ids = &EntityIdentifiers_DeviceIDs{DeviceIDs: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "gateway_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_GatewayIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_GatewayIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'gateway_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *GatewayIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_GatewayIds).GatewayIds
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_GatewayIds).GatewayIds
						} else {
							newDst = &GatewayIdentifiers{}
							dst.Ids = &EntityIdentifiers_GatewayIds{GatewayIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "organization_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_OrganizationIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_OrganizationIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'organization_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *OrganizationIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_OrganizationIds).OrganizationIds
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_OrganizationIds).OrganizationIds
						} else {
							newDst = &OrganizationIdentifiers{}
							dst.Ids = &EntityIdentifiers_OrganizationIds{OrganizationIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}
				case "user_ids":
					_, srcOk := src.Ids.(*EntityIdentifiers_UserIds)
					if !srcOk && src.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in source")
					}
					_, dstOk := dst.Ids.(*EntityIdentifiers_UserIds)
					if !dstOk && dst.Ids != nil {
						return fmt.Errorf("attempt to set oneof 'user_ids', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *UserIdentifiers
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Ids.(*EntityIdentifiers_UserIds).UserIds
						}
						if dstOk {
							newDst = dst.Ids.(*EntityIdentifiers_UserIds).UserIds
						} else {
							newDst = &UserIdentifiers{}
							dst.Ids = &EntityIdentifiers_UserIds{UserIds: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Ids = src.Ids
						} else {
							dst.Ids = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *EndDeviceVersionIdentifiers) SetFields(src *EndDeviceVersionIdentifiers, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "brand_id":
			if len(subs) > 0 {
				return fmt.Errorf("'brand_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandID = src.BrandID
			} else {
				var zero string
				dst.BrandID = zero
			}
		case "model_id":
			if len(subs) > 0 {
				return fmt.Errorf("'model_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ModelID = src.ModelID
			} else {
				var zero string
				dst.ModelID = zero
			}
		case "hardware_version":
			if len(subs) > 0 {
				return fmt.Errorf("'hardware_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HardwareVersion = src.HardwareVersion
			} else {
				var zero string
				dst.HardwareVersion = zero
			}
		case "firmware_version":
			if len(subs) > 0 {
				return fmt.Errorf("'firmware_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FirmwareVersion = src.FirmwareVersion
			} else {
				var zero string
				dst.FirmwareVersion = zero
			}
		case "band_id":
			if len(subs) > 0 {
				return fmt.Errorf("'band_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BandID = src.BandID
			} else {
				var zero string
				dst.BandID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
