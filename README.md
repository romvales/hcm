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

- HCMWorkerService

1. createWorker(name: { firstName: string, middleName: string, lastName: string }, email: string, mobileNumber: string)
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
22. hasOverridenStandardsShift()
23. isWorkerHired(worker: Worker)
24. isWorkerOnLeave(worker: Worker)
25. isWorkerRemote(worker: Worker)
26. isWorkerOnline(worker: Worker)
27. isWorkerRemotelyOnline(worker: Worker)
28. isWorkerOffline(worker: Worker)
29. isWorkerAway(worker: Worker)
30. isWorkerSuspended(worker: Worker)
31. isWorkerOnCall(worker: Worker)

2. Shifts

- StandardShift
- OverrideShift
- enum:StandardShiftDay

- HCMWorkerShiftService

3. Attendance

- Attendance
- enum:AttendanceType
- enum:AttendanceStatus

- HCMAttendanceService

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

- HCMPendingJoinRequestService(Entity = unknown)

1. sendRequest(entity: Entity, recepientId: number)
2. cancelRequest(entity: Entity, recepientId: number)
3. getPendingRequests(entity: Entity)
4. acceptPendingRequest(entity: Entity, requestId: number)
5. declinePendingRequest(entity: Entity, requestId: number)