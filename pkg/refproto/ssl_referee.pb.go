// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_referee.proto

package refproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are the "coarse" stages of the game.
type Referee_Stage int32

const (
	// The first half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_NORMAL_FIRST_HALF_PRE Referee_Stage = 0
	// The first half of the normal game, before half time.
	Referee_NORMAL_FIRST_HALF Referee_Stage = 1
	// Half time between first and second halves.
	Referee_NORMAL_HALF_TIME Referee_Stage = 2
	// The second half is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_NORMAL_SECOND_HALF_PRE Referee_Stage = 3
	// The second half of the normal game, after half time.
	Referee_NORMAL_SECOND_HALF Referee_Stage = 4
	// The break before extra time.
	Referee_EXTRA_TIME_BREAK Referee_Stage = 5
	// The first half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_EXTRA_FIRST_HALF_PRE Referee_Stage = 6
	// The first half of extra time.
	Referee_EXTRA_FIRST_HALF Referee_Stage = 7
	// Half time between first and second extra halves.
	Referee_EXTRA_HALF_TIME Referee_Stage = 8
	// The second half of extra time is about to start.
	// A kickoff is called within this stage.
	// This stage ends with the NORMAL_START.
	Referee_EXTRA_SECOND_HALF_PRE Referee_Stage = 9
	// The second half of extra time.
	Referee_EXTRA_SECOND_HALF Referee_Stage = 10
	// The break before penalty shootout.
	Referee_PENALTY_SHOOTOUT_BREAK Referee_Stage = 11
	// The penalty shootout.
	Referee_PENALTY_SHOOTOUT Referee_Stage = 12
	// The game is over.
	Referee_POST_GAME Referee_Stage = 13
)

var Referee_Stage_name = map[int32]string{
	0:  "NORMAL_FIRST_HALF_PRE",
	1:  "NORMAL_FIRST_HALF",
	2:  "NORMAL_HALF_TIME",
	3:  "NORMAL_SECOND_HALF_PRE",
	4:  "NORMAL_SECOND_HALF",
	5:  "EXTRA_TIME_BREAK",
	6:  "EXTRA_FIRST_HALF_PRE",
	7:  "EXTRA_FIRST_HALF",
	8:  "EXTRA_HALF_TIME",
	9:  "EXTRA_SECOND_HALF_PRE",
	10: "EXTRA_SECOND_HALF",
	11: "PENALTY_SHOOTOUT_BREAK",
	12: "PENALTY_SHOOTOUT",
	13: "POST_GAME",
}
var Referee_Stage_value = map[string]int32{
	"NORMAL_FIRST_HALF_PRE":  0,
	"NORMAL_FIRST_HALF":      1,
	"NORMAL_HALF_TIME":       2,
	"NORMAL_SECOND_HALF_PRE": 3,
	"NORMAL_SECOND_HALF":     4,
	"EXTRA_TIME_BREAK":       5,
	"EXTRA_FIRST_HALF_PRE":   6,
	"EXTRA_FIRST_HALF":       7,
	"EXTRA_HALF_TIME":        8,
	"EXTRA_SECOND_HALF_PRE":  9,
	"EXTRA_SECOND_HALF":      10,
	"PENALTY_SHOOTOUT_BREAK": 11,
	"PENALTY_SHOOTOUT":       12,
	"POST_GAME":              13,
}

func (x Referee_Stage) Enum() *Referee_Stage {
	p := new(Referee_Stage)
	*p = x
	return p
}
func (x Referee_Stage) String() string {
	return proto.EnumName(Referee_Stage_name, int32(x))
}
func (x *Referee_Stage) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Referee_Stage_value, data, "Referee_Stage")
	if err != nil {
		return err
	}
	*x = Referee_Stage(value)
	return nil
}
func (Referee_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{0, 0}
}

// These are the "fine" states of play on the field.
type Referee_Command int32

const (
	// All robots should completely stop moving.
	Referee_HALT Referee_Command = 0
	// Robots must keep 50 cm from the ball.
	Referee_STOP Referee_Command = 1
	// A prepared kickoff or penalty may now be taken.
	Referee_NORMAL_START Referee_Command = 2
	// The ball is dropped and free for either team.
	Referee_FORCE_START Referee_Command = 3
	// The yellow team may move into kickoff position.
	Referee_PREPARE_KICKOFF_YELLOW Referee_Command = 4
	// The blue team may move into kickoff position.
	Referee_PREPARE_KICKOFF_BLUE Referee_Command = 5
	// The yellow team may move into penalty position.
	Referee_PREPARE_PENALTY_YELLOW Referee_Command = 6
	// The blue team may move into penalty position.
	Referee_PREPARE_PENALTY_BLUE Referee_Command = 7
	// The yellow team may take a direct free kick.
	Referee_DIRECT_FREE_YELLOW Referee_Command = 8
	// The blue team may take a direct free kick.
	Referee_DIRECT_FREE_BLUE Referee_Command = 9
	// The yellow team may take an indirect free kick.
	Referee_INDIRECT_FREE_YELLOW Referee_Command = 10
	// The blue team may take an indirect free kick.
	Referee_INDIRECT_FREE_BLUE Referee_Command = 11
	// The yellow team is currently in a timeout.
	Referee_TIMEOUT_YELLOW Referee_Command = 12
	// The blue team is currently in a timeout.
	Referee_TIMEOUT_BLUE Referee_Command = 13
	// The yellow team just scored a goal.
	// For information only.
	// For rules compliance, teams must treat as STOP.
	// Deprecated: Use the score field from the team infos instead. That way, you can also detect revoked goals.
	Referee_GOAL_YELLOW Referee_Command = 14 // Deprecated: Do not use.
	// The blue team just scored a goal. See also GOAL_YELLOW.
	Referee_GOAL_BLUE Referee_Command = 15 // Deprecated: Do not use.
	// Equivalent to STOP, but the yellow team must pick up the ball and
	// drop it in the Designated Position.
	Referee_BALL_PLACEMENT_YELLOW Referee_Command = 16
	// Equivalent to STOP, but the blue team must pick up the ball and drop
	// it in the Designated Position.
	Referee_BALL_PLACEMENT_BLUE Referee_Command = 17
)

var Referee_Command_name = map[int32]string{
	0:  "HALT",
	1:  "STOP",
	2:  "NORMAL_START",
	3:  "FORCE_START",
	4:  "PREPARE_KICKOFF_YELLOW",
	5:  "PREPARE_KICKOFF_BLUE",
	6:  "PREPARE_PENALTY_YELLOW",
	7:  "PREPARE_PENALTY_BLUE",
	8:  "DIRECT_FREE_YELLOW",
	9:  "DIRECT_FREE_BLUE",
	10: "INDIRECT_FREE_YELLOW",
	11: "INDIRECT_FREE_BLUE",
	12: "TIMEOUT_YELLOW",
	13: "TIMEOUT_BLUE",
	14: "GOAL_YELLOW",
	15: "GOAL_BLUE",
	16: "BALL_PLACEMENT_YELLOW",
	17: "BALL_PLACEMENT_BLUE",
}
var Referee_Command_value = map[string]int32{
	"HALT":                   0,
	"STOP":                   1,
	"NORMAL_START":           2,
	"FORCE_START":            3,
	"PREPARE_KICKOFF_YELLOW": 4,
	"PREPARE_KICKOFF_BLUE":   5,
	"PREPARE_PENALTY_YELLOW": 6,
	"PREPARE_PENALTY_BLUE":   7,
	"DIRECT_FREE_YELLOW":     8,
	"DIRECT_FREE_BLUE":       9,
	"INDIRECT_FREE_YELLOW":   10,
	"INDIRECT_FREE_BLUE":     11,
	"TIMEOUT_YELLOW":         12,
	"TIMEOUT_BLUE":           13,
	"GOAL_YELLOW":            14,
	"GOAL_BLUE":              15,
	"BALL_PLACEMENT_YELLOW":  16,
	"BALL_PLACEMENT_BLUE":    17,
}

func (x Referee_Command) Enum() *Referee_Command {
	p := new(Referee_Command)
	*p = x
	return p
}
func (x Referee_Command) String() string {
	return proto.EnumName(Referee_Command_name, int32(x))
}
func (x *Referee_Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Referee_Command_value, data, "Referee_Command")
	if err != nil {
		return err
	}
	*x = Referee_Command(value)
	return nil
}
func (Referee_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{0, 1}
}

// Each UDP packet contains one of these messages.
type Referee struct {
	// The UNIX timestamp when the packet was sent, in microseconds.
	// Divide by 1,000,000 to get a time_t.
	PacketTimestamp *uint64        `protobuf:"varint,1,req,name=packet_timestamp,json=packetTimestamp" json:"packet_timestamp,omitempty"`
	Stage           *Referee_Stage `protobuf:"varint,2,req,name=stage,enum=Referee_Stage" json:"stage,omitempty"`
	// The number of microseconds left in the stage.
	// The following stages have this value; the rest do not:
	// NORMAL_FIRST_HALF
	// NORMAL_HALF_TIME
	// NORMAL_SECOND_HALF
	// EXTRA_TIME_BREAK
	// EXTRA_FIRST_HALF
	// EXTRA_HALF_TIME
	// EXTRA_SECOND_HALF
	// PENALTY_SHOOTOUT_BREAK
	//
	// If the stage runs over its specified time, this value
	// becomes negative.
	StageTimeLeft *int32           `protobuf:"zigzag32,3,opt,name=stage_time_left,json=stageTimeLeft" json:"stage_time_left,omitempty"`
	Command       *Referee_Command `protobuf:"varint,4,req,name=command,enum=Referee_Command" json:"command,omitempty"`
	// The number of commands issued since startup (mod 2^32).
	CommandCounter *uint32 `protobuf:"varint,5,req,name=command_counter,json=commandCounter" json:"command_counter,omitempty"`
	// The UNIX timestamp when the command was issued, in microseconds.
	// This value changes only when a new command is issued, not on each packet.
	CommandTimestamp *uint64 `protobuf:"varint,6,req,name=command_timestamp,json=commandTimestamp" json:"command_timestamp,omitempty"`
	// Information about the two teams.
	Yellow             *Referee_TeamInfo `protobuf:"bytes,7,req,name=yellow" json:"yellow,omitempty"`
	Blue               *Referee_TeamInfo `protobuf:"bytes,8,req,name=blue" json:"blue,omitempty"`
	DesignatedPosition *Referee_Point    `protobuf:"bytes,9,opt,name=designated_position,json=designatedPosition" json:"designated_position,omitempty"`
	// Information about the direction of play.
	// True, if the blue team will have it's goal on the positive x-axis of the ssl-vision coordinate system
	// Obviously, the yellow team will play on the opposide half
	// For compatibility, this field is optional
	BlueTeamOnPositiveHalf *bool `protobuf:"varint,10,opt,name=blueTeamOnPositiveHalf" json:"blueTeamOnPositiveHalf,omitempty"`
	// The game event that caused the referee command
	// deprecated in favor for current_game_event
	GameEvent *Game_Event `protobuf:"bytes,11,opt,name=gameEvent" json:"gameEvent,omitempty"` // Deprecated: Do not use.
	// the game event that is currently processed
	CurrentGameEvent *GameEvent `protobuf:"bytes,12,opt,name=current_game_event,json=currentGameEvent" json:"current_game_event,omitempty"`
	// all non-finished proposed game events that may be processed next
	ProposedGameEvents   []*ProposedGameEvent `protobuf:"bytes,13,rep,name=proposed_game_events,json=proposedGameEvents" json:"proposed_game_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Referee) Reset()         { *m = Referee{} }
func (m *Referee) String() string { return proto.CompactTextString(m) }
func (*Referee) ProtoMessage()    {}
func (*Referee) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{0}
}
func (m *Referee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee.Unmarshal(m, b)
}
func (m *Referee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee.Marshal(b, m, deterministic)
}
func (dst *Referee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee.Merge(dst, src)
}
func (m *Referee) XXX_Size() int {
	return xxx_messageInfo_Referee.Size(m)
}
func (m *Referee) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee.DiscardUnknown(m)
}

var xxx_messageInfo_Referee proto.InternalMessageInfo

func (m *Referee) GetPacketTimestamp() uint64 {
	if m != nil && m.PacketTimestamp != nil {
		return *m.PacketTimestamp
	}
	return 0
}

func (m *Referee) GetStage() Referee_Stage {
	if m != nil && m.Stage != nil {
		return *m.Stage
	}
	return Referee_NORMAL_FIRST_HALF_PRE
}

func (m *Referee) GetStageTimeLeft() int32 {
	if m != nil && m.StageTimeLeft != nil {
		return *m.StageTimeLeft
	}
	return 0
}

func (m *Referee) GetCommand() Referee_Command {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return Referee_HALT
}

func (m *Referee) GetCommandCounter() uint32 {
	if m != nil && m.CommandCounter != nil {
		return *m.CommandCounter
	}
	return 0
}

func (m *Referee) GetCommandTimestamp() uint64 {
	if m != nil && m.CommandTimestamp != nil {
		return *m.CommandTimestamp
	}
	return 0
}

func (m *Referee) GetYellow() *Referee_TeamInfo {
	if m != nil {
		return m.Yellow
	}
	return nil
}

func (m *Referee) GetBlue() *Referee_TeamInfo {
	if m != nil {
		return m.Blue
	}
	return nil
}

func (m *Referee) GetDesignatedPosition() *Referee_Point {
	if m != nil {
		return m.DesignatedPosition
	}
	return nil
}

func (m *Referee) GetBlueTeamOnPositiveHalf() bool {
	if m != nil && m.BlueTeamOnPositiveHalf != nil {
		return *m.BlueTeamOnPositiveHalf
	}
	return false
}

// Deprecated: Do not use.
func (m *Referee) GetGameEvent() *Game_Event {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func (m *Referee) GetCurrentGameEvent() *GameEvent {
	if m != nil {
		return m.CurrentGameEvent
	}
	return nil
}

func (m *Referee) GetProposedGameEvents() []*ProposedGameEvent {
	if m != nil {
		return m.ProposedGameEvents
	}
	return nil
}

// Information about a single team.
type Referee_TeamInfo struct {
	// The team's name (empty string if operator has not typed anything).
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// The number of goals scored by the team during normal play and overtime.
	Score *uint32 `protobuf:"varint,2,req,name=score" json:"score,omitempty"`
	// The number of red cards issued to the team since the beginning of the game.
	RedCards *uint32 `protobuf:"varint,3,req,name=red_cards,json=redCards" json:"red_cards,omitempty"`
	// The amount of time (in microseconds) left on each yellow card issued to the team.
	// If no yellow cards are issued, this array has no elements.
	// Otherwise, times are ordered from smallest to largest.
	YellowCardTimes []uint32 `protobuf:"varint,4,rep,packed,name=yellow_card_times,json=yellowCardTimes" json:"yellow_card_times,omitempty"`
	// The total number of yellow cards ever issued to the team.
	YellowCards *uint32 `protobuf:"varint,5,req,name=yellow_cards,json=yellowCards" json:"yellow_cards,omitempty"`
	// The number of timeouts this team can still call.
	// If in a timeout right now, that timeout is excluded.
	Timeouts *uint32 `protobuf:"varint,6,req,name=timeouts" json:"timeouts,omitempty"`
	// The number of microseconds of timeout this team can use.
	TimeoutTime *uint32 `protobuf:"varint,7,req,name=timeout_time,json=timeoutTime" json:"timeout_time,omitempty"`
	// The pattern number of this team's goalie.
	Goalie *uint32 `protobuf:"varint,8,req,name=goalie" json:"goalie,omitempty"`
	// The total number of bot collisions committed by this team
	BotCollisions *uint32 `protobuf:"varint,9,opt,name=bot_collisions,json=botCollisions" json:"bot_collisions,omitempty"`
	// The number of consecutive ball placement failures of this team
	BallPlacementFailures *uint32 `protobuf:"varint,10,opt,name=ball_placement_failures,json=ballPlacementFailures" json:"ball_placement_failures,omitempty"`
	// The total number of infringements where robots were too fast during game stoppage
	BotSpeedInfringements *uint32 `protobuf:"varint,11,opt,name=bot_speed_infringements,json=botSpeedInfringements" json:"bot_speed_infringements,omitempty"`
	// Indicate if the team is able and allowed to place the ball
	CanPlaceBall         *bool    `protobuf:"varint,12,opt,name=can_place_ball,json=canPlaceBall" json:"can_place_ball,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Referee_TeamInfo) Reset()         { *m = Referee_TeamInfo{} }
func (m *Referee_TeamInfo) String() string { return proto.CompactTextString(m) }
func (*Referee_TeamInfo) ProtoMessage()    {}
func (*Referee_TeamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{0, 0}
}
func (m *Referee_TeamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee_TeamInfo.Unmarshal(m, b)
}
func (m *Referee_TeamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee_TeamInfo.Marshal(b, m, deterministic)
}
func (dst *Referee_TeamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee_TeamInfo.Merge(dst, src)
}
func (m *Referee_TeamInfo) XXX_Size() int {
	return xxx_messageInfo_Referee_TeamInfo.Size(m)
}
func (m *Referee_TeamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee_TeamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Referee_TeamInfo proto.InternalMessageInfo

func (m *Referee_TeamInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Referee_TeamInfo) GetScore() uint32 {
	if m != nil && m.Score != nil {
		return *m.Score
	}
	return 0
}

func (m *Referee_TeamInfo) GetRedCards() uint32 {
	if m != nil && m.RedCards != nil {
		return *m.RedCards
	}
	return 0
}

func (m *Referee_TeamInfo) GetYellowCardTimes() []uint32 {
	if m != nil {
		return m.YellowCardTimes
	}
	return nil
}

func (m *Referee_TeamInfo) GetYellowCards() uint32 {
	if m != nil && m.YellowCards != nil {
		return *m.YellowCards
	}
	return 0
}

func (m *Referee_TeamInfo) GetTimeouts() uint32 {
	if m != nil && m.Timeouts != nil {
		return *m.Timeouts
	}
	return 0
}

func (m *Referee_TeamInfo) GetTimeoutTime() uint32 {
	if m != nil && m.TimeoutTime != nil {
		return *m.TimeoutTime
	}
	return 0
}

func (m *Referee_TeamInfo) GetGoalie() uint32 {
	if m != nil && m.Goalie != nil {
		return *m.Goalie
	}
	return 0
}

func (m *Referee_TeamInfo) GetBotCollisions() uint32 {
	if m != nil && m.BotCollisions != nil {
		return *m.BotCollisions
	}
	return 0
}

func (m *Referee_TeamInfo) GetBallPlacementFailures() uint32 {
	if m != nil && m.BallPlacementFailures != nil {
		return *m.BallPlacementFailures
	}
	return 0
}

func (m *Referee_TeamInfo) GetBotSpeedInfringements() uint32 {
	if m != nil && m.BotSpeedInfringements != nil {
		return *m.BotSpeedInfringements
	}
	return 0
}

func (m *Referee_TeamInfo) GetCanPlaceBall() bool {
	if m != nil && m.CanPlaceBall != nil {
		return *m.CanPlaceBall
	}
	return false
}

// The coordinates of the Designated Position. These are measured in
// millimetres and correspond to SSL-Vision coordinates. These fields are
// always either both present (in the case of a ball placement command) or
// both absent (in the case of any other command).
type Referee_Point struct {
	X                    *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                    *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Referee_Point) Reset()         { *m = Referee_Point{} }
func (m *Referee_Point) String() string { return proto.CompactTextString(m) }
func (*Referee_Point) ProtoMessage()    {}
func (*Referee_Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{0, 1}
}
func (m *Referee_Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Referee_Point.Unmarshal(m, b)
}
func (m *Referee_Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Referee_Point.Marshal(b, m, deterministic)
}
func (dst *Referee_Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Referee_Point.Merge(dst, src)
}
func (m *Referee_Point) XXX_Size() int {
	return xxx_messageInfo_Referee_Point.Size(m)
}
func (m *Referee_Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Referee_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Referee_Point proto.InternalMessageInfo

func (m *Referee_Point) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Referee_Point) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type ProposedGameEvent struct {
	// the UNIX timestamp when the game event was proposed
	Timestamp *uint64 `protobuf:"varint,1,req,name=timestamp" json:"timestamp,omitempty"`
	// the identifier of the proposer
	ProposerId *string `protobuf:"bytes,2,req,name=proposer_id,json=proposerId" json:"proposer_id,omitempty"`
	// the actual game event
	GameEvent            *GameEvent `protobuf:"bytes,3,req,name=game_event,json=gameEvent" json:"game_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProposedGameEvent) Reset()         { *m = ProposedGameEvent{} }
func (m *ProposedGameEvent) String() string { return proto.CompactTextString(m) }
func (*ProposedGameEvent) ProtoMessage()    {}
func (*ProposedGameEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ssl_referee_f2db4d8b0445c395, []int{1}
}
func (m *ProposedGameEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposedGameEvent.Unmarshal(m, b)
}
func (m *ProposedGameEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposedGameEvent.Marshal(b, m, deterministic)
}
func (dst *ProposedGameEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedGameEvent.Merge(dst, src)
}
func (m *ProposedGameEvent) XXX_Size() int {
	return xxx_messageInfo_ProposedGameEvent.Size(m)
}
func (m *ProposedGameEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedGameEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedGameEvent proto.InternalMessageInfo

func (m *ProposedGameEvent) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ProposedGameEvent) GetProposerId() string {
	if m != nil && m.ProposerId != nil {
		return *m.ProposerId
	}
	return ""
}

func (m *ProposedGameEvent) GetGameEvent() *GameEvent {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func init() {
	proto.RegisterType((*Referee)(nil), "Referee")
	proto.RegisterType((*Referee_TeamInfo)(nil), "Referee.TeamInfo")
	proto.RegisterType((*Referee_Point)(nil), "Referee.Point")
	proto.RegisterType((*ProposedGameEvent)(nil), "ProposedGameEvent")
	proto.RegisterEnum("Referee_Stage", Referee_Stage_name, Referee_Stage_value)
	proto.RegisterEnum("Referee_Command", Referee_Command_name, Referee_Command_value)
}

func init() { proto.RegisterFile("ssl_referee.proto", fileDescriptor_ssl_referee_f2db4d8b0445c395) }

var fileDescriptor_ssl_referee_f2db4d8b0445c395 = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x5e, 0x1b, 0x08, 0x70, 0x8c, 0xc1, 0x9c, 0xfc, 0xac, 0x43, 0x2b, 0x95, 0xa6, 0xbb, 0x2d,
	0x69, 0x55, 0xd4, 0xe6, 0x62, 0xd5, 0x5e, 0x55, 0x86, 0x98, 0x04, 0xc5, 0x89, 0xd1, 0xe0, 0x55,
	0xbb, 0x57, 0x23, 0x07, 0x86, 0xc8, 0xaa, 0xb1, 0x91, 0xed, 0x6c, 0x37, 0x17, 0x7d, 0x85, 0x3e,
	0x4f, 0xef, 0xfb, 0x04, 0x7d, 0xa3, 0x6a, 0xc6, 0x63, 0xcc, 0x92, 0xee, 0xdd, 0xcc, 0xf7, 0x7d,
	0xe7, 0x9b, 0x33, 0x73, 0x8e, 0x8f, 0xa1, 0x9b, 0xa6, 0x21, 0x4d, 0xd8, 0x8a, 0x25, 0x8c, 0x0d,
	0x37, 0x49, 0x9c, 0xc5, 0xbd, 0x23, 0x0e, 0x3d, 0xf8, 0x6b, 0x46, 0xd9, 0x7b, 0x16, 0x65, 0x12,
	0x3d, 0xfd, 0x18, 0xa5, 0x17, 0x3f, 0xfc, 0xf8, 0x73, 0x4e, 0x9d, 0xfd, 0xd3, 0x86, 0x3a, 0xc9,
	0x2d, 0xf0, 0x1c, 0x8c, 0x8d, 0xbf, 0xf8, 0x9d, 0x65, 0x34, 0x0b, 0xd6, 0x2c, 0xcd, 0xfc, 0xf5,
	0xc6, 0x54, 0xfa, 0xea, 0xa0, 0x4a, 0x3a, 0x39, 0xee, 0x15, 0x30, 0xbe, 0x82, 0x5a, 0x9a, 0xf9,
	0x0f, 0xcc, 0x54, 0xfb, 0xea, 0xa0, 0x7d, 0xd1, 0x1e, 0x4a, 0x8f, 0xe1, 0x9c, 0xa3, 0x24, 0x27,
	0xf1, 0x6b, 0xe8, 0x88, 0x85, 0xf0, 0xa3, 0x21, 0x5b, 0x65, 0x66, 0xa5, 0xaf, 0x0c, 0xba, 0x44,
	0x17, 0x30, 0xb7, 0x73, 0xd8, 0x2a, 0xc3, 0x6f, 0xa1, 0xbe, 0x88, 0xd7, 0x6b, 0x3f, 0x5a, 0x9a,
	0x55, 0xe1, 0x67, 0x6c, 0xfd, 0xc6, 0x39, 0x4e, 0x0a, 0x01, 0x7e, 0x03, 0x1d, 0xb9, 0xa4, 0x8b,
	0xf8, 0x31, 0xca, 0x58, 0x62, 0xd6, 0xfa, 0xea, 0x40, 0x27, 0x6d, 0x09, 0x8f, 0x73, 0x14, 0xbf,
	0x83, 0x6e, 0x21, 0x2c, 0xaf, 0x73, 0x20, 0xae, 0x63, 0x48, 0xa2, 0xbc, 0xcf, 0x39, 0x1c, 0x3c,
	0xb1, 0x30, 0x8c, 0xff, 0x30, 0xeb, 0x7d, 0x75, 0xa0, 0x5d, 0x74, 0xb7, 0x09, 0x78, 0xcc, 0x5f,
	0x4f, 0xa3, 0x55, 0x4c, 0xa4, 0x00, 0x5f, 0x43, 0xf5, 0x3e, 0x7c, 0x64, 0x66, 0xe3, 0x53, 0x42,
	0x41, 0xe3, 0x2f, 0x70, 0xb8, 0x64, 0x69, 0xf0, 0x10, 0xf9, 0x19, 0x5b, 0xd2, 0x4d, 0x9c, 0x06,
	0x59, 0x10, 0x47, 0x66, 0xb3, 0xaf, 0x0c, 0xb4, 0x9d, 0xf7, 0x9a, 0xc5, 0x41, 0x94, 0x11, 0x2c,
	0xa5, 0x33, 0xa9, 0xc4, 0x37, 0x70, 0xc2, 0x8d, 0xb8, 0xad, 0x1b, 0xe5, 0xe8, 0x7b, 0x76, 0xed,
	0x87, 0x2b, 0x13, 0xfa, 0xca, 0xa0, 0x41, 0x3e, 0xc1, 0xe2, 0xf7, 0xd0, 0xe4, 0xa5, 0xb6, 0x79,
	0xa5, 0x4d, 0x4d, 0x1c, 0xa7, 0x0d, 0xaf, 0x78, 0xf1, 0x05, 0x34, 0x52, 0x4d, 0x85, 0x94, 0x0a,
	0xfc, 0x09, 0x70, 0xf1, 0x98, 0x24, 0xbc, 0x2d, 0xca, 0x0e, 0x31, 0x5b, 0x22, 0x0e, 0x44, 0x9c,
	0xd0, 0x11, 0x43, 0xaa, 0xb6, 0x08, 0x5e, 0xc2, 0xd1, 0x26, 0x89, 0x37, 0x71, 0xca, 0x96, 0x3b,
	0xa1, 0xa9, 0xa9, 0xf7, 0x2b, 0x03, 0xed, 0x02, 0x87, 0x33, 0x49, 0x96, 0x1e, 0xb8, 0xd9, 0x87,
	0xd2, 0xde, 0xdf, 0x15, 0x68, 0x14, 0x4f, 0x87, 0x08, 0xd5, 0xc8, 0x5f, 0x33, 0xd1, 0x75, 0x4d,
	0x22, 0xd6, 0x78, 0x04, 0xb5, 0x74, 0x11, 0x27, 0x79, 0xab, 0xe9, 0x24, 0xdf, 0xe0, 0x67, 0xd0,
	0x4c, 0xd8, 0x92, 0x2e, 0xfc, 0x64, 0x99, 0x9a, 0x15, 0xc1, 0x34, 0x12, 0xb6, 0x1c, 0xf3, 0x3d,
	0x0e, 0xa1, 0x9b, 0x17, 0x4b, 0xf0, 0x79, 0xf9, 0xcd, 0x6a, 0xbf, 0x32, 0xd0, 0x47, 0xaa, 0xa1,
	0x90, 0x4e, 0x4e, 0x72, 0xad, 0xe8, 0x00, 0xfc, 0x12, 0x5a, 0x3b, 0xfa, 0x54, 0x36, 0x94, 0x56,
	0xca, 0x52, 0xec, 0x41, 0x83, 0xdb, 0xc4, 0x8f, 0x59, 0x2a, 0x9a, 0x48, 0x27, 0xdb, 0x3d, 0x0f,
	0x97, 0x6b, 0x71, 0x94, 0x68, 0x21, 0x9d, 0x68, 0x12, 0xe3, 0x47, 0xe0, 0x09, 0x1c, 0x3c, 0xc4,
	0x7e, 0x18, 0xe4, 0x6d, 0xa3, 0x13, 0xb9, 0xc3, 0xd7, 0xd0, 0xbe, 0x8f, 0x33, 0xba, 0x88, 0xc3,
	0x30, 0x48, 0x83, 0x38, 0x4a, 0x45, 0x83, 0xe8, 0x44, 0xbf, 0x8f, 0xb3, 0xf1, 0x16, 0xc4, 0x37,
	0xf0, 0xf2, 0xde, 0x0f, 0x43, 0xba, 0x09, 0xfd, 0x05, 0x5b, 0xf3, 0x5a, 0xad, 0xfc, 0x20, 0x7c,
	0x4c, 0x58, 0x2a, 0x9a, 0x41, 0x27, 0xc7, 0x9c, 0x9e, 0x15, 0xec, 0x44, 0x92, 0x22, 0x2e, 0xce,
	0x68, 0xba, 0x61, 0x6c, 0x49, 0x83, 0x68, 0x95, 0x04, 0xd1, 0x83, 0x50, 0xa4, 0xa2, 0x33, 0x78,
	0x5c, 0x9c, 0xcd, 0x39, 0x3b, 0xdd, 0x25, 0xf1, 0x15, 0xb4, 0x17, 0x7e, 0x94, 0x1f, 0x47, 0xb9,
	0xb5, 0x68, 0x88, 0x06, 0x69, 0x2d, 0xfc, 0x48, 0x9c, 0x32, 0xf2, 0xc3, 0xb0, 0xf7, 0x15, 0xd4,
	0x44, 0xfb, 0x62, 0x0b, 0x94, 0x0f, 0xa2, 0x66, 0x2a, 0x51, 0x3e, 0xf0, 0xdd, 0x93, 0x28, 0x96,
	0x4a, 0x94, 0xa7, 0xb3, 0x7f, 0x55, 0xa8, 0x89, 0xa1, 0x80, 0xa7, 0x70, 0x7c, 0xe7, 0x92, 0x5b,
	0xcb, 0xa1, 0x93, 0x29, 0x99, 0x7b, 0xf4, 0xda, 0x72, 0x26, 0x74, 0x46, 0x6c, 0xe3, 0x05, 0x1e,
	0x43, 0xf7, 0x19, 0x65, 0x28, 0x78, 0x04, 0x86, 0x84, 0x85, 0xd6, 0x9b, 0xde, 0xda, 0x86, 0x8a,
	0x3d, 0x38, 0x91, 0xe8, 0xdc, 0x1e, 0xbb, 0x77, 0x97, 0xa5, 0x51, 0x05, 0x4f, 0x00, 0x9f, 0x73,
	0x46, 0x95, 0x3b, 0xd9, 0xbf, 0x79, 0xc4, 0x12, 0x1e, 0x74, 0x44, 0x6c, 0xeb, 0xc6, 0xa8, 0xa1,
	0x09, 0x47, 0x39, 0xba, 0x97, 0xd0, 0x41, 0xa9, 0xdf, 0xc9, 0xa7, 0x8e, 0x87, 0xd0, 0xc9, 0xd1,
	0x32, 0x9d, 0x06, 0xbf, 0x56, 0x0e, 0xee, 0x67, 0xd3, 0xe4, 0xd7, 0x7a, 0x46, 0x19, 0xc0, 0x2f,
	0x30, 0xb3, 0xef, 0x2c, 0xc7, 0x7b, 0x47, 0xe7, 0xd7, 0xae, 0xeb, 0xb9, 0x6f, 0x3d, 0x99, 0x92,
	0xc6, 0x0f, 0xde, 0xe7, 0x8c, 0x16, 0xea, 0xd0, 0x9c, 0xb9, 0x73, 0x8f, 0x5e, 0x59, 0xb7, 0xb6,
	0xa1, 0x9f, 0xfd, 0x55, 0x81, 0xba, 0x1c, 0x8c, 0xd8, 0x80, 0xea, 0xb5, 0xe5, 0x78, 0xc6, 0x0b,
	0xbe, 0x9a, 0x7b, 0xee, 0xcc, 0x50, 0xd0, 0x80, 0x56, 0xf1, 0x0a, 0x9e, 0x45, 0x3c, 0x43, 0xc5,
	0x0e, 0x68, 0x13, 0x97, 0x8c, 0x6d, 0x09, 0x54, 0x44, 0x0e, 0xc4, 0x9e, 0x59, 0xc4, 0xa6, 0x37,
	0xd3, 0xf1, 0x8d, 0x3b, 0x99, 0xd0, 0x77, 0xb6, 0xe3, 0xb8, 0xbf, 0x1a, 0x55, 0xfe, 0x2c, 0xfb,
	0xdc, 0xc8, 0x79, 0x6b, 0x1b, 0xb5, 0xdd, 0xa8, 0x22, 0x4b, 0x19, 0x75, 0xb0, 0x1b, 0x55, 0x70,
	0x22, 0xaa, 0xce, 0x8b, 0x72, 0x39, 0x25, 0xf6, 0xd8, 0xa3, 0x13, 0x62, 0xdb, 0x45, 0x44, 0x83,
	0xdf, 0x75, 0x17, 0x17, 0xea, 0x26, 0xf7, 0x99, 0xde, 0xfd, 0x8f, 0x1e, 0xb8, 0xcf, 0xc7, 0x8c,
	0x88, 0xd0, 0x10, 0xa1, 0xcd, 0x6b, 0xc1, 0x9f, 0x51, 0x6a, 0x5b, 0xfc, 0x09, 0x0a, 0x4c, 0xa8,
	0x74, 0x3c, 0x04, 0xed, 0xca, 0xb5, 0x9c, 0x42, 0xd2, 0xee, 0xa9, 0x0d, 0x05, 0xbb, 0xd0, 0x14,
	0xa0, 0xd0, 0x74, 0x04, 0x74, 0x0a, 0xc7, 0x23, 0xcb, 0x71, 0xe8, 0xcc, 0xb1, 0xc6, 0xf6, 0xad,
	0x7d, 0xb7, 0x35, 0x35, 0xf0, 0x25, 0x1c, 0xee, 0x51, 0x22, 0xae, 0x7b, 0xf6, 0x27, 0x74, 0x9f,
	0x4d, 0x3b, 0xfc, 0x1c, 0x9a, 0xfb, 0xff, 0xd1, 0x12, 0xc0, 0x2f, 0x40, 0x93, 0xd3, 0x30, 0xa1,
	0xc1, 0x52, 0x7c, 0x2f, 0x4d, 0x02, 0x05, 0x34, 0x5d, 0xe2, 0x39, 0xc0, 0xce, 0x40, 0xae, 0x88,
	0xbf, 0xcd, 0xee, 0x40, 0x2e, 0x67, 0xf8, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0x57, 0xdc,
	0x3e, 0x09, 0x08, 0x00, 0x00,
}