// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var UplinkMessageFieldPathsNested = []string{
	"consumed_airtime",
	"correlation_ids",
	"device_channel_index",
	"payload",
	"payload.Payload",
	"payload.Payload.join_accept_payload",
	"payload.Payload.join_accept_payload.cf_list",
	"payload.Payload.join_accept_payload.cf_list.ch_masks",
	"payload.Payload.join_accept_payload.cf_list.freq",
	"payload.Payload.join_accept_payload.cf_list.type",
	"payload.Payload.join_accept_payload.dev_addr",
	"payload.Payload.join_accept_payload.dl_settings",
	"payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"payload.Payload.join_accept_payload.encrypted",
	"payload.Payload.join_accept_payload.join_nonce",
	"payload.Payload.join_accept_payload.net_id",
	"payload.Payload.join_accept_payload.rx_delay",
	"payload.Payload.join_request_payload",
	"payload.Payload.join_request_payload.dev_eui",
	"payload.Payload.join_request_payload.dev_nonce",
	"payload.Payload.join_request_payload.join_eui",
	"payload.Payload.mac_payload",
	"payload.Payload.mac_payload.decoded_payload",
	"payload.Payload.mac_payload.f_hdr",
	"payload.Payload.mac_payload.f_hdr.dev_addr",
	"payload.Payload.mac_payload.f_hdr.f_cnt",
	"payload.Payload.mac_payload.f_hdr.f_ctrl",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"payload.Payload.mac_payload.f_hdr.f_opts",
	"payload.Payload.mac_payload.f_port",
	"payload.Payload.mac_payload.frm_payload",
	"payload.Payload.mac_payload.full_f_cnt",
	"payload.Payload.rejoin_request_payload",
	"payload.Payload.rejoin_request_payload.dev_eui",
	"payload.Payload.rejoin_request_payload.join_eui",
	"payload.Payload.rejoin_request_payload.net_id",
	"payload.Payload.rejoin_request_payload.rejoin_cnt",
	"payload.Payload.rejoin_request_payload.rejoin_type",
	"payload.m_hdr",
	"payload.m_hdr.m_type",
	"payload.m_hdr.major",
	"payload.mic",
	"raw_payload",
	"received_at",
	"rx_metadata",
	"settings",
	"settings.coding_rate",
	"settings.data_rate",
	"settings.data_rate.modulation",
	"settings.data_rate.modulation.fsk",
	"settings.data_rate.modulation.fsk.bit_rate",
	"settings.data_rate.modulation.lora",
	"settings.data_rate.modulation.lora.bandwidth",
	"settings.data_rate.modulation.lora.spreading_factor",
	"settings.data_rate_index",
	"settings.downlink",
	"settings.downlink.antenna_index",
	"settings.downlink.invert_polarization",
	"settings.downlink.tx_power",
	"settings.enable_crc",
	"settings.frequency",
	"settings.time",
	"settings.timestamp",
}

var UplinkMessageFieldPathsTopLevel = []string{
	"consumed_airtime",
	"correlation_ids",
	"device_channel_index",
	"payload",
	"raw_payload",
	"received_at",
	"rx_metadata",
	"settings",
}
var DownlinkMessageFieldPathsNested = []string{
	"correlation_ids",
	"end_device_ids",
	"end_device_ids.application_ids",
	"end_device_ids.application_ids.application_id",
	"end_device_ids.dev_addr",
	"end_device_ids.dev_eui",
	"end_device_ids.device_id",
	"end_device_ids.join_eui",
	"payload",
	"payload.Payload",
	"payload.Payload.join_accept_payload",
	"payload.Payload.join_accept_payload.cf_list",
	"payload.Payload.join_accept_payload.cf_list.ch_masks",
	"payload.Payload.join_accept_payload.cf_list.freq",
	"payload.Payload.join_accept_payload.cf_list.type",
	"payload.Payload.join_accept_payload.dev_addr",
	"payload.Payload.join_accept_payload.dl_settings",
	"payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"payload.Payload.join_accept_payload.encrypted",
	"payload.Payload.join_accept_payload.join_nonce",
	"payload.Payload.join_accept_payload.net_id",
	"payload.Payload.join_accept_payload.rx_delay",
	"payload.Payload.join_request_payload",
	"payload.Payload.join_request_payload.dev_eui",
	"payload.Payload.join_request_payload.dev_nonce",
	"payload.Payload.join_request_payload.join_eui",
	"payload.Payload.mac_payload",
	"payload.Payload.mac_payload.decoded_payload",
	"payload.Payload.mac_payload.f_hdr",
	"payload.Payload.mac_payload.f_hdr.dev_addr",
	"payload.Payload.mac_payload.f_hdr.f_cnt",
	"payload.Payload.mac_payload.f_hdr.f_ctrl",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"payload.Payload.mac_payload.f_hdr.f_opts",
	"payload.Payload.mac_payload.f_port",
	"payload.Payload.mac_payload.frm_payload",
	"payload.Payload.mac_payload.full_f_cnt",
	"payload.Payload.rejoin_request_payload",
	"payload.Payload.rejoin_request_payload.dev_eui",
	"payload.Payload.rejoin_request_payload.join_eui",
	"payload.Payload.rejoin_request_payload.net_id",
	"payload.Payload.rejoin_request_payload.rejoin_cnt",
	"payload.Payload.rejoin_request_payload.rejoin_type",
	"payload.m_hdr",
	"payload.m_hdr.m_type",
	"payload.m_hdr.major",
	"payload.mic",
	"raw_payload",
	"settings",
	"settings.request",
	"settings.request.absolute_time",
	"settings.request.advanced",
	"settings.request.class",
	"settings.request.downlink_paths",
	"settings.request.frequency_plan_id",
	"settings.request.priority",
	"settings.request.rx1_data_rate_index",
	"settings.request.rx1_delay",
	"settings.request.rx1_frequency",
	"settings.request.rx2_data_rate_index",
	"settings.request.rx2_frequency",
	"settings.scheduled",
	"settings.scheduled.coding_rate",
	"settings.scheduled.data_rate",
	"settings.scheduled.data_rate.modulation",
	"settings.scheduled.data_rate.modulation.fsk",
	"settings.scheduled.data_rate.modulation.fsk.bit_rate",
	"settings.scheduled.data_rate.modulation.lora",
	"settings.scheduled.data_rate.modulation.lora.bandwidth",
	"settings.scheduled.data_rate.modulation.lora.spreading_factor",
	"settings.scheduled.data_rate_index",
	"settings.scheduled.downlink",
	"settings.scheduled.downlink.antenna_index",
	"settings.scheduled.downlink.invert_polarization",
	"settings.scheduled.downlink.tx_power",
	"settings.scheduled.enable_crc",
	"settings.scheduled.frequency",
	"settings.scheduled.time",
	"settings.scheduled.timestamp",
}

var DownlinkMessageFieldPathsTopLevel = []string{
	"correlation_ids",
	"end_device_ids",
	"payload",
	"raw_payload",
	"settings",
}
var TxAcknowledgmentFieldPathsNested = []string{
	"correlation_ids",
	"downlink_message",
	"downlink_message.correlation_ids",
	"downlink_message.end_device_ids",
	"downlink_message.end_device_ids.application_ids",
	"downlink_message.end_device_ids.application_ids.application_id",
	"downlink_message.end_device_ids.dev_addr",
	"downlink_message.end_device_ids.dev_eui",
	"downlink_message.end_device_ids.device_id",
	"downlink_message.end_device_ids.join_eui",
	"downlink_message.payload",
	"downlink_message.payload.Payload",
	"downlink_message.payload.Payload.join_accept_payload",
	"downlink_message.payload.Payload.join_accept_payload.cf_list",
	"downlink_message.payload.Payload.join_accept_payload.cf_list.ch_masks",
	"downlink_message.payload.Payload.join_accept_payload.cf_list.freq",
	"downlink_message.payload.Payload.join_accept_payload.cf_list.type",
	"downlink_message.payload.Payload.join_accept_payload.dev_addr",
	"downlink_message.payload.Payload.join_accept_payload.dl_settings",
	"downlink_message.payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"downlink_message.payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"downlink_message.payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"downlink_message.payload.Payload.join_accept_payload.encrypted",
	"downlink_message.payload.Payload.join_accept_payload.join_nonce",
	"downlink_message.payload.Payload.join_accept_payload.net_id",
	"downlink_message.payload.Payload.join_accept_payload.rx_delay",
	"downlink_message.payload.Payload.join_request_payload",
	"downlink_message.payload.Payload.join_request_payload.dev_eui",
	"downlink_message.payload.Payload.join_request_payload.dev_nonce",
	"downlink_message.payload.Payload.join_request_payload.join_eui",
	"downlink_message.payload.Payload.mac_payload",
	"downlink_message.payload.Payload.mac_payload.decoded_payload",
	"downlink_message.payload.Payload.mac_payload.f_hdr",
	"downlink_message.payload.Payload.mac_payload.f_hdr.dev_addr",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_cnt",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"downlink_message.payload.Payload.mac_payload.f_hdr.f_opts",
	"downlink_message.payload.Payload.mac_payload.f_port",
	"downlink_message.payload.Payload.mac_payload.frm_payload",
	"downlink_message.payload.Payload.mac_payload.full_f_cnt",
	"downlink_message.payload.Payload.rejoin_request_payload",
	"downlink_message.payload.Payload.rejoin_request_payload.dev_eui",
	"downlink_message.payload.Payload.rejoin_request_payload.join_eui",
	"downlink_message.payload.Payload.rejoin_request_payload.net_id",
	"downlink_message.payload.Payload.rejoin_request_payload.rejoin_cnt",
	"downlink_message.payload.Payload.rejoin_request_payload.rejoin_type",
	"downlink_message.payload.m_hdr",
	"downlink_message.payload.m_hdr.m_type",
	"downlink_message.payload.m_hdr.major",
	"downlink_message.payload.mic",
	"downlink_message.raw_payload",
	"downlink_message.settings",
	"downlink_message.settings.request",
	"downlink_message.settings.request.absolute_time",
	"downlink_message.settings.request.advanced",
	"downlink_message.settings.request.class",
	"downlink_message.settings.request.downlink_paths",
	"downlink_message.settings.request.frequency_plan_id",
	"downlink_message.settings.request.priority",
	"downlink_message.settings.request.rx1_data_rate_index",
	"downlink_message.settings.request.rx1_delay",
	"downlink_message.settings.request.rx1_frequency",
	"downlink_message.settings.request.rx2_data_rate_index",
	"downlink_message.settings.request.rx2_frequency",
	"downlink_message.settings.scheduled",
	"downlink_message.settings.scheduled.coding_rate",
	"downlink_message.settings.scheduled.data_rate",
	"downlink_message.settings.scheduled.data_rate.modulation",
	"downlink_message.settings.scheduled.data_rate.modulation.fsk",
	"downlink_message.settings.scheduled.data_rate.modulation.fsk.bit_rate",
	"downlink_message.settings.scheduled.data_rate.modulation.lora",
	"downlink_message.settings.scheduled.data_rate.modulation.lora.bandwidth",
	"downlink_message.settings.scheduled.data_rate.modulation.lora.spreading_factor",
	"downlink_message.settings.scheduled.data_rate_index",
	"downlink_message.settings.scheduled.downlink",
	"downlink_message.settings.scheduled.downlink.antenna_index",
	"downlink_message.settings.scheduled.downlink.invert_polarization",
	"downlink_message.settings.scheduled.downlink.tx_power",
	"downlink_message.settings.scheduled.enable_crc",
	"downlink_message.settings.scheduled.frequency",
	"downlink_message.settings.scheduled.time",
	"downlink_message.settings.scheduled.timestamp",
	"result",
}

var TxAcknowledgmentFieldPathsTopLevel = []string{
	"correlation_ids",
	"downlink_message",
	"result",
}
var GatewayTxAcknowledgmentFieldPathsNested = []string{
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"tx_ack",
	"tx_ack.correlation_ids",
	"tx_ack.downlink_message",
	"tx_ack.downlink_message.correlation_ids",
	"tx_ack.downlink_message.end_device_ids",
	"tx_ack.downlink_message.end_device_ids.application_ids",
	"tx_ack.downlink_message.end_device_ids.application_ids.application_id",
	"tx_ack.downlink_message.end_device_ids.dev_addr",
	"tx_ack.downlink_message.end_device_ids.dev_eui",
	"tx_ack.downlink_message.end_device_ids.device_id",
	"tx_ack.downlink_message.end_device_ids.join_eui",
	"tx_ack.downlink_message.payload",
	"tx_ack.downlink_message.payload.Payload",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.cf_list",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.cf_list.ch_masks",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.cf_list.freq",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.cf_list.type",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.dev_addr",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.dl_settings",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.encrypted",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.join_nonce",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.net_id",
	"tx_ack.downlink_message.payload.Payload.join_accept_payload.rx_delay",
	"tx_ack.downlink_message.payload.Payload.join_request_payload",
	"tx_ack.downlink_message.payload.Payload.join_request_payload.dev_eui",
	"tx_ack.downlink_message.payload.Payload.join_request_payload.dev_nonce",
	"tx_ack.downlink_message.payload.Payload.join_request_payload.join_eui",
	"tx_ack.downlink_message.payload.Payload.mac_payload",
	"tx_ack.downlink_message.payload.Payload.mac_payload.decoded_payload",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.dev_addr",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_cnt",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_hdr.f_opts",
	"tx_ack.downlink_message.payload.Payload.mac_payload.f_port",
	"tx_ack.downlink_message.payload.Payload.mac_payload.frm_payload",
	"tx_ack.downlink_message.payload.Payload.mac_payload.full_f_cnt",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload.dev_eui",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload.join_eui",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload.net_id",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload.rejoin_cnt",
	"tx_ack.downlink_message.payload.Payload.rejoin_request_payload.rejoin_type",
	"tx_ack.downlink_message.payload.m_hdr",
	"tx_ack.downlink_message.payload.m_hdr.m_type",
	"tx_ack.downlink_message.payload.m_hdr.major",
	"tx_ack.downlink_message.payload.mic",
	"tx_ack.downlink_message.raw_payload",
	"tx_ack.downlink_message.settings",
	"tx_ack.downlink_message.settings.request",
	"tx_ack.downlink_message.settings.request.absolute_time",
	"tx_ack.downlink_message.settings.request.advanced",
	"tx_ack.downlink_message.settings.request.class",
	"tx_ack.downlink_message.settings.request.downlink_paths",
	"tx_ack.downlink_message.settings.request.frequency_plan_id",
	"tx_ack.downlink_message.settings.request.priority",
	"tx_ack.downlink_message.settings.request.rx1_data_rate_index",
	"tx_ack.downlink_message.settings.request.rx1_delay",
	"tx_ack.downlink_message.settings.request.rx1_frequency",
	"tx_ack.downlink_message.settings.request.rx2_data_rate_index",
	"tx_ack.downlink_message.settings.request.rx2_frequency",
	"tx_ack.downlink_message.settings.scheduled",
	"tx_ack.downlink_message.settings.scheduled.coding_rate",
	"tx_ack.downlink_message.settings.scheduled.data_rate",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation.fsk",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation.fsk.bit_rate",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation.lora",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation.lora.bandwidth",
	"tx_ack.downlink_message.settings.scheduled.data_rate.modulation.lora.spreading_factor",
	"tx_ack.downlink_message.settings.scheduled.data_rate_index",
	"tx_ack.downlink_message.settings.scheduled.downlink",
	"tx_ack.downlink_message.settings.scheduled.downlink.antenna_index",
	"tx_ack.downlink_message.settings.scheduled.downlink.invert_polarization",
	"tx_ack.downlink_message.settings.scheduled.downlink.tx_power",
	"tx_ack.downlink_message.settings.scheduled.enable_crc",
	"tx_ack.downlink_message.settings.scheduled.frequency",
	"tx_ack.downlink_message.settings.scheduled.time",
	"tx_ack.downlink_message.settings.scheduled.timestamp",
	"tx_ack.result",
}

var GatewayTxAcknowledgmentFieldPathsTopLevel = []string{
	"gateway_ids",
	"tx_ack",
}
var GatewayUplinkMessageFieldPathsNested = []string{
	"band_id",
	"message",
	"message.consumed_airtime",
	"message.correlation_ids",
	"message.device_channel_index",
	"message.payload",
	"message.payload.Payload",
	"message.payload.Payload.join_accept_payload",
	"message.payload.Payload.join_accept_payload.cf_list",
	"message.payload.Payload.join_accept_payload.cf_list.ch_masks",
	"message.payload.Payload.join_accept_payload.cf_list.freq",
	"message.payload.Payload.join_accept_payload.cf_list.type",
	"message.payload.Payload.join_accept_payload.dev_addr",
	"message.payload.Payload.join_accept_payload.dl_settings",
	"message.payload.Payload.join_accept_payload.dl_settings.opt_neg",
	"message.payload.Payload.join_accept_payload.dl_settings.rx1_dr_offset",
	"message.payload.Payload.join_accept_payload.dl_settings.rx2_dr",
	"message.payload.Payload.join_accept_payload.encrypted",
	"message.payload.Payload.join_accept_payload.join_nonce",
	"message.payload.Payload.join_accept_payload.net_id",
	"message.payload.Payload.join_accept_payload.rx_delay",
	"message.payload.Payload.join_request_payload",
	"message.payload.Payload.join_request_payload.dev_eui",
	"message.payload.Payload.join_request_payload.dev_nonce",
	"message.payload.Payload.join_request_payload.join_eui",
	"message.payload.Payload.mac_payload",
	"message.payload.Payload.mac_payload.decoded_payload",
	"message.payload.Payload.mac_payload.f_hdr",
	"message.payload.Payload.mac_payload.f_hdr.dev_addr",
	"message.payload.Payload.mac_payload.f_hdr.f_cnt",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl.ack",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl.adr_ack_req",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl.class_b",
	"message.payload.Payload.mac_payload.f_hdr.f_ctrl.f_pending",
	"message.payload.Payload.mac_payload.f_hdr.f_opts",
	"message.payload.Payload.mac_payload.f_port",
	"message.payload.Payload.mac_payload.frm_payload",
	"message.payload.Payload.mac_payload.full_f_cnt",
	"message.payload.Payload.rejoin_request_payload",
	"message.payload.Payload.rejoin_request_payload.dev_eui",
	"message.payload.Payload.rejoin_request_payload.join_eui",
	"message.payload.Payload.rejoin_request_payload.net_id",
	"message.payload.Payload.rejoin_request_payload.rejoin_cnt",
	"message.payload.Payload.rejoin_request_payload.rejoin_type",
	"message.payload.m_hdr",
	"message.payload.m_hdr.m_type",
	"message.payload.m_hdr.major",
	"message.payload.mic",
	"message.raw_payload",
	"message.received_at",
	"message.rx_metadata",
	"message.settings",
	"message.settings.coding_rate",
	"message.settings.data_rate",
	"message.settings.data_rate.modulation",
	"message.settings.data_rate.modulation.fsk",
	"message.settings.data_rate.modulation.fsk.bit_rate",
	"message.settings.data_rate.modulation.lora",
	"message.settings.data_rate.modulation.lora.bandwidth",
	"message.settings.data_rate.modulation.lora.spreading_factor",
	"message.settings.data_rate_index",
	"message.settings.downlink",
	"message.settings.downlink.antenna_index",
	"message.settings.downlink.invert_polarization",
	"message.settings.downlink.tx_power",
	"message.settings.enable_crc",
	"message.settings.frequency",
	"message.settings.time",
	"message.settings.timestamp",
}

var GatewayUplinkMessageFieldPathsTopLevel = []string{
	"band_id",
	"message",
}
var ApplicationUplinkFieldPathsNested = []string{
	"app_s_key",
	"app_s_key.encrypted_key",
	"app_s_key.kek_label",
	"app_s_key.key",
	"confirmed",
	"consumed_airtime",
	"decoded_payload",
	"decoded_payload_warnings",
	"f_cnt",
	"f_port",
	"frm_payload",
	"last_a_f_cnt_down",
	"locations",
	"received_at",
	"rx_metadata",
	"session_key_id",
	"settings",
	"settings.coding_rate",
	"settings.data_rate",
	"settings.data_rate.modulation",
	"settings.data_rate.modulation.fsk",
	"settings.data_rate.modulation.fsk.bit_rate",
	"settings.data_rate.modulation.lora",
	"settings.data_rate.modulation.lora.bandwidth",
	"settings.data_rate.modulation.lora.spreading_factor",
	"settings.data_rate_index",
	"settings.downlink",
	"settings.downlink.antenna_index",
	"settings.downlink.invert_polarization",
	"settings.downlink.tx_power",
	"settings.enable_crc",
	"settings.frequency",
	"settings.time",
	"settings.timestamp",
	"version_ids",
	"version_ids.band_id",
	"version_ids.brand_id",
	"version_ids.firmware_version",
	"version_ids.hardware_version",
	"version_ids.model_id",
}

var ApplicationUplinkFieldPathsTopLevel = []string{
	"app_s_key",
	"confirmed",
	"consumed_airtime",
	"decoded_payload",
	"decoded_payload_warnings",
	"f_cnt",
	"f_port",
	"frm_payload",
	"last_a_f_cnt_down",
	"locations",
	"received_at",
	"rx_metadata",
	"session_key_id",
	"settings",
	"version_ids",
}
var ApplicationLocationFieldPathsNested = []string{
	"attributes",
	"location",
	"location.accuracy",
	"location.altitude",
	"location.latitude",
	"location.longitude",
	"location.source",
	"service",
}

var ApplicationLocationFieldPathsTopLevel = []string{
	"attributes",
	"location",
	"service",
}
var ApplicationJoinAcceptFieldPathsNested = []string{
	"app_s_key",
	"app_s_key.encrypted_key",
	"app_s_key.kek_label",
	"app_s_key.key",
	"invalidated_downlinks",
	"pending_session",
	"received_at",
	"session_key_id",
}

var ApplicationJoinAcceptFieldPathsTopLevel = []string{
	"app_s_key",
	"invalidated_downlinks",
	"pending_session",
	"received_at",
	"session_key_id",
}
var ApplicationDownlinkFieldPathsNested = []string{
	"class_b_c",
	"class_b_c.absolute_time",
	"class_b_c.gateways",
	"confirmed",
	"correlation_ids",
	"decoded_payload",
	"decoded_payload_warnings",
	"f_cnt",
	"f_port",
	"frm_payload",
	"priority",
	"session_key_id",
}

var ApplicationDownlinkFieldPathsTopLevel = []string{
	"class_b_c",
	"confirmed",
	"correlation_ids",
	"decoded_payload",
	"decoded_payload_warnings",
	"f_cnt",
	"f_port",
	"frm_payload",
	"priority",
	"session_key_id",
}
var ApplicationDownlinksFieldPathsNested = []string{
	"downlinks",
}

var ApplicationDownlinksFieldPathsTopLevel = []string{
	"downlinks",
}
var ApplicationDownlinkFailedFieldPathsNested = []string{
	"downlink",
	"downlink.class_b_c",
	"downlink.class_b_c.absolute_time",
	"downlink.class_b_c.gateways",
	"downlink.confirmed",
	"downlink.correlation_ids",
	"downlink.decoded_payload",
	"downlink.decoded_payload_warnings",
	"downlink.f_cnt",
	"downlink.f_port",
	"downlink.frm_payload",
	"downlink.priority",
	"downlink.session_key_id",
	"error",
	"error.attributes",
	"error.cause",
	"error.cause.attributes",
	"error.cause.correlation_id",
	"error.cause.message_format",
	"error.cause.name",
	"error.cause.namespace",
	"error.code",
	"error.correlation_id",
	"error.details",
	"error.message_format",
	"error.name",
	"error.namespace",
}

var ApplicationDownlinkFailedFieldPathsTopLevel = []string{
	"downlink",
	"error",
}
var ApplicationInvalidatedDownlinksFieldPathsNested = []string{
	"downlinks",
	"last_f_cnt_down",
	"session_key_id",
}

var ApplicationInvalidatedDownlinksFieldPathsTopLevel = []string{
	"downlinks",
	"last_f_cnt_down",
	"session_key_id",
}
var DownlinkQueueOperationErrorDetailsFieldPathsNested = []string{
	"dev_addr",
	"min_f_cnt_down",
	"pending_dev_addr",
	"pending_min_f_cnt_down",
	"pending_session_key_id",
	"session_key_id",
}

var DownlinkQueueOperationErrorDetailsFieldPathsTopLevel = []string{
	"dev_addr",
	"min_f_cnt_down",
	"pending_dev_addr",
	"pending_min_f_cnt_down",
	"pending_session_key_id",
	"session_key_id",
}
var ApplicationServiceDataFieldPathsNested = []string{
	"data",
	"service",
}

var ApplicationServiceDataFieldPathsTopLevel = []string{
	"data",
	"service",
}
var ApplicationUpFieldPathsNested = []string{
	"correlation_ids",
	"end_device_ids",
	"end_device_ids.application_ids",
	"end_device_ids.application_ids.application_id",
	"end_device_ids.dev_addr",
	"end_device_ids.dev_eui",
	"end_device_ids.device_id",
	"end_device_ids.join_eui",
	"received_at",
	"simulated",
	"up",
	"up.downlink_ack",
	"up.downlink_ack.class_b_c",
	"up.downlink_ack.class_b_c.absolute_time",
	"up.downlink_ack.class_b_c.gateways",
	"up.downlink_ack.confirmed",
	"up.downlink_ack.correlation_ids",
	"up.downlink_ack.decoded_payload",
	"up.downlink_ack.decoded_payload_warnings",
	"up.downlink_ack.f_cnt",
	"up.downlink_ack.f_port",
	"up.downlink_ack.frm_payload",
	"up.downlink_ack.priority",
	"up.downlink_ack.session_key_id",
	"up.downlink_failed",
	"up.downlink_failed.downlink",
	"up.downlink_failed.downlink.class_b_c",
	"up.downlink_failed.downlink.class_b_c.absolute_time",
	"up.downlink_failed.downlink.class_b_c.gateways",
	"up.downlink_failed.downlink.confirmed",
	"up.downlink_failed.downlink.correlation_ids",
	"up.downlink_failed.downlink.decoded_payload",
	"up.downlink_failed.downlink.decoded_payload_warnings",
	"up.downlink_failed.downlink.f_cnt",
	"up.downlink_failed.downlink.f_port",
	"up.downlink_failed.downlink.frm_payload",
	"up.downlink_failed.downlink.priority",
	"up.downlink_failed.downlink.session_key_id",
	"up.downlink_failed.error",
	"up.downlink_failed.error.attributes",
	"up.downlink_failed.error.cause",
	"up.downlink_failed.error.cause.attributes",
	"up.downlink_failed.error.cause.correlation_id",
	"up.downlink_failed.error.cause.message_format",
	"up.downlink_failed.error.cause.name",
	"up.downlink_failed.error.cause.namespace",
	"up.downlink_failed.error.code",
	"up.downlink_failed.error.correlation_id",
	"up.downlink_failed.error.details",
	"up.downlink_failed.error.message_format",
	"up.downlink_failed.error.name",
	"up.downlink_failed.error.namespace",
	"up.downlink_nack",
	"up.downlink_nack.class_b_c",
	"up.downlink_nack.class_b_c.absolute_time",
	"up.downlink_nack.class_b_c.gateways",
	"up.downlink_nack.confirmed",
	"up.downlink_nack.correlation_ids",
	"up.downlink_nack.decoded_payload",
	"up.downlink_nack.decoded_payload_warnings",
	"up.downlink_nack.f_cnt",
	"up.downlink_nack.f_port",
	"up.downlink_nack.frm_payload",
	"up.downlink_nack.priority",
	"up.downlink_nack.session_key_id",
	"up.downlink_queue_invalidated",
	"up.downlink_queue_invalidated.downlinks",
	"up.downlink_queue_invalidated.last_f_cnt_down",
	"up.downlink_queue_invalidated.session_key_id",
	"up.downlink_queued",
	"up.downlink_queued.class_b_c",
	"up.downlink_queued.class_b_c.absolute_time",
	"up.downlink_queued.class_b_c.gateways",
	"up.downlink_queued.confirmed",
	"up.downlink_queued.correlation_ids",
	"up.downlink_queued.decoded_payload",
	"up.downlink_queued.decoded_payload_warnings",
	"up.downlink_queued.f_cnt",
	"up.downlink_queued.f_port",
	"up.downlink_queued.frm_payload",
	"up.downlink_queued.priority",
	"up.downlink_queued.session_key_id",
	"up.downlink_sent",
	"up.downlink_sent.class_b_c",
	"up.downlink_sent.class_b_c.absolute_time",
	"up.downlink_sent.class_b_c.gateways",
	"up.downlink_sent.confirmed",
	"up.downlink_sent.correlation_ids",
	"up.downlink_sent.decoded_payload",
	"up.downlink_sent.decoded_payload_warnings",
	"up.downlink_sent.f_cnt",
	"up.downlink_sent.f_port",
	"up.downlink_sent.frm_payload",
	"up.downlink_sent.priority",
	"up.downlink_sent.session_key_id",
	"up.join_accept",
	"up.join_accept.app_s_key",
	"up.join_accept.app_s_key.encrypted_key",
	"up.join_accept.app_s_key.kek_label",
	"up.join_accept.app_s_key.key",
	"up.join_accept.invalidated_downlinks",
	"up.join_accept.pending_session",
	"up.join_accept.received_at",
	"up.join_accept.session_key_id",
	"up.location_solved",
	"up.location_solved.attributes",
	"up.location_solved.location",
	"up.location_solved.location.accuracy",
	"up.location_solved.location.altitude",
	"up.location_solved.location.latitude",
	"up.location_solved.location.longitude",
	"up.location_solved.location.source",
	"up.location_solved.service",
	"up.service_data",
	"up.service_data.data",
	"up.service_data.service",
	"up.uplink_message",
	"up.uplink_message.app_s_key",
	"up.uplink_message.app_s_key.encrypted_key",
	"up.uplink_message.app_s_key.kek_label",
	"up.uplink_message.app_s_key.key",
	"up.uplink_message.confirmed",
	"up.uplink_message.consumed_airtime",
	"up.uplink_message.decoded_payload",
	"up.uplink_message.decoded_payload_warnings",
	"up.uplink_message.f_cnt",
	"up.uplink_message.f_port",
	"up.uplink_message.frm_payload",
	"up.uplink_message.last_a_f_cnt_down",
	"up.uplink_message.locations",
	"up.uplink_message.received_at",
	"up.uplink_message.rx_metadata",
	"up.uplink_message.session_key_id",
	"up.uplink_message.settings",
	"up.uplink_message.settings.coding_rate",
	"up.uplink_message.settings.data_rate",
	"up.uplink_message.settings.data_rate.modulation",
	"up.uplink_message.settings.data_rate.modulation.fsk",
	"up.uplink_message.settings.data_rate.modulation.fsk.bit_rate",
	"up.uplink_message.settings.data_rate.modulation.lora",
	"up.uplink_message.settings.data_rate.modulation.lora.bandwidth",
	"up.uplink_message.settings.data_rate.modulation.lora.spreading_factor",
	"up.uplink_message.settings.data_rate_index",
	"up.uplink_message.settings.downlink",
	"up.uplink_message.settings.downlink.antenna_index",
	"up.uplink_message.settings.downlink.invert_polarization",
	"up.uplink_message.settings.downlink.tx_power",
	"up.uplink_message.settings.enable_crc",
	"up.uplink_message.settings.frequency",
	"up.uplink_message.settings.time",
	"up.uplink_message.settings.timestamp",
	"up.uplink_message.version_ids",
	"up.uplink_message.version_ids.band_id",
	"up.uplink_message.version_ids.brand_id",
	"up.uplink_message.version_ids.firmware_version",
	"up.uplink_message.version_ids.hardware_version",
	"up.uplink_message.version_ids.model_id",
}

var ApplicationUpFieldPathsTopLevel = []string{
	"correlation_ids",
	"end_device_ids",
	"received_at",
	"simulated",
	"up",
}
var MessagePayloadFormattersFieldPathsNested = []string{
	"down_formatter",
	"down_formatter_parameter",
	"up_formatter",
	"up_formatter_parameter",
}

var MessagePayloadFormattersFieldPathsTopLevel = []string{
	"down_formatter",
	"down_formatter_parameter",
	"up_formatter",
	"up_formatter_parameter",
}
var DownlinkQueueRequestFieldPathsNested = []string{
	"downlinks",
	"end_device_ids",
	"end_device_ids.application_ids",
	"end_device_ids.application_ids.application_id",
	"end_device_ids.dev_addr",
	"end_device_ids.dev_eui",
	"end_device_ids.device_id",
	"end_device_ids.join_eui",
}

var DownlinkQueueRequestFieldPathsTopLevel = []string{
	"downlinks",
	"end_device_ids",
}
var ApplicationDownlink_ClassBCFieldPathsNested = []string{
	"absolute_time",
	"gateways",
}

var ApplicationDownlink_ClassBCFieldPathsTopLevel = []string{
	"absolute_time",
	"gateways",
}
