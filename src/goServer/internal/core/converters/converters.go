package converters

import (
	"fmt"
	"goServer/internal/core/pb"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/relvacode/iso8601"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertIso8601ToTimestamppb(val any) *timestamppb.Timestamp {
	if val == nil {
		return nil
	}

	timeStr := val.(string)

	if len(timeStr) == 0 {
		return nil
	}

	ts, err := iso8601.ParseString(timeStr)
	if err != nil {
		log.Panicln(err)
	}

	return timestamppb.New(ts)
}

type Worker struct {
	Id            int64            `json:"id,omitempty"`
	UserId        *string          `json:"userId,omitempty"`
	CreatedById   *int64           `json:"createdById,omitempty"`
	UpdatedById   *int64           `json:"updatedById,omitempty"`
	CreatedAt     string           `json:"createdAt,omitempty"`
	LastUpdatedAt string           `json:"lastUpdatedAt,omitempty"`
	PictureUrl    *string          `json:"pictureUrl,omitempty"`
	FirstName     string           `json:"firstName,omitempty"`
	LastName      string           `json:"lastName,omitempty"`
	MiddleName    *string          `json:"middleName,omitempty"`
	Birthdate     string           `json:"birthdate,omitempty"`
	Username      string           `json:"username,omitempty"`
	Email         string           `json:"email,omitempty"`
	Mobile        *string          `json:"mobile,omitempty"`
	Addresses     []map[string]any `json:"addresses,omitempty"`
	Flags         uint32           `json:"flags,omitempty"`
	Uuid          string           `json:"uuid,omitempty"`
}

func (w *Worker) TranslatePb(pb *pb.Worker) *Worker {
	w.Id = pb.Id
	w.UserId = pb.UserId
	w.CreatedById = pb.CreatedById
	w.UpdatedById = pb.UpdatedById
	w.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	w.PictureUrl = pb.PictureUrl
	w.FirstName = pb.FirstName
	w.LastName = pb.LastName
	w.MiddleName = pb.MiddleName
	w.Birthdate = pb.Birthdate.AsTime().Format(time.RFC3339Nano)
	w.Username = pb.Username
	w.Email = pb.Email
	w.Mobile = pb.Mobile
	w.Flags = pb.Flags
	w.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		w.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.Addresses != nil {
		addresses := []map[string]any{}

		for _, addr := range pb.Addresses {
			addresses = append(addresses, addr.AsMap())
		}

		w.Addresses = addresses
	}

	return w
}

type WorkerRoleRelation struct {
	Id             int64 `json:"id,omitempty"`
	RoleId         int64 `json:"roleId,omitempty"`
	WorkerId       int64 `json:"workerId,omitempty"`
	OrganizationId int64 `json:"organizationId,omitempty"`
}

func (w *WorkerRoleRelation) TranslatePb(pb *pb.WorkerRoleRelation) *WorkerRoleRelation {
	w.Id = pb.Id
	w.RoleId = pb.RoleId
	w.WorkerId = pb.WorkerId
	w.OrganizationId = pb.OrganizationId

	return w
}

type WorkerOrganizationRelation struct {
	Id             int64 `json:"id,omitempty"`
	WorkerId       int64 `json:"workerId,omitempty"`
	OrganizationId int64 `json:"organizationId,omitempty"`
}

func (w *WorkerOrganizationRelation) TranslatePb(pb *pb.WorkerOrganizationRelation) *WorkerOrganizationRelation {
	w.Id = pb.Id
	w.WorkerId = pb.WorkerId
	w.OrganizationId = pb.OrganizationId

	return w
}

type WorkerIdentityCardRelation struct {
	Id      int64 `json:"id,omitempty"`
	OwnerId int64 `json:"ownerId,omitempty"`
	CardId  int64 `json:"cardId,omitempty"`
}

func (w *WorkerIdentityCardRelation) TranslatePb(pb *pb.WorkerIdentityCardRelation) *WorkerIdentityCardRelation {
	w.Id = pb.Id
	w.OwnerId = pb.OwnerId
	w.CardId = pb.CardId

	return w
}

type WorkerIdentityCard struct {
	Id            int64          `json:"id,omitempty"`
	WorkerId      int64          `json:"workerId,omitempty"`
	CreatedById   *int64         `json:"createdById,omitempty"`
	UpdatedById   *int64         `json:"updatedById,omitempty"`
	CreatedAt     string         `json:"createdAt,omitempty"`
	LastUpdatedAt string         `json:"lastUpdatedAt,omitempty"`
	FrontImageUrl string         `json:"frontImageUrl,omitempty"`
	BackImageUrl  string         `json:"backImageUrl,omitempty"`
	Name          string         `json:"name,omitempty"`
	ExtractedInfo map[string]any `json:"extractedInfo,omitempty"`
	Flags         uint32         `json:"flags,omitempty"`
	Uuid          string         `json:"uuid,omitempty"`
}

func (id *WorkerIdentityCard) TranslatePb(pb *pb.WorkerIdentityCard) *WorkerIdentityCard {
	id.Id = pb.Id
	id.WorkerId = pb.WorkerId
	id.CreatedById = pb.CreatedById
	id.UpdatedById = pb.UpdatedById
	id.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	id.FrontImageUrl = pb.FrontImageUrl
	id.BackImageUrl = pb.BackImageUrl
	id.Name = pb.Name
	id.ExtractedInfo = pb.ExtractedInfo.AsMap()
	id.Flags = pb.Flags
	id.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		id.LastUpdatedAt = pb.LastUpdatedAt.String()
	}

	return id
}

type VerifiedWorkerIdentityCardRelation struct {
	Id           int64 `json:"id,omitempty"`
	OwnerId      int64 `json:"ownerId,omitempty"`
	VerifiedById int64 `json:"verifiedById,omitempty"`
	CardId       int64 `json:"cardId,omitempty"`
}

func (id *VerifiedWorkerIdentityCardRelation) TranslatePb(pb *pb.VerifiedWorkerIdentityCardRelation) *VerifiedWorkerIdentityCardRelation {
	id.Id = pb.Id
	id.OwnerId = pb.OwnerId
	id.VerifiedById = pb.VerifiedById
	id.CardId = pb.CardId

	return id
}

type Team struct {
	Id             int64  `json:"id,omitempty"`
	CreatedById    *int64 `json:"createdById,omitempty"`
	UpdatedById    *int64 `json:"updatedById,omitempty"`
	OrganizationId int64  `json:"organizationId,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	LastUpdatedAt  string `json:"lastUpdatedAt,omitempty"`
	Name           string `json:"name,omitempty"`
	Flags          uint32 `json:"flags,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}

func (t *Team) TranslatePb(pb *pb.Team) *Team {
	t.Id = pb.Id
	t.CreatedById = pb.CreatedById
	t.UpdatedById = pb.UpdatedById
	t.OrganizationId = pb.OrganizationId
	t.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	t.Name = pb.Name
	t.Flags = pb.Flags
	t.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		t.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	return t
}

type TeamMemberRelation struct {
	Id       int64 `json:"id,omitempty"`
	TeamId   int64 `json:"teamId,omitempty"`
	WorkerId int64 `json:"workerId,omitempty"`
}

func (t *TeamMemberRelation) TranslatePb(pb *pb.TeamMemberRelation) *TeamMemberRelation {
	t.Id = pb.Id
	t.TeamId = pb.TeamId
	t.WorkerId = pb.WorkerId

	return t
}

type Shift struct {
	Id             int64  `json:"id,omitempty"`
	CreatedById    *int64 `json:"createdById,omitempty"`
	UpdatedById    *int64 `json:"updatedById,omitempty"`
	OrganizationId int64  `json:"organizationId,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	LastUpdatedAt  string `json:"lastUpdatedAt,omitempty"`
	Name           string `json:"name,omitempty"`
	Day            int32  `json:"day,omitempty"` // Assuming ShiftDay is an int32 enum
	ClockIn        string `json:"clockIn,omitempty"`
	ClockOut       string `json:"clockOut,omitempty"`
	GroupId        string `json:"groupId,omitempty"`
	Flags          uint32 `json:"flags,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}

func (s *Shift) TranslatePb(pb *pb.Shift) *Shift {
	s.Id = pb.Id
	s.CreatedById = pb.CreatedById
	s.UpdatedById = pb.UpdatedById
	s.OrganizationId = pb.OrganizationId
	s.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	s.Name = pb.Name
	s.Day = int32(pb.Day.Number())
	s.ClockIn = pb.ClockIn.AsTime().Format(time.RFC3339Nano)
	s.ClockOut = pb.ClockOut.AsTime().Format(time.RFC3339Nano)
	s.GroupId = pb.GroupId
	s.Flags = pb.Flags
	s.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		s.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	return s
}

type OverrideShift struct {
	Id               int64  `json:"id,omitempty"`
	OrganizationId   int64  `json:"organizationId,omitempty"`
	CreatedById      *int64 `json:"createdById,omitempty"`
	UpdatedById      *int64 `json:"updatedById,omitempty"`
	WorkerId         *int64 `json:"workerId,omitempty"`
	CreatedAt        string `json:"createdAt,omitempty"`
	LastUpdatedAt    string `json:"lastUpdatedAt,omitempty"`
	VerifiedAt       string `json:"verifiedAt,omitempty"`
	CompletedAt      string `json:"completedAt,omitempty"`
	StartsOn         string `json:"startsOn,omitempty"`
	EndsOn           string `json:"endsOn,omitempty"`
	Name             string `json:"name,omitempty"`
	Day              int32  `json:"day,omitempty"`
	OverrideClockIn  string `json:"overrideClockIn,omitempty"`
	OverrideClockOut string `json:"overrideClockOut,omitempty"`
	GroupId          string `json:"groupId,omitempty"`
	Flags            uint32 `json:"flags,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
}

func (s *OverrideShift) TranslatePb(pb *pb.OverrideShift) *OverrideShift {
	s.Id = pb.Id
	s.OrganizationId = pb.OrganizationId
	s.CreatedById = pb.CreatedById
	s.UpdatedById = pb.UpdatedById
	s.WorkerId = pb.WorkerId
	s.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	s.Name = pb.Name
	s.Day = int32(pb.Day.Number())
	s.OverrideClockIn = pb.OverrideClockIn.AsTime().Format(time.RFC3339Nano)
	s.OverrideClockOut = pb.OverrideClockOut.AsTime().Format(time.RFC3339Nano)
	s.GroupId = pb.GroupId
	s.Flags = pb.Flags
	s.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		s.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.VerifiedAt != nil {
		s.VerifiedAt = pb.VerifiedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.CompletedAt != nil {
		s.CompletedAt = pb.CompletedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.StartsOn != nil {
		s.StartsOn = pb.StartsOn.AsTime().Format(time.RFC3339Nano)
	}

	if pb.EndsOn != nil {
		s.EndsOn = pb.EndsOn.AsTime().Format(time.RFC3339Nano)
	}

	return s
}

type Role struct {
	Id             int64  `json:"id,omitempty"`
	CreatedById    *int64 `json:"createdById,omitempty"`
	UpdatedById    *int64 `json:"updatedById,omitempty"`
	OrganizationId int64  `json:"organizationId,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	LastUpdatedAt  string `json:"lastUpdatedAt,omitempty"`
	Name           string `json:"name,omitempty"`
	Flags          uint32 `json:"flags,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}

func (r *Role) TranslatePb(pb *pb.Role) *Role {
	r.Id = pb.Id
	r.CreatedById = pb.CreatedById
	r.UpdatedById = pb.UpdatedById
	r.OrganizationId = pb.OrganizationId
	r.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	r.Name = pb.Name
	r.Flags = pb.Flags
	r.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		r.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	return r
}

type RoleShiftRelation struct {
	Id              int64 `json:"id,omitempty"`
	RoleId          int64 `json:"roleId,omitempty"`
	StandardShiftId int64 `json:"standardShiftId,omitempty"`
}

func (r *RoleShiftRelation) TranslatePb(pb *pb.RoleShiftRelation) *RoleShiftRelation {
	r.Id = pb.Id
	r.RoleId = pb.RoleId
	r.StandardShiftId = pb.StandardShiftId

	return r
}

type Payroll struct {
	Id             int64   `json:"id,omitempty"`
	CreatedById    int64   `json:"createdById,omitempty"`
	UpdatedById    int64   `json:"updatedById,omitempty"`
	VerifiedById   *int64  `json:"verifiedById,omitempty"`
	OrganizationId int64   `json:"organizationId,omitempty"`
	CreatedAt      string  `json:"createdAt,omitempty"`
	LastUpdatedAt  string  `json:"lastUpdatedAt,omitempty"`
	VerifiedAt     *string `json:"verifiedAt,omitempty"`
	Total          float32 `json:"total,omitempty"`
	Flags          uint32  `json:"flags,omitempty"`
	Uuid           string  `json:"uuid,omitempty"`
}

func (p *Payroll) TranslatePb(pb *pb.Payroll) *Payroll {
	p.Id = pb.Id
	p.CreatedById = pb.CreatedById
	p.UpdatedById = pb.UpdatedById
	p.VerifiedById = pb.VerifiedById
	p.OrganizationId = pb.OrganizationId
	p.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	p.Total = pb.Total
	p.Flags = pb.Flags
	p.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		p.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.VerifiedAt != nil {
		verifiedAt := pb.VerifiedAt.AsTime().Format(time.RFC3339Nano)
		p.VerifiedAt = &verifiedAt
	}

	return p
}

type PayrollCompensationRelation struct {
	Id             int64 `json:"id,omitempty"`
	PayrollId      int64 `json:"payrollId,omitempty"`
	CompensationId int64 `json:"compensationId,omitempty"`
}

func (p *PayrollCompensationRelation) TranslatePb(pb *pb.PayrollCompensationRelation) *PayrollCompensationRelation {
	p.Id = pb.Id
	p.PayrollId = pb.PayrollId
	p.CompensationId = pb.CompensationId

	return p
}

type Organization struct {
	Id               int64   `json:"id,omitempty"`
	CreatedById      *int64  `json:"createdById,omitempty"`
	UpdatedById      *int64  `json:"updatedById,omitempty"`
	CreatedAt        string  `json:"createdAt,omitempty"`
	LastUpdatedAt    string  `json:"lastUpdatedAt,omitempty"`
	Industry         int32   `json:"industry,omitempty"`
	OverrideIndustry *string `json:"overrideIndustry,omitempty"`
	Name             string  `json:"name,omitempty"`
	Username         string  `json:"username,omitempty"`
	Flags            uint32  `json:"flags,omitempty"`
	Uuid             string  `json:"uuid,omitempty"`
}

func (o *Organization) TranslatePb(pb *pb.Organization) *Organization {
	o.Id = pb.Id
	o.CreatedById = pb.CreatedById
	o.UpdatedById = pb.UpdatedById
	o.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	o.Industry = int32(pb.Industry.Number())
	o.OverrideIndustry = pb.OverrideIndustry
	o.Name = pb.Name
	o.Flags = pb.Flags
	o.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		o.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	// Whenever the pb.Username field is empty, generate a random uuid as its new username.
	if pb.GetUsername() == "" {
		uuidStr := strings.ReplaceAll(uuid.NewString(), "-", "")[0:16]
		id, _ := strconv.ParseUint(uuidStr, 16, 64)
		o.Username = fmt.Sprintf("%d", id)
	}

	return o
}

type Member struct {
	Id                    int64  `json:"id,omitempty"`
	OrganizationId        int64  `json:"organizationId,omitempty"`
	WorkerId              int64  `json:"workerId,omitempty"`
	HiredById             *int64 `json:"hiredBy,omitempty"`
	CreatedAt             string `json:"createdAt,omitempty"`
	LastUpdatedAt         string `json:"lastUpdatedAt,omitempty"`
	HiredAt               string `json:"hiredAt,omitempty"`
	SuspendedAt           string `json:"suspendedAt,omitempty"`
	LeaveAt               string `json:"leaveAt,omitempty"`
	TerminatedAt          string `json:"terminatedAt,omitempty"`
	ReturnedAt            string `json:"returnedAt,omitempty"`
	ScheduledSuspensionAt string `json:"scheduledSuspensionAt,omitempty"`
	Flags                 uint32 `json:"flags,omitempty"`
	Uuid                  string `json:"uuid,omitempty"`
}

func (m *Member) TranslatePb(pb *pb.Member) *Member {
	m.Id = pb.Id
	m.OrganizationId = pb.OrganizationId
	m.WorkerId = pb.WorkerId
	m.HiredById = pb.HiredById
	m.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	m.Flags = pb.Flags
	m.Uuid = pb.Uuid

	if pb.LastUpdatedAt != nil {
		m.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.HiredAt != nil {
		m.HiredAt = pb.HiredAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.SuspendedAt != nil {
		m.SuspendedAt = pb.HiredAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.LeaveAt != nil {
		m.LeaveAt = pb.LeaveAt.AsTime().String()
	}

	if pb.TerminatedAt != nil {
		m.TerminatedAt = pb.TerminatedAt.AsTime().String()
	}

	if pb.ReturnedAt != nil {
		m.ReturnedAt = pb.ReturnedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.ScheduledSuspensionAt != nil {
		m.ScheduledSuspensionAt = pb.ScheduledSuspensionAt.AsTime().String()
	}

	return m
}

type JoinRequest struct {
	Id             int64  `json:"id,omitempty"`
	WorkerId       int64  `json:"workerId,omitempty"`
	OrganizationId int64  `json:"organizationId,omitempty"`
	SenderType     int32  `json:"senderType"`
	CreatedAt      string `json:"createdAt,omitempty"`
	ExpiredAt      string `json:"expiredAt,omitempty"`
	Flags          uint32 `json:"flags,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}

func (j *JoinRequest) TranslatePb(pb *pb.JoinRequest) *JoinRequest {
	j.Id = pb.Id
	j.WorkerId = pb.WorkerId
	j.OrganizationId = pb.OrganizationId
	j.SenderType = int32(pb.SenderType)
	j.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	j.ExpiredAt = pb.ExpiredAt.AsTime().Format(time.RFC3339Nano)
	j.Flags = pb.Flags
	j.Uuid = pb.Uuid

	return j
}

type OrganizationPendingRequestRelation struct {
	Id             int64 `json:"id,omitempty"`
	OrganizationId int64 `json:"organizationId,omitempty"`
	RequestId      int64 `json:"requestId,omitempty"`
}

func (o *OrganizationPendingRequestRelation) TranslatePb(pb *pb.OrganizationPendingRequestRelation) *OrganizationPendingRequestRelation {
	o.Id = pb.Id
	o.OrganizationId = pb.OrganizationId
	o.RequestId = pb.RequestId

	return o
}

type WorkerPendingRequestRelation struct {
	Id        int64 `json:"id,omitempty"`
	WorkerId  int64 `json:"workerId,omitempty"`
	RequestId int64 `json:"requestId,omitempty"`
}

func (w *WorkerPendingRequestRelation) TranslatePb(pb *pb.WorkerPendingRequestRelation) *WorkerPendingRequestRelation {
	w.Id = pb.Id
	w.WorkerId = pb.WorkerId
	w.RequestId = pb.RequestId

	return w
}

type Deduction struct {
	Id            int64   `json:"id,omitempty"`
	CreatedById   *int64  `json:"createdById,omitempty"`
	UpdatedById   *int64  `json:"updatedById,omitempty"`
	WorkerId      *int64  `json:"workerId,omitempty"`
	CreatedAt     string  `json:"createdAt,omitempty"`
	LastUpdatedAt string  `json:"lastUpdatedAt,omitempty"`
	EffectiveAt   string  `json:"effectiveAt,omitempty"`
	Name          string  `json:"name,omitempty"`
	Value         float32 `json:"value,omitempty"`
	Flags         uint32  `json:"flags,omitempty"`
	Uuid          string  `json:"uuid,omitempty"`
	Typ           int32   `json:"typ,omitempty"`
}

func (w *Deduction) TranslatePb(pb *pb.Deduction) *Deduction {
	w.Id = pb.Id
	w.CreatedById = pb.CreatedById
	w.UpdatedById = pb.UpdatedById
	w.WorkerId = pb.WorkerId
	w.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	w.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	w.Name = pb.Name
	w.Value = pb.Value
	w.Flags = pb.Flags
	w.Uuid = pb.Uuid
	w.Typ = int32(pb.Typ)

	return w
}

type Compensation struct {
	Id             int64   `json:"id,omitempty"`
	OrganizationId int64   `json:"organizationId,omitempty"`
	WorkerId       int64   `json:"workerId,omitempty"`
	CreatedById    *int64  `json:"createdById,omitempty"`
	UpdatedById    *int64  `json:"updatedById,omitempty"`
	CreatedAt      string  `json:"createdAt,omitempty"`
	LastUpdatedAt  string  `json:"lastUpdatedAt,omitempty"`
	PaidAt         string  `json:"paidAt,omitempty"`
	ApprovedAt     string  `json:"approvedAt,omitempty"`
	RejectedAt     string  `json:"rejectedAt,omitempty"`
	PeriodStart    string  `json:"periodStart,omitempty"`
	PeriodEnd      string  `json:"periodEnd,omitempty"`
	Gvalue         float32 `json:"gvalue,omitempty"`
	Avalue         float32 `json:"avalue,omitempty"`
	Dvalue         float32 `json:"dvalue,omitempty"`
	Value          float32 `json:"value,omitempty"`
	Flags          uint32  `json:"flags,omitempty"`
	Uuid           string  `json:"uuid,omitempty"`
	Scheme         int32   `json:"scheme,omitempty"`
}

func (c *Compensation) TranslatePb(pb *pb.Compensation) *Compensation {
	c.Id = pb.Id
	c.OrganizationId = pb.OrganizationId
	c.WorkerId = pb.WorkerId
	c.CreatedById = pb.CreatedById
	c.UpdatedById = pb.UpdatedById
	c.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	c.PeriodStart = pb.PeriodStart.AsTime().Format(time.RFC3339Nano)
	c.PeriodEnd = pb.PeriodEnd.AsTime().Format(time.RFC3339Nano)
	c.Gvalue = pb.Gvalue
	c.Avalue = pb.Avalue
	c.Dvalue = pb.Dvalue
	c.Value = pb.Value
	c.Flags = pb.Flags
	c.Uuid = pb.Uuid
	c.Scheme = int32(pb.Scheme)

	if pb.LastUpdatedAt != nil {
		c.LastUpdatedAt = pb.LastUpdatedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.PaidAt != nil {
		c.PaidAt = pb.PaidAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.ApprovedAt != nil {
		c.ApprovedAt = pb.ApprovedAt.AsTime().Format(time.RFC3339Nano)
	}

	if pb.RejectedAt != nil {
		c.RejectedAt = pb.RejectedAt.AsTime().Format(time.RFC3339Nano)
	}

	return c
}

type CompensationAdditionRelation struct {
	Id             int64 `json:"id,omitempty"`
	CompensationId int64 `json:"compensationId,omitempty"`
	AdditionId     int64 `json:"additionId,omitempty"`
}

func (c *CompensationAdditionRelation) TranslatePb(pb *pb.CompensationAdditionRelation) *CompensationAdditionRelation {
	c.Id = pb.Id
	c.CompensationId = pb.CompensationId
	c.AdditionId = pb.AdditionId

	return c
}

type CompensationDeductionRelation struct {
	Id             int64 `json:"id,omitempty"`
	CompensationId int64 `json:"compensationId,omitempty"`
	DeductionId    int64 `json:"deductionId,omitempty"`
}

func (c *CompensationDeductionRelation) TranslatePb(pb *pb.CompensationDeductionRelation) *CompensationDeductionRelation {
	c.Id = pb.Id
	c.CompensationId = pb.CompensationId
	c.DeductionId = pb.DeductionId

	return c
}

type Attendance struct {
	Id            int64  `json:"id,omitempty"`
	WorkerId      int64  `json:"workerId,omitempty"`
	ShiftId       *int64 `json:"shiftId,omitempty"`
	OshiftId      *int64 `json:"oshiftId,omitempty"`
	CreatedById   *int64 `json:"createdById,omitempty"`
	UpdatedById   *int64 `json:"updatedById,omitempty"`
	CreatedAt     string `json:"createdAt,omitempty"`
	LastUpdatedAt string `json:"lastUpdatedAt,omitempty"`
	ClockIn       string `json:"clockIn,omitempty"`
	ClockOut      string `json:"clockOut,omitempty"`
	Computed      int64  `json:"computed,omitempty"`
	UnderTime     int64  `json:"underTime,omitempty"`
	OverTime      int64  `json:"overTime,omitempty"`
	LateTime      int64  `json:"lateTime,omitempty"`
	BreakTime     int64  `json:"breakTime,omitempty"`
	Flags         uint32 `json:"flags,omitempty"`
	Uuid          string `json:"uuid,omitempty"`
}

func (a *Attendance) TranslatePb(pb *pb.Attendance) *Attendance {
	a.Id = pb.Id
	a.WorkerId = pb.WorkerId
	a.ShiftId = pb.ShiftId
	a.OshiftId = pb.OshiftId
	a.CreatedById = pb.CreatedById
	a.UpdatedById = pb.UpdatedById
	a.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	a.ClockIn = pb.ClockIn.AsTime().Format(time.RFC3339Nano)
	a.ClockOut = pb.ClockOut.AsTime().Format(time.RFC3339Nano)
	a.Computed = pb.Computed
	a.UnderTime = pb.UnderTime
	a.OverTime = pb.OverTime
	a.LateTime = pb.LateTime
	a.BreakTime = pb.BreakTime
	a.Flags = pb.Flags
	a.Uuid = pb.Uuid

	return a
}

type Addition struct {
	Id            int64   `json:"id"`
	CreatedById   *int64  `json:"createdById,omitempty"`
	UpdatedById   *int64  `json:"updatedByid,omitempty"`
	WorkerId      *int64  `json:"workerId,omitempty"`
	CreatedAt     string  `json:"createdAt,omitempty"`
	LastUpdatedAt string  `json:"lastUpdatedAt,omitempty"`
	EffectiveAt   string  `json:"effectiveAt,omitempty"`
	Name          string  `json:"name,omitempty"`
	Value         float32 `json:"value,omitempty"`
	Flags         uint32  `json:"flags,omitempty"`
	Uuid          string  `json:"uuid,omitempty"`
	Typ           int32   `json:"typ,omitempty"`
}

func (a *Addition) TranslatePb(pb *pb.Addition) *Addition {
	a.Id = pb.Id
	a.CreatedById = pb.CreatedById
	a.UpdatedById = pb.UpdatedById
	a.WorkerId = pb.WorkerId
	a.CreatedAt = pb.CreatedAt.AsTime().Format(time.RFC3339Nano)
	a.Name = pb.Name
	a.Value = pb.Value
	a.Flags = pb.Flags
	a.Uuid = pb.Uuid
	a.Typ = int32(pb.Typ)

	if pb.LastUpdatedAt != nil {
		a.LastUpdatedAt = pb.LastUpdatedAt.AsTime().String()
	}

	if pb.EffectiveAt != nil {
		a.EffectiveAt = pb.EffectiveAt.AsTime().String()
	}

	return a
}
