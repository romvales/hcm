# A specification document for the HRMS

### A collection of interfaces to where our system will be depending on

## Table of contents

1. Terminologies
2. A General Rule of Thumb

## Terminologies

Human Resource Management System (HRMS), is a software designed to automate various HR functions within an organization. **It reduces the manual labor in managing employee data, payroll, benefits administration, recruitment, performance management, time tracking, and other functions related to human resource.**

## A General Rule of Thumb

Each interfaces that will be described in this source must be followed accordingly so that there will be no confusions throughout the development cycle and helps make your HRMS robust.

Below are the types, enums and interfaces that will be outlined:

1. Worker

- Role
- Team
- Worker
- WorkerIdentityCard
- WorkerAddress
- enum:WorkerPayCycle
- enum:WorkerAddressType
- enum:WorkerGender
- enum:WorkerStatus
- enum:WorkerType
- enum:TeamStatus
- enum:RoleStatus

- HCMRoleService

1. createRole(name: string)
2. getRoleById()
3. deleteRoleById()
4. saveRole()
5. changeRoleName()
6. changeRoleStatus()
7. getRoleShifts()
8. addStandardShift()
9. updateStandardShift()
10. containsCompleteStandardShift()
11. isActive()
12. isInactive()
13. isOnReview()
14. isTerminated()  
15. isDisabled()

- HCMTeamService

1. createTeam(name: string)
2. getTeamById(teamId: number)
3. deleteTeamById(teamId: number)
4. saveTeam(team: Team)
5. changeTeamName(name: string)
6. changeTeamStatus(status: TeamStatus)
7. getWorkerMembers()
8. addWorkerToTeam(team: Team, worker: Worker)
9. removeWorkerFromTeam(team: Team, worker: Worker)
10. isActive()
11. isInactive()
12. isOnReview()
13. isTerminated()
14. isDisabled()

- HCMWorkerService

1. createWorker(name: { firstName: string, middleName: string, lastName: string }, email: string, mobileNumber?: string, birthdate?: number)
2. getWorkerById(workerId: number)
3. deleteWorkerById(workerId: number)
4. saveWorker(worker: Worker)
5. changeWorkerStatus(status: WorkerStatus)
6. changeWorkerType(type: WorkerType)
7. changeWorkerRole(roleId: number)
8. changeWorkerTeam(teamId: number)
9. changeWorkerPayCycle(cycle: WorkerPayCycle)
10. suspend(worker: Worker)
11. terminate(worker: Worker)
12. resign(worker: Worker)
13. addWorkerAddress(address: WorkerAddress)
14. addIdentityCards(cards: WorkerIdentityCard[])
15. getWorkerType()
16. getWorkerStatus()
17. getAddresses()
18. getIdentityCards()
19. getOrganization()
20. getRole()
21. getTeam()
22. hasOverridenStandardRoleShift()
23. isWorkerHired(worker: Worker)
24. isWorkerOnLeave(worker: Worker)
25. isWorkerRemote(worker: Worker)
26. isWorkerOnline(worker: Worker)
27. isWorkerRemotelyOnline(worker: Worker)
28. isWorkerOffline(worker: Worker)
29. isWorkerAway(worker: Worker)
30. isWorkerSuspended(worker: Worker)
31. isWorkerOnCall(worker: Worker)

- #BaseOrganizationEntityStatusChecker(Entity = unknown)

1. isActive(entity: Entity)
2. isInactive(entity: Entity)
3. isOnReview(entity: Entity)
4. isTerminated(entity: Entity)

2. Shifts

- StandardShift
- OverrideShift
- enum:StandardShiftDay

- HCMWorkerShiftService

3. Attendance

- Attendance
- enum:AttendanceType
- enum:AttendanceStatus
- enum:AttendancePerformanceLabel
- enum:AttendanceClockInType
- enum:AttendanceClockOutType

- HCMAttendanceService

1. createAttendance(worker: Worker, clockIn: number)
2. getAttendanceById(attendanceId: number)
3. removeAttendanceById(attendanceId: number)
4. saveAttendance(attendance: Attendance)
4. changeStatus(attendance: Attendance, status: AttendanceStatus)
5. changeType(attendance: Attendance, type: AttendanceType)
6. changePerfLabel(attendance: Attendance, label: AttendancePerfomanceLabel)
7. changeClockInType(attendance: Attendance, type: AttendanceClockInType)
8. changeClockOutType(attendance: Attendance, type: AttendanceClockOutType)
9. clockIn(worker: Worker, type: AttendanceClockInType)
10. clockOut(worker: Worker, type: AttendanceClockOutType)
11. getShift(attendance: Attendance)
12. isLate(attendance: Attendance)
13. isOverride(attendance: Attendance)
14. isHoliday(attendance: Attendance)
15. isBreak(attendance: Attendance)

4. Payroll

- Payroll
- enum:PayrollStatus
- enum:PayrollPayCycleType

5. Compute





6. Compensation

- WorkerPayInfo
- WorkerPayInfoOverride
- Compensation
- Addition
- Deduction
- enum:AdditionStatus
- enum:AdditionType
- enum:DeductionStatus
- enum:DeductionType
- enum:WorkerPayInfoType

- HCMWorkerPayInfoService

- HCMCompensationService

- HCMAdditionService

- HCMDeductionService

6. Organization

- Organization
- PendingJoinRequest
- enum:OrganizationStatus
- enum:OrganizationIndustry
- enum:PendingJoinRequestInvitationType
- enum:PendingJoinRequestStatus

- HCMOrganizationService

TODO: Every organization should have a configuration that enables leader to change and modify certain conditions within the organization.

1. createOrg(name: string, industry: OrganizationIndustry, overrideIndustry?: string)
2. getOrgById(organizationId: number)
3. removeOrgById(organizationId: number)
4. saveOrg(org: Organization)
5. changeOrgName(org: Organization, name: string)
6. changeOrgIndustry(org: Organization, industry: OrganizationIndustry)
7. changeOrgStatus(org: Organization, status: OrganizationStatus)
8. removeWorkerById(org: Organization, workerId: number)
9. addToOrgbyId(org: Organization, workerId: number)
10. getOrgCreator()
11. isActive()
12. isInactive()
13. isSuspended()
14. isDissolved()

- #HCMPendingJoinRequestService(Entity = unknown)

1. sendRequest(entity: Entity, recepientId: number)
2. cancelRequest(entity: Entity, recepientId: number)
3. getPendingRequests(entity: Entity)
4. acceptPendingRequest(entity: Entity, requestId: number)
5. declinePendingRequest(entity: Entity, requestId: number)

7. Client