package datamodel

const (
	LOCALPATH          string = "/tmp/PacketVortoj.db"
	PACKET_TABLENAME   string = "packet"
	DISTANCE_TABLENAME string = "distance"
	DBTYPE             string = "sqlite3"
)

type DBPacket struct {
	ID        int64  `db:"id"`
	DeviceID  string `db:"deviceid"`
	SrcMAC    string `db:"src_mac"`
	DstMAC    string `db:"dst_mac"`
	SrcIP     string `db:"src_ip"`
	DstIP     string `db:"dst_ip"`
	SrcPort   string `db:"src_port"`
	DstPort   string `db:"dst_port"`
	SYN       bool   `db:"syn"`
	ACK       bool   `db:"ack"`
	Sequence  int64  `db:"sequence"`
	Protocol  string `db:"protocol"`
	Length    int64  `db:"length"`
	DataChank []byte `db:"datachank"`
}

type DistPacket struct {
	ID       int64  `db:"id"`
	MACaddr  string `db:"macaddr"`
	Pwr      int64  `db:"pwr"`
	Distance int64  `db:"distance"`
}
